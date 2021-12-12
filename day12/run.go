package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type job struct {
	nodes         map[string][]string
	numberOfPaths int
	twice         bool
}

func main() {
	start := time.Now()

	file, err := os.Open("day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nodes := map[string][]string{}

	for scanner.Scan() {
		if scanner.Text() != "" {
			parts := strings.Split(scanner.Text(), "-")

			if node, ok := nodes[parts[0]]; ok {
				nodes[parts[0]] = append(node, parts[1])
			} else {
				nodes[parts[0]] = []string{parts[1]}
			}

			if node, ok := nodes[parts[1]]; ok {
				nodes[parts[1]] = append(node, parts[0])
			} else {
				nodes[parts[1]] = []string{parts[0]}
			}
		}
	}

	job := job{
		nodes:         nodes,
		numberOfPaths: 0,
	}

	job.findPaths("start", []string{})

	fmt.Println(fmt.Sprintf("Part1: %d", job.numberOfPaths))

	job.numberOfPaths = 0
	job.twice = true

	job.findPaths("start", []string{})

	fmt.Println(fmt.Sprintf("Part2: %d", job.numberOfPaths))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func (j *job) findPaths(node string, path []string) {
	if node == "end" {
		j.numberOfPaths++
		return
	}
	if len(path) != 0 && node == "start" {
		return
	}
	if previouslyVisited(node, path, j.twice) {
		return
	}
	path = append(path, node)
	for _, n := range j.nodes[node] {
		j.findPaths(n, path)
	}
}

func previouslyVisited(node string, path []string, twice bool) bool {
	if strings.ToUpper(node) == node {
		return false
	}
	if twice {
		visited := map[string]bool{}
		twiceVisited := false
		if node != "start" && node != "end" {
			for _, p := range path {
				if strings.ToUpper(p) != p {
					if visited[p] {
						twiceVisited = true
					}
					visited[p] = true
				}
			}
		}
		if twiceVisited {
			for _, p := range path {
				if p == node {
					return true
				}
			}
		}
		return false
	} else {
		for _, p := range path {
			if p == node {
				return true
			}
		}
		return false
	}
}
