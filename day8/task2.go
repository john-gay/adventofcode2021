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

	file, err := os.Open("day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		if scanner.Text() != "" {
			numStrList := []string{}
			p1 := strings.Split(scanner.Text(), " | ")

			valueMap := createMap(p1[0])

			p2 := strings.Split(p1[1], " ")
			for _, part := range p2 {
				partArr := strings.Split(part, "")
				sort.Strings(partArr)
				p := strings.Join(partArr, "")
				numStrList = append(numStrList, valueMap[p])
			}
			numStr := strings.Join(numStrList, "")
			v, _ := strconv.ParseInt(numStr, 10, 64)
			count += int(v)
		}
	}
	log.Println(fmt.Sprintf("Count: %d", count))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func createMap(input string) map[string]string {
	valueMap := map[string]string{}
	inverseMap := map[int]string{}

	parts := strings.Split(input, " ")

	var unknown5 []string
	var unknown6 []string

	for _, part := range parts {
		partArr := strings.Split(part, "")
		sort.Strings(partArr)
		sortedPart := strings.Join(partArr, "")

		switch len(part) {
		case 2:
			valueMap[sortedPart] = "1"
			inverseMap[1] = sortedPart
		case 3:
			valueMap[sortedPart] = "7"
			inverseMap[7] = sortedPart
		case 4:
			valueMap[sortedPart] = "4"
			inverseMap[4] = sortedPart
		case 7:
			valueMap[sortedPart] = "8"
			inverseMap[8] = sortedPart
		case 5:
			unknown5 = append(unknown5, sortedPart)
		case 6:
			unknown6 = append(unknown6, sortedPart)
		}
	}

	// Find 3
	newUnknown5 := []string{}
	for _, unknown := range unknown5 {
		matches := 0
		matched := false
		for _, c := range unknown {
			if strings.Contains(inverseMap[1], string(c)) {
				matches++
			}
			if matches == 2 {
				valueMap[unknown] = "3"
				inverseMap[3] = unknown
				matched = true
				break
			}
		}
		if !matched {
			newUnknown5 = append(newUnknown5, unknown)
		}
	}
	unknown5 = newUnknown5

	// Find 5
	newUnknown5 = []string{}
	for _, unknown := range unknown5 {
		matches := 0
		matched := false
		for _, c := range unknown {
			if strings.Contains(inverseMap[4], string(c)) {
				matches++
			}
			if matches == 3 {
				valueMap[unknown] = "5"
				inverseMap[5] = unknown
				matched = true
				break
			}
		}
		if !matched {
			newUnknown5 = append(newUnknown5, unknown)
		}
	}
	unknown5 = newUnknown5

	// Find 2
	valueMap[unknown5[0]] = "2"
	inverseMap[2] = unknown5[0]

	// Find 6
	newUnknown6 := []string{}
	for _, unknown := range unknown6 {
		matches := 0
		matched := false
		for _, c := range unknown {
			if strings.Contains(inverseMap[1], string(c)) {
				matches++
			}
		}
		if matches == 1 {
			valueMap[unknown] = "6"
			inverseMap[6] = unknown
			matched = true
		}
		if !matched {
			newUnknown6 = append(newUnknown6, unknown)
		}
	}
	unknown6 = newUnknown6

	// Find 9
	newUnknown6 = []string{}
	for _, unknown := range unknown6 {
		matches := 0
		matched := false
		for _, c := range unknown {
			if strings.Contains(inverseMap[4], string(c)) {
				matches++
			}
		}
		if matches == 4 {
			valueMap[unknown] = "9"
			inverseMap[9] = unknown
			matched = true
		}
		if !matched {
			newUnknown6 = append(newUnknown6, unknown)
		}
	}
	unknown6 = newUnknown6

	// Find 0
	valueMap[unknown6[0]] = "0"
	inverseMap[0] = unknown6[0]

	return valueMap
}
