package main

func allNewNodes() []Node {
	var n []Node
	n = append(n, nodesSnlAlumni()...)
	n = append(n, nodesPodcastStandup()...)
	n = append(n, nodesUkComedy()...)
	n = append(n, nodesDailyShowPolitical()...)
	n = append(n, nodesMainstreamTouring()...)
	n = append(n, nodesAltAndMisc()...)
	return n
}

func nodesSnlAlumni() []Node {
	return []Node{
		{ID: "amy-schumer", Name: "Amy Schumer", Aka: []string{}, BornYear: yr(1981), ActiveStartYear: yr(2004), Tags: []string{"standup", "tv", "film"}, Notability: 4, Links: wiki("Amy_Schumer")},
		{ID: "hasan-minhaj", Name: "Hasan Minhaj", Aka: []string{}, BornYear: yr(1985), ActiveStartYear: yr(2011), Tags: []string{"standup", "tv", "latenight"}, Notability: 4, Links: wiki("Hasan_Minhaj")},
		{ID: "pete-davidson", Name: "Pete Davidson", Aka: []string{}, BornYear: yr(1993), ActiveStartYear: yr(2013), Tags: []string{"standup", "snl", "film"}, Notability: 4, Links: wiki("Pete_Davidson")},
		{ID: "bill-hader", Name: "Bill Hader", Aka: []string{}, BornYear: yr(1978), ActiveStartYear: yr(2003), Tags: []string{"snl", "sketch", "tv", "film"}, Notability: 5, Links: wiki("Bill_Hader")},
		{ID: "kristen-wiig", Name: "Kristen Wiig", Aka: []string{}, BornYear: yr(1973), ActiveStartYear: yr(2003), Tags: []string{"snl", "sketch", "film", "improv"}, Notability: 5, Links: wiki("Kristen_Wiig")},
		{ID: "kate-mckinnon", Name: "Kate McKinnon", Aka: []string{}, BornYear: yr(1984), ActiveStartYear: yr(2007), Tags: []string{"snl", "sketch", "improv"}, Notability: 5, Links: wiki("Kate_McKinnon")},
		{ID: "andy-samberg", Name: "Andy Samberg", Aka: []string{}, BornYear: yr(1978), ActiveStartYear: yr(2001), Tags: []string{"snl", "sketch", "film", "tv"}, Notability: 4, Links: wiki("Andy_Samberg")},
		{ID: "seth-meyers", Name: "Seth Meyers", Aka: []string{}, BornYear: yr(1973), ActiveStartYear: yr(2001), Tags: []string{"snl", "latenight", "sketch"}, Notability: 4, Links: wiki("Seth_Meyers")},
		{ID: "jimmy-fallon", Name: "Jimmy Fallon", Aka: []string{}, BornYear: yr(1974), ActiveStartYear: yr(1998), Tags: []string{"snl", "latenight", "sketch"}, Notability: 5, Links: wiki("Jimmy_Fallon")},
		{ID: "tracy-morgan", Name: "Tracy Morgan", Aka: []string{}, BornYear: yr(1968), ActiveStartYear: yr(1996), Tags: []string{"snl", "tv", "standup"}, Notability: 4, Links: wiki("Tracy_Morgan")},
		{ID: "will-ferrell", Name: "Will Ferrell", Aka: []string{}, BornYear: yr(1967), ActiveStartYear: yr(1995), Tags: []string{"snl", "film", "sketch", "improv"}, Notability: 5, Links: wiki("Will_Ferrell")},
		{ID: "maya-rudolph", Name: "Maya Rudolph", Aka: []string{}, BornYear: yr(1972), ActiveStartYear: yr(1999), Tags: []string{"snl", "sketch", "film", "improv"}, Notability: 4, Links: wiki("Maya_Rudolph")},
		{ID: "kenan-thompson", Name: "Kenan Thompson", Aka: []string{}, BornYear: yr(1978), ActiveStartYear: yr(1994), Tags: []string{"snl", "sketch", "tv"}, Notability: 4, Links: wiki("Kenan_Thompson")},
		{ID: "bowen-yang", Name: "Bowen Yang", Aka: []string{}, BornYear: yr(1990), ActiveStartYear: yr(2018), Tags: []string{"snl", "sketch", "podcast"}, Notability: 3, Links: wiki("Bowen_Yang")},
		{ID: "colin-jost", Name: "Colin Jost", Aka: []string{}, BornYear: yr(1982), ActiveStartYear: yr(2005), Tags: []string{"snl", "sketch"}, Notability: 3, Links: wiki("Colin_Jost")},
		{ID: "michael-che", Name: "Michael Che", Aka: []string{}, BornYear: yr(1983), ActiveStartYear: yr(2009), Tags: []string{"snl", "standup", "sketch"}, Notability: 3, Links: wiki("Michael_Che")},
		{ID: "cecily-strong", Name: "Cecily Strong", Aka: []string{}, BornYear: yr(1984), ActiveStartYear: yr(2012), Tags: []string{"snl", "sketch", "improv"}, Notability: 3, Links: wiki("Cecily_Strong")},
		{ID: "fred-armisen", Name: "Fred Armisen", Aka: []string{}, BornYear: yr(1966), ActiveStartYear: yr(1998), Tags: []string{"snl", "sketch", "tv"}, Notability: 4, Links: wiki("Fred_Armisen")},
		{ID: "chris-farley", Name: "Chris Farley", Aka: []string{}, BornYear: yr(1964), DiedYear: yr(1997), ActiveStartYear: yr(1990), ActiveEndYear: yr(1997), Tags: []string{"snl", "sketch", "film"}, Notability: 5, Links: wiki("Chris_Farley")},
		{ID: "mike-myers", Name: "Mike Myers", Aka: []string{}, BornYear: yr(1963), ActiveStartYear: yr(1988), Tags: []string{"snl", "sketch", "film"}, Notability: 5, Links: wiki("Mike_Myers")},
		{ID: "jason-sudeikis", Name: "Jason Sudeikis", Aka: []string{}, BornYear: yr(1975), ActiveStartYear: yr(2003), Tags: []string{"snl", "improv", "tv", "film"}, Notability: 4, Links: wiki("Jason_Sudeikis")},
		{ID: "will-forte", Name: "Will Forte", Aka: []string{}, BornYear: yr(1970), ActiveStartYear: yr(2002), Tags: []string{"snl", "sketch", "tv", "film"}, Notability: 3, Links: wiki("Will_Forte")},
		{ID: "david-spade", Name: "David Spade", Aka: []string{}, BornYear: yr(1964), ActiveStartYear: yr(1988), Tags: []string{"snl", "standup", "tv", "film"}, Notability: 4, Links: wiki("David_Spade")},
	}
}
