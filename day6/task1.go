package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fish []int

	for scanner.Scan() {
		if scanner.Text() != "" {
			for _, c := range scanner.Text() {
				if string(c) != "," {
					d, _ := strconv.ParseInt(string(c), 10, 32)
					fish = append(fish, int(d))
				}
			}
		}
	}

	fmt.Println(fmt.Sprintf("input: %v", fish))

	var newFish []int
	for i := 0; i < 80; i++ {
		newFish = []int{}
		for _, f := range fish {
			if f == 0 {
				newFish = append(newFish, 6, 8)
			} else {
				newFish = append(newFish, f-1)
			}
		}
		fish = newFish
	}

	fmt.Println(fmt.Sprintf("Number of fish: %d", len(fish)))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
