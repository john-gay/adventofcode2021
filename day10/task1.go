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

	file, err := os.Open("day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var opens []string
	sum := 0

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
						sum += 3
						valid = false
					}
				case "]":
					if opens[len(opens)-1] == "[" {
						opens = opens[:len(opens)-1]
					} else {
						sum += 57
						valid = false
					}
				case "}":
					if opens[len(opens)-1] == "{" {
						opens = opens[:len(opens)-1]
					} else {
						sum += 1197
						valid = false
					}
				case ">":
					if opens[len(opens)-1] == "<" {
						opens = opens[:len(opens)-1]
					} else {
						sum += 25137
						valid = false
					}
				default:
					opens = append(opens, part)
				}
			}
		}
	}
	fmt.Println(sum)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
