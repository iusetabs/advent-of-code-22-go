package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
	"strings"
)

func getLowerAndUpper(elfString string) (int, int) {
	elfSplit := strings.Split(elfString, "-")
	a, _ := strconv.Atoi(elfSplit[0])
	b, _ := strconv.Atoi(elfSplit[1])
	return a, b
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	p1 := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		ranges := strings.Split(line, ",")
		elfA, elfB := ranges[0], ranges[1]
		aLower, aUpper := getLowerAndUpper(elfA)
		bLower, bUpper := getLowerAndUpper(elfB)
		if aLower >= bLower && aUpper <= bUpper {
			p1+=1
		} else if bLower >= aLower && bUpper <= aUpper  {
			p1+=1
		}
	}

	log.Println(fmt.Sprintf("%s%d%s", "Result P1: ", p1, "."))
}
