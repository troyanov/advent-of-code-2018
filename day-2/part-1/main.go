package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	twos := 0
	threes := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()

		stwos, sthrees := false, false
		for _, x := range s {
			c := strings.Count(s, string(x))
			switch c {
			case 2:
				if stwos {
					continue
				}
				stwos = true
				twos++
			case 3:
				if sthrees {
					continue
				}
				sthrees = true
				threes++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("What is the checksum? %d", twos*threes)
}
