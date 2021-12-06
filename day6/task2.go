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

	fishMap := map[int]int{}
	for _, f := range fish {
		if _, ok := fishMap[f]; ok {
			fishMap[f] = fishMap[f] + 1
		} else {
			fishMap[f] = 1
		}
	}

	fmt.Println(fmt.Sprintf("map: %v", fishMap))

	newFishMap := map[int]int{}
	for i := 0; i < 256; i++ {
		newFishMap = map[int]int{}
		for key, value := range fishMap {
			if key == 0 {
				newFishMap[6] = newFishMap[6] + value
				newFishMap[8] = newFishMap[8] + value
			} else {
				newFishMap[key-1] = newFishMap[key-1] + value
			}
		}
		fishMap = newFishMap
	}

	fmt.Println(fmt.Sprintf("map: %v", fishMap))

	count := 0
	for _, value := range fishMap {
		count = count + value
	}
	fmt.Println(fmt.Sprintf("Number of fish: %d", count))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
