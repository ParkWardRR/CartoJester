package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

type WikiSummary struct {
	Title       string `json:"title"`
	Extract     string `json:"extract"`
	Description string `json:"description"`
}

func intPtr(v int) *int { return &v }

func fetchSummary(name string) (*WikiSummary, error) {
	encoded := url.PathEscape(strings.ReplaceAll(name, " ", "_"))
	u := fmt.Sprintf("https://en.wikipedia.org/api/rest_v1/page/summary/%s", encoded)

	client := &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("GET", u, nil)
	req.Header.Set("User-Agent", "CartoJester/1.0 (twesh@users.noreply.github.com)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var ws WikiSummary
	json.Unmarshal(body, &ws)
	return &ws, nil
}

func extractYears(text string) (born *int, died *int) {
	// "born April 5, 1975" or "(1975–"
	reBorn := regexp.MustCompile(`(?:born\s+(?:\w+\s+\d+,\s+)?|[\(]\s*)(\d{4})`)
	reDied := regexp.MustCompile(`(?:died|–\s*)(\d{4})\s*\)`)

	if m := reBorn.FindStringSubmatch(text); m != nil {
		y := 0
		fmt.Sscanf(m[1], "%d", &y)
		if y > 1800 && y < 2020 {
			born = intPtr(y)
		}
	}
	if m := reDied.FindStringSubmatch(text); m != nil {
		y := 0
		fmt.Sscanf(m[1], "%d", &y)
		if y > 1800 && y < 2030 {
			died = intPtr(y)
		}
	}
	return
}

func guessTags(desc, extract string) []string {
	text := strings.ToLower(desc + " " + extract)
	tags := []string{}
	checks := map[string][]string{
		"standup":   {"stand-up", "standup", "comedian"},
		"film":      {"actor", "actress", "film", "movie"},
		"tv":        {"television", "sitcom", "tv show", "series"},
		"snl":       {"saturday night live", "snl"},
		"improv":    {"improv", "second city", "groundlings", "ucb"},
		"sketch":    {"sketch"},
		"podcast":   {"podcast"},
		"animation": {"voice actor", "animated", "voice role"},
		"panel":     {"panel show", "panel"},
		"latenight": {"late-night", "late night", "talk show host"},
	}

	for tag, keywords := range checks {
		for _, kw := range keywords {
			if strings.Contains(text, kw) {
				tags = append(tags, tag)
				break
			}
		}
	}

	if len(tags) == 0 {
		tags = append(tags, "standup")
	}
	return tags
}

func slugify(name string) string {
	s := strings.ToLower(name)
	s = strings.ReplaceAll(s, "'", "")
	s = strings.ReplaceAll(s, "'", "")
	s = strings.ReplaceAll(s, ".", "")

	reg := regexp.MustCompile(`[^a-z0-9]+`)
	s = reg.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

type NewComedian struct {
	WikiName   string
	Slug       string
	Notability int
}

func main() {
	seedPath := "app/src/lib/data/seed.json"

	raw, _ := os.ReadFile(seedPath)
	var seed SeedData
	json.Unmarshal(raw, &seed)

	existingIDs := map[string]bool{}
	for _, n := range seed.Nodes {
		existingIDs[n.ID] = true
	}

	// Missing comedians we need for edges
	newComics := []NewComedian{
		{"John C. Reilly", "john-c-reilly", 4},
		{"Rob Schneider", "rob-schneider", 3},
		{"Kevin James", "kevin-james", 3},
		{"Chevy Chase", "chevy-chase", 4},
		{"Molly Shannon", "molly-shannon", 3},
		{"Rachel Dratch", "rachel-dratch", 3},
		{"Andy Richter", "andy-richter", 3},
		{"Louis C.K.", "louis-ck", 5},
		{"Brendan Schaub", "brendan-schaub", 2},
		{"Bill Hicks", "bill-hicks", 5},
		{"Stephen Merchant", "stephen-merchant", 3},
		{"Karl Pilkington", "karl-pilkington", 3},
		{"Tim Vine", "tim-vine", 3},
		{"Larry Charles", "larry-charles", 3},
		{"Randall Park", "randall-park", 3},
		{"Seth Rogen", "seth-rogen", 5},
		{"James Franco", "james-franco", 3},
		{"Jonah Hill", "jonah-hill", 4},
		{"Evan Goldberg", "evan-goldberg", 3},
		{"Abbi Jacobson", "abbi-jacobson", 3},
		{"Ilana Glazer", "ilana-glazer", 3},
		{"Chelsea Peretti", "chelsea-peretti", 3},
		{"Carl Reiner", "carl-reiner", 5},
		{"Gene Wilder", "gene-wilder", 5},
		{"Bing Crosby", "bing-crosby", 4},
		{"Matt Rife", "matt-rife", 3},
		{"Martin Short", "martin-short", 4},
		{"Jeff Daniels", "jeff-daniels", 3},
		{"Sean Lock", "sean-lock", 4},
		{"Jon Richardson", "jon-richardson", 3},
		{"Ed Gamble", "ed-gamble", 3},
		{"Bernie Mac", "bernie-mac", 5},
		{"D.L. Hughley", "dl-hughley", 4},
		{"Joan Rivers", "joan-rivers", 5},
		{"Phyllis Diller", "phyllis-diller", 4},
		{"Lily Tomlin", "lily-tomlin", 5},
		{"Chris Pratt", "chris-pratt", 4},
		{"Nick Offerman", "nick-offerman", 3},
		{"Megan Mullally", "megan-mullally", 3},
		{"Zach Galifianakis", "zach-galifianakis", 4},
		{"Mel Brooks", "mel-brooks", 5},
		{"Justin Timberlake", "justin-timberlake", 3},
		{"Rob Beckett", "rob-beckett", 3},
		{"Ed Byrne", "ed-byrne", 3},
		{"Julian Barratt", "julian-barratt", 3},
		{"Steve Carell", "steve-carell", 5},
		{"Jeff Ross", "jeff-ross", 3},
		{"Neal Brennan", "neal-brennan", 3},
		{"Isla Fisher", "isla-fisher", 3},
		{"Jane Fonda", "jane-fonda", 4},
		{"Judd Apatow", "judd-apatow", 4},
		{"Jordan Klepper", "jordan-klepper", 3},
		{"Nick Cannon", "nick-cannon", 3},
		{"Ricky Gervais", "ricky-gervais", 5},
		{"Russell Brand", "russell-brand", 4},
		{"Jim Carrey", "jim-carrey", 5},
		{"Carrie Brownstein", "carrie-brownstein", 3},
		{"Larry Wilmore", "larry-wilmore", 3},
	}

	addedNodes := 0
	for _, nc := range newComics {
		if existingIDs[nc.Slug] {
			fmt.Printf("  SKIP (exists): %s\n", nc.WikiName)
			continue
		}

		summary, err := fetchSummary(nc.WikiName)
		if err != nil {
			fmt.Printf("  ERR: %s: %v\n", nc.WikiName, err)
			continue
		}

		born, died := extractYears(summary.Extract)
		tags := guessTags(summary.Description, summary.Extract)

		node := Node{
			ID:         nc.Slug,
			Name:       nc.WikiName,
			Aka:        []string{},
			BornYear:   born,
			DiedYear:   died,
			Tags:       tags,
			Notability: nc.Notability,
			Links: []Link{
				{Label: "Wikipedia", URL: fmt.Sprintf("https://en.wikipedia.org/wiki/%s", url.PathEscape(strings.ReplaceAll(nc.WikiName, " ", "_")))},
			},
		}

		seed.Nodes = append(seed.Nodes, node)
		existingIDs[nc.Slug] = true
		addedNodes++
		fmt.Printf("  + %s (born=%v, tags=%v)\n", nc.WikiName, born, tags)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("\nAdded %d new comedian nodes\n", addedNodes)
	fmt.Printf("Total nodes now: %d\n", len(seed.Nodes))

	out, _ := json.MarshalIndent(seed, "", "  ")
	os.WriteFile(seedPath, append(out, '\n'), 0644)
	fmt.Println("Wrote seed.json")
}
