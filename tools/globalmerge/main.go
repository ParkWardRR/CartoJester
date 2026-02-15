package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
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

type NewComedian struct {
	WikiName   string
	Slug       string
	Tags       []string
	Notability int
	BornYear   *int
	DiedYear   *int
	WikiURL    string
}

type NewEdge struct {
	Source  string
	Target  string
	Type    string
	Summary string
	Weight  int
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

	existingEdges := map[string]bool{}
	for _, e := range seed.Edges {
		existingEdges[e.SourceID+"|"+e.TargetID] = true
		existingEdges[e.TargetID+"|"+e.SourceID] = true
	}

	// ═══════════════════════════════════════
	// INDIA comedians
	// ═══════════════════════════════════════
	indiaComics := []NewComedian{
		// Historical Bollywood
		{WikiName: "Johnny Lever", Slug: "johnny-lever", Tags: []string{"india", "film", "standup"}, Notability: 5, WikiURL: "https://en.wikipedia.org/wiki/Johnny_Lever"},
		{WikiName: "Kishore Kumar", Slug: "kishore-kumar", Tags: []string{"india", "film"}, Notability: 5, BornYear: intPtr(1929), DiedYear: intPtr(1987), WikiURL: "https://en.wikipedia.org/wiki/Kishore_Kumar"},
		{WikiName: "Deven Verma", Slug: "deven-verma", Tags: []string{"india", "film"}, Notability: 3, BornYear: intPtr(1937), DiedYear: intPtr(2014), WikiURL: "https://en.wikipedia.org/wiki/Deven_Verma"},
		{WikiName: "I. S. Johar", Slug: "i-s-johar", Tags: []string{"india", "film"}, Notability: 3, BornYear: intPtr(1920), DiedYear: intPtr(1984), WikiURL: "https://en.wikipedia.org/wiki/I._S._Johar"},
		{WikiName: "Jagdeep", Slug: "jagdeep", Tags: []string{"india", "film"}, Notability: 4, BornYear: intPtr(1939), DiedYear: intPtr(2020), WikiURL: "https://en.wikipedia.org/wiki/Jagdeep"},
		{WikiName: "Jaspal Bhatti", Slug: "jaspal-bhatti", Tags: []string{"india", "tv", "film"}, Notability: 4, BornYear: intPtr(1955), DiedYear: intPtr(2012), WikiURL: "https://en.wikipedia.org/wiki/Jaspal_Bhatti"},
		{WikiName: "Johnny Walker", Slug: "johnny-walker-actor", Tags: []string{"india", "film"}, Notability: 4, BornYear: intPtr(1926), DiedYear: intPtr(2003), WikiURL: "https://en.wikipedia.org/wiki/Johnny_Walker_(actor)"},
		{WikiName: "Kader Khan", Slug: "kader-khan", Tags: []string{"india", "film"}, Notability: 4, BornYear: intPtr(1937), DiedYear: intPtr(2018), WikiURL: "https://en.wikipedia.org/wiki/Kader_Khan"},
		{WikiName: "Keshto Mukherjee", Slug: "keshto-mukherjee", Tags: []string{"india", "film"}, Notability: 3, BornYear: intPtr(1934), DiedYear: intPtr(2012), WikiURL: "https://en.wikipedia.org/wiki/Keshto_Mukherjee"},
		{WikiName: "Mehmood", Slug: "mehmood", Tags: []string{"india", "film"}, Notability: 5, BornYear: intPtr(1932), DiedYear: intPtr(2004), WikiURL: "https://en.wikipedia.org/wiki/Mehmood_(actor)"},
		// Modern India standup
		{WikiName: "Vir Das", Slug: "vir-das", Tags: []string{"india", "standup", "film"}, Notability: 5, BornYear: intPtr(1979), WikiURL: "https://en.wikipedia.org/wiki/Vir_Das"},
		{WikiName: "Zakir Khan", Slug: "zakir-khan", Tags: []string{"india", "standup"}, Notability: 4, BornYear: intPtr(1987), WikiURL: "https://en.wikipedia.org/wiki/Zakir_Khan_(comedian)"},
		{WikiName: "Kapil Sharma", Slug: "kapil-sharma", Tags: []string{"india", "tv", "standup"}, Notability: 5, BornYear: intPtr(1981), WikiURL: "https://en.wikipedia.org/wiki/Kapil_Sharma_(comedian)"},
		{WikiName: "Sunil Grover", Slug: "sunil-grover", Tags: []string{"india", "tv"}, Notability: 4, BornYear: intPtr(1977), WikiURL: "https://en.wikipedia.org/wiki/Sunil_Grover"},
		{WikiName: "Tanmay Bhat", Slug: "tanmay-bhat", Tags: []string{"india", "standup", "podcast"}, Notability: 4, BornYear: intPtr(1987), WikiURL: "https://en.wikipedia.org/wiki/Tanmay_Bhat"},
		{WikiName: "Munawar Faruqui", Slug: "munawar-faruqui", Tags: []string{"india", "standup"}, Notability: 3, BornYear: intPtr(1992), WikiURL: "https://en.wikipedia.org/wiki/Munawar_Faruqui"},
		{WikiName: "Papa CJ", Slug: "papa-cj", Tags: []string{"india", "standup"}, Notability: 3, WikiURL: "https://en.wikipedia.org/wiki/Papa_CJ"},
		{WikiName: "Vadivelu", Slug: "vadivelu", Tags: []string{"india", "film"}, Notability: 4, BornYear: intPtr(1961), WikiURL: "https://en.wikipedia.org/wiki/Vadivelu"},
		{WikiName: "Yogi Babu", Slug: "yogi-babu", Tags: []string{"india", "film"}, Notability: 3, BornYear: intPtr(1985), WikiURL: "https://en.wikipedia.org/wiki/Yogi_Babu"},
		{WikiName: "Varun Grover", Slug: "varun-grover-writer", Tags: []string{"india", "standup"}, Notability: 3, WikiURL: "https://en.wikipedia.org/wiki/Varun_Grover_(writer)"},
		{WikiName: "Sorabh Pant", Slug: "sorabh-pant", Tags: []string{"india", "standup"}, Notability: 2, WikiURL: "https://en.wikipedia.org/wiki/Sorabh_Pant"},
		{WikiName: "Vasu Primlani", Slug: "vasu-primlani", Tags: []string{"india", "standup"}, Notability: 2, WikiURL: "https://en.wikipedia.org/wiki/Vasu_Primlani"},
	}

	// ═══════════════════════════════════════
	// SE ASIA comedians
	// ═══════════════════════════════════════
	seAsiaComics := []NewComedian{
		// Philippines
		{WikiName: "Dolphy", Slug: "dolphy", Tags: []string{"philippines", "film"}, Notability: 5, BornYear: intPtr(1928), DiedYear: intPtr(2012), WikiURL: "https://en.wikipedia.org/wiki/Dolphy"},
		{WikiName: "German Moreno", Slug: "german-moreno", Tags: []string{"philippines", "tv"}, Notability: 3, BornYear: intPtr(1933), DiedYear: intPtr(2016), WikiURL: "https://en.wikipedia.org/wiki/German_Moreno"},
		{WikiName: "Don Pepot", Slug: "don-pepot", Tags: []string{"philippines", "film"}, Notability: 3, WikiURL: "https://en.wikipedia.org/wiki/Don_Pepot"},
		{WikiName: "Elizabeth Ramsey", Slug: "elizabeth-ramsey", Tags: []string{"philippines", "standup"}, Notability: 3, WikiURL: "https://en.wikipedia.org/wiki/Elizabeth_Ramsey"},
		{WikiName: "Larry Silva", Slug: "larry-silva", Tags: []string{"philippines", "film"}, Notability: 2, WikiURL: ""},
		// Indonesia
		{WikiName: "Ateng", Slug: "ateng-actor", Tags: []string{"indonesia", "film"}, Notability: 3, WikiURL: "https://en.wikipedia.org/wiki/Ateng_(actor)"},
		{WikiName: "Dono", Slug: "dono", Tags: []string{"indonesia", "film"}, Notability: 3, WikiURL: "https://en.wikipedia.org/wiki/Dono"},
		{WikiName: "Tukul Arwana", Slug: "tukul-arwana", Tags: []string{"indonesia", "tv"}, Notability: 3, WikiURL: "https://en.wikipedia.org/wiki/Tukul_Arwana"},
		{WikiName: "Indra Birowo", Slug: "indra-birowo", Tags: []string{"indonesia", "tv"}, Notability: 2, WikiURL: "https://en.wikipedia.org/wiki/Indra_Birowo"},
		{WikiName: "Denny Cagur", Slug: "denny-cagur", Tags: []string{"indonesia", "tv"}, Notability: 2, WikiURL: "https://en.wikipedia.org/wiki/Denny_Cagur"},
		// Malaysia
		{WikiName: "Harith Iskander", Slug: "harith-iskander", Tags: []string{"malaysia", "standup"}, Notability: 4, WikiURL: "https://en.wikipedia.org/wiki/Harith_Iskander"},
		{WikiName: "Jason Leong", Slug: "jason-leong", Tags: []string{"malaysia", "standup"}, Notability: 3, WikiURL: "https://en.wikipedia.org/wiki/Jason_Leong"},
		{WikiName: "Douglas Lim", Slug: "douglas-lim", Tags: []string{"malaysia", "standup"}, Notability: 3, WikiURL: "https://en.wikipedia.org/wiki/Douglas_Lim"},
		{WikiName: "Nigel Ng", Slug: "nigel-ng", Tags: []string{"malaysia", "standup"}, Notability: 4, WikiURL: "https://en.wikipedia.org/wiki/Nigel_Ng"},
		{WikiName: "Hannan Azlan", Slug: "hannan-azlan", Tags: []string{"malaysia", "standup"}, Notability: 2, WikiURL: ""},
		// Singapore
		{WikiName: "Rishi Budhrani", Slug: "rishi-budhrani", Tags: []string{"singapore", "standup"}, Notability: 2, WikiURL: "https://en.wikipedia.org/wiki/Rishi_Budhrani"},
		{WikiName: "Marcus Chin", Slug: "marcus-chin", Tags: []string{"singapore", "tv"}, Notability: 3, WikiURL: "https://en.wikipedia.org/wiki/Marcus_Chin"},
		{WikiName: "Irene Ang", Slug: "irene-ang", Tags: []string{"singapore", "tv"}, Notability: 3, WikiURL: "https://en.wikipedia.org/wiki/Irene_Ang"},
		{WikiName: "Michelle Chong", Slug: "michelle-chong", Tags: []string{"singapore", "tv", "film"}, Notability: 3, WikiURL: "https://en.wikipedia.org/wiki/Michelle_Chong"},
	}

	allNewComics := append(indiaComics, seAsiaComics...)

	// ═══════════════════════════════════════
	// Add nodes
	// ═══════════════════════════════════════
	addedNodes := 0
	for _, nc := range allNewComics {
		if existingIDs[nc.Slug] {
			fmt.Printf("  SKIP node (exists): %s\n", nc.WikiName)
			continue
		}

		node := Node{
			ID:         nc.Slug,
			Name:       nc.WikiName,
			Aka:        []string{},
			BornYear:   nc.BornYear,
			DiedYear:   nc.DiedYear,
			Tags:       nc.Tags,
			Notability: nc.Notability,
			Links:      []Link{},
		}

		if nc.WikiURL != "" {
			node.Links = append(node.Links, Link{Label: "Wikipedia", URL: nc.WikiURL})
		}

		// Try Wikipedia enrichment for birth year if missing
		if nc.BornYear == nil && nc.WikiName != "" {
			summary, err := fetchSummary(nc.WikiName)
			if err == nil && summary != nil {
				// Simple year extraction from extract
				extract := summary.Extract
				for y := 1880; y < 2005; y++ {
					yStr := fmt.Sprintf("%d", y)
					if strings.Contains(extract, "born "+yStr) || strings.Contains(extract, "(born "+yStr) {
						year := y
						node.BornYear = &year
						break
					}
				}
			}
			time.Sleep(80 * time.Millisecond)
		}

		seed.Nodes = append(seed.Nodes, node)
		existingIDs[nc.Slug] = true
		addedNodes++
		fmt.Printf("  + %s (%v)\n", nc.WikiName, nc.Tags)
	}

	// ═══════════════════════════════════════
	// Add edges (India cohort + SE Asia cohort + World cohort)
	// ═══════════════════════════════════════
	newEdges := []NewEdge{
		// India Historical hub (johnny-lever as hub)
		{"deven-verma", "johnny-lever", "troupe", "India historical Bollywood comedy cohort", 1},
		{"i-s-johar", "johnny-lever", "troupe", "India historical Bollywood comedy cohort", 1},
		{"jagdeep", "johnny-lever", "troupe", "India historical Bollywood comedy cohort", 1},
		{"jaspal-bhatti", "johnny-lever", "troupe", "India historical Bollywood comedy cohort", 1},
		{"johnny-walker-actor", "johnny-lever", "troupe", "India historical Bollywood comedy cohort", 1},
		{"kader-khan", "johnny-lever", "troupe", "India historical Bollywood comedy cohort", 1},
		{"keshto-mukherjee", "johnny-lever", "troupe", "India historical Bollywood comedy cohort", 1},
		{"kishore-kumar", "johnny-lever", "troupe", "India historical Bollywood comedy cohort", 1},
		{"mehmood", "johnny-lever", "troupe", "India historical Bollywood comedy cohort", 1},
		// India cross-era links
		{"mehmood", "jagdeep", "collaboration", "Frequent Bollywood comedy co-stars in Hindi cinema", 3},
		{"kader-khan", "johnny-lever", "collaboration", "Bollywood comedy duo in numerous 90s films", 4},
		{"kapil-sharma", "sunil-grover", "collaboration", "Co-stars on The Kapil Sharma Show", 5},
		// India Modern hub (vir-das as hub)
		{"zakir-khan", "vir-das", "troupe", "India modern stand-up comedy cohort", 1},
		{"kapil-sharma", "vir-das", "troupe", "India modern comedy cohort", 1},
		{"sunil-grover", "vir-das", "troupe", "India modern comedy cohort", 1},
		{"tanmay-bhat", "vir-das", "troupe", "India modern stand-up cohort", 1},
		{"munawar-faruqui", "vir-das", "troupe", "India modern stand-up cohort", 1},
		{"papa-cj", "vir-das", "troupe", "India modern stand-up cohort", 1},
		{"vadivelu", "vir-das", "troupe", "India comedy cohort", 1},
		{"yogi-babu", "vir-das", "troupe", "India comedy cohort", 1},
		{"varun-grover-writer", "vir-das", "troupe", "India comedy cohort", 1},
		// India within-group real connections
		{"tanmay-bhat", "zakir-khan", "collaboration", "Both part of the Indian stand-up comedy circuit; Amazon specials", 3},
		{"sorabh-pant", "vir-das", "collaboration", "Both pioneered Indian stand-up comedy scene", 2},
		{"vasu-primlani", "vir-das", "collaboration", "Both pioneered Indian stand-up comedy scene", 2},

		// SE Asia Historical hub (dolphy as hub)
		{"german-moreno", "dolphy", "troupe", "SE Asia historical comedy cohort", 1},
		{"don-pepot", "dolphy", "troupe", "SE Asia historical comedy cohort", 1},
		{"elizabeth-ramsey", "dolphy", "troupe", "SE Asia historical comedy cohort", 1},
		{"larry-silva", "dolphy", "troupe", "SE Asia historical comedy cohort", 1},
		{"ateng-actor", "dolphy", "troupe", "SE Asia historical comedy cohort", 1},
		{"dono", "dolphy", "troupe", "SE Asia historical comedy cohort", 1},
		{"tukul-arwana", "dolphy", "troupe", "SE Asia historical comedy cohort", 1},
		{"indra-birowo", "dolphy", "troupe", "SE Asia historical comedy cohort", 1},
		{"denny-cagur", "dolphy", "troupe", "SE Asia historical comedy cohort", 1},

		// SE Asia Modern hub (ronny-chieng as hub, already exists)
		{"harith-iskander", "ronny-chieng", "troupe", "SE Asia modern comedy cohort", 1},
		{"jason-leong", "ronny-chieng", "troupe", "SE Asia modern comedy cohort", 1},
		{"douglas-lim", "ronny-chieng", "troupe", "SE Asia modern comedy cohort", 1},
		{"nigel-ng", "ronny-chieng", "troupe", "SE Asia modern comedy cohort", 1},
		{"hannan-azlan", "ronny-chieng", "troupe", "SE Asia modern comedy cohort", 1},
		{"rishi-budhrani", "ronny-chieng", "troupe", "SE Asia modern comedy cohort", 1},
		{"marcus-chin", "ronny-chieng", "troupe", "SE Asia modern comedy cohort", 1},
		{"irene-ang", "ronny-chieng", "troupe", "SE Asia modern comedy cohort", 1},
		{"michelle-chong", "ronny-chieng", "troupe", "SE Asia modern comedy cohort", 1},
		// SE Asia real relationships
		{"harith-iskander", "douglas-lim", "collaboration", "Malaysian stand-up comedy scene co-founders; frequent co-performers", 3},
		{"harith-iskander", "nigel-ng", "collaboration", "Malaysian comedians; Uncle Roger's predecessor on the world stage", 2},
		{"irene-ang", "michelle-chong", "collaboration", "Both Singapore comedy actresses in local TV industry", 2},

		// World Historical hub (chaplin) - connect to existing nodes
		{"keaton", "chaplin", "troupe", "World historical top comedy cohort", 1},
		{"laurel", "chaplin", "troupe", "World historical comedy cohort", 1},
		{"hardy", "chaplin", "troupe", "World historical comedy cohort", 1},
		{"groucho", "chaplin", "troupe", "World historical comedy cohort", 1},
		{"harpo", "chaplin", "troupe", "World historical comedy cohort", 1},
		{"chico", "chaplin", "troupe", "World historical comedy cohort", 1},
		{"ball", "chaplin", "troupe", "World historical comedy cohort", 1},
		{"hope", "chaplin", "troupe", "World historical comedy cohort", 1},
		{"bruce", "chaplin", "troupe", "World historical comedy cohort", 1},
		// World Modern hub (chappelle) - connect to existing nodes
		{"hart", "chappelle", "troupe", "World modern top comedy cohort", 1},
		{"rock", "chappelle", "troupe", "World modern top comedy cohort", 1},
		{"bill-burr", "chappelle", "troupe", "World modern top comedy cohort", 1},
		{"mulaney", "chappelle", "troupe", "World modern top comedy cohort", 1},
		{"wong", "chappelle", "troupe", "World modern top comedy cohort", 1},
		{"trevor-noah", "chappelle", "troupe", "World modern top comedy cohort", 1},
		{"ricky-gervais", "chappelle", "troupe", "World modern top comedy cohort", 1},
		{"hannah-gadsby", "chappelle", "troupe", "World modern top comedy cohort", 1},
		{"joe-rogan", "chappelle", "troupe", "World modern top comedy cohort", 1},

		// Cross-region bridges
		{"vir-das", "trevor-noah", "collaboration", "Both Netflix-era international comics; mutual appearances", 2},
		{"ronny-chieng", "trevor-noah", "collaboration", "Ronny was Daily Show correspondent under Noah", 4},
		{"nigel-ng", "mrbean", "influence", "Uncle Roger's physical comedy influenced by British comedy", 1},
	}

	addedEdges := 0
	skippedEdges := 0
	missingEdges := 0
	nextEdgeNum := len(seed.Edges) + 1

	for _, ne := range newEdges {
		if !existingIDs[ne.Source] {
			fmt.Printf("  MISSING SRC: %s\n", ne.Source)
			missingEdges++
			continue
		}
		if !existingIDs[ne.Target] {
			fmt.Printf("  MISSING TGT: %s\n", ne.Target)
			missingEdges++
			continue
		}

		key := ne.Source + "|" + ne.Target
		if existingEdges[key] {
			skippedEdges++
			continue
		}

		e := Edge{
			ID:       fmt.Sprintf("eg%03d", nextEdgeNum),
			SourceID: ne.Source,
			TargetID: ne.Target,
			Type:     ne.Type,
			Weight:   ne.Weight,
			Summary:  ne.Summary,
			Evidence: []Link{},
		}

		seed.Edges = append(seed.Edges, e)
		existingEdges[key] = true
		existingEdges[ne.Target+"|"+ne.Source] = true
		nextEdgeNum++
		addedEdges++
	}

	fmt.Printf("\n═══ Results ═══\n")
	fmt.Printf("  Added nodes: %d\n", addedNodes)
	fmt.Printf("  Added edges: %d\n", addedEdges)
	fmt.Printf("  Skipped edges (dupe): %d\n", skippedEdges)
	fmt.Printf("  Missing edge nodes: %d\n", missingEdges)
	fmt.Printf("  Total nodes: %d\n", len(seed.Nodes))
	fmt.Printf("  Total edges: %d\n", len(seed.Edges))

	out, _ := json.MarshalIndent(seed, "", "  ")
	os.WriteFile(seedPath, append(out, '\n'), 0644)
	fmt.Println("  Wrote seed.json")
}
