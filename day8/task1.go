package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		if scanner.Text() != "" {
			p1 := strings.Split(scanner.Text(), " | ")
			p2 := strings.Split(p1[1], " ")
			for _, part := range p2 {
				if len(part) == 2 || len(part) == 3 || len(part) == 4 || len(part) == 7 {
					count++
				}
			}
		}
	}
	log.Println(fmt.Sprintf("Count: %d", count))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
