package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m := make(map[string]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		_, x, y, w, h := parse(s)
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				key := fmt.Sprintf("%d,%d", i, j)
				m[key]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	i := 0
	for _, v := range m {
		if v > 1 {
			i++
		}
	}

	f, err = os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		id, x, y, w, h := parse(s)
		ok := true

		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				key := fmt.Sprintf("%d,%d", i, j)
				if m[key] > 1 {
					ok = false
				}
			}
		}
		if ok {
			fmt.Printf("What is the ID of the only claim that doesn't overlap? Answer: %s", id)
		}
	}
}

func parse(s string) (id string, x, y, w, h int) {
	d := strings.Split(strings.Split(s, "@")[1], ":")

	l := strings.Split(strings.TrimSpace(d[0]), ",")
	r := strings.Split(strings.TrimSpace(d[1]), "x")

	x, _ = strconv.Atoi(l[0])
	y, _ = strconv.Atoi(l[1])
	w, _ = strconv.Atoi(r[0])
	h, _ = strconv.Atoi(r[1])

	id = strings.Split(s, "@")[0][1:]
	return
}
