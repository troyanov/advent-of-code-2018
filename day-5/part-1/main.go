package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ioutil.ReadFile("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	s := reduce(string(b))
	fmt.Printf("How many units remain after fully reacting the polymer you scanned?? %d", len([]rune(s)))
}

func reduce(s string) string {
	var r string
	for i := range s {
		if i >= len(s)-1 {
			return s
		}
		if s[i] < 97 {
			if s[i] == s[i+1]-32 {
				r = s[:i] + s[i+2:]
				return reduce(r)
			}
		}
		if s[i] >= 97 {
			if s[i] == s[i+1]+32 {
				r = s[:i] + s[i+2:]
				return reduce(r)
			}
		}
	}
	if r == s {
		return s
	}
	return reduce(r)
}
