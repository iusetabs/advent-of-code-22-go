package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"log"
	"strconv"
)

func lineAt(fileName string, indx int) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		if indx == i {
			return strings.TrimSpace(scanner.Text())
		}
		i++
	}
	return ""
}

func isTreeVisibleInRow(trees []string, v int, i int, indx int) bool {
	visible := true
	visibleLoop:
	for i < indx {
		val, _ := strconv.Atoi(trees[i])
		if val >= v {
			visible = false
			break visibleLoop
		}
		i+=1
	}
	return visible
}

func isTreeVisibeInColumn(fileName string, v int, currentLineNum int, lineNum int, indx int) bool {
	visible := true
	visibleLoop:
	for {
		line := strings.Split(lineAt(fileName, currentLineNum), "")
		if len(line) == 0 || lineNum == currentLineNum {
			break visibleLoop
		} 
		val, _ := strconv.Atoi(line[indx])
		if val >= v {
			visible = false 
			break visibleLoop

		}
		currentLineNum+=1	
	}
	return visible
}

func main() {
	fileName := "./input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Assuming length is the same, the trees on the first and last line can be seen.
	lineLength := len(strings.Split(lineAt(fileName, 0), ""))
	p1 := 0

	lineNumber := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		treeSizes := strings.Split(line, "")
		maxLeft := 0
		for i, v := range treeSizes {
			vAsInt, _ := strconv.Atoi(v)
			// If already on the edge then this tree can be seen
			if i == 0 || i == len(treeSizes)-1 {
				p1 += 1
			} else if vAsInt > maxLeft && isTreeVisibleInRow(treeSizes, vAsInt, 0, i) {
					p1 += 1
					maxLeft = vAsInt
			} else if isTreeVisibleInRow(treeSizes, vAsInt, i+1, lineLength) {
					p1 += 1
					maxLeft = vAsInt
			} else if isTreeVisibeInColumn(fileName, vAsInt, 0, lineNumber, i) {
					p1 += 1
			} else if isTreeVisibeInColumn(fileName, vAsInt, lineNumber+1, lineNumber, i) {
					p1 += 1
			}
		}
		lineNumber+=1
	}
	fmt.Println(fmt.Sprintf("%s%d", "P1: ", p1))
	
}
