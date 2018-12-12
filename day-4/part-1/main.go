package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []string
	guards := make(map[int][]int)
	guardsTotal := make(map[int]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Strings(lines)

	var ID, start, end, diff int
	for _, l := range lines {
		d := strings.Trim(strings.Split(l, "]")[1], " ")
		if strings.HasPrefix(d, "Guard") {
			ID, _ = strconv.Atoi(strings.Split(d, " ")[1][1:])
			_, exists := guards[ID]
			if !exists {
				guards[ID] = make([]int, 60)
			}
		}
		if strings.HasPrefix(d, "falls") {
			dp := strings.Split(l, "]")
			m := strings.Split(dp[0], ":")[1]
			start, _ = strconv.Atoi(m)
		}
		if strings.HasPrefix(d, "wakes") {
			dp := strings.Split(l, "]")
			m := strings.Split(dp[0], ":")[1]
			end, _ = strconv.Atoi(m)

			diff = end - start
			guardsTotal[ID] += diff

			for i := start; i <= end; i++ {
				guards[ID][i]++
			}
		}
	}

	max := 0
	ID = 0
	for k, v := range guardsTotal {
		if v > max {
			max = v
			ID = k
		}
	}

	max = 0
	minute := 0
	for i, m := range guards[ID] {
		if m > max {
			max = m
			minute = i
		}
	}
	fmt.Printf("What is the ID of the guard you chose multiplied by the minute you chose? %d*%d = %d", ID, minute, ID*minute)
}
