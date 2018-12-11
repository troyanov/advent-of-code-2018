package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var values []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		i, _ := strconv.Atoi(s)
		values = append(values, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	total := 0
	totals := make(map[int]struct{})
	for {
		for _, v := range values {
			total += v
			_, exists := totals[total]
			if exists {
				fmt.Printf("What is the first frequency your device reaches twice?? %d", total)
				return
			}
			totals[total] = struct{}{}
		}
	}
}
