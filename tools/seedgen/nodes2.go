package main

func nodesPodcastStandup() []Node {
	return []Node{
		{ID: "bill-burr", Name: "Bill Burr", Aka: []string{}, BornYear: yr(1968), ActiveStartYear: yr(1992), Tags: []string{"standup", "podcast", "tv"}, Notability: 5, Links: wiki("Bill_Burr")},
		{ID: "marc-maron", Name: "Marc Maron", Aka: []string{}, BornYear: yr(1963), ActiveStartYear: yr(1987), Tags: []string{"standup", "podcast", "tv"}, Notability: 4, Links: wiki("Marc_Maron")},
		{ID: "joe-rogan", Name: "Joe Rogan", Aka: []string{}, BornYear: yr(1967), ActiveStartYear: yr(1988), Tags: []string{"standup", "podcast", "tv"}, Notability: 5, Links: wiki("Joe_Rogan")},
		{ID: "andrew-santino", Name: "Andrew Santino", Aka: []string{}, BornYear: yr(1983), ActiveStartYear: yr(2007), Tags: []string{"standup", "podcast", "tv"}, Notability: 3, Links: wiki("Andrew_Santino")},
		{ID: "bobby-lee", Name: "Bobby Lee", Aka: []string{}, BornYear: yr(1971), ActiveStartYear: yr(2001), Tags: []string{"standup", "podcast", "sketch", "tv"}, Notability: 3, Links: wiki("Bobby_Lee")},
		{ID: "theo-von", Name: "Theo Von", Aka: []string{}, BornYear: yr(1980), ActiveStartYear: yr(2003), Tags: []string{"standup", "podcast"}, Notability: 3, Links: wiki("Theo_Von")},
		{ID: "nikki-glaser", Name: "Nikki Glaser", Aka: []string{}, BornYear: yr(1984), ActiveStartYear: yr(2006), Tags: []string{"standup", "tv"}, Notability: 4, Links: wiki("Nikki_Glaser")},
		{ID: "whitney-cummings", Name: "Whitney Cummings", Aka: []string{}, BornYear: yr(1982), ActiveStartYear: yr(2004), Tags: []string{"standup", "tv"}, Notability: 3, Links: wiki("Whitney_Cummings")},
		{ID: "tom-segura", Name: "Tom Segura", Aka: []string{}, BornYear: yr(1979), ActiveStartYear: yr(2007), Tags: []string{"standup", "podcast"}, Notability: 4, Links: wiki("Tom_Segura")},
		{ID: "bert-kreischer", Name: "Bert Kreischer", Aka: []string{"The Machine"}, BornYear: yr(1972), ActiveStartYear: yr(1997), Tags: []string{"standup", "podcast"}, Notability: 4, Links: wiki("Bert_Kreischer")},
		{ID: "christina-p", Name: "Christina Pazsitzky", Aka: []string{"Christina P"}, BornYear: yr(1976), ActiveStartYear: yr(2002), Tags: []string{"standup", "podcast"}, Notability: 3, Links: wiki("Christina_Pazsitzky")},
		{ID: "shane-gillis", Name: "Shane Gillis", Aka: []string{}, BornYear: yr(1987), ActiveStartYear: yr(2012), Tags: []string{"standup", "podcast", "sketch"}, Notability: 4, Links: wiki("Shane_Gillis")},
		{ID: "mark-normand", Name: "Mark Normand", Aka: []string{}, BornYear: yr(1983), ActiveStartYear: yr(2006), Tags: []string{"standup", "podcast"}, Notability: 3, Links: wiki("Mark_Normand")},
		{ID: "sam-morril", Name: "Sam Morril", Aka: []string{}, BornYear: yr(1986), ActiveStartYear: yr(2008), Tags: []string{"standup", "podcast"}, Notability: 3, Links: wiki("Sam_Morril")},
	}
}

func nodesAltAndMisc() []Node {
	return []Node{
		{ID: "sarah-silverman", Name: "Sarah Silverman", Aka: []string{}, BornYear: yr(1970), ActiveStartYear: yr(1992), Tags: []string{"standup", "tv", "film"}, Notability: 4, Links: wiki("Sarah_Silverman")},
		{ID: "patton-oswalt", Name: "Patton Oswalt", Aka: []string{}, BornYear: yr(1969), ActiveStartYear: yr(1988), Tags: []string{"standup", "tv", "film"}, Notability: 4, Links: wiki("Patton_Oswalt")},
		{ID: "tig-notaro", Name: "Tig Notaro", Aka: []string{}, BornYear: yr(1971), ActiveStartYear: yr(1997), Tags: []string{"standup", "tv"}, Notability: 4, Links: wiki("Tig_Notaro")},
		{ID: "neal-brennan", Name: "Neal Brennan", Aka: []string{}, BornYear: yr(1973), ActiveStartYear: yr(1996), Tags: []string{"standup", "tv", "sketch"}, Notability: 3, Links: wiki("Neal_Brennan")},
		{ID: "judd-apatow", Name: "Judd Apatow", Aka: []string{}, BornYear: yr(1967), ActiveStartYear: yr(1985), Tags: []string{"film", "tv"}, Notability: 5, Links: wiki("Judd_Apatow")},
		{ID: "eric-andre", Name: "Eric Andre", Aka: []string{}, BornYear: yr(1983), ActiveStartYear: yr(2009), Tags: []string{"tv", "standup", "sketch"}, Notability: 3, Links: wiki("Eric_André")},
		{ID: "nathan-fielder", Name: "Nathan Fielder", Aka: []string{}, BornYear: yr(1983), ActiveStartYear: yr(2007), Tags: []string{"tv"}, Notability: 4, Links: wiki("Nathan_Fielder")},
		{ID: "sacha-baron-cohen", Name: "Sacha Baron Cohen", Aka: []string{"Borat", "Ali G", "Bruno"}, BornYear: yr(1971), ActiveStartYear: yr(1995), Tags: []string{"film", "tv", "sketch"}, Notability: 5, Links: wiki("Sacha_Baron_Cohen")},
		{ID: "keegan-michael-key", Name: "Keegan-Michael Key", Aka: []string{}, BornYear: yr(1971), ActiveStartYear: yr(2003), Tags: []string{"sketch", "improv", "tv", "film"}, Notability: 4, Links: wiki("Keegan-Michael_Key")},
		{ID: "jordan-peele", Name: "Jordan Peele", Aka: []string{}, BornYear: yr(1979), ActiveStartYear: yr(2003), Tags: []string{"sketch", "improv", "film"}, Notability: 5, Links: wiki("Jordan_Peele")},
		{ID: "donald-glover", Name: "Donald Glover", Aka: []string{"Childish Gambino"}, BornYear: yr(1983), ActiveStartYear: yr(2006), Tags: []string{"sketch", "tv", "standup", "film"}, Notability: 5, Links: wiki("Donald_Glover")},
		{ID: "bo-burnham", Name: "Bo Burnham", Aka: []string{}, BornYear: yr(1990), ActiveStartYear: yr(2006), Tags: []string{"standup", "film"}, Notability: 4, Links: wiki("Bo_Burnham")},
		{ID: "maria-bamford", Name: "Maria Bamford", Aka: []string{}, BornYear: yr(1970), ActiveStartYear: yr(1990), Tags: []string{"standup", "tv"}, Notability: 4, Links: wiki("Maria_Bamford")},
		{ID: "hannibal-buress", Name: "Hannibal Buress", Aka: []string{}, BornYear: yr(1983), ActiveStartYear: yr(2009), Tags: []string{"standup", "tv", "podcast"}, Notability: 3, Links: wiki("Hannibal_Buress")},
		{ID: "nick-kroll", Name: "Nick Kroll", Aka: []string{}, BornYear: yr(1978), ActiveStartYear: yr(2003), Tags: []string{"sketch", "standup", "tv", "improv"}, Notability: 4, Links: wiki("Nick_Kroll")},
		{ID: "aziz-ansari", Name: "Aziz Ansari", Aka: []string{}, BornYear: yr(1983), ActiveStartYear: yr(2005), Tags: []string{"standup", "tv"}, Notability: 4, Links: wiki("Aziz_Ansari")},
		{ID: "aubrey-plaza", Name: "Aubrey Plaza", Aka: []string{}, BornYear: yr(1984), ActiveStartYear: yr(2004), Tags: []string{"tv", "film", "improv"}, Notability: 4, Links: wiki("Aubrey_Plaza")},
		{ID: "jason-mantzoukas", Name: "Jason Mantzoukas", Aka: []string{}, BornYear: yr(1972), ActiveStartYear: yr(2003), Tags: []string{"improv", "podcast", "tv", "film"}, Notability: 3, Links: wiki("Jason_Mantzoukas")},
		{ID: "paul-f-tompkins", Name: "Paul F. Tompkins", Aka: []string{}, BornYear: yr(1968), ActiveStartYear: yr(1988), Tags: []string{"standup", "improv", "podcast"}, Notability: 3, Links: wiki("Paul_F._Tompkins")},
		{ID: "lauren-lapkus", Name: "Lauren Lapkus", Aka: []string{}, BornYear: yr(1985), ActiveStartYear: yr(2010), Tags: []string{"improv", "podcast", "tv"}, Notability: 2, Links: wiki("Lauren_Lapkus")},
		{ID: "kate-berlant", Name: "Kate Berlant", Aka: []string{}, BornYear: yr(1988), ActiveStartYear: yr(2012), Tags: []string{"standup", "tv"}, Notability: 2, Links: wiki("Kate_Berlant")},
		{ID: "aparna-nancherla", Name: "Aparna Nancherla", Aka: []string{}, BornYear: yr(1982), ActiveStartYear: yr(2008), Tags: []string{"standup", "tv"}, Notability: 2, Links: wiki("Aparna_Nancherla")},
		{ID: "jim-jefferies", Name: "Jim Jefferies", Aka: []string{}, BornYear: yr(1977), ActiveStartYear: yr(1998), Tags: []string{"standup", "tv"}, Notability: 4, Links: wiki("Jim_Jefferies")},
		{ID: "demetri-martin", Name: "Demetri Martin", Aka: []string{}, BornYear: yr(1973), ActiveStartYear: yr(1997), Tags: []string{"standup", "tv"}, Notability: 3, Links: wiki("Demetri_Martin")},
		{ID: "david-cross", Name: "David Cross", Aka: []string{}, BornYear: yr(1964), ActiveStartYear: yr(1991), Tags: []string{"standup", "sketch", "tv"}, Notability: 4, Links: wiki("David_Cross")},
		{ID: "bob-odenkirk", Name: "Bob Odenkirk", Aka: []string{}, BornYear: yr(1962), ActiveStartYear: yr(1987), Tags: []string{"sketch", "tv", "film"}, Notability: 5, Links: wiki("Bob_Odenkirk")},
		{ID: "jb-smoove", Name: "JB Smoove", Aka: []string{"Jerry Angelo Brooks"}, BornYear: yr(1965), ActiveStartYear: yr(1995), Tags: []string{"standup", "tv", "improv"}, Notability: 3, Links: wiki("J._B._Smoove")},
		{ID: "larry-david", Name: "Larry David", Aka: []string{}, BornYear: yr(1947), ActiveStartYear: yr(1977), Tags: []string{"tv", "standup"}, Notability: 5, Links: wiki("Larry_David")},
		{ID: "hannah-gadsby", Name: "Hannah Gadsby", Aka: []string{}, BornYear: yr(1978), ActiveStartYear: yr(2006), Tags: []string{"standup"}, Notability: 4, Links: wiki("Hannah_Gadsby")},
		{ID: "mindy-kaling", Name: "Mindy Kaling", Aka: []string{}, BornYear: yr(1979), ActiveStartYear: yr(2002), Tags: []string{"tv", "film"}, Notability: 4, Links: wiki("Mindy_Kaling")},
		{ID: "kumail-nanjiani", Name: "Kumail Nanjiani", Aka: []string{}, BornYear: yr(1978), ActiveStartYear: yr(2008), Tags: []string{"standup", "film", "tv"}, Notability: 4, Links: wiki("Kumail_Nanjiani")},
		{ID: "ken-jeong", Name: "Ken Jeong", Aka: []string{}, BornYear: yr(1969), ActiveStartYear: yr(1995), Tags: []string{"film", "standup", "tv"}, Notability: 4, Links: wiki("Ken_Jeong")},
	}
}
