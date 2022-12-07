package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"bufio"
	"log"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sizes := make(map[string]int)
	path := []string{}
	p1 := 0
	p2 := 30000000

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.Contains(line, "$ cd") {
			name := strings.Split(line, " ")[2]
			if name == ".." {
				path = path[:len(path)-1]
			} else if name == "/" {
				path = append(path, name)
			} else {
				path = append(path, name + "/")
			}

		} else if !(strings.Contains(line, "$") || strings.Contains(line, "dir ")) {
			lineSplit := strings.Split(line, " ")
			size, _ := strconv.Atoi(lineSplit[0])
			for i := 1; i <= len(path); i++ {
				key := strings.Join(path[:i], "")
				sizes[key] = sizes[key] + size
			}
		}
	}
	
	memoryNeeded := 30000000 - (70000000 - sizes["/"]) // Needed memory - currently unused memory
	for _, v := range sizes {
		if v <= 100000 {
			p1 += v
		} else if v >= memoryNeeded && v <= p2 {
			p2 = v
		}
	}

	fmt.Println(fmt.Sprintf("%s%d", "P1: ", p1))
	fmt.Println(fmt.Sprintf("%s%d", "P2: ", p2))
}
