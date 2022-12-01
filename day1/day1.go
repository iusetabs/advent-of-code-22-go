package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
	"sort"
)

func scanUntilEmptyLine(scanner *bufio.Scanner, returnList []string) []string {
	scanner.Scan()
	nextLine := strings.TrimSpace(scanner.Text())
	if len(nextLine) == 0 {
		return returnList
	} 
	//log.Println(fmt.Sprintf("%s%s", "nextLine: ", nextLine))
	return scanUntilEmptyLine(scanner, append(returnList, nextLine))
}

func countCalories(lines []string, returnNum int) int {
	if len(lines) == 0 {
		return returnNum
	}
	nextCalorie, err := strconv.Atoi(lines[0])
	if err != nil  {
		panic(fmt.Sprintf("%s%v%s", "Could not convert non string: ", lines[0], " to int"))
	}

	return countCalories(lines[1:], returnNum+nextCalorie)
}

func getIndividualElfCalories(scanner *bufio.Scanner, returnList []int, i int) []int {
	lines := scanUntilEmptyLine(scanner, []string{})
	if len(lines) == 0 {
		return returnList
	} 
	log.Println(fmt.Sprintf("%s%d%s%v", "elf #", i, " :", lines))
	totalCalories := countCalories(lines, 0)
	return getIndividualElfCalories(scanner, append(returnList, totalCalories), i+1)
}

func max(array []int, currentMax int) int {
	if len(array) == 0 {
		return currentMax
	}
	updatedMax := currentMax
	if array[0] > currentMax {
		updatedMax = array[0]
	}

	return max(array[1:], updatedMax)
}


func reverseInts(input []int, returnList []int) []int {
	if len(input) == 0 {
		return returnList
	}

	return reverseInts(input[:len(input)-1], append(returnList, input[len(input)-1]))
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elfCalorieCounts := getIndividualElfCalories(scanner, []int{}, 1)
	log.Println(fmt.Sprintf("%s%d", "Max calorie: ", max((append([]int(nil), elfCalorieCounts...)), 0)))

	sort.Ints(elfCalorieCounts)

	reversedCalorieCounts := reverseInts(append([]int(nil), elfCalorieCounts...), []int{})
	countOfFirstThreeMaxes := reversedCalorieCounts[0] + reversedCalorieCounts[1] + reversedCalorieCounts[2] 
	log.Println(fmt.Sprintf("%s%d", "Count of first three maxes is: ", countOfFirstThreeMaxes)) 




}
