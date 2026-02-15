package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

func kdlStr(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	return `"` + s + `"`
}

func intOrNull(p *int) string {
	if p == nil {
		return "null"
	}
	return fmt.Sprintf("%d", *p)
}

func main() {
	raw, err := os.ReadFile("app/src/lib/data/seed.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "read: %v\n", err)
		os.Exit(1)
	}
	var data SeedData
	if err := json.Unmarshal(raw, &data); err != nil {
		fmt.Fprintf(os.Stderr, "json: %v\n", err)
		os.Exit(1)
	}

	var b strings.Builder

	// Header
	b.WriteString("// ─────────────────────────────────────────────────────────\n")
	b.WriteString("// CartoJester — Comedian Alliances Atlas\n")
	b.WriteString("// Full KDL Export\n")
	b.WriteString(fmt.Sprintf("// Generated: 2026-02-15  |  %d comedians  |  %d alliances\n", len(data.Nodes), len(data.Edges)))
	b.WriteString("// ─────────────────────────────────────────────────────────\n\n")

	// Metadata
	b.WriteString("metadata {\n")
	b.WriteString(fmt.Sprintf("    total-comedians %d\n", len(data.Nodes)))
	b.WriteString(fmt.Sprintf("    total-alliances %d\n", len(data.Edges)))
	b.WriteString("    source \"Wikipedia List of Comedians + manual research\"\n")
	b.WriteString("    project \"CartoJester\"\n")
	b.WriteString("    url \"https://parkwardrr.github.io/CartoJester/\"\n")
	b.WriteString("}\n\n")

	// Groups
	b.WriteString("// ═══════════════════════════════════════════════════════\n")
	b.WriteString("// Comedy Groups / Movements\n")
	b.WriteString("// ═══════════════════════════════════════════════════════\n\n")

	type Group struct {
		ID              string
		Name            string
		Era             string
		YearStart       int
		YearEnd         int
		Description     string
		Characteristics []string
		MemberIDs       []string
		Color           string
		Icon            string
	}

	groups := []Group{
		{"silent-pioneers", "Silent Era Pioneers", "1895–1930", 1895, 1930, "The original physical comedians who invented the language of screen comedy.", []string{"Physical comedy", "Slapstick", "Pantomime", "Visual storytelling", "Studio system"}, []string{"chaplin", "keaton", "lloyd", "laurel", "hardy", "mack-sennett", "arbuckle", "langdon"}, "#94a3b8", "🎬"},
		{"vaudeville-stars", "Vaudeville & Stage Comics", "1900–1950", 1900, 1950, "Performers who honed their craft on the vaudeville circuit.", []string{"Timing", "Wordplay", "Musical comedy", "Audience rapport", "Variety acts"}, []string{"groucho", "harpo", "chico", "wc-fields", "mae-west", "burns", "jack-benny", "three-stooges-moe", "three-stooges-larry", "three-stooges-curly", "abbott", "costello"}, "#d97706", "🎭"},
		{"golden-age-tv", "Golden Age of Television", "1950–1970", 1950, 1970, "The first generation of TV comedians who defined the sitcom format.", []string{"Sitcom format", "Family-friendly", "Physical comedy on TV", "Live audiences", "Network era"}, []string{"ball", "hope", "skelton", "sid-caesar", "carl-reiner", "dick-van-dyke", "don-rickles", "phyllis-diller", "jack-benny", "burns", "newhart"}, "#8b5cf6", "📺"},
		{"counterculture-rebels", "Counterculture Comedy", "1960–1980", 1960, 1980, "Rebels who tore down comedy conventions with unflinching honesty.", []string{"Social commentary", "Boundary-pushing", "Political satire", "Confessional style", "Anti-establishment"}, []string{"bruce", "pryor", "carlin", "joan-rivers", "lily-tomlin", "dick-gregory", "mort-sahl", "bill-hicks", "moms-mabley"}, "#ef4444", "✊"},
		{"snl-original", "SNL & Sketch Revolution", "1975–1995", 1975, 1995, "Saturday Night Live launched careers and defined a generation.", []string{"Sketch comedy", "Improv roots", "Character work", "Political satire", "Ensemble energy"}, []string{"belushi", "aykroyd", "murray", "murphy", "martin-short", "chevy-chase", "gilda-radner", "mike-myers", "dana-carvey", "phil-hartman", "chris-farley", "norm-macdonald", "david-spade", "molly-shannon", "will-ferrell", "tracy-morgan", "jimmy-fallon", "seth-meyers", "bill-hader", "kristen-wiig", "andy-samberg", "kate-mckinnon", "maya-rudolph", "kenan-thompson", "fred-armisen", "pete-davidson", "bowen-yang", "colin-jost", "michael-che", "cecily-strong", "jason-sudeikis", "will-forte", "leslie-jones"}, "#3b82f6", "📡"},
		{"standup-boom", "80s/90s Stand-Up Boom", "1980–2000", 1980, 2000, "The comedy club explosion and HBO specials era.", []string{"Comedy clubs", "HBO specials", "Observational humor", "Arena tours", "Late-night crossover"}, []string{"seinfeld", "rock", "martin", "williams", "goldberg", "ellen", "ray-romano", "tim-allen", "roseanne", "sinbad", "bernie-mac", "cedric", "steve-harvey", "dl-hughley", "wanda-sykes", "bill-burr", "jeff-dunham", "jim-gaffigan", "brian-regan", "dane-cook", "dave-attell", "norm-macdonald", "sarah-silverman", "katt-williams", "mitch-hedberg", "larry-david"}, "#f97316", "🎤"},
		{"alt-comedy", "Alternative Comedy", "1995–2015", 1995, 2015, "Rejecting traditional punchline-driven comedy for quirky, meta styles.", []string{"Anti-humor", "Meta-comedy", "Absurdism", "Indie venues", "Podcast culture"}, []string{"zach-galifianakis", "maria-bamford", "patton-oswalt", "bo-burnham", "reggie-watts", "tig-notaro", "hannibal-buress", "eric-andre", "demetri-martin", "david-cross", "bob-odenkirk", "doug-stanhope", "kate-berlant", "aparna-nancherla", "nathan-fielder", "stewart-lee"}, "#14b8a6", "🌀"},
		{"improv-ucb", "Improv & UCB Movement", "1990–2020", 1990, 2020, "The improv pipeline from Second City, UCB, iO, and Groundlings.", []string{"Yes, and", "Long-form improv", "Writers room pipeline", "Ensemble focus", "Character work"}, []string{"fey", "poehler", "will-ferrell", "kate-mckinnon", "keegan-michael-key", "jordan-peele", "abbi-jacobson", "ilana-glazer", "nick-kroll", "jason-sudeikis", "aubrey-plaza", "jason-mantzoukas", "paul-f-tompkins", "lauren-lapkus"}, "#06b6d4", "🎪"},
		{"kings-of-comedy", "Kings & Queens of Comedy", "2000–2020", 2000, 2020, "Brought Black comedy to arena-scale audiences and launched media empires.", []string{"Arena comedy", "Cultural commentary", "Media empires", "Netflix specials", "Tour economics"}, []string{"rock", "hart", "chappelle", "bernie-mac", "cedric", "steve-harvey", "dl-hughley", "tiffany-haddish", "wanda-sykes", "mike-epps", "katt-williams", "ali-siddiq", "deon-cole"}, "#eab308", "👑"},
		{"netflix-era", "Streaming & Netflix Era", "2015–2025", 2015, 2025, "Comedy specials went global via streaming.", []string{"Global reach", "Diverse voices", "Social media crossover", "Special-driven careers", "International comedy"}, []string{"mulaney", "wong", "bo-burnham", "hasan-minhaj", "trevor-noah", "hannah-gadsby", "nate-bargatze", "taylor-tomlinson", "sam-morril", "mark-normand", "shane-gillis", "matt-rife", "nikki-glaser", "jerrod-carmichael", "iliza-shlesinger", "ronny-chieng", "michelle-wolf", "james-acaster", "mo-gilligan", "ali-siddiq"}, "#ec4899", "📱"},
		{"podcast-comedy", "Podcast Comedy Sphere", "2009–2025", 2009, 2025, "The podcast revolution created a new comedy ecosystem.", []string{"Long-form conversation", "Fan communities", "Cross-promotion", "Touring synergy", "YouTube crossover"}, []string{"joe-rogan", "marc-maron", "bill-burr", "tom-segura", "bert-kreischer", "christina-p", "theo-von", "andrew-santino", "bobby-lee", "shane-gillis", "mark-normand", "sam-morril", "conan-obrien", "paul-f-tompkins", "jason-mantzoukas", "doug-stanhope", "bowen-yang"}, "#10b981", "🎙️"},
		{"daily-show-political", "Daily Show & Political Comedy", "1999–2025", 1999, 2025, "The Daily Show launched a generation of political comedians.", []string{"Political satire", "Correspondent pipeline", "Late-night crossover", "Institutional comedy", "News parody"}, []string{"jon-stewart", "stephen-colbert", "john-oliver", "samantha-bee", "trevor-noah", "hasan-minhaj", "larry-wilmore", "jordan-klepper", "ronny-chieng", "michelle-wolf"}, "#f472b6", "🗳️"},
		{"uk-panel-shows", "UK Panel Show Circuit", "2000–2025", 2000, 2025, "British panel shows created a unique comedy ecosystem.", []string{"Quick wit", "Panel banter", "Audience rapport", "Topical humor", "Cross-show regulars"}, []string{"jimmy-carr", "dara-o-briain", "david-mitchell", "lee-mack", "katherine-ryan", "richard-ayoade", "noel-fielding", "robert-webb", "james-acaster", "romesh-ranganathan", "nish-kumar", "frankie-boyle", "sarah-millican", "jack-whitehall", "rob-brydon", "aisling-bea", "bill-bailey", "mo-gilligan"}, "#e879f9", "🇬🇧"},
	}

	for _, g := range groups {
		b.WriteString(fmt.Sprintf("group %s {\n", kdlStr(g.ID)))
		b.WriteString(fmt.Sprintf("    name %s\n", kdlStr(g.Name)))
		b.WriteString(fmt.Sprintf("    icon %s\n", kdlStr(g.Icon)))
		b.WriteString(fmt.Sprintf("    era %s\n", kdlStr(g.Era)))
		b.WriteString(fmt.Sprintf("    year-start %d\n", g.YearStart))
		b.WriteString(fmt.Sprintf("    year-end %d\n", g.YearEnd))
		b.WriteString(fmt.Sprintf("    color %s\n", kdlStr(g.Color)))
		b.WriteString(fmt.Sprintf("    description %s\n", kdlStr(g.Description)))
		b.WriteString("    characteristics {\n")
		for _, c := range g.Characteristics {
			b.WriteString(fmt.Sprintf("        - %s\n", kdlStr(c)))
		}
		b.WriteString("    }\n")
		b.WriteString("    members {\n")
		for _, m := range g.MemberIDs {
			b.WriteString(fmt.Sprintf("        - %s\n", kdlStr(m)))
		}
		b.WriteString("    }\n")
		b.WriteString("}\n\n")
	}

	// Nodes
	b.WriteString("// ═══════════════════════════════════════════════════════\n")
	b.WriteString(fmt.Sprintf("// Comedians (%d total)\n", len(data.Nodes)))
	b.WriteString("// ═══════════════════════════════════════════════════════\n\n")

	for _, n := range data.Nodes {
		b.WriteString(fmt.Sprintf("comedian %s {\n", kdlStr(n.ID)))
		b.WriteString(fmt.Sprintf("    name %s\n", kdlStr(n.Name)))

		if len(n.Aka) > 0 {
			b.WriteString("    aka {\n")
			for _, a := range n.Aka {
				b.WriteString(fmt.Sprintf("        - %s\n", kdlStr(a)))
			}
			b.WriteString("    }\n")
		}

		b.WriteString(fmt.Sprintf("    born %s\n", intOrNull(n.BornYear)))
		b.WriteString(fmt.Sprintf("    died %s\n", intOrNull(n.DiedYear)))
		b.WriteString(fmt.Sprintf("    active-start %s\n", intOrNull(n.ActiveStartYear)))
		b.WriteString(fmt.Sprintf("    active-end %s\n", intOrNull(n.ActiveEndYear)))

		if len(n.Tags) > 0 {
			var tqs []string
			for _, t := range n.Tags {
				tqs = append(tqs, kdlStr(t))
			}
			b.WriteString(fmt.Sprintf("    tags %s\n", strings.Join(tqs, " ")))
		}

		b.WriteString(fmt.Sprintf("    notability %d\n", n.Notability))

		if len(n.Links) > 0 {
			b.WriteString("    links {\n")
			for _, l := range n.Links {
				b.WriteString(fmt.Sprintf("        link %s url=%s\n", kdlStr(l.Label), kdlStr(l.URL)))
			}
			b.WriteString("    }\n")
		}

		b.WriteString("}\n\n")
	}

	// Edges
	b.WriteString("// ═══════════════════════════════════════════════════════\n")
	b.WriteString(fmt.Sprintf("// Alliances / Relationships (%d total)\n", len(data.Edges)))
	b.WriteString("// ═══════════════════════════════════════════════════════\n\n")

	for _, e := range data.Edges {
		b.WriteString(fmt.Sprintf("alliance %s {\n", kdlStr(e.ID)))
		b.WriteString(fmt.Sprintf("    source %s\n", kdlStr(e.SourceID)))
		b.WriteString(fmt.Sprintf("    target %s\n", kdlStr(e.TargetID)))
		b.WriteString(fmt.Sprintf("    type %s\n", kdlStr(e.Type)))
		b.WriteString(fmt.Sprintf("    start-year %s\n", intOrNull(e.StartYr)))
		b.WriteString(fmt.Sprintf("    end-year %s\n", intOrNull(e.EndYr)))
		b.WriteString(fmt.Sprintf("    weight %d\n", e.Weight))
		b.WriteString(fmt.Sprintf("    summary %s\n", kdlStr(e.Summary)))

		if len(e.Evidence) > 0 {
			b.WriteString("    evidence {\n")
			for _, ev := range e.Evidence {
				b.WriteString(fmt.Sprintf("        source %s url=%s\n", kdlStr(ev.Label), kdlStr(ev.URL)))
			}
			b.WriteString("    }\n")
		}

		b.WriteString("}\n\n")
	}

	// Write output
	out := b.String()
	if err := os.WriteFile("comedians_atlas.kdl", []byte(out), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "write: %v\n", err)
		os.Exit(1)
	}

	// Stats
	lines := strings.Count(out, "\n")
	fmt.Printf("Exported %d comedians + %d alliances + %d groups to comedians_atlas.kdl\n", len(data.Nodes), len(data.Edges), len(groups))
	fmt.Printf("File: %d lines, %d bytes\n", lines, len(out))
}
