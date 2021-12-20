package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("day20/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	enhancementAlg := []string{}

	image := [][]string{}

	for scanner.Scan() {
		if scanner.Text() != "" {
			if len(enhancementAlg) == 0 {
				enhancementAlg = strings.Split(scanner.Text(), "")
			} else {
				image = append(image, strings.Split(scanner.Text(), ""))
			}
		}
	}

	image = expandImageOnce(image)

	turns := 50
	for turn := 1; turn <= turns; turn++ {
		image = enhanceImage(image, enhancementAlg)

		if turn == 2 {
			fmt.Println(fmt.Sprintf("Part 1: %d", countLights(image)))
		}
	}

	fmt.Println(fmt.Sprintf("Part 2: %d", countLights(image)))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func expandImageOnce(image [][]string) [][]string {
	v := []string{}
	for i := 0; i < 350; i++ {
		v = append(v, ".")
	}

	for i := range image {
		image[i] = append(v, image[i]...)
		image[i] = append(image[i], v...)
	}

	v = []string{}
	for range image[0] {
		v = append(v, ".")
	}

	for range image[0] {
		image = append([][]string{v}, image...)
		image = append(image, v)
	}

	return image
}

func countLights(image [][]string) int {
	count := 0
	removeEnds := 50
	for i := range image {
		if i > removeEnds && i < len(image)-removeEnds+1 {
			for j := range image[i] {
				if j > removeEnds && j < len(image[i])-removeEnds+1 {
					if image[i][j] == "#" {
						count++
					}
				}
			}
		}
	}
	return count
}

func enhanceImage(image [][]string, alg []string) [][]string {
	newImage := make([][]string, len(image))
	for i := range newImage {
		row := []string{}
		for range image[i] {
			row = append(row, ".")
		}
		newImage[i] = row
	}
	for i := range image {
		if i > 0 && i < len(image)-1 {
			for j := range image[i] {
				if j > 0 && j < len(image[i])-1 {
					encoded := []string{}
					for _, ii := range []int{i - 1, i, i + 1} {
						for _, jj := range []int{j - 1, j, j + 1} {
							encoded = append(encoded, image[ii][jj])
						}
						bEncoded := []string{}
						for _, char := range encoded {
							if char == "#" {
								bEncoded = append(bEncoded, "1")
							} else {
								bEncoded = append(bEncoded, "0")
							}
						}
						algIndex, _ := strconv.ParseInt(strings.Join(bEncoded, ""), 2, 64)
						newImage[i][j] = alg[algIndex]
					}
				}
			}
		}
	}
	return newImage
}

func expandImage(image [][]string, turn int) [][]string {
	image = expandHorizontal(true, image, turn)
	image = expandHorizontal(false, image, turn)
	image = expandVertical(true, image, turn)
	image = expandVertical(false, image, turn)

	return image
}

func expandHorizontal(left bool, image [][]string, turn int) [][]string {
	v := "."
	if turn%2 == 0 {
		v = "#"
	}

	requireExpand := false
	indexRange := []int{0}
	if !left {
		indexRange = []int{len(image[0]) - 1}
	}

	for _, index := range indexRange {
		for _, row := range image {
			if row[index] == "#" {
				requireExpand = true
				break
			}
		}
	}

	if requireExpand {
		for i := range image {
			for range indexRange {
				if left {
					image[i] = append([]string{v}, image[i]...)
				} else {
					image[i] = append(image[i], v)
				}
			}
		}
	}
	return image
}

func expandVertical(bot bool, image [][]string, turn int) [][]string {
	v := "."
	if turn%2 == 0 {
		v = "#"
	}

	requireExpand := false
	indexRange := []int{0}
	if !bot {
		indexRange = []int{len(image) - 1}
	}

	for _, index := range indexRange {
		for _, cell := range image[index] {
			if cell == "#" {
				requireExpand = true
				break
			}
		}
	}
	if requireExpand {
		newRow := []string{}
		for range image[0] {
			newRow = append(newRow, v)
		}
		for range indexRange {
			if bot {
				image = append([][]string{newRow}, image...)
			} else {
				image = append(image, newRow)
			}
		}
	}
	return image
}

func printImage(image [][]string) {
	fmt.Println("")
	for _, row := range image {
		fmt.Println(row)
	}
}
