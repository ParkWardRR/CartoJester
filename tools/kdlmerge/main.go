package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Seed data structures
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

// Parsed KDL alliance
type KDLAlliance struct {
	ID         string
	Source     string
	Target     string
	Type       string
	Weight     float64
	StartYr    string
	EndYr      string
	BasisGroup string
	BasisTag   string
	Derived    bool
	Summary    string
}

func intPtr(v int) *int { return &v }

func main() {
	kdlPath := "kdl import/comedians_atlas_fixed.kdl"
	seedPath := "app/src/lib/data/seed.json"

	// Load existing seed
	seedRaw, err := os.ReadFile(seedPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read seed: %v\n", err)
		os.Exit(1)
	}
	var seed SeedData
	if err := json.Unmarshal(seedRaw, &seed); err != nil {
		fmt.Fprintf(os.Stderr, "parse seed: %v\n", err)
		os.Exit(1)
	}

	// Build node ID set
	nodeIDs := map[string]bool{}
	for _, n := range seed.Nodes {
		nodeIDs[n.ID] = true
	}

	// Build existing edge key set (source-target or target-source)
	existingEdges := map[string]bool{}
	for _, e := range seed.Edges {
		existingEdges[e.SourceID+"|"+e.TargetID] = true
		existingEdges[e.TargetID+"|"+e.SourceID] = true
	}

	// Parse KDL for derived alliances
	file, err := os.Open(kdlPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open kdl: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var alliances []KDLAlliance
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	reAlliance := regexp.MustCompile(`^alliance\s+"([^"]+)"\s*\{`)
	reProp := regexp.MustCompile(`^\s+(\S+)\s+(.+)$`)

	var current *KDLAlliance

	for scanner.Scan() {
		line := scanner.Text()

		if m := reAlliance.FindStringSubmatch(line); m != nil {
			if current != nil {
				alliances = append(alliances, *current)
			}
			current = &KDLAlliance{ID: m[1]}
			continue
		}

		if current != nil {
			if strings.TrimSpace(line) == "}" {
				alliances = append(alliances, *current)
				current = nil
				continue
			}

			if m := reProp.FindStringSubmatch(line); m != nil {
				key := m[1]
				val := strings.Trim(m[2], `"`)
				switch key {
				case "source":
					current.Source = val
				case "target":
					current.Target = val
				case "type":
					current.Type = val
				case "weight":
					w, _ := strconv.ParseFloat(val, 64)
					current.Weight = w
				case "year-start", "start-year":
					current.StartYr = val
				case "year-end", "end-year":
					current.EndYr = val
				case "basis-group":
					current.BasisGroup = val
				case "basis-tag":
					current.BasisTag = val
				case "derived":
					current.Derived = val == "true"
				case "summary":
					current.Summary = val
				}
			}
		}
	}

	fmt.Printf("Parsed %d alliances from KDL\n", len(alliances))

	// Filter: only accept derived edges that are group-based (not tag-based)
	// Tag-based edges (shared-tag) are too low quality
	var newGroupEdges []KDLAlliance
	var newTagEdges int
	var existingSkipped int
	var missingNodes int

	for _, a := range alliances {
		if !a.Derived {
			continue // skip non-derived (already in our seed)
		}

		// Skip tag-based edges - too noisy
		if a.Type == "shared-tag" || a.BasisTag != "" {
			newTagEdges++
			continue
		}

		// Skip if source or target doesn't exist
		if !nodeIDs[a.Source] || !nodeIDs[a.Target] {
			missingNodes++
			continue
		}

		// Skip if edge already exists
		key := a.Source + "|" + a.Target
		if existingEdges[key] {
			existingSkipped++
			continue
		}

		newGroupEdges = append(newGroupEdges, a)
		existingEdges[key] = true
		existingEdges[a.Target+"|"+a.Source] = true
	}

	fmt.Printf("\nFiltering results:\n")
	fmt.Printf("  Group-based edges (GOOD): %d\n", len(newGroupEdges))
	fmt.Printf("  Tag-based edges (SKIPPED): %d\n", newTagEdges)
	fmt.Printf("  Already existing (SKIPPED): %d\n", existingSkipped)
	fmt.Printf("  Missing nodes (SKIPPED): %d\n", missingNodes)

	// Convert and add new edges
	nextEdgeNum := len(seed.Edges) + 1
	added := 0
	for _, a := range newGroupEdges {
		e := Edge{
			ID:       fmt.Sprintf("eg%03d", nextEdgeNum),
			SourceID: a.Source,
			TargetID: a.Target,
			Type:     "troupe", // Map same-movement to troupe (closest existing type)
			Weight:   1,
			Summary:  a.Summary,
			Evidence: []Link{},
		}

		// Parse years
		if a.StartYr != "" && a.StartYr != "null" {
			y, _ := strconv.Atoi(a.StartYr)
			if y > 0 {
				e.StartYr = intPtr(y)
			}
		}
		if a.EndYr != "" && a.EndYr != "null" {
			y, _ := strconv.Atoi(a.EndYr)
			if y > 0 {
				e.EndYr = intPtr(y)
			}
		}

		seed.Edges = append(seed.Edges, e)
		nextEdgeNum++
		added++
	}

	fmt.Printf("\nAdded %d new group-based edges\n", added)
	fmt.Printf("Total edges now: %d\n", len(seed.Edges))

	// Write updated seed
	out, _ := json.MarshalIndent(seed, "", "  ")
	if err := os.WriteFile(seedPath, append(out, '\n'), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "write seed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Wrote updated seed.json (%d nodes, %d edges)\n", len(seed.Nodes), len(seed.Edges))
}
