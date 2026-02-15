package main

func nodesMainstreamTouring() []Node {
	return []Node{
		{ID: "gabriel-iglesias", Name: "Gabriel Iglesias", Aka: []string{"Fluffy"}, BornYear: yr(1976), ActiveStartYear: yr(1997), Tags: []string{"standup"}, Notability: 4, Links: wiki("Gabriel_Iglesias")},
		{ID: "jo-koy", Name: "Jo Koy", Aka: []string{"Joseph Glenn Herbert"}, BornYear: yr(1971), ActiveStartYear: yr(1994), Tags: []string{"standup"}, Notability: 4, Links: wiki("Jo_Koy")},
		{ID: "russell-peters", Name: "Russell Peters", Aka: []string{}, BornYear: yr(1970), ActiveStartYear: yr(1989), Tags: []string{"standup"}, Notability: 4, Links: wiki("Russell_Peters")},
		{ID: "nate-bargatze", Name: "Nate Bargatze", Aka: []string{}, BornYear: yr(1979), ActiveStartYear: yr(2003), Tags: []string{"standup"}, Notability: 4, Links: wiki("Nate_Bargatze")},
		{ID: "taylor-tomlinson", Name: "Taylor Tomlinson", Aka: []string{}, BornYear: yr(1993), ActiveStartYear: yr(2015), Tags: []string{"standup", "tv"}, Notability: 3, Links: wiki("Taylor_Tomlinson")},
		{ID: "jerrod-carmichael", Name: "Jerrod Carmichael", Aka: []string{}, BornYear: yr(1987), ActiveStartYear: yr(2011), Tags: []string{"standup", "tv"}, Notability: 4, Links: wiki("Jerrod_Carmichael")},
		{ID: "tiffany-haddish", Name: "Tiffany Haddish", Aka: []string{}, BornYear: yr(1979), ActiveStartYear: yr(2005), Tags: []string{"standup", "film", "tv"}, Notability: 4, Links: wiki("Tiffany_Haddish")},
		{ID: "wanda-sykes", Name: "Wanda Sykes", Aka: []string{}, BornYear: yr(1964), ActiveStartYear: yr(1987), Tags: []string{"standup", "tv"}, Notability: 4, Links: wiki("Wanda_Sykes")},
		{ID: "chelsea-handler", Name: "Chelsea Handler", Aka: []string{}, BornYear: yr(1975), ActiveStartYear: yr(2002), Tags: []string{"standup", "tv", "latenight"}, Notability: 4, Links: wiki("Chelsea_Handler")},
		{ID: "leslie-jones", Name: "Leslie Jones", Aka: []string{}, BornYear: yr(1967), ActiveStartYear: yr(1987), Tags: []string{"snl", "standup", "tv", "film"}, Notability: 3, Links: wiki("Leslie_Jones_(comedian)")},
		{ID: "iliza-shlesinger", Name: "Iliza Shlesinger", Aka: []string{}, BornYear: yr(1983), ActiveStartYear: yr(2007), Tags: []string{"standup"}, Notability: 3, Links: wiki("Iliza_Shlesinger")},
		{ID: "daniel-tosh", Name: "Daniel Tosh", Aka: []string{}, BornYear: yr(1975), ActiveStartYear: yr(1998), Tags: []string{"standup", "tv"}, Notability: 4, Links: wiki("Daniel_Tosh")},
		{ID: "ali-siddiq", Name: "Ali Siddiq", Aka: []string{}, BornYear: yr(1976), ActiveStartYear: yr(2007), Tags: []string{"standup"}, Notability: 3, Links: wiki("Ali_Siddiq")},
		{ID: "deon-cole", Name: "Deon Cole", Aka: []string{}, BornYear: yr(1971), ActiveStartYear: yr(1998), Tags: []string{"standup", "tv"}, Notability: 3, Links: wiki("Deon_Cole")},
		{ID: "mike-epps", Name: "Mike Epps", Aka: []string{}, BornYear: yr(1970), ActiveStartYear: yr(1995), Tags: []string{"standup", "film"}, Notability: 3, Links: wiki("Mike_Epps")},
		{ID: "katt-williams", Name: "Katt Williams", Aka: []string{}, BornYear: yr(1971), ActiveStartYear: yr(1999), Tags: []string{"standup", "film"}, Notability: 4, Links: wiki("Katt_Williams")},
		{ID: "cedric", Name: "Cedric the Entertainer", Aka: []string{"Cedric Antonio Kyles"}, BornYear: yr(1964), ActiveStartYear: yr(1987), Tags: []string{"standup", "film", "tv"}, Notability: 4, Links: wiki("Cedric_the_Entertainer")},
		{ID: "steve-harvey", Name: "Steve Harvey", Aka: []string{}, BornYear: yr(1957), ActiveStartYear: yr(1985), Tags: []string{"standup", "tv"}, Notability: 5, Links: wiki("Steve_Harvey")},
		{ID: "dave-attell", Name: "Dave Attell", Aka: []string{}, BornYear: yr(1965), ActiveStartYear: yr(1988), Tags: []string{"standup"}, Notability: 4, Links: wiki("Dave_Attell")},
		{ID: "jeff-ross", Name: "Jeff Ross", Aka: []string{"Roastmaster General"}, BornYear: yr(1965), ActiveStartYear: yr(1990), Tags: []string{"standup", "tv"}, Notability: 3, Links: wiki("Jeff_Ross")},
		{ID: "jeff-dunham", Name: "Jeff Dunham", Aka: []string{}, BornYear: yr(1962), ActiveStartYear: yr(1976), Tags: []string{"standup"}, Notability: 4, Links: wiki("Jeff_Dunham")},
		{ID: "jim-gaffigan", Name: "Jim Gaffigan", Aka: []string{}, BornYear: yr(1966), ActiveStartYear: yr(1991), Tags: []string{"standup", "tv", "film"}, Notability: 4, Links: wiki("Jim_Gaffigan")},
		{ID: "brian-regan", Name: "Brian Regan", Aka: []string{}, BornYear: yr(1958), ActiveStartYear: yr(1980), Tags: []string{"standup"}, Notability: 4, Links: wiki("Brian_Regan_(comedian)")},
		{ID: "doug-stanhope", Name: "Doug Stanhope", Aka: []string{}, BornYear: yr(1967), ActiveStartYear: yr(1990), Tags: []string{"standup", "podcast"}, Notability: 3, Links: wiki("Doug_Stanhope")},
		{ID: "dane-cook", Name: "Dane Cook", Aka: []string{}, BornYear: yr(1972), ActiveStartYear: yr(1990), Tags: []string{"standup", "film"}, Notability: 4, Links: wiki("Dane_Cook")},
		{ID: "norm-macdonald", Name: "Norm Macdonald", Aka: []string{}, BornYear: yr(1959), DiedYear: yr(2021), ActiveStartYear: yr(1986), ActiveEndYear: yr(2021), Tags: []string{"standup", "snl", "tv"}, Notability: 5, Links: wiki("Norm_Macdonald")},
		{ID: "mitch-hedberg", Name: "Mitch Hedberg", Aka: []string{}, BornYear: yr(1968), DiedYear: yr(2005), ActiveStartYear: yr(1990), ActiveEndYear: yr(2005), Tags: []string{"standup"}, Notability: 4, Links: wiki("Mitch_Hedberg")},
	}
}
