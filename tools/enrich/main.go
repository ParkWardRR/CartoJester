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

type seedLink struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type seedNode struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Aka             []string   `json:"aka"`
	BornYear        *int       `json:"bornYear"`
	DiedYear        *int       `json:"diedYear"`
	ActiveStartYear *int       `json:"activeStartYear"`
	ActiveEndYear   *int       `json:"activeEndYear"`
	Tags            []string   `json:"tags"`
	Notability      int        `json:"notability"`
	Links           []seedLink `json:"links"`
}

func intPtr(v int) *int { return &v }

// Use Wikipedia REST API to get basic page info + extract
type wikiSummary struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Extract     string `json:"extract"`
}

func fetchWikiSummary(name string) (*wikiSummary, error) {
	slug := strings.ReplaceAll(name, " ", "_")
	url := fmt.Sprintf("https://en.wikipedia.org/api/rest_v1/page/summary/%s", slug)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "CartoJester/1.0 (https://github.com/ParkWardRR/CartoJester)")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	var ws wikiSummary
	json.Unmarshal(body, &ws)
	return &ws, nil
}

var reYear = regexp.MustCompile(`\b(born|Born)\b[^)]*?\b((?:19|20)\d{2})\b`)
var reDied = regexp.MustCompile(`\b((?:19|20)\d{2})\s*[–-]\s*((?:19|20)\d{2})\b`)
var reSlug = regexp.MustCompile(`[^a-z0-9]+`)

func extractYears(extract string) (*int, *int) {
	// Try "1950 – 2020" pattern first (for deceased)
	if m := reDied.FindStringSubmatch(extract); len(m) == 3 {
		var born, died int
		fmt.Sscanf(m[1], "%d", &born)
		fmt.Sscanf(m[2], "%d", &died)
		return intPtr(born), intPtr(died)
	}
	// Try "born 1985" pattern
	if m := reYear.FindStringSubmatch(extract); len(m) == 3 {
		var born int
		fmt.Sscanf(m[2], "%d", &born)
		return intPtr(born), nil
	}
	return nil, nil
}

func guessTags(desc, extract string) []string {
	lower := strings.ToLower(desc + " " + extract)
	tags := []string{}
	if strings.Contains(lower, "stand-up") || strings.Contains(lower, "standup") || strings.Contains(lower, "stand up") {
		tags = append(tags, "standup")
	}
	if strings.Contains(lower, "actor") || strings.Contains(lower, "actress") {
		tags = append(tags, "film")
	}
	if strings.Contains(lower, "television") || strings.Contains(lower, " tv ") || strings.Contains(lower, "sitcom") {
		tags = append(tags, "tv")
	}
	if strings.Contains(lower, "sketch") || strings.Contains(lower, "snl") || strings.Contains(lower, "saturday night live") {
		tags = append(tags, "snl")
	}
	if strings.Contains(lower, "improv") || strings.Contains(lower, "improvisation") {
		tags = append(tags, "improv")
	}
	if strings.Contains(lower, "podcast") {
		tags = append(tags, "podcast")
	}
	if strings.Contains(lower, "panel") || strings.Contains(lower, "panel show") {
		tags = append(tags, "panel")
	}
	if strings.Contains(lower, "writer") || strings.Contains(lower, "screenwriter") {
		tags = append(tags, "tv")
	}
	if strings.Contains(lower, "animation") || strings.Contains(lower, "animated") || strings.Contains(lower, "voice actor") {
		tags = append(tags, "animation")
	}
	if len(tags) == 0 {
		tags = append(tags, "standup")
	}
	// Deduplicate
	seen := map[string]bool{}
	var unique []string
	for _, t := range tags {
		if !seen[t] {
			seen[t] = true
			unique = append(unique, t)
		}
	}
	return unique
}

func main() {
	raw, err := os.ReadFile("comedians_500_new.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "read names: %v\n", err)
		os.Exit(1)
	}
	var allNames []string
	for _, line := range strings.Split(string(raw), "\n") {
		name := strings.TrimSpace(line)
		if name != "" {
			allNames = append(allNames, name)
		}
	}
	if len(allNames) > 500 {
		allNames = allNames[:500]
	}

	// Load existing seed
	seedRaw, _ := os.ReadFile("app/src/lib/data/seed.json")
	var seed struct {
		Nodes []json.RawMessage `json:"nodes"`
		Edges []json.RawMessage `json:"edges"`
	}
	json.Unmarshal(seedRaw, &seed)
	existingNames := map[string]bool{}
	existingIDs := map[string]bool{}
	for _, n := range seed.Nodes {
		var node struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		}
		json.Unmarshal(n, &node)
		existingNames[node.Name] = true
		existingIDs[node.ID] = true
	}

	// Filter out already-existing
	var toProcess []string
	for _, name := range allNames {
		if !existingNames[name] {
			toProcess = append(toProcess, name)
		}
	}
	fmt.Printf("Processing %d new names (of %d total, %d already in seed)\n", len(toProcess), len(allNames), len(allNames)-len(toProcess))

	slugify := func(name string) string {
		s := strings.ToLower(name)
		s = reSlug.ReplaceAllString(s, "-")
		return strings.Trim(s, "-")
	}

	wikiSlug := func(name string) string {
		return strings.ReplaceAll(name, " ", "_")
	}

	added := 0
	errors := 0
	for i, name := range toProcess {
		id := slugify(name)
		if existingIDs[id] {
			id = id + "-2"
		}
		n := seedNode{
			ID:         id,
			Name:       name,
			Aka:        []string{},
			Tags:       []string{"standup"},
			Notability: 2,
			Links:      []seedLink{{Label: "Wikipedia", URL: "https://en.wikipedia.org/wiki/" + wikiSlug(name)}},
		}

		// Try to enrich from Wikipedia REST API
		ws, err := fetchWikiSummary(name)
		if err == nil && ws != nil && ws.Extract != "" {
			born, died := extractYears(ws.Extract)
			n.BornYear = born
			n.DiedYear = died
			if born != nil {
				n.ActiveStartYear = intPtr(*born + 20)
			}
			if died != nil {
				n.ActiveEndYear = died
			}
			n.Tags = guessTags(ws.Description, ws.Extract)
		} else {
			errors++
		}

		raw, _ := json.Marshal(n)
		seed.Nodes = append(seed.Nodes, raw)
		existingNames[name] = true
		existingIDs[id] = true
		added++

		if (i+1)%25 == 0 || i == len(toProcess)-1 {
			fmt.Printf("  Progress: %d/%d (added %d, errors %d)\n", i+1, len(toProcess), added, errors)
		}
		// Rate limit: ~10 req/sec
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("\nAdding %d new nodes to seed.json\n", added)
	outSeed, _ := json.MarshalIndent(seed, "", "  ")
	os.WriteFile("app/src/lib/data/seed.json", append(outSeed, '\n'), 0644)
	fmt.Printf("Done! seed.json now has %d total nodes\n", len(seed.Nodes))
}
