package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Node represents a comedian in the graph
type Node struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	AKA             []string          `json:"aka"`
	BornYear        *int              `json:"bornYear"`
	DiedYear        *int              `json:"diedYear"`
	ActiveStartYear *int              `json:"activeStartYear"`
	ActiveEndYear   *int              `json:"activeEndYear"`
	Tags            []string          `json:"tags"`
	Notability      int               `json:"notability"`
	Links           []Link            `json:"links"`
	Source          string            `json:"source,omitempty"`
	ExternalIDs     map[string]string `json:"externalIds,omitempty"`
}

// Edge represents an alliance between two comedians
type Edge struct {
	ID        string `json:"id"`
	SourceID  string `json:"sourceId"`
	TargetID  string `json:"targetId"`
	Type      string `json:"type"`
	StartYear *int   `json:"startYear"`
	EndYear   *int   `json:"endYear"`
	Weight    int    `json:"weight"`
	Summary   string `json:"summary"`
	Evidence  []Link `json:"evidence"`
	Source    string `json:"source,omitempty"`
}

// Link is a labeled URL
type Link struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

// Dataset is the top-level data structure
type Dataset struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "validate":
		cmdValidate()
	case "merge":
		cmdMerge()
	case "stats":
		cmdStats()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("🃏 bunt — CartoJester Data Pipeline")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Usage: bunt <command> [args]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  validate <file>          Validate a JSON dataset")
	fmt.Println("  merge <seed> <auto> <out> Merge seed and auto datasets")
	fmt.Println("  stats <file>             Show dataset statistics")
}

func loadDataset(path string) (*Dataset, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	var ds Dataset
	if err := json.Unmarshal(data, &ds); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}
	return &ds, nil
}

func cmdValidate() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: bunt validate <file>")
		os.Exit(1)
	}
	path := os.Args[2]
	ds, err := loadDataset(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ %v\n", err)
		os.Exit(1)
	}

	errors := validate(ds)
	if len(errors) > 0 {
		fmt.Printf("❌ %d validation errors in %s:\n", len(errors), path)
		for _, e := range errors {
			fmt.Printf("  • %s\n", e)
		}
		os.Exit(1)
	}
	fmt.Printf("✅ %s is valid (%d nodes, %d edges)\n", path, len(ds.Nodes), len(ds.Edges))
}

func validate(ds *Dataset) []string {
	var errs []string
	ids := map[string]bool{}
	for _, n := range ds.Nodes {
		if n.ID == "" {
			errs = append(errs, "node with empty id")
		}
		if ids[n.ID] {
			errs = append(errs, fmt.Sprintf("duplicate node id: %s", n.ID))
		}
		ids[n.ID] = true
		if n.Name == "" {
			errs = append(errs, fmt.Sprintf("node %s has empty name", n.ID))
		}
		if n.Notability < 1 || n.Notability > 5 {
			errs = append(errs, fmt.Sprintf("node %s notability %d out of range [1,5]", n.ID, n.Notability))
		}
	}
	for _, e := range ds.Edges {
		if !ids[e.SourceID] {
			errs = append(errs, fmt.Sprintf("edge %s references unknown source: %s", e.ID, e.SourceID))
		}
		if !ids[e.TargetID] {
			errs = append(errs, fmt.Sprintf("edge %s references unknown target: %s", e.ID, e.TargetID))
		}
		if e.Weight < 1 || e.Weight > 5 {
			errs = append(errs, fmt.Sprintf("edge %s weight %d out of range [1,5]", e.ID, e.Weight))
		}
	}
	return errs
}

func cmdMerge() {
	if len(os.Args) < 5 {
		fmt.Fprintln(os.Stderr, "Usage: bunt merge <seed> <auto> <out>")
		os.Exit(1)
	}
	seedPath, autoPath, outPath := os.Args[2], os.Args[3], os.Args[4]

	seed, err := loadDataset(seedPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ %v\n", err)
		os.Exit(1)
	}
	auto, err := loadDataset(autoPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ %v\n", err)
		os.Exit(1)
	}

	merged := mergeDatasets(seed, auto)
	out, _ := json.MarshalIndent(merged, "", "  ")
	if err := os.WriteFile(outPath, out, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "❌ writing %s: %v\n", outPath, err)
		os.Exit(1)
	}
	fmt.Printf("✅ Merged → %s (%d nodes, %d edges)\n", outPath, len(merged.Nodes), len(merged.Edges))
}

func mergeDatasets(seed, auto *Dataset) *Dataset {
	nodeMap := map[string]Node{}
	for _, n := range seed.Nodes {
		n.Source = "seed"
		nodeMap[n.ID] = n
	}
	for _, n := range auto.Nodes {
		if _, exists := nodeMap[n.ID]; !exists {
			n.Source = "auto"
			nodeMap[n.ID] = n
		}
	}
	edgeMap := map[string]Edge{}
	for _, e := range seed.Edges {
		e.Source = "seed"
		edgeMap[e.ID] = e
	}
	for _, e := range auto.Edges {
		if _, exists := edgeMap[e.ID]; !exists {
			e.Source = "auto"
			edgeMap[e.ID] = e
		}
	}

	merged := &Dataset{}
	for _, n := range nodeMap {
		merged.Nodes = append(merged.Nodes, n)
	}
	for _, e := range edgeMap {
		merged.Edges = append(merged.Edges, e)
	}
	return merged
}

func cmdStats() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: bunt stats <file>")
		os.Exit(1)
	}
	ds, err := loadDataset(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ %v\n", err)
		os.Exit(1)
	}
	fmt.Println("🃏 Dataset Statistics")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("  Nodes: %d\n", len(ds.Nodes))
	fmt.Printf("  Edges: %d\n", len(ds.Edges))

	tagCounts := map[string]int{}
	for _, n := range ds.Nodes {
		for _, t := range n.Tags {
			tagCounts[t]++
		}
	}
	fmt.Println("  Tags:")
	for t, c := range tagCounts {
		fmt.Printf("    %s: %d\n", t, c)
	}

	typeCounts := map[string]int{}
	for _, e := range ds.Edges {
		typeCounts[e.Type]++
	}
	fmt.Println("  Edge types:")
	for t, c := range typeCounts {
		fmt.Printf("    %s: %d\n", t, c)
	}
}
