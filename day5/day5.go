package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func takeN[T comparable](a []T, N int, doReverse bool) ([]T, []T) {
	b := make([]T, N)

	for i := 0; i < N; i++ {
		if doReverse { // Part 1
			b[N-i-1] = a[i]
		} else { // Part 2
			b[i] = a[i]
		}
	}
	return a[N:], b

}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	hash := make(map[int][]string)
	hashP2 := make(map[int][]string)
	p1, p2 := []string{}, []string{}

	// Parsing of the initial table.
	parseStacks:
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break parseStacks
		} else if strings.Contains(line, "[") {
			for i := 1; i <= len(line); i+=4 {
				indx := max((i/4)+1, 1)
				char := line[i]
				if string(char) != " " {
					hash[indx] = append(hash[indx], string(char))	
					hashP2[indx] = append(hashP2[indx], string(char))	
				}
			}
		}
	}

	// Actual changing of the crates
	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), " ")
		N, _ := strconv.Atoi(lineSplit[1])

		from, _ := strconv.Atoi(lineSplit[3])
		to, _ := strconv.Atoi(lineSplit[5])

		// P1
		updatedFrom, movedList := takeN[string](hash[from], N, true)
		hash[from] = updatedFrom
		hash[to] = append(movedList, hash[to]...)

		// P2
		updatedFrom, movedList = takeN[string](hashP2[from], N, false)
		hashP2[from] = updatedFrom
		hashP2[to] = append(movedList, hashP2[to]...)
	}

	// Final state of crates. Note the first element of each list contained within the hashmap corresponds to the last element.
	log.Println(fmt.Sprintf("%s%v", "p1: ", hash))	
	log.Println(fmt.Sprintf("%s%v", "p2: ", hashP2))	
	
	// Creating the message
	for i := 1; i <= len(hash); i++ {
		if len(hash[i]) != 0 {
			p1 = append(p1, hash[i][0])
			p2 = append(p2, hashP2[i][0])
		}
	}

	// Printing the message
	fmt.Println(fmt.Sprintf("%s%v", "P1: ", strings.Join(p1,"")))
	fmt.Println(fmt.Sprintf("%s%v", "P2: ", strings.Join(p2,"")))
	
}
