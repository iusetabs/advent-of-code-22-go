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

func readLeftToRight(a int, b int) bool {
	return a < b
}

func counterLeftToRight(i int) int {
	return i + 1
}

func readRightToLeft(a int, b int) bool {
	return a >= b
}

func counterRightToLeft(i int) int {
	return i - 1
}

func isTreeVisibleInRow(trees []string, v int, i int, indx int, forCondition func(int, int) bool, counterChange func(int) int) (bool, int) {
	visible := true
	numOfVisibleTrees := 0
	visibleLoop:
	for forCondition(i, indx) {
		val, _ := strconv.Atoi(trees[i])
		numOfVisibleTrees+=1
		if val >= v {
			visible = false
			break visibleLoop
		}
		i = counterChange(i)
	}
	return visible, numOfVisibleTrees
}

func isTreeVisibeInColumn(fileName string, v int, currentLineNum int, lineNum int, indx int, counterChange func(int) int) (bool, int) {
	visible := true
	numOfVisibleTrees := 0
	visibleLoop:
	for {
		currentLineNum = counterChange(currentLineNum)
		line := strings.Split(lineAt(fileName, currentLineNum), "")
		if len(line) == 0  {
                        break visibleLoop
                }
		numOfVisibleTrees+=1
		val, _ := strconv.Atoi(line[indx])
		if val >= v {
			visible = false 
			break visibleLoop

		} else if lineNum == currentLineNum {
			break visibleLoop
		}
	}
	return visible, numOfVisibleTrees
}

func main() {
	fileName := "./input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lineLength := len(strings.Split(lineAt(fileName, 0), ""))
	p1, p2, lineNumber := 0, 0, 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		treeSizes := strings.Split(line, "")
		for i, v := range treeSizes {
			vAsInt, _ := strconv.Atoi(v)
			isVisibleFromLeft, visibleTreesOnLeft := isTreeVisibleInRow(treeSizes, vAsInt, i-1, 0, readRightToLeft, counterRightToLeft)
			isVisibleFromRight, visibleTreesOnRight := isTreeVisibleInRow(treeSizes, vAsInt, i+1, lineLength, readLeftToRight, counterLeftToRight)
			isVisibleFromTop, visibleTreesOnTop := isTreeVisibeInColumn(fileName, vAsInt, lineNumber, 0, i, counterRightToLeft)
			isVisibleFromBottom, visibleTreesOnBottom :=  isTreeVisibeInColumn(fileName, vAsInt, lineNumber, lineNumber, i, counterLeftToRight)
			if isVisibleFromLeft  || isVisibleFromRight || isVisibleFromTop || isVisibleFromBottom {
					p1 += 1
			}
			scenicValue := visibleTreesOnLeft * visibleTreesOnRight * visibleTreesOnTop * visibleTreesOnBottom
			if scenicValue > p2 {
				p2 = scenicValue
			}
		}
		lineNumber+=1
	}
	fmt.Println(fmt.Sprintf("%s%d", "P1: ", p1))
	fmt.Println(fmt.Sprintf("%s%d", "P2: ", p2))
}
