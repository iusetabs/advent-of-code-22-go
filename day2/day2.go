package main

import (
	"fmt"
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

func decideMyMove(opponent string, me string) string {
	if me == "X" { // lose 
		if opponent == "C" {
			return "Y"
		} else if opponent == "B" {
			return "X"
		} else if opponent == "A" {
			return "Z"
		}
	} else if me == "Y" { // draw 
		if opponent == "A" {
			return "X"
		} else if opponent == "B" {
			return "Y"
		} else if opponent == "C" {
			return "Z"
		}
	} else if me == "Z" { // win
		if opponent == "B" {
			return "Z"
		} else if opponent == "A" {
			return "Y"
		} else if opponent == "C" {
			return "X"
		}
	}
	panic("Could not match character to known strategy decideMyMove")

}

func getMyScore(opponent string, me string) int {
	//log.Println(fmt.Sprintf("%s%s%s%s", "Opponent: ", opponent, ". Me: ", me))

	if me == "X" {
		if opponent == "C" {
			return 1 + 6
		} else if opponent == "B" {
			return 1 + 0
		} else if opponent == "A" {
			return 1 + 3
		}
	} else if me == "Y" {
		if opponent == "A" {
			return 2 + 6
		} else if opponent == "C" {
			return 2 + 0
		} else if opponent == "B" {
			return 2 + 3
		}
	} else if me == "Z" {
		if opponent == "B" {
			return 3 + 6
		} else if opponent == "A" {
			return 3 + 0
		} else if opponent == "C" {
			return 3 + 3
		}
	} 
	panic("Could not match character to known strategy getMyScore")
}

func playGameAndReturnScore(scanner *bufio.Scanner, finalScore int) int {
	scanner.Scan()
	nextLine := strings.TrimSpace(scanner.Text())
	if len(nextLine) == 0 {
		return finalScore
	}
	strategies := strings.Split(nextLine, " ")
// part 1	return playGameAndReturnScore(scanner, finalScore+getMyScore(strategies[0], strategies[1]))
        return playGameAndReturnScore(scanner, finalScore+getMyScore(strategies[0], decideMyMove(strategies[0], strategies[1])))
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	finalScore := playGameAndReturnScore(scanner, 0)
	log.Println(fmt.Sprintf("%s%d%s", "Your final score is: ", finalScore, "."))

}
