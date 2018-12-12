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

	var ID, start, end int
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

			for i := start; i <= end; i++ {
				guards[ID][i] = guards[ID][i] + 1
			}
		}
	}

	guardID := 0
	minute := 0
	frequency := 0

	for ID := range guards {
		for i, v := range guards[ID] {
			if v > frequency {
				frequency = v
				minute = i
				guardID = ID
			}
		}
	}
	fmt.Printf("What is the ID of the guard you chose multiplied by the minute you chose? %d*%d = %d",
		guardID, minute, guardID*minute)
}
