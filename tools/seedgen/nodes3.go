package main

func nodesUkComedy() []Node {
	return []Node{
		{ID: "ricky-gervais", Name: "Ricky Gervais", Aka: []string{}, BornYear: yr(1961), ActiveStartYear: yr(1997), Tags: []string{"standup", "tv"}, Notability: 5, Links: wiki("Ricky_Gervais")},
		{ID: "steve-coogan", Name: "Steve Coogan", Aka: []string{"Alan Partridge"}, BornYear: yr(1965), ActiveStartYear: yr(1988), Tags: []string{"tv", "film", "standup"}, Notability: 4, Links: wiki("Steve_Coogan")},
		{ID: "james-acaster", Name: "James Acaster", Aka: []string{}, BornYear: yr(1985), ActiveStartYear: yr(2008), Tags: []string{"standup", "panel"}, Notability: 4, Links: wiki("James_Acaster")},
		{ID: "dara-o-briain", Name: "Dara Ó Briain", Aka: []string{}, BornYear: yr(1972), ActiveStartYear: yr(1994), Tags: []string{"standup", "panel", "tv"}, Notability: 3, Links: wiki("Dara_Ó_Briain")},
		{ID: "jimmy-carr", Name: "Jimmy Carr", Aka: []string{}, BornYear: yr(1972), ActiveStartYear: yr(2000), Tags: []string{"standup", "panel", "tv"}, Notability: 4, Links: wiki("Jimmy_Carr")},
		{ID: "katherine-ryan", Name: "Katherine Ryan", Aka: []string{}, BornYear: yr(1983), ActiveStartYear: yr(2008), Tags: []string{"standup", "panel", "tv"}, Notability: 3, Links: wiki("Katherine_Ryan")},
		{ID: "richard-ayoade", Name: "Richard Ayoade", Aka: []string{}, BornYear: yr(1977), ActiveStartYear: yr(2001), Tags: []string{"tv", "panel", "film"}, Notability: 4, Links: wiki("Richard_Ayoade")},
		{ID: "noel-fielding", Name: "Noel Fielding", Aka: []string{}, BornYear: yr(1973), ActiveStartYear: yr(1998), Tags: []string{"tv", "panel", "sketch"}, Notability: 4, Links: wiki("Noel_Fielding")},
		{ID: "david-mitchell", Name: "David Mitchell", Aka: []string{}, BornYear: yr(1974), ActiveStartYear: yr(1995), Tags: []string{"tv", "panel", "sketch"}, Notability: 4, Links: wiki("David_Mitchell_(comedian)")},
		{ID: "robert-webb", Name: "Robert Webb", Aka: []string{}, BornYear: yr(1972), ActiveStartYear: yr(1995), Tags: []string{"tv", "panel", "sketch"}, Notability: 3, Links: wiki("Robert_Webb")},
		{ID: "olivia-colman", Name: "Olivia Colman", Aka: []string{}, BornYear: yr(1974), ActiveStartYear: yr(1998), Tags: []string{"tv", "film"}, Notability: 5, Links: wiki("Olivia_Colman")},
		{ID: "rowan-atkinson", Name: "Rowan Atkinson", Aka: []string{"Mr. Bean"}, BornYear: yr(1955), ActiveStartYear: yr(1978), Tags: []string{"tv", "sketch", "film"}, Notability: 5, Links: wiki("Rowan_Atkinson")},
		{ID: "phoebe-waller-bridge", Name: "Phoebe Waller-Bridge", Aka: []string{}, BornYear: yr(1985), ActiveStartYear: yr(2009), Tags: []string{"tv", "film"}, Notability: 5, Links: wiki("Phoebe_Waller-Bridge")},
		{ID: "aisling-bea", Name: "Aisling Bea", Aka: []string{}, BornYear: yr(1984), ActiveStartYear: yr(2010), Tags: []string{"standup", "panel", "tv"}, Notability: 3, Links: wiki("Aisling_Bea")},
		{ID: "james-corden", Name: "James Corden", Aka: []string{}, BornYear: yr(1978), ActiveStartYear: yr(1996), Tags: []string{"tv", "latenight"}, Notability: 4, Links: wiki("James_Corden")},
		{ID: "sarah-millican", Name: "Sarah Millican", Aka: []string{}, BornYear: yr(1975), ActiveStartYear: yr(2004), Tags: []string{"standup", "panel", "tv"}, Notability: 3, Links: wiki("Sarah_Millican")},
		{ID: "lee-mack", Name: "Lee Mack", Aka: []string{}, BornYear: yr(1968), ActiveStartYear: yr(1994), Tags: []string{"standup", "panel", "tv"}, Notability: 4, Links: wiki("Lee_Mack")},
		{ID: "rob-brydon", Name: "Rob Brydon", Aka: []string{}, BornYear: yr(1965), ActiveStartYear: yr(1992), Tags: []string{"panel", "tv", "standup"}, Notability: 3, Links: wiki("Rob_Brydon")},
		{ID: "michael-mcintyre", Name: "Michael McIntyre", Aka: []string{}, BornYear: yr(1976), ActiveStartYear: yr(1999), Tags: []string{"standup", "tv"}, Notability: 4, Links: wiki("Michael_McIntyre")},
		{ID: "jack-whitehall", Name: "Jack Whitehall", Aka: []string{}, BornYear: yr(1988), ActiveStartYear: yr(2007), Tags: []string{"standup", "panel", "tv"}, Notability: 3, Links: wiki("Jack_Whitehall")},
		{ID: "romesh-ranganathan", Name: "Romesh Ranganathan", Aka: []string{}, BornYear: yr(1978), ActiveStartYear: yr(2010), Tags: []string{"standup", "panel", "tv"}, Notability: 3, Links: wiki("Romesh_Ranganathan")},
		{ID: "nish-kumar", Name: "Nish Kumar", Aka: []string{}, BornYear: yr(1985), ActiveStartYear: yr(2010), Tags: []string{"standup", "panel", "tv"}, Notability: 3, Links: wiki("Nish_Kumar")},
		{ID: "frankie-boyle", Name: "Frankie Boyle", Aka: []string{}, BornYear: yr(1972), ActiveStartYear: yr(1996), Tags: []string{"standup", "panel", "tv"}, Notability: 3, Links: wiki("Frankie_Boyle")},
		{ID: "stewart-lee", Name: "Stewart Lee", Aka: []string{}, BornYear: yr(1968), ActiveStartYear: yr(1988), Tags: []string{"standup"}, Notability: 4, Links: wiki("Stewart_Lee")},
		{ID: "dylan-moran", Name: "Dylan Moran", Aka: []string{}, BornYear: yr(1971), ActiveStartYear: yr(1993), Tags: []string{"standup", "tv"}, Notability: 4, Links: wiki("Dylan_Moran")},
		{ID: "bill-bailey", Name: "Bill Bailey", Aka: []string{}, BornYear: yr(1964), ActiveStartYear: yr(1984), Tags: []string{"standup", "panel", "tv"}, Notability: 4, Links: wiki("Bill_Bailey")},
		{ID: "simon-amstell", Name: "Simon Amstell", Aka: []string{}, BornYear: yr(1979), ActiveStartYear: yr(1998), Tags: []string{"standup", "tv"}, Notability: 3, Links: wiki("Simon_Amstell")},
		{ID: "mo-gilligan", Name: "Mo Gilligan", Aka: []string{}, BornYear: yr(1988), ActiveStartYear: yr(2017), Tags: []string{"standup", "panel", "tv"}, Notability: 3, Links: wiki("Mo_Gilligan")},
	}
}

func nodesDailyShowPolitical() []Node {
	return []Node{
		{ID: "stephen-colbert", Name: "Stephen Colbert", Aka: []string{}, BornYear: yr(1964), ActiveStartYear: yr(1988), Tags: []string{"tv", "latenight", "sketch", "improv"}, Notability: 5, Links: wiki("Stephen_Colbert")},
		{ID: "jon-stewart", Name: "Jon Stewart", Aka: []string{"Jonathan Stuart Leibowitz"}, BornYear: yr(1962), ActiveStartYear: yr(1986), Tags: []string{"tv", "latenight", "standup"}, Notability: 5, Links: wiki("Jon_Stewart")},
		{ID: "trevor-noah", Name: "Trevor Noah", Aka: []string{}, BornYear: yr(1984), ActiveStartYear: yr(2002), Tags: []string{"standup", "tv", "latenight"}, Notability: 5, Links: wiki("Trevor_Noah")},
		{ID: "samantha-bee", Name: "Samantha Bee", Aka: []string{}, BornYear: yr(1969), ActiveStartYear: yr(2003), Tags: []string{"tv", "latenight"}, Notability: 3, Links: wiki("Samantha_Bee")},
		{ID: "john-oliver", Name: "John Oliver", Aka: []string{}, BornYear: yr(1977), ActiveStartYear: yr(2001), Tags: []string{"tv", "latenight", "standup"}, Notability: 5, Links: wiki("John_Oliver")},
		{ID: "larry-wilmore", Name: "Larry Wilmore", Aka: []string{}, BornYear: yr(1961), ActiveStartYear: yr(1987), Tags: []string{"tv"}, Notability: 3, Links: wiki("Larry_Wilmore")},
		{ID: "jordan-klepper", Name: "Jordan Klepper", Aka: []string{}, BornYear: yr(1979), ActiveStartYear: yr(2008), Tags: []string{"tv", "improv"}, Notability: 2, Links: wiki("Jordan_Klepper")},
		{ID: "desus-nice", Name: "Desus Nice", Aka: []string{"Daniel Baker"}, BornYear: yr(1983), ActiveStartYear: yr(2013), Tags: []string{"tv", "podcast"}, Notability: 3, Links: wiki("Desus_Nice")},
		{ID: "the-kid-mero", Name: "The Kid Mero", Aka: []string{"Joel Martinez"}, BornYear: yr(1983), ActiveStartYear: yr(2013), Tags: []string{"tv", "podcast"}, Notability: 3, Links: wiki("The_Kid_Mero")},
		{ID: "ronny-chieng", Name: "Ronny Chieng", Aka: []string{}, BornYear: yr(1985), ActiveStartYear: yr(2009), Tags: []string{"standup", "tv"}, Notability: 3, Links: wiki("Ronny_Chieng")},
		{ID: "michelle-wolf", Name: "Michelle Wolf", Aka: []string{}, BornYear: yr(1985), ActiveStartYear: yr(2011), Tags: []string{"standup", "tv"}, Notability: 3, Links: wiki("Michelle_Wolf")},
		{ID: "conan-obrien", Name: "Conan O'Brien", Aka: []string{"Coco"}, BornYear: yr(1963), ActiveStartYear: yr(1988), Tags: []string{"tv", "latenight", "podcast"}, Notability: 5, Links: wiki("Conan_O'Brien")},
	}
}
