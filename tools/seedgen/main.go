package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Link struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}
type Node struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Aka             []string `json:"aka"`
	BornYear        *int     `json:"bornYear"`
	DiedYear        *int     `json:"diedYear"`
	ActiveStartYear *int     `json:"activeStartYear"`
	ActiveEndYear   *int     `json:"activeEndYear"`
	Tags            []string `json:"tags"`
	Notability      int      `json:"notability"`
	Links           []Link   `json:"links"`
}
type Evidence struct {
	URL   string `json:"url"`
	Label string `json:"label"`
}
type Edge struct {
	ID       string     `json:"id"`
	SourceID string     `json:"sourceId"`
	TargetID string     `json:"targetId"`
	Type     string     `json:"type"`
	StartYr  *int       `json:"startYear"`
	EndYr    *int       `json:"endYear"`
	Weight   int        `json:"weight"`
	Summary  string     `json:"summary"`
	Evidence []Evidence `json:"evidence"`
}
type SeedData struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

func yr(y int) *int { return &y }

func wiki(name string) []Link {
	return []Link{{Label: "Wikipedia", URL: "https://en.wikipedia.org/wiki/" + name}}
}

func ev(name string) []Evidence {
	return []Evidence{{URL: "https://en.wikipedia.org/wiki/" + name, Label: "Wikipedia"}}
}

func main() {
	// Load existing seed.json
	existing, err := os.ReadFile("../../app/src/lib/data/seed.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading seed.json: %v\n", err)
		os.Exit(1)
	}
	var data SeedData
	if err := json.Unmarshal(existing, &data); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing seed.json: %v\n", err)
		os.Exit(1)
	}

	existingIDs := map[string]bool{}
	for _, n := range data.Nodes {
		existingIDs[n.ID] = true
	}

	newNodes := allNewNodes()
	for _, n := range newNodes {
		if !existingIDs[n.ID] {
			data.Nodes = append(data.Nodes, n)
			existingIDs[n.ID] = true
		}
	}

	newEdges := allNewEdges(len(data.Edges))
	data.Edges = append(data.Edges, newEdges...)

	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile("../../app/src/lib/data/seed.json", append(out, '\n'), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Done: %d nodes, %d edges\n", len(data.Nodes), len(data.Edges))
}
