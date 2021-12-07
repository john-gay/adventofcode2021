package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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

	total := 0
	for _, crab := range crabs {
		total += crab
	}

	mean := float64(total) / float64(len(crabs))
	fmt.Println(fmt.Sprintf("Mean: %f", mean))

	rounded := int(math.Round(mean))
	fmt.Println(fmt.Sprintf("Mean: %d", rounded))

	boundaries := []int{int(math.Ceil(mean)), int(math.Floor(mean))}

	minFuel := -1
	for _, mean := range boundaries {
		fuel := 0
		for _, crab := range crabs {
			if mean > crab {
				fuel += toAdd(mean - crab)
			} else if mean < crab {
				fuel += toAdd(crab - mean)
			}
		}
		if minFuel == -1 || fuel < minFuel {
			minFuel = fuel
		}
	}

	log.Println(fmt.Sprintf("Fuel: %d", minFuel))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func toAdd(n int) int {
	sum := float64(n) / 2 * float64(n+1)
	return int(sum)
}
