package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var crabs []int

	for scanner.Scan() {
		if scanner.Text() != "" {
			parts := strings.Split(scanner.Text(), ",")
			for _, part := range parts {
				d, _ := strconv.ParseInt(part, 10, 32)
				crabs = append(crabs, int(d))
			}
		}
	}

	sort.Ints(crabs)

	medium := crabs[len(crabs)/2]

	fuel := 0
	for _, crab := range crabs {
		if medium > crab {
			fuel += medium - crab
		} else if medium < crab {
			fuel += crab - medium
		}
	}

	log.Println(fmt.Sprintf("Fuel: %d", fuel))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
