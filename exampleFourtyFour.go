package main

import(
	s "strings"
	"fmt"
)

var p = fmt.Println

func FourtyFour() {

	p("Contains: ", s.Contains("tests", "es"))
	p("Count: ", s.Count("tests", "s"))
	p("HasPrefix: ", s.HasPrefix("tests", "te"))
	p("HasSuffix: ", s.HasSuffix("tests", "ts"))
	p("Index: ", s.Index("tests", "e"))
	p("Join: ", s.Join([]string{"tests", "12"}, "-"))
	p("Repeat: ", s.Repeat("lance", 5))
	p("Replace: ", s.Replace("lance", "a", "A", -1))
	p("Replace: ", s.Replace("lance", "a", "A", 1))
	p("Split: ", s.Split("lance", ""))
	p("ToLower: ", s.ToLower("LANCE"))
	p("ToUpper: ", s.ToUpper("lance"))


	p("Length: ", len("Lance"))
	p("Char: ", "lance"[4])



}