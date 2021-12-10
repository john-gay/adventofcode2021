package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var opens []string
	var scores []int

	for scanner.Scan() {
		opens = []string{}
		if scanner.Text() != "" {
			parts := strings.Split(scanner.Text(), "")
			valid := true
			for _, part := range parts {
				if !valid {
					break
				}
				switch part {
				case ")":
					if opens[len(opens)-1] == "(" {
						opens = opens[:len(opens)-1]
					} else {
						valid = false
					}
				case "]":
					if opens[len(opens)-1] == "[" {
						opens = opens[:len(opens)-1]
					} else {
						valid = false
					}
				case "}":
					if opens[len(opens)-1] == "{" {
						opens = opens[:len(opens)-1]
					} else {
						valid = false
					}
				case ">":
					if opens[len(opens)-1] == "<" {
						opens = opens[:len(opens)-1]
					} else {
						valid = false
					}
				default:
					opens = append(opens, part)
				}
			}
			if valid {
				score := 0
				for i := len(opens) - 1; i >= 0; i-- {
					score *= 5
					switch opens[i] {
					case "(":
						score += 1
					case "[":
						score += 2
					case "{":
						score += 3
					case "<":
						score += 4
					}
				}
				scores = append(scores, score)
			}
		}
	}

	sort.Ints(scores)

	fmt.Println(scores[(len(scores)-1)/2])

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
