package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"log"
)

// Half of a line is one compartment of a rucksack. Equal in size.
// A and a represent a different item. 
// Match alphabet to characters.
// Use ASCII values and match to priority value.
// Decimal A = 65
// Decimal Z = 90
// Decimal a = 97
// Decimal z = 122
// Find common character between two comparments

// Credit for this function goes to https://github.com/juliangruber/go-intersect/blob/master
func getCommonCharacter[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}
	return set
}


func convertFromASCIIToPriority(value int) int {
	if value < 97 { 
		return value - 38 // 65 decimal val of A - 27 priority of A = 38
	} else {
		return value - 96 // 97 decimal val of a - 1 priority of a = 96 
	}
}

func getPoints(scanner *bufio.Scanner, points int) int {
	scanner.Scan()
	nextLine := strings.TrimSpace(scanner.Text())
	if len(nextLine) == 0 {
		return points
	}

	firstCompartment := strings.Split(nextLine[:len(nextLine)/2], "")
	secondCompartment := strings.Split(nextLine[len(nextLine)/2:], "")

	sharedCharacterAsDecimal := int([]rune(getCommonCharacter[string](firstCompartment, secondCompartment)[0])[0])

	return getPoints(scanner, points+convertFromASCIIToPriority(sharedCharacterAsDecimal))
}

func getPointsPart2(scanner *bufio.Scanner, points int) int {
	scanner.Scan()
	lineOne := strings.TrimSpace(scanner.Text())
	if len(lineOne) == 0 {
		return points
	}
	scanner.Scan()
	lineTwo := strings.TrimSpace(scanner.Text())
	scanner.Scan()
	lineThree := strings.TrimSpace(scanner.Text())

	nextLines := []string{lineOne, lineTwo, lineThree}
	commonBetweenFirstTwoLines := getCommonCharacter[string](strings.Split(nextLines[0], ""), strings.Split(nextLines[1], ""))

	commonBetweenFirstTwoAndThird := getCommonCharacter[string](commonBetweenFirstTwoLines, strings.Split(nextLines[2], ""))[0]
	sharedCharacterAsDecimal := int([]rune(commonBetweenFirstTwoAndThird)[0])

	return getPointsPart2(scanner, points+convertFromASCIIToPriority(sharedCharacterAsDecimal))

}

func part2() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	result := getPointsPart2(bufio.NewScanner(file), 0)
	log.Println(fmt.Sprintf("%s%d%s", "Sum of priorites: ", result, "."))
}

func part1() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	result := getPoints(bufio.NewScanner(file), 0)
	log.Println(fmt.Sprintf("%s%d%s", "Sum of priorites: ", result, "."))
}


func main() {
	part2()
}
