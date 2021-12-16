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

type job struct {
	versionSum int
}

func main() {
	start := time.Now()

	file, err := os.Open("day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input string

	for scanner.Scan() {
		if scanner.Text() != "" {
			input = scanner.Text()
		}
	}

	binary := toBinary(input)

	packet := strings.Split(binary, "")

	j := job{}
	value, _ := j.processPacket(packet)

	fmt.Println(fmt.Sprintf("Part 1: %d", j.versionSum))
	fmt.Println(fmt.Sprintf("Part 2: %d", value))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func (j *job) processPacket(packet []string) (int, []string) {
	version := binaryToInt(packet[:3])
	packet = packet[3:]
	j.versionSum += version

	packetType := binaryToInt(packet[:3])
	packet = packet[3:]

	if packetType == 4 {
		sections := []string{}
		running := true
		for ok := true; ok; ok = running {
			section := strings.Join(packet[:5], "")
			packet = packet[5:]
			sections = append(sections, section[1:])
			if section[:1] == "0" {
				running = false
			}
		}
		return binaryToInt(sections), packet
	}

	lengthTypeId := packet[:1]
	packet = packet[1:]
	subValues := []int{}
	if lengthTypeId[0] == "0" {
		length := binaryToInt(packet[:15])
		packet = packet[15:]
		subPacket := packet[:length]
		packet = packet[length:]

		running := true
		for ok := true; ok; ok = running {
			var subValue int
			subValue, subPacket = j.processPacket(subPacket)
			subValues = append(subValues, subValue)
			if len(subPacket) == 0 {
				running = false
			}
		}
	} else {
		length := binaryToInt(packet[:11])
		packet = packet[11:]

		for i := 0; i < length; i++ {
			var subValue int
			subValue, packet = j.processPacket(packet)
			subValues = append(subValues, subValue)
		}
	}

	value := 0
	switch packetType {
	case 0:
		value = sum(subValues)
	case 1:
		value = product(subValues)
	case 2:
		value = minimum(subValues)
	case 3:
		value = maximum(subValues)
	case 5:
		value = greaterThan(subValues)
	case 6:
		value = lessThan(subValues)
	case 7:
		value = equalTo(subValues)
	}
	return value, packet
}

func equalTo(values []int) int {
	if values[0] == values[1] {
		return 1
	}
	return 0
}

func lessThan(values []int) int {
	if values[0] < values[1] {
		return 1
	}
	return 0
}

func greaterThan(values []int) int {
	if values[0] > values[1] {
		return 1
	}
	return 0
}

func maximum(values []int) int {
	max := -1
	for _, value := range values {
		if max < value {
			max = value
		}
	}
	return max
}

func minimum(values []int) int {
	min := -1
	for _, value := range values {
		if min == -1 || min > value {
			min = value
		}
	}
	return min
}

func product(values []int) int {
	total := 1
	for _, value := range values {
		total *= value
	}
	return total
}

func sum(values []int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func binaryToInt(packet []string) int {
	d, _ := strconv.ParseInt(strings.Join(packet, ""), 2, 64)
	return int(d)
}

func toBinary(input string) string {
	output := ""
	character := strings.Split(input, "")
	for _, char := range character {
		switch char {
		case "0":
			output += "0000"
		case "1":
			output += "0001"
		case "2":
			output += "0010"
		case "3":
			output += "0011"
		case "4":
			output += "0100"
		case "5":
			output += "0101"
		case "6":
			output += "0110"
		case "7":
			output += "0111"
		case "8":
			output += "1000"
		case "9":
			output += "1001"
		case "A":
			output += "1010"
		case "B":
			output += "1011"
		case "C":
			output += "1100"
		case "D":
			output += "1101"
		case "E":
			output += "1110"
		case "F":
			output += "1111"
		}
	}
	return output
}
