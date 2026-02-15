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
	Title   string `json:"title"`
	Extract string `json:"extract"`
}

func intPtr(v int) *int { return &v }

func fetchExtract(name string) (string, error) {
	encoded := url.PathEscape(strings.ReplaceAll(name, " ", "_"))
	u := fmt.Sprintf("https://en.wikipedia.org/api/rest_v1/page/summary/%s", encoded)

	client := &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("GET", u, nil)
	req.Header.Set("User-Agent", "CartoJester-Research/1.0 (https://parkwardrr.github.io/CartoJester; twesh@users.noreply.github.com)")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var ws WikiSummary
	json.Unmarshal(body, &ws)
	return ws.Extract, nil
}

// Known comedian pairings — researched relationships
type ResearchedEdge struct {
	Source  string
	Target  string
	Type    string
	Summary string
	Weight  int
	StartYr *int
	EndYr   *int
}

func main() {
	seedPath := "app/src/lib/data/seed.json"

	raw, err := os.ReadFile(seedPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read: %v\n", err)
		os.Exit(1)
	}
	var seed SeedData
	json.Unmarshal(raw, &seed)

	nodeByID := map[string]*Node{}
	nodeByName := map[string]string{} // lowercase name -> id
	for i := range seed.Nodes {
		n := &seed.Nodes[i]
		nodeByID[n.ID] = n
		nodeByName[strings.ToLower(n.Name)] = n.ID
	}

	existingEdges := map[string]bool{}
	for _, e := range seed.Edges {
		existingEdges[e.SourceID+"|"+e.TargetID] = true
		existingEdges[e.TargetID+"|"+e.SourceID] = true
	}

	// Hand-researched well-known relationships
	researched := []ResearchedEdge{
		// SNL cast connections
		{"will-ferrell", "john-c-reilly", "collaboration", "Frequent film partners: Talladega Nights, Step Brothers, Holmes & Watson", 4, intPtr(2006), nil},
		{"adam-sandler", "rob-schneider", "collaboration", "Long-time friends and frequent film collaborators since SNL", 4, intPtr(1990), nil},
		{"adam-sandler", "david-spade", "collaboration", "SNL castmates and recurring film collaborators", 3, intPtr(1990), nil},
		{"adam-sandler", "kevin-james", "collaboration", "Frequent Happy Madison collaborators", 3, intPtr(2007), nil},
		{"chris-farley", "david-spade", "collaboration", "SNL duo and co-starred in Tommy Boy, Black Sheep", 5, intPtr(1993), intPtr(1997)},
		{"chevy-chase", "bill-murray", "rivalry", "Famous behind-the-scenes feud at SNL", 3, intPtr(1977), nil},
		{"will-ferrell", "molly-shannon", "collaboration", "SNL castmates and frequent scene partners", 3, intPtr(1995), intPtr(2002)},
		{"tina-fey", "amy-poehler", "collaboration", "SNL Weekend Update co-anchors; co-hosted Golden Globes", 5, intPtr(2004), nil},
		{"tina-fey", "rachel-dratch", "collaboration", "Second City and SNL collaborators, longtime friends", 4, intPtr(1994), nil},

		// Late night
		{"conan-obrien", "andy-richter", "collaboration", "Conan's sidekick and collaborative partner for decades", 5, intPtr(1993), nil},
		{"conan-obrien", "jon-stewart", "collaboration", "Fellow late-night hosts, mutual respect and crossover appearances", 3, intPtr(1999), nil},
		{"jon-stewart", "stephen-colbert", "mentor", "Colbert developed his persona on Stewart's Daily Show", 5, intPtr(1997), intPtr(2014)},
		{"stephen-colbert", "steve-carell", "collaboration", "Both emerged from Daily Show; Second City alumni", 3, intPtr(1997), intPtr(2005)},
		{"jimmy-fallon", "justin-timberlake", "collaboration", "Frequent SNL and Tonight Show sketch partners", 3, intPtr(2002), nil},

		// Podcast world
		{"joe-rogan", "bert-kreischer", "collaboration", "Close friends, podcast guests, and sober October challenge partners", 4, intPtr(2010), nil},
		{"joe-rogan", "tom-segura", "collaboration", "Regular podcast guests and touring comedian friends", 3, intPtr(2011), nil},
		{"tom-segura", "bert-kreischer", "collaboration", "Co-host 2 Bears, 1 Cave podcast", 5, intPtr(2019), nil},
		{"tom-segura", "christina-p", "family", "Married couple, co-host Your Mom's House podcast", 5, intPtr(2008), nil},
		{"marc-maron", "louis-ck", "collaboration", "Famous WTF podcast episode; longtime comedy peers", 3, intPtr(2010), nil},
		{"bobby-lee", "andrew-santino", "collaboration", "Co-host Bad Friends podcast", 5, intPtr(2020), nil},
		{"theo-von", "brendan-schaub", "collaboration", "Co-hosted King and the Sting podcast", 4, intPtr(2018), intPtr(2022)},
		{"mark-normand", "sam-morril", "collaboration", "Co-host We Might Be Drunk podcast; frequent collaborators", 5, intPtr(2021), nil},

		// Stand-up mentors and influences
		{"richard-pryor", "dave-chappelle", "influence", "Pryor's raw honesty deeply influenced Chappelle's comedy", 4, intPtr(1990), intPtr(2005)},
		{"george-carlin", "bill-hicks", "influence", "Carlin's fearless social criticism inspired Hicks", 4, intPtr(1985), intPtr(1994)},
		{"george-carlin", "louis-ck", "influence", "Carlin's work-ethic of discarding material annually inspired Louis CK", 4, intPtr(1990), intPtr(2008)},
		{"chris-rock", "dave-chappelle", "collaboration", "Close friends; both top-tier stand-ups who frequently reference each other", 4, intPtr(1990), nil},
		{"chris-rock", "kevin-hart", "mentor", "Rock helped launch Hart's career; longtime friends", 4, intPtr(2003), nil},
		{"jerry-seinfeld", "larry-david", "collaboration", "Co-created Seinfeld, one of the most successful sitcoms ever", 5, intPtr(1989), intPtr(1998)},
		{"jerry-seinfeld", "bill-burr", "collaboration", "Comedians in Cars Getting Coffee; mutual respect among stand-ups", 2, intPtr(2012), nil},
		{"dave-chappelle", "neal-brennan", "collaboration", "Co-created Chappelle's Show", 5, intPtr(2003), intPtr(2006)},

		// Key & Peele / comedy partnerships
		{"keegan-michael-key", "jordan-peele", "collaboration", "Key & Peele sketch show; longtime comedy duo from Second City", 5, intPtr(2012), intPtr(2015)},

		// UK connections
		{"ricky-gervais", "stephen-merchant", "collaboration", "Co-created The Office (UK) and Extras", 5, intPtr(2001), intPtr(2007)},
		{"ricky-gervais", "karl-pilkington", "collaboration", "Radio show and An Idiot Abroad collaboration", 4, intPtr(2001), intPtr(2012)},
		{"david-mitchell", "robert-webb", "collaboration", "Peep Show and That Mitchell and Webb Look duo", 5, intPtr(2003), nil},
		{"lee-mack", "tim-vine", "collaboration", "Not Going Out and frequent panel show interactions", 3, intPtr(2006), nil},
		{"sacha-baron-cohen", "larry-charles", "collaboration", "Director of Borat, Bruno, and The Dictator", 4, intPtr(2006), intPtr(2012)},

		// Modern Netflix era
		{"john-mulaney", "nick-kroll", "collaboration", "Co-created Oh, Hello on Broadway; longtime friends", 5, intPtr(2015), nil},
		{"bo-burnham", "jerrod-carmichael", "collaboration", "Fellow Netflix-era stand-ups who push the form", 2, intPtr(2016), nil},
		{"ali-wong", "randall-park", "collaboration", "Co-starred in Always Be My Maybe", 3, intPtr(2019), intPtr(2019)},
		{"nate-bargatze", "shane-gillis", "collaboration", "Fellow touring stand-ups with Netflix specials; mutual admiration", 2, intPtr(2020), nil},

		// Improv connections
		{"fred-armisen", "carrie-brownstein", "collaboration", "Co-created Portlandia", 5, intPtr(2011), intPtr(2018)},

		// Sacha Baron Cohen
		{"sacha-baron-cohen", "isla-fisher", "family", "Married couple", 4, intPtr(2010), nil},

		// Roast connections
		{"nikki-glaser", "jeff-ross", "collaboration", "Frequent Comedy Central Roast performers", 2, intPtr(2015), nil},
		{"gabriel-iglesias", "jeff-dunham", "collaboration", "Fellow arena-level touring comedians; both in comedy specials circuit", 2, intPtr(2010), nil},

		// Whitney Cummings connections
		{"whitney-cummings", "chelsea-handler", "collaboration", "Both had late-night/talk shows; frequent collaborators", 3, intPtr(2012), nil},

		// Film comedy trios
		{"seth-rogen", "james-franco", "collaboration", "Frequent film partners: Pineapple Express, The Interview, Freaks and Geeks", 5, intPtr(1999), intPtr(2017)},
		{"seth-rogen", "jonah-hill", "collaboration", "Superbad co-stars and Judd Apatow film regulars", 4, intPtr(2007), nil},
		{"seth-rogen", "evan-goldberg", "collaboration", "Writing partners since high school; co-wrote Superbad, This Is the End", 5, intPtr(2001), nil},

		// Judd Apatow connections
		{"judd-apatow", "adam-sandler", "collaboration", "Apatow and Sandler were roommates; Funny People", 4, intPtr(1990), nil},

		// Broad City
		{"abbi-jacobson", "ilana-glazer", "collaboration", "Co-created and starred in Broad City", 5, intPtr(2014), intPtr(2019)},

		// Chelsea Peretti
		{"chelsea-peretti", "jordan-peele", "family", "Married couple", 4, intPtr(2016), nil},

		// Russell Brand
		{"russell-brand", "noel-fielding", "collaboration", "British comedy scene collaborators and friends", 2, intPtr(2006), nil},

		// International connections
		{"trevor-noah", "hasan-minhaj", "collaboration", "Both Daily Show correspondents who got their own shows", 3, intPtr(2015), nil},

		// Classic connections
		{"mel-brooks", "carl-reiner", "collaboration", "2000 Year Old Man routine; lifeling friends for 70+ years", 5, intPtr(1950), intPtr(2020)},
		{"mel-brooks", "gene-wilder", "collaboration", "Director-star duo: The Producers, Young Frankenstein, Blazing Saddles", 5, intPtr(1967), intPtr(1974)},
		{"lucille-ball", "carol-burnett", "mentor", "Ball mentored Burnett; Burnett cited Ball as primary inspiration", 4, intPtr(1960), intPtr(1989)},
		{"bob-hope", "bing-crosby", "collaboration", "Road to... film series partnership", 5, intPtr(1940), intPtr(1962)},
		{"dan-aykroyd", "john-belushi", "collaboration", "Blues Brothers and SNL co-stars; best friends", 5, intPtr(1975), intPtr(1982)},

		// More modern
		{"taylor-tomlinson", "nate-bargatze", "collaboration", "Fellow touring Netflix-era stand-ups", 2, intPtr(2020), nil},
		{"iliza-shlesinger", "whitney-cummings", "collaboration", "Fellow female stand-ups in Netflix era; podcast appearances", 2, intPtr(2015), nil},
		{"matt-rife", "nick-cannon", "collaboration", "Wild 'N Out castmates early in career", 2, intPtr(2015), intPtr(2018)},
		{"shane-gillis", "mark-normand", "collaboration", "Regular podcast appearances; touring friends", 3, intPtr(2019), nil},

		// Writing partnerships
		{"bob-odenkirk", "david-cross", "collaboration", "Co-created Mr. Show with Bob and David", 5, intPtr(1995), intPtr(1998)},
		{"nick-kroll", "jason-mantzoukas", "collaboration", "The League co-stars; frequent collaborators", 4, intPtr(2009), nil},

		// Canadian comedy
		{"martin-short", "steve-martin", "collaboration", "Three Amigos alumni; Father of the Bride and Only Murders co-stars", 4, intPtr(1986), nil},
		{"mike-myers", "dana-carvey", "collaboration", "Wayne's World duo from SNL", 5, intPtr(1989), intPtr(1993)},

		// Jim Carrey connections
		{"jim-carrey", "jeff-daniels", "collaboration", "Dumb and Dumber duo", 4, intPtr(1994), intPtr(2014)},
		{"jim-carrey", "robin-williams", "collaboration", "Fellow physical-comedy stand-ups; mutual admiration", 3, intPtr(1990), intPtr(2014)},

		// British panel show connections
		{"jimmy-carr", "sean-lock", "collaboration", "8 Out of 10 Cats co-regulars for many years", 4, intPtr(2005), intPtr(2021)},
		{"jimmy-carr", "jon-richardson", "collaboration", "8 Out of 10 Cats team captains", 3, intPtr(2012), nil},
		{"dara-o-briain", "ed-byrne", "collaboration", "Mock the Week regulars and Irish comedy touring circuit", 3, intPtr(2005), nil},
		{"lee-mack", "david-mitchell", "collaboration", "Would I Lie to You? team captains for over 15 years", 5, intPtr(2007), nil},
		{"noel-fielding", "julian-barratt", "collaboration", "The Mighty Boosh duo", 5, intPtr(2004), intPtr(2009)},
		{"james-acaster", "ed-gamble", "collaboration", "Off Menu podcast co-hosts; longtime friends", 4, intPtr(2018), nil},
		{"romesh-ranganathan", "rob-beckett", "collaboration", "Co-present Rob & Romesh Vs...", 4, intPtr(2019), nil},
		{"sarah-millican", "katherine-ryan", "collaboration", "UK panel show circuit regulars; 8 Out of 10 Cats appearances", 2, intPtr(2012), nil},

		// Daily Show alumni
		{"jon-stewart", "john-oliver", "mentor", "Oliver was Daily Show correspondent before Last Week Tonight", 4, intPtr(2006), intPtr(2014)},
		{"jon-stewart", "samantha-bee", "mentor", "Bee was a Daily Show correspondent before Full Frontal", 4, intPtr(2003), intPtr(2015)},
		{"jon-stewart", "trevor-noah", "mentor", "Noah succeeded Stewart as Daily Show host", 4, intPtr(2015), intPtr(2015)},
		{"jon-stewart", "hasan-minhaj", "mentor", "Minhaj was a Daily Show correspondent before Patriot Act", 3, intPtr(2014), intPtr(2018)},
		{"jon-stewart", "larry-wilmore", "mentor", "Wilmore was contributor/correspondent on Daily Show before his own show", 3, intPtr(2006), intPtr(2015)},

		// More Kings/Queens of comedy connections
		{"bernie-mac", "cedric", "collaboration", "Original Kings of Comedy tour together", 5, intPtr(2000), intPtr(2000)},
		{"bernie-mac", "steve-harvey", "collaboration", "Original Kings of Comedy tour together", 5, intPtr(2000), intPtr(2000)},
		{"bernie-mac", "dl-hughley", "collaboration", "Original Kings of Comedy tour together", 5, intPtr(2000), intPtr(2000)},
		{"cedric", "steve-harvey", "collaboration", "Original Kings of Comedy tour together", 5, intPtr(2000), intPtr(2000)},
		{"cedric", "dl-hughley", "collaboration", "Original Kings of Comedy tour together", 5, intPtr(2000), intPtr(2000)},
		{"steve-harvey", "dl-hughley", "collaboration", "Original Kings of Comedy tour together", 5, intPtr(2000), intPtr(2000)},

		// Wanda Sykes
		{"wanda-sykes", "chris-rock", "collaboration", "Writers and performers in comedy scene; Sykes wrote for Rock's show", 3, intPtr(1999), nil},

		// Women in comedy mentorships
		{"joan-rivers", "sarah-silverman", "influence", "Rivers paved the way for provocative female comedy", 3, intPtr(1990), intPtr(2014)},
		{"phyllis-diller", "joan-rivers", "mentor", "Diller mentored and encouraged Rivers early in her career", 3, intPtr(1960), intPtr(1975)},
		{"lily-tomlin", "jane-fonda", "collaboration", "9 to 5 and Grace and Frankie co-stars", 4, intPtr(1980), nil},

		// Kings of Comedy → next gen
		{"dave-chappelle", "ali-siddiq", "collaboration", "Siddiq opened for Chappelle tours", 2, intPtr(2017), nil},
		{"kevin-hart", "tiffany-haddish", "collaboration", "Night School co-stars; Hart helped boost Haddish's career", 3, intPtr(2018), nil},

		// Jo Koy, Russell Peters, Gabriel Iglesias - touring comedians
		{"jo-koy", "gabriel-iglesias", "collaboration", "Fellow touring arena comedians with Netflix specials", 2, intPtr(2015), nil},
		{"russell-peters", "gabriel-iglesias", "collaboration", "Both arena-level touring comics with cultural comedy", 2, intPtr(2010), nil},

		// Improv connections
		{"aubrey-plaza", "chris-pratt", "collaboration", "Parks and Recreation castmates", 4, intPtr(2009), intPtr(2015)},
		{"nick-offerman", "megan-mullally", "family", "Married couple, both comedy actors", 4, intPtr(2003), nil},
		{"amy-poehler", "nick-kroll", "collaboration", "Kroll Show appearances; Parks & Rec era collaborators", 2, intPtr(2013), nil},

		// Second City links
		{"john-belushi", "bill-murray", "collaboration", "Second City and SNL founding cast connections", 4, intPtr(1973), intPtr(1982)},

		// Chelsea Handler connections
		{"chelsea-handler", "sarah-silverman", "collaboration", "Fellow provocative female comics; competitive peers", 2, intPtr(2007), nil},
		{"chelsea-handler", "jo-koy", "family", "Former romantic partners; remained collaborators", 3, intPtr(2021), intPtr(2022)},

		// Daniel Tosh
		{"daniel-tosh", "jeff-ross", "collaboration", "Comedy Central peers; roast circuit", 2, intPtr(2010), nil},

		// Jim Jefferies
		{"jim-jefferies", "joe-rogan", "collaboration", "Fellow stand-ups; frequent podcast appearances", 2, intPtr(2014), nil},

		// Ken Jeong
		{"ken-jeong", "zach-galifianakis", "collaboration", "The Hangover co-stars", 4, intPtr(2009), intPtr(2013)},
	}

	// Map old IDs to actual IDs in our database (some might differ)
	idMap := map[string]string{
		"richard-pryor":          "pryor",
		"george-carlin":          "carlin",
		"chris-rock":             "rock",
		"kevin-hart":             "hart",
		"jerry-seinfeld":         "seinfeld",
		"larry-david":            "larry-david",
		"eddie-murphy":           "murphy",
		"robin-williams":         "williams",
		"steve-martin":           "martin",
		"bill-murray":            "murray",
		"adam-sandler":           "sandler",
		"tina-fey":               "fey",
		"amy-poehler":            "poehler",
		"dave-chappelle":         "chappelle",
		"john-belushi":           "belushi",
		"dan-aykroyd":            "aykroyd",
		"lucille-ball":           "ball",
		"bob-hope":               "hope",
		"carl-reiner":            "carl-reiner",
		"dave-attell":            "dave-attell",
		"john-mulaney":           "mulaney",
		"ali-wong":               "wong",
		"cedric-the-entertainer": "cedric",
	}

	resolveID := func(name string) string {
		if mapped, ok := idMap[name]; ok {
			if _, exists := nodeByID[mapped]; exists {
				return mapped
			}
		}
		if _, exists := nodeByID[name]; exists {
			return name
		}
		if id, ok := nodeByName[strings.ToLower(name)]; ok {
			return id
		}
		return ""
	}

	added := 0
	skipped := 0
	missing := 0
	nextNum := len(seed.Edges) + 1

	for _, r := range researched {
		srcID := resolveID(r.Source)
		tgtID := resolveID(r.Target)

		if srcID == "" || tgtID == "" {
			if srcID == "" {
				fmt.Printf("  MISSING NODE: %s\n", r.Source)
			}
			if tgtID == "" {
				fmt.Printf("  MISSING NODE: %s\n", r.Target)
			}
			missing++
			continue
		}

		key := srcID + "|" + tgtID
		if existingEdges[key] {
			skipped++
			continue
		}

		e := Edge{
			ID:       fmt.Sprintf("er%03d", nextNum),
			SourceID: srcID,
			TargetID: tgtID,
			Type:     r.Type,
			Weight:   r.Weight,
			Summary:  r.Summary,
			Evidence: []Link{},
		}
		e.StartYr = r.StartYr
		e.EndYr = r.EndYr

		seed.Edges = append(seed.Edges, e)
		existingEdges[key] = true
		existingEdges[tgtID+"|"+srcID] = true
		nextNum++
		added++
	}

	fmt.Printf("\nResults:\n")
	fmt.Printf("  Added: %d new researched edges\n", added)
	fmt.Printf("  Skipped (already exist): %d\n", skipped)
	fmt.Printf("  Missing nodes: %d\n", missing)
	fmt.Printf("  Total edges now: %d\n", len(seed.Edges))

	// Write
	out, _ := json.MarshalIndent(seed, "", "  ")
	os.WriteFile(seedPath, append(out, '\n'), 0644)
	fmt.Printf("  Wrote updated seed.json\n")
}
