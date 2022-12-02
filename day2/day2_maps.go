package main

import (
//	"fmt"
	"bufio"
	"log"
	"strings"
	"os"
)

// A Rock 1
// B Paper 2 
// C Scissors 3
// 0 if lost 6 if won


// Paper > Rock
// Scissors > Paper
// Rock > Scissors

func decideMyMoveImproved(opponent string, me string) string {
	winningCharMap := map[string]string{"A": "Y", "B": "Z", "C": "X"} 
	drawingCharMap := map[string]string{"A": "X", "B": "Y", "C": "Z"}
	losingCharMap := map[string]string{"A": "Z", "B": "X", "C": "Y"}
	decisionMap := map[string]string{"X": losingCharMap[opponent], "Y": drawingCharMap[opponent], "Z": winningCharMap[opponent]}
	
	return decisionMap[me]
}

func getMyScoreImproved(opponent string, me string) int {
	scoreMap := map[string]int{"X": 1, "Y": 2, "Z": 3}
	winningCharMap := map[string]string{"A": "Y", "B": "Z", "C": "X"} 
	drawingCharMap := map[string]string{"A": "X", "B": "Y", "C": "Z"}
	winningScore := 6
	drawingScore := 3

	if winningCharMap[opponent] == me { 
		return scoreMap[me] + winningScore
	} else if drawingCharMap[opponent] == me {
		return scoreMap[me] + drawingScore
	} else {
		return scoreMap[me]
	}	

	panic("Could not match character to known strategy getMyScore")
}

func playGameAndReturnScoreImproved(scanner *bufio.Scanner, finalScore int) int {
	scanner.Scan()
	nextLine := strings.TrimSpace(scanner.Text())
	if len(nextLine) == 0 {
		return finalScore
	}
	strategies := strings.Split(nextLine, " ")
	// PART1 return playGameAndReturnScore(scanner, finalScore+getMyScore(strategies[0], strategies[1]))
        return playGameAndReturnScoreImproved(scanner, finalScore+getMyScoreImproved(strategies[0], decideMyMoveImproved(strategies[0], strategies[1])))
}

func main_improved() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	playGameAndReturnScoreImproved(scanner, 0)

//	finalScore := playGameAndReturnScoreImproved(scanner, 0)
//	log.Println(fmt.Sprintf("%s%d%s", "Your final score is: ", finalScore, "."))

}
