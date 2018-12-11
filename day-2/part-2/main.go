package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var ids []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		ids = append(ids, s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, h := range ids {
		for _, j := range ids {
			var diff, d int
			if h == j {
				continue
			}

			for i := range h {
				if h[i] != j[i] {
					diff++
					d = i
				}
			}
			if diff == 1 {
				common := h[:d] + h[d+1:]
				fmt.Printf("What letters are common between the two correct box IDs? %s", common)
				return
			}
		}
	}
}
