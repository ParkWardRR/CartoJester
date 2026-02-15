package main

import "fmt"

func allNewEdges(startIdx int) []Edge {
	var edges []Edge
	i := startIdx
	add := func(src, tgt, typ string, sy, ey *int, w int, sum, wikiRef string) {
		i++
		edges = append(edges, Edge{
			ID: fmt.Sprintf("e%03d", i), SourceID: src, TargetID: tgt,
			Type: typ, StartYr: sy, EndYr: ey, Weight: w, Summary: sum,
			Evidence: ev(wikiRef),
		})
	}

	// SNL troupe connections
	add("will-ferrell", "chris-farley", "troupe", yr(1995), yr(1997), 4, "Overlapping SNL cast members 1995-97", "Saturday_Night_Live")
	add("will-ferrell", "maya-rudolph", "troupe", yr(2000), yr(2002), 3, "SNL cast members together", "Saturday_Night_Live")
	add("will-ferrell", "jimmy-fallon", "troupe", yr(1998), yr(2002), 3, "SNL cast members together", "Saturday_Night_Live")
	add("will-ferrell", "tracy-morgan", "troupe", yr(1996), yr(2002), 3, "SNL cast members together", "Saturday_Night_Live")
	add("fey", "jimmy-fallon", "collaboration", yr(2000), yr(2004), 5, "Weekend Update co-anchors on SNL", "Saturday_Night_Live")
	add("fey", "seth-meyers", "collaboration", yr(2001), yr(2006), 4, "Fey mentored Meyers as head writer at SNL", "Saturday_Night_Live")
	add("fey", "tracy-morgan", "collaboration", yr(2006), yr(2013), 5, "30 Rock co-stars", "30_Rock")
	add("seth-meyers", "jimmy-fallon", "troupe", yr(2001), yr(2004), 3, "SNL cast members together", "Saturday_Night_Live")
	add("bill-hader", "kristen-wiig", "troupe", yr(2005), yr(2012), 5, "Core SNL cast together, frequent sketch partners", "Saturday_Night_Live")
	add("bill-hader", "andy-samberg", "troupe", yr(2005), yr(2012), 4, "SNL cast members together", "Saturday_Night_Live")
	add("bill-hader", "fred-armisen", "troupe", yr(2005), yr(2012), 4, "SNL cast & Portlandia/IFC comedy sphere", "Saturday_Night_Live")
	add("bill-hader", "kate-mckinnon", "troupe", yr(2012), yr(2019), 3, "Overlapping SNL cast", "Saturday_Night_Live")
	add("kristen-wiig", "maya-rudolph", "troupe", yr(2005), yr(2007), 4, "SNL cast together; Bridesmaids co-stars", "Bridesmaids_(2011_film)")
	add("kate-mckinnon", "cecily-strong", "troupe", yr(2012), yr(2022), 5, "Core SNL cast together for decade", "Saturday_Night_Live")
	add("kate-mckinnon", "pete-davidson", "troupe", yr(2014), yr(2022), 4, "SNL cast members together", "Saturday_Night_Live")
	add("kate-mckinnon", "bowen-yang", "troupe", yr(2019), yr(2022), 3, "SNL cast members together", "Saturday_Night_Live")
	add("colin-jost", "michael-che", "collaboration", yr(2014), nil, 5, "Weekend Update co-anchors on SNL", "Saturday_Night_Live")
	add("andy-samberg", "seth-meyers", "troupe", yr(2005), yr(2012), 3, "SNL cast together", "Saturday_Night_Live")
	add("pete-davidson", "mulaney", "collaboration", yr(2018), nil, 4, "Close friends; co-toured and co-starred together", "Pete_Davidson")
	add("kenan-thompson", "tracy-morgan", "troupe", yr(2003), yr(2025), 3, "Longest-tenured SNL cast overlap", "Saturday_Night_Live")
	add("mike-myers", "dana-carvey", "collaboration", yr(1989), yr(1993), 5, "Wayne's World duo on SNL and film", "Wayne's_World_(film)")

	// Daily Show ecosystem
	add("jon-stewart", "stephen-colbert", "collaboration", yr(1999), yr(2005), 5, "Colbert was Daily Show correspondent; Stewart produced Colbert Report", "The_Daily_Show")
	add("jon-stewart", "john-oliver", "collaboration", yr(2006), yr(2013), 4, "Oliver was Daily Show correspondent, guest-hosted", "The_Daily_Show")
	add("jon-stewart", "samantha-bee", "collaboration", yr(2003), yr(2015), 4, "Bee was longest-serving Daily Show correspondent", "The_Daily_Show")
	add("jon-stewart", "trevor-noah", "mentor", yr(2015), yr(2022), 4, "Stewart selected Noah as Daily Show successor", "Trevor_Noah")
	add("jon-stewart", "hasan-minhaj", "collaboration", yr(2014), yr(2018), 3, "Minhaj was Daily Show correspondent", "Hasan_Minhaj")
	add("jon-stewart", "larry-wilmore", "collaboration", yr(1999), yr(2015), 3, "Wilmore was Daily Show contributor; Stewart produced Nightly Show", "Larry_Wilmore")
	add("jon-stewart", "jordan-klepper", "collaboration", yr(2014), nil, 3, "Klepper as Daily Show correspondent", "Jordan_Klepper")
	add("jon-stewart", "ronny-chieng", "collaboration", yr(2015), nil, 3, "Chieng as Daily Show correspondent", "Ronny_Chieng")
	add("stephen-colbert", "john-oliver", "troupe", yr(2006), yr(2013), 3, "Daily Show correspondents together", "The_Daily_Show")
	add("stephen-colbert", "samantha-bee", "troupe", yr(2003), yr(2005), 3, "Daily Show correspondents together", "The_Daily_Show")

	// Podcast comedy sphere
	add("joe-rogan", "tom-segura", "collaboration", yr(2010), nil, 4, "Frequent podcast guests on each other's shows", "The_Joe_Rogan_Experience")
	add("joe-rogan", "bert-kreischer", "collaboration", yr(2010), nil, 4, "Close friends; frequent podcast collaborators", "The_Joe_Rogan_Experience")
	add("joe-rogan", "theo-von", "collaboration", yr(2017), nil, 3, "Regular podcast guest; Rogan helped boost Von's career", "The_Joe_Rogan_Experience")
	add("joe-rogan", "bill-burr", "collaboration", yr(2010), nil, 4, "Frequent podcast guests; standup contemporaries", "The_Joe_Rogan_Experience")
	add("joe-rogan", "marc-maron", "rivalry", yr(2010), nil, 3, "Competing podcast pioneers; occasional friction", "Marc_Maron")
	add("joe-rogan", "shane-gillis", "collaboration", yr(2019), nil, 3, "Rogan supported Gillis; frequent podcast guest", "Shane_Gillis")
	add("tom-segura", "bert-kreischer", "collaboration", yr(2014), nil, 5, "Co-hosts of 2 Bears 1 Cave podcast", "2_Bears,_1_Cave")
	add("tom-segura", "christina-p", "family", yr(2008), nil, 5, "Married; co-host Your Mom's House podcast", "Your_Mom's_House")
	add("bert-kreischer", "mark-normand", "collaboration", yr(2020), nil, 3, "Frequent podcast/touring collaborators", "Bert_Kreischer")
	add("mark-normand", "sam-morril", "collaboration", yr(2015), nil, 5, "Best friends; co-tour and co-create content", "Mark_Normand")
	add("mark-normand", "shane-gillis", "collaboration", yr(2019), nil, 4, "Close friends; co-headline and podcast together", "Shane_Gillis")
	add("andrew-santino", "bobby-lee", "collaboration", yr(2020), nil, 5, "Co-hosts of Bad Friends podcast", "Bad_Friends")
	add("bill-burr", "joe-rogan", "collaboration", yr(2005), nil, 4, "Long-time standup friends; mutual podcast appearances", "Bill_Burr")
	add("marc-maron", "sarah-silverman", "collaboration", yr(1992), nil, 3, "NYC standup contemporaries; former couple", "Marc_Maron")
	add("marc-maron", "patton-oswalt", "collaboration", yr(1990), nil, 4, "Alt-comedy peers; WTF podcast guests", "WTF_with_Marc_Maron")

	// Chappelle ecosystem
	add("chappelle", "neal-brennan", "collaboration", yr(1996), yr(2006), 5, "Co-created and co-wrote Chappelle's Show", "Chappelle's_Show")
	add("chappelle", "bill-burr", "collaboration", yr(2000), nil, 3, "Standup contemporaries; mutual respect", "Dave_Chappelle")
	add("chappelle", "jon-stewart", "collaboration", yr(2020), nil, 3, "Co-toured and appeared together", "Dave_Chappelle")

	// Key & Peele
	add("keegan-michael-key", "jordan-peele", "collaboration", yr(2003), yr(2015), 5, "Key & Peele sketch show duo", "Key_%26_Peele")

	// UK comedy connections
	add("ricky-gervais", "steve-coogan", "collaboration", yr(2005), nil, 3, "UK comedy contemporaries; shared comedy sphere", "Ricky_Gervais")
	add("david-mitchell", "robert-webb", "collaboration", yr(1995), nil, 5, "Peep Show duo; Mitchell and Webb Look", "Peep_Show_(TV_series)")
	add("david-mitchell", "lee-mack", "collaboration", yr(2009), nil, 4, "Co-captains on Would I Lie to You", "Would_I_Lie_to_You%3F_(TV_series)")
	add("noel-fielding", "richard-ayoade", "collaboration", yr(2004), nil, 4, "The IT Crowd & Mighty Boosh comedy sphere", "The_IT_Crowd")
	add("jimmy-carr", "dara-o-briain", "troupe", yr(2005), nil, 3, "UK panel show circuit regulars", "8_Out_of_10_Cats")
	add("jimmy-carr", "katherine-ryan", "troupe", yr(2012), nil, 3, "8 Out of 10 Cats and panel show circuit", "8_Out_of_10_Cats")
	add("jimmy-carr", "romesh-ranganathan", "troupe", yr(2015), nil, 3, "UK panel show regulars together", "8_Out_of_10_Cats")
	add("james-acaster", "nish-kumar", "collaboration", yr(2015), nil, 3, "UK standup contemporaries; frequent panel show appearances together", "James_Acaster")
	add("lee-mack", "frankie-boyle", "troupe", yr(2007), nil, 3, "Mock the Week and panel show regulars", "Mock_the_Week")
	add("dara-o-briain", "frankie-boyle", "collaboration", yr(2005), yr(2019), 4, "Host and regular on Mock the Week", "Mock_the_Week")
	add("rob-brydon", "steve-coogan", "collaboration", yr(2010), nil, 5, "The Trip film/series duo", "The_Trip_(2010_film)")

	// Apatow ecosystem
	add("judd-apatow", "will-ferrell", "studio", yr(2004), yr(2008), 4, "Apatow produced Anchorman and Talladega Nights", "Judd_Apatow")
	add("judd-apatow", "sandler", "collaboration", yr(2009), yr(2009), 3, "Funny People starring Sandler, directed by Apatow", "Funny_People")
	add("judd-apatow", "amy-schumer", "studio", yr(2015), yr(2015), 4, "Directed Trainwreck starring Schumer", "Trainwreck_(film)")
	add("judd-apatow", "pete-davidson", "studio", yr(2020), yr(2020), 3, "Co-wrote and directed King of Staten Island", "The_King_of_Staten_Island")
	add("judd-apatow", "kumail-nanjiani", "studio", yr(2017), yr(2017), 4, "Produced The Big Sick", "The_Big_Sick")

	// Larry David / Curb
	add("larry-david", "seinfeld", "collaboration", yr(1989), yr(1998), 5, "Co-created Seinfeld", "Seinfeld")
	add("larry-david", "jb-smoove", "collaboration", yr(2007), nil, 4, "JB Smoove as Leon on Curb Your Enthusiasm", "Curb_Your_Enthusiasm")

	// Late night
	add("conan-obrien", "andy-samberg", "collaboration", yr(2009), yr(2010), 3, "Samberg was vocal Conan supporter during late-night conflict", "Conan_O'Brien")
	add("conan-obrien", "bill-hader", "collaboration", yr(2005), nil, 3, "Frequent late-night guest; comedy peers", "Conan_O'Brien")
	add("conan-obrien", "jimmy-fallon", "rivalry", yr(2009), yr(2014), 3, "Competing late-night hosts", "Late-night_wars")

	// Desus & Mero
	add("desus-nice", "the-kid-mero", "collaboration", yr(2013), yr(2022), 5, "Desus & Mero podcast and TV show duo", "Desus_%26_Mero")

	// Cross-connections
	add("sarah-silverman", "amy-schumer", "influence", yr(2010), nil, 3, "Silverman paved way for Schumer's provocative style", "Amy_Schumer")
	add("fred-armisen", "bill-hader", "collaboration", yr(2014), yr(2019), 4, "Co-created Documentary Now!", "Documentary_Now!")
	add("aziz-ansari", "hart", "collaboration", yr(2009), nil, 3, "Contemporary standup peers; comedy touring circuit", "Aziz_Ansari")
	add("nick-kroll", "mulaney", "collaboration", yr(2008), nil, 5, "Oh, Hello co-creators and performers; close friends", "Oh,_Hello")
	add("donald-glover", "fey", "collaboration", yr(2006), yr(2009), 4, "Glover wrote for 30 Rock under Fey", "30_Rock")
	add("poehler", "aubrey-plaza", "collaboration", yr(2009), yr(2015), 4, "Parks and Recreation co-stars", "Parks_and_Recreation")
	add("poehler", "aziz-ansari", "collaboration", yr(2009), yr(2015), 4, "Parks and Recreation co-stars", "Parks_and_Recreation")
	add("poehler", "nick-kroll", "collaboration", yr(2004), nil, 3, "UCB improv community connections", "Upright_Citizens_Brigade")
	add("david-cross", "bob-odenkirk", "collaboration", yr(1995), nil, 5, "Mr. Show co-creators", "Mr._Show_with_Bob_and_David")
	add("chris-farley", "david-spade", "collaboration", yr(1990), yr(1997), 5, "SNL and film duo; Tommy Boy, Black Sheep", "Tommy_Boy")
	add("chris-farley", "sandler", "troupe", yr(1990), yr(1995), 5, "Close friends and SNL cast members together", "Saturday_Night_Live")
	add("chris-farley", "rock", "troupe", yr(1990), yr(1993), 4, "SNL cast members together in early 90s", "Saturday_Night_Live")
	add("mike-myers", "murphy", "collaboration", yr(2001), yr(2010), 4, "Shrek franchise co-stars", "Shrek")
	add("norm-macdonald", "david-spade", "troupe", yr(1994), yr(1998), 4, "SNL cast members together; close friends", "Saturday_Night_Live")
	add("norm-macdonald", "conan-obrien", "collaboration", yr(1998), nil, 4, "Legendary recurring late-night guest appearances", "Norm_Macdonald")
	add("norm-macdonald", "sandler", "troupe", yr(1993), yr(1998), 3, "SNL cast members; appeared in Sandler films", "Saturday_Night_Live")
	add("phoebe-waller-bridge", "olivia-colman", "collaboration", yr(2016), yr(2019), 5, "Fleabag co-stars; Waller-Bridge wrote for Colman", "Fleabag")
	add("phoebe-waller-bridge", "aisling-bea", "collaboration", yr(2016), nil, 3, "UK comedy contemporaries; mutual supporters", "Phoebe_Waller-Bridge")
	add("mindy-kaling", "ricky-gervais", "collaboration", yr(2005), yr(2013), 3, "The Office (US) connect; Gervais created UK original", "The_Office_(American_TV_series)")
	add("bo-burnham", "jerrod-carmichael", "collaboration", yr(2022), nil, 3, "Contemporary special-driven comedians; mutual admirers", "Bo_Burnham")
	add("wanda-sykes", "rock", "collaboration", yr(1999), nil, 3, "Standup contemporaries; Chris Rock Show writers", "The_Chris_Rock_Show")
	add("steve-harvey", "cedric", "collaboration", yr(2000), yr(2001), 5, "The Original Kings of Comedy tour together", "The_Original_Kings_of_Comedy")
	add("cedric", "rock", "troupe", yr(1990), nil, 3, "Black comedy circuit; Kings of Comedy era", "The_Original_Kings_of_Comedy")
	add("katt-williams", "mike-epps", "collaboration", yr(2005), nil, 3, "Comedy touring circuit contemporaries", "Katt_Williams")
	add("dave-attell", "jeff-ross", "collaboration", yr(1995), nil, 4, "NYC standup veterans; Bumping Mics co-tour", "Bumping_Mics")
	add("jeff-ross", "nikki-glaser", "collaboration", yr(2015), nil, 3, "Comedy Central roast circuit", "Comedy_Central_Roast")
	add("jim-gaffigan", "nate-bargatze", "influence", yr(2010), nil, 3, "Gaffigan's clean observational style influenced Bargatze", "Nate_Bargatze")
	add("mitch-hedberg", "dave-attell", "collaboration", yr(1995), yr(2005), 4, "NYC standup circuit contemporaries; close friends", "Mitch_Hedberg")
	add("james-corden", "jack-whitehall", "collaboration", yr(2010), nil, 3, "UK comedy contemporaries; A League of Their Own", "A_League_of_Their_Own_(panel_show)")
	add("dylan-moran", "bill-bailey", "collaboration", yr(2000), yr(2004), 5, "Black Books co-stars", "Black_Books")
	add("rowan-atkinson", "ricky-gervais", "influence", yr(1997), nil, 3, "Atkinson's character comedy influenced Gervais's approach", "Ricky_Gervais")
	add("stewart-lee", "ricky-gervais", "rivalry", yr(2005), nil, 3, "Contrasting comedy philosophies; UK comedy spectrum", "Stewart_Lee")
	add("leslie-jones", "kate-mckinnon", "troupe", yr(2014), yr(2019), 4, "SNL cast; Ghostbusters (2016) co-stars", "Ghostbusters_(2016_film)")
	add("michelle-wolf", "jon-stewart", "collaboration", yr(2016), yr(2017), 3, "Wolf wrote for The Daily Show", "Michelle_Wolf")
	add("jason-mantzoukas", "nick-kroll", "collaboration", yr(2009), nil, 4, "The League co-stars; How Did This Get Made podcast", "The_League")
	add("paul-f-tompkins", "marc-maron", "collaboration", yr(2009), nil, 3, "Alt-comedy/podcast pioneers together", "Paul_F._Tompkins")
	add("lauren-lapkus", "paul-f-tompkins", "collaboration", yr(2014), nil, 3, "UCB improv and podcast community", "Comedy_Bang!_Bang!")
	add("jason-sudeikis", "will-forte", "troupe", yr(2003), yr(2010), 4, "SNL cast members together", "Saturday_Night_Live")
	add("jason-sudeikis", "kristen-wiig", "troupe", yr(2005), yr(2010), 3, "SNL cast members together", "Saturday_Night_Live")

	return edges
}
