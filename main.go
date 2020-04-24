package main

import (
	"github.com/whcass/sudoku-solver/puzzle"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fileContent, err := ioutil.ReadFile("puzzles/sudoku_puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	values := strings.Fields(string(fileContent))
	solver := puzzle.NewSolver(values)
	solver.Solve()
}
