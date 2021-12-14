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

	file, err := os.Open("day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var template []string
	rules := map[string]string{}

	for scanner.Scan() {
		if scanner.Text() != "" {
			if len(template) == 0 {
				parts := strings.Split(scanner.Text(), "")
				for _, part := range parts {
					template = append(template, part)
				}
			} else {
				parts := strings.Split(scanner.Text(), " -> ")
				rules[parts[0]] = parts[1]
			}
		}
	}

	templateMap := map[string]int{}
	for i := 0; i < len(template)-1; i++ {
		templateMap[strings.Join(template[i:i+2], "")]++
	}

	fmt.Println(fmt.Sprintf("Part 1: %d", makeTurns(10, templateMap, rules)))
	fmt.Println(fmt.Sprintf("Part 2: %d", makeTurns(40, templateMap, rules)))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func makeTurns(turns int, template map[string]int, rules map[string]string) int {
	for turn := 0; turn < turns; turn++ {
		nextTemplate := map[string]int{}
		for key, value := range template {
			keyParts := strings.Split(key, "")
			nextTemplate[keyParts[0]+rules[key]] += value
			nextTemplate[rules[key]+keyParts[1]] += value
		}
		template = nextTemplate
	}

	return countElements(template)
}

func countElements(template map[string]int) int {
	counts := map[string]int{}
	for key, value := range template {
		keyParts := strings.Split(key, "")
		for _, keyPart := range keyParts {
			counts[keyPart] += value
		}
	}

	min := -1
	max := 0
	for _, v := range counts {
		value := (v + 1) / 2
		if min < 0 || min > value {
			min = value
		}
		if max < value {
			max = value
		}
	}

	return max - min
}
