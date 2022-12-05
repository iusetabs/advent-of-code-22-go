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

func takeN[T comparable](a []T, N int) ([]T, []T) {
	log.Println(fmt.Sprintf("%s%v%s%d", "a: ", a, " N: ", N))	
	b := make([]T, N)

	for i := 0; i < N; i++ {
		log.Println(fmt.Sprintf("%s%v", "b[i]", b[i]))	
		log.Println(fmt.Sprintf("%s%v", "a[len]", a[len(a)-1-i]))	
		b[N-i-1] = a[i]
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
	p1 := []string{}

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
				}
			}
		}
	}

	for scanner.Scan() {
		log.Println(fmt.Sprintf("%s%v", "hash: ", hash))	

		lineSplit := strings.Split(scanner.Text(), " ")
		N, _ := strconv.Atoi(lineSplit[1])
		log.Println(fmt.Sprintf("%v", lineSplit))	

		from, _ := strconv.Atoi(lineSplit[3])
		to, _ := strconv.Atoi(lineSplit[5])
		log.Println(fmt.Sprintf("%s%d%s%d%s%d", "\nN: ", N, " from: ", from, " to: ", to))	

		updatedFrom, movedList := takeN[string](hash[from], N)
		log.Println(fmt.Sprintf("%s%v%s%v", "updatedFrom: ", updatedFrom, " movedList: ", movedList))	
		hash[from] = updatedFrom
		hash[to] = append(movedList, hash[to]...)
	}

	log.Println(fmt.Sprintf("%s%v", "p1: ", hash))	
	
	for i := 1; i <= len(hash); i++ {
		if len(hash[i]) != 0 {
			p1 = append(p1, hash[i][0])
		}
	}

	fmt.Println(fmt.Sprintf("%v", strings.Join(p1,"")))
	
}
