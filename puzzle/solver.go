package puzzle

import (
	"fmt"
	"log"
	"strconv"
)

type Solver struct {
	columns []*CellGroup
	rows    []*CellGroup
	blocks  []*CellGroup
	board   [][]*Cell
}

func (s Solver) InitSolver(values []string) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			value, err := strconv.Atoi(values[9*r+c])
			if err != nil {
				log.Fatal(err)
			}
			blockIndex := (c / 3) + (r/3)*3
			blockPos := (c % 3) + (r%3)*3
			predefined := false
			if value != 0 {
				predefined = true
			}
			cell := NewCell(value, r, c, blockIndex, predefined)
			s.columns[c].Add(cell, r)
			s.rows[r].Add(cell, c)
			s.board[r][c] = cell

			s.blocks[blockIndex].Add(cell, blockPos)
		}
	}
}

func (s Solver) GetCandidates(c *Cell) (bool, *Cell) {
	row := s.rows[c.row]
	col := s.columns[c.col]
	block := s.blocks[c.block]
	changed := false
	if c.preDefined {
		return false, c
	}
	for i := 0; i < 9; i++ {
		if row.cells[i].value != 0 {
			//changed = true
			c.RemoveCandidate(row.cells[i].value)
		}
		if col.cells[i].value != 0 {
			//changed = true
			c.RemoveCandidate(col.cells[i].value)
		}
		if block.cells[i].value != 0 {
			//changed = true
			c.RemoveCandidate(block.cells[i].value)
		}
	}
	candidates := c.candidate
	hiddenSingle := -1
	for _, v := range candidates {
		candidateFound := false
		for i := 0; i < 9; i++ {

			candidateFound = row.cells[i].CheckForCandidate(v)
			candidateFound = col.cells[i].CheckForCandidate(v)
			candidateFound = block.cells[i].CheckForCandidate(v)
			if candidateFound {
				break
			}
		}
		if candidateFound {
			hiddenSingle = v
			break
		}
	}
	changed = c.UpdateValue()
	if hiddenSingle != -1 {
		//c.SetValue(hiddenSingle)
		c.value = hiddenSingle
		changed = true
	}
	return changed, c
}

func (s Solver) PrintBoard() {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			fmt.Printf("%d ", s.board[r][c].value)
		}
		fmt.Println()
	}
}

func (s Solver) Solve() {
	for {
		changed := false
		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				cell := s.board[r][c]
				changed, cell = s.GetCandidates(cell)
			}
		}
		fmt.Println("======================================")
		s.PrintBoard()
		if !changed {
			break
		}
	}
}

func NewSolver(values []string) *Solver {
	columns := make([]*CellGroup, 9)
	rows := make([]*CellGroup, 9)
	blocks := make([]*CellGroup, 9)
	board := make([][]*Cell, 9)
	for i := 0; i < 9; i++ {
		board[i] = make([]*Cell, 9)
	}
	for i := 0; i < 9; i++ {
		columns[i] = NewCellGroup()
		rows[i] = NewCellGroup()
		blocks[i] = NewCellGroup()
	}
	solver := Solver{columns: columns, rows: rows, blocks: blocks, board: board}
	solver.InitSolver(values)

	return &solver
}
