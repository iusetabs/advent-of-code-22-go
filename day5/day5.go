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
	log.Println(fmt.Sprintf("%s%v%s%d", "a: ", a, " N: ", N))	
	b := make([]T, N)

	for i := 0; i < N; i++ {
		log.Println(fmt.Sprintf("%s%v", "b[i]", b[i]))	
		log.Println(fmt.Sprintf("%s%v", "a[len]", a[len(a)-1-i]))	
		if doReverse {
			b[N-i-1] = a[i]
		} else {
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


	for scanner.Scan() {
		log.Println(fmt.Sprintf("%s%v", "hashP2: ", hashP2))	

		lineSplit := strings.Split(scanner.Text(), " ")
		N, _ := strconv.Atoi(lineSplit[1])
		log.Println(fmt.Sprintf("%v", lineSplit))	

		from, _ := strconv.Atoi(lineSplit[3])
		to, _ := strconv.Atoi(lineSplit[5])
		log.Println(fmt.Sprintf("%s%d%s%d%s%d", "\nN: ", N, " from: ", from, " to: ", to))	

		updatedFrom, movedList := takeN[string](hash[from], N, true)
		log.Println(fmt.Sprintf("%s%v%s%v", "P1 updatedFrom: ", updatedFrom, " movedList: ", movedList))	
		hash[from] = updatedFrom
		hash[to] = append(movedList, hash[to]...)

		updatedFrom, movedList = takeN[string](hashP2[from], N, false)
		log.Println(fmt.Sprintf("%s%v%s%v", "P2 updatedFrom: ", updatedFrom, " movedList: ", movedList))	
		hashP2[from] = updatedFrom
		hashP2[to] = append(movedList, hashP2[to]...)
	}

	log.Println(fmt.Sprintf("%s%v", "p1: ", hash))	
	log.Println(fmt.Sprintf("%s%v", "p2: ", hashP2))	
	
	for i := 1; i <= len(hash); i++ {
		if len(hash[i]) != 0 {
			p1 = append(p1, hash[i][0])
			p2 = append(p2, hashP2[i][0])
		}
	}

	fmt.Println(fmt.Sprintf("%s%v", "P1: ", strings.Join(p1,"")))
	fmt.Println(fmt.Sprintf("%s%v", "P2: ", strings.Join(p2,"")))
	
}
