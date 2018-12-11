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

	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		i, _ := strconv.Atoi(s)
		total += i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("What is the resulting frequency after all of the changes in frequency have been applied? %d", total)
}
