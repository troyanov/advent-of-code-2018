package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"sync"
)

func main() {
	b, err := ioutil.ReadFile("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	var units [][]string
	for i := 65; i <= 90; i++ {
		units = append(units, []string{string(rune(i)), string(rune(i + 32))})
	}

	s := string(b)
	var lengths []int

	wg := sync.WaitGroup{}
	for _, u := range units {
		wg.Add(1)
		go func(u []string) {
			defer wg.Done()
			short := remove(s, u)
			reduced := reduce(short)
			lengths = append(lengths, len([]rune(reduced)))
		}(u)
	}
	wg.Wait()
	sort.Ints(lengths)

	fmt.Printf("What is the length of the shortest polymer you can produce? %d", lengths[0])
}

func remove(s string, units []string) string {
	for _, u := range units {
		s = strings.Replace(s, u, "", -1)
	}
	return s
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
