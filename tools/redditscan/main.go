package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// Reddit API structures (public, no auth needed)
type RedditResponse struct {
	Data RedditData `json:"data"`
}

type RedditData struct {
	Children []RedditChild `json:"children"`
	After    string        `json:"after"`
}

type RedditChild struct {
	Data RedditPost `json:"data"`
}

type RedditPost struct {
	Title       string  `json:"title"`
	Selftext    string  `json:"selftext"`
	URL         string  `json:"url"`
	Permalink   string  `json:"permalink"`
	Score       int     `json:"score"`
	Created     float64 `json:"created_utc"`
	Author      string  `json:"author"`
	NumComments int     `json:"num_comments"`
}

type Link struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Node struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Aka         []string `json:"aka"`
	BornYear    *int     `json:"bornYear"`
	DiedYear    *int     `json:"diedYear"`
	ActiveStart *int     `json:"activeStartYear"`
	ActiveEnd   *int     `json:"activeEndYear"`
	Tags        []string `json:"tags"`
	Notability  int      `json:"notability"`
	Links       []Link   `json:"links"`
}

type Edge struct {
	ID       string `json:"id"`
	SourceID string `json:"sourceId"`
	TargetID string `json:"targetId"`
	Type     string `json:"type"`
	StartYr  *int   `json:"startYear"`
	EndYr    *int   `json:"endYear"`
	Weight   int    `json:"weight"`
	Summary  string `json:"summary"`
	Evidence []Link `json:"evidence"`
}

type SeedData struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type DiscoveredRelation struct {
	Comic1  string
	Comic2  string
	Context string
	Source  string
	Score   int
}

func fetchSubreddit(sub string, sort string, limit int) ([]RedditPost, error) {
	url := fmt.Sprintf("https://old.reddit.com/r/%s/%s.json?limit=%d&raw_json=1", sub, sort, limit)

	client := &http.Client{Timeout: 15 * time.Second}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "CartoJester-Scanner/1.0 (comedy research bot)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(body[:200]))
	}

	body, _ := io.ReadAll(resp.Body)
	var rr RedditResponse
	json.Unmarshal(body, &rr)

	posts := []RedditPost{}
	for _, c := range rr.Data.Children {
		posts = append(posts, c.Data)
	}
	return posts, nil
}

func searchSubreddit(sub string, query string, limit int) ([]RedditPost, error) {
	url := fmt.Sprintf("https://old.reddit.com/r/%s/search.json?q=%s&restrict_sr=on&sort=relevance&t=all&limit=%d&raw_json=1",
		sub, strings.ReplaceAll(query, " ", "+"), limit)

	client := &http.Client{Timeout: 15 * time.Second}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "CartoJester-Scanner/1.0 (comedy research bot)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var rr RedditResponse
	json.Unmarshal(body, &rr)

	posts := []RedditPost{}
	for _, c := range rr.Data.Children {
		posts = append(posts, c.Data)
	}
	return posts, nil
}

func main() {
	seedPath := "app/src/lib/data/seed.json"
	raw, _ := os.ReadFile(seedPath)
	var seed SeedData
	json.Unmarshal(raw, &seed)

	// Build name→ID lookup
	nameToID := map[string]string{}
	idToName := map[string]string{}
	for _, n := range seed.Nodes {
		nameToID[strings.ToLower(n.Name)] = n.ID
		idToName[n.ID] = n.Name
		for _, aka := range n.Aka {
			nameToID[strings.ToLower(aka)] = n.ID
		}
	}

	existingEdges := map[string]bool{}
	for _, e := range seed.Edges {
		existingEdges[e.SourceID+"|"+e.TargetID] = true
		existingEdges[e.TargetID+"|"+e.SourceID] = true
	}

	// Subreddits to scan
	subreddits := []string{
		"StandUpComedy",
		"Standup",
		"comedy",
		"comedians",
	}

	allPosts := []RedditPost{}

	// 1. Fetch top posts from comedy subreddits
	for _, sub := range subreddits {
		fmt.Printf("Scanning r/%s (top/all)...\n", sub)
		posts, err := fetchSubreddit(sub, "top", 100)
		if err != nil {
			fmt.Printf("  Error: %v\n", err)
		} else {
			allPosts = append(allPosts, posts...)
			fmt.Printf("  Got %d posts\n", len(posts))
		}
		time.Sleep(2 * time.Second) // Reddit rate limiting
	}

	// 2. Search for specific relationship queries
	searchQueries := []string{
		"comedy duo",
		"open for",
		"toured with",
		"collaboration special",
		"podcast together",
		"writing partner",
		"comedy friends",
		"mentor comedy",
	}

	for _, q := range searchQueries {
		fmt.Printf("Searching: \"%s\"...\n", q)
		posts, err := searchSubreddit("StandUpComedy", q, 25)
		if err != nil {
			fmt.Printf("  Error: %v\n", err)
		} else {
			allPosts = append(allPosts, posts...)
			fmt.Printf("  Got %d posts\n", len(posts))
		}
		time.Sleep(2 * time.Second) // Rate limiting
	}

	fmt.Printf("\nTotal posts to analyze: %d\n", len(allPosts))

	// 3. Scan all posts for comedian name co-occurrences
	// Build a fast name matcher
	namePatterns := map[string]*regexp.Regexp{}
	for _, n := range seed.Nodes {
		if n.Notability >= 3 {
			// Build pattern: match full name, case insensitive
			escapedName := regexp.QuoteMeta(n.Name)
			namePatterns[n.ID] = regexp.MustCompile(`(?i)\b` + escapedName + `\b`)
		}
	}

	discovered := []DiscoveredRelation{}
	pairCount := map[string]int{}

	for _, post := range allPosts {
		text := post.Title + " " + post.Selftext

		// Find all comedians mentioned in this post
		mentioned := []string{}
		for id, pattern := range namePatterns {
			if pattern.MatchString(text) {
				mentioned = append(mentioned, id)
			}
		}

		// If 2+ comedians mentioned, record relationships
		for i := 0; i < len(mentioned); i++ {
			for j := i + 1; j < len(mentioned); j++ {
				a, b := mentioned[i], mentioned[j]
				if a > b {
					a, b = b, a
				}
				key := a + "|" + b

				if existingEdges[key] {
					continue // Already have this edge
				}

				pairCount[key]++

				if pairCount[key] == 1 {
					discovered = append(discovered, DiscoveredRelation{
						Comic1:  a,
						Comic2:  b,
						Context: truncate(post.Title, 120),
						Source:  "https://reddit.com" + post.Permalink,
						Score:   post.Score,
					})
				}
			}
		}
	}

	// Sort by frequency (multi-post mentions are higher confidence)
	fmt.Printf("\n═══ Discovered relationships ═══\n")
	fmt.Printf("Unique pairs found: %d\n\n", len(discovered))

	// Only add relationships mentioned in 2+ posts or with high scores
	addedEdges := 0
	nextNum := len(seed.Edges) + 1

	for _, d := range discovered {
		key := d.Comic1 + "|" + d.Comic2
		count := pairCount[key]

		n1 := idToName[d.Comic1]
		n2 := idToName[d.Comic2]

		if count >= 2 || d.Score >= 100 {
			fmt.Printf("  ✅ %s ↔ %s (mentioned %dx, score %d)\n", n1, n2, count, d.Score)
			fmt.Printf("     Context: %s\n", d.Context)

			e := Edge{
				ID:       fmt.Sprintf("erd%03d", nextNum),
				SourceID: d.Comic1,
				TargetID: d.Comic2,
				Type:     "collaboration",
				Weight:   min(count, 3),
				Summary:  fmt.Sprintf("Reddit: co-mentioned in %d+ posts (e.g. \"%s\")", count, truncate(d.Context, 80)),
				Evidence: []Link{{Label: "Reddit", URL: d.Source}},
			}
			seed.Edges = append(seed.Edges, e)
			existingEdges[key] = true
			nextNum++
			addedEdges++
		} else {
			fmt.Printf("  ⚠️  %s ↔ %s (only %d mention, score %d) — skipped\n", n1, n2, count, d.Score)
		}
	}

	fmt.Printf("\n═══ Reddit scan results ═══\n")
	fmt.Printf("  Posts scanned: %d\n", len(allPosts))
	fmt.Printf("  Pairs discovered: %d\n", len(discovered))
	fmt.Printf("  Edges added (high confidence): %d\n", addedEdges)
	fmt.Printf("  Total edges now: %d\n", len(seed.Edges))

	out, _ := json.MarshalIndent(seed, "", "  ")
	os.WriteFile(seedPath, append(out, '\n'), 0644)
	fmt.Println("  Wrote seed.json")
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
