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
		x, y, w, h := parse(s)
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

	fmt.Printf("How many square inches of fabric are within two or more claims? Answer: %d\n", i)
}

func parse(s string) (x, y, w, h int) {
	d := strings.Split(strings.Split(s, "@")[1], ":")

	l := strings.Split(strings.TrimSpace(d[0]), ",")
	r := strings.Split(strings.TrimSpace(d[1]), "x")

	x, _ = strconv.Atoi(l[0])
	y, _ = strconv.Atoi(l[1])
	w, _ = strconv.Atoi(r[0])
	h, _ = strconv.Atoi(r[1])
	return
}
