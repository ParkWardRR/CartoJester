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

type NewsSource struct {
	Name string
	URL  string
	Type string // rss, html
}

type DiscoveredRelation struct {
	Comic1  string
	Comic2  string
	Context string
	Source  string
}

func fetchPage(url string) (string, error) {
	client := &http.Client{Timeout: 15 * time.Second}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "CartoJester-NewsScanner/1.0 (comedy research)")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
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
	}

	existingEdges := map[string]bool{}
	for _, e := range seed.Edges {
		existingEdges[e.SourceID+"|"+e.TargetID] = true
		existingEdges[e.TargetID+"|"+e.SourceID] = true
	}

	// Build name regex patterns for notable comedians (3+ notability)
	namePatterns := map[string]*regexp.Regexp{}
	for _, n := range seed.Nodes {
		if n.Notability >= 3 {
			escapedName := regexp.QuoteMeta(n.Name)
			namePatterns[n.ID] = regexp.MustCompile(`(?i)\b` + escapedName + `\b`)
		}
	}

	// Comedy news sources to scan
	sources := []NewsSource{
		// Dead Frog - Comedy news aggregator
		{Name: "Dead Frog", URL: "https://deadfrog.us/", Type: "html"},
		// Chortle - UK's comedy guide
		{Name: "Chortle", URL: "https://www.chortle.co.uk/news", Type: "html"},
		{Name: "Chortle Features", URL: "https://www.chortle.co.uk/features", Type: "html"},
		// The Comedy Bureau
		{Name: "Comedy Bureau", URL: "https://thecomedybureau.com/", Type: "html"},
		// Vulture comedy section
		{Name: "Vulture Comedy", URL: "https://www.vulture.com/comedy/", Type: "html"},
		// Paste Magazine comedy
		{Name: "Paste Comedy", URL: "https://www.pastemagazine.com/comedy", Type: "html"},
		// Laugh Button
		{Name: "Laugh Button", URL: "https://thelaughbutton.com/", Type: "html"},
		// Too Many Comics
		{Name: "Splitsider", URL: "https://www.vulture.com/tags/comedy-pairings/", Type: "html"},
	}

	allContent := map[string]string{}

	for _, src := range sources {
		fmt.Printf("Scanning %s (%s)...\n", src.Name, src.URL)
		content, err := fetchPage(src.URL)
		if err != nil {
			fmt.Printf("  ❌ Error: %v\n", err)
			continue
		}
		allContent[src.Name] = content
		fmt.Printf("  ✅ Got %d bytes\n", len(content))
		time.Sleep(1 * time.Second)
	}

	// Scan each page for comedian co-occurrences
	pairCount := map[string]int{}
	pairContext := map[string]string{}
	pairSource := map[string]string{}

	for srcName, content := range allContent {
		// Find all comedians mentioned
		mentioned := []string{}
		for id, pattern := range namePatterns {
			if pattern.MatchString(content) {
				mentioned = append(mentioned, id)
			}
		}

		fmt.Printf("\n%s: found %d comedians mentioned\n", srcName, len(mentioned))
		if len(mentioned) > 0 && len(mentioned) <= 30 {
			for _, id := range mentioned {
				fmt.Printf("  - %s\n", idToName[id])
			}
		}

		// Look for co-mentions in nearby text
		// Split content into paragraphs/segments for proximity analysis
		stripHTML := regexp.MustCompile(`<[^>]*>`)
		plainText := stripHTML.ReplaceAllString(content, " ")
		// Split into ~500 char windows
		windows := splitIntoWindows(plainText, 500, 100)

		for _, window := range windows {
			windowMentions := []string{}
			for id, pattern := range namePatterns {
				if pattern.MatchString(window) {
					windowMentions = append(windowMentions, id)
				}
			}

			// Record pairs from this window
			for i := 0; i < len(windowMentions); i++ {
				for j := i + 1; j < len(windowMentions); j++ {
					a, b := windowMentions[i], windowMentions[j]
					if a > b {
						a, b = b, a
					}
					key := a + "|" + b
					if existingEdges[key] {
						continue
					}
					pairCount[key]++
					if _, ok := pairContext[key]; !ok {
						pairContext[key] = truncate(window, 200)
						pairSource[key] = srcName
					}
				}
			}
		}
	}

	// Process discovered pairs
	fmt.Printf("\n═══ Comedy News Discoveries ═══\n")
	addedEdges := 0
	nextNum := len(seed.Edges) + 1

	for key, count := range pairCount {
		parts := strings.SplitN(key, "|", 2)
		n1 := idToName[parts[0]]
		n2 := idToName[parts[1]]

		if count >= 2 {
			fmt.Printf("  ✅ %s ↔ %s (co-mentioned %dx in %s)\n", n1, n2, count, pairSource[key])
			fmt.Printf("     Context: %s\n", truncate(pairContext[key], 120))

			e := Edge{
				ID:       fmt.Sprintf("ens%03d", nextNum),
				SourceID: parts[0],
				TargetID: parts[1],
				Type:     "collaboration",
				Weight:   min(count, 3),
				Summary:  fmt.Sprintf("News: co-mentioned in %s (%d+ contexts)", pairSource[key], count),
				Evidence: []Link{},
			}
			seed.Edges = append(seed.Edges, e)
			existingEdges[key] = true
			nextNum++
			addedEdges++
		} else {
			fmt.Printf("  ⚠️  %s ↔ %s (1 mention) — skipped\n", n1, n2)
		}
	}

	fmt.Printf("\n═══ News scan results ═══\n")
	fmt.Printf("  Sources scanned: %d\n", len(allContent))
	fmt.Printf("  Pairs discovered: %d\n", len(pairCount))
	fmt.Printf("  Edges added (2+ mentions): %d\n", addedEdges)
	fmt.Printf("  Total edges now: %d\n", len(seed.Edges))

	out, _ := json.MarshalIndent(seed, "", "  ")
	os.WriteFile(seedPath, append(out, '\n'), 0644)
	fmt.Println("  Wrote seed.json")
}

func splitIntoWindows(text string, windowSize, step int) []string {
	windows := []string{}
	for i := 0; i < len(text)-windowSize; i += step {
		end := i + windowSize
		if end > len(text) {
			end = len(text)
		}
		windows = append(windows, text[i:end])
	}
	if len(text) <= windowSize {
		windows = append(windows, text)
	}
	return windows
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
