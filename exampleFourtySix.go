package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func FourtySix() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("punch peach"))
	fmt.Println(r.FindStringIndex("punch peach"))
	fmt.Println(r.FindStringSubmatch("punch peach"))
	fmt.Println(r.FindStringSubmatchIndex("punch peach"))
	fmt.Println(r.FindAllString("punch peach pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("punch peach pinch", -1))
	fmt.Println(r.FindAllString("punch peach pinch", 2))
	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
