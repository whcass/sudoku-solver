package puzzle

import (
	"reflect"
	"strconv"
)

type Cell struct {
	value      int
	preDefined bool
	candidate  map[int]int
	row        int
	col        int
	block      int
}

func (c Cell) RemoveCandidate(value int) {
	delete(c.candidate, value)
}

func (c Cell) CheckForCandidate(value int) bool {
	for _, v := range c.candidate {
		if v == value {
			return true
		}
	}

	return false
}

func (c Cell) UpdateValue() bool {
	if len(c.candidate) == 1 {
		key := reflect.ValueOf(c.candidate).MapKeys()
		index, _ := strconv.Atoi(key[0].String())
		c.value = c.candidate[index]
		return true
	}

	return false
}

func (c Cell) SetValue(single int) {
	c.value = single
}

func NewCell(v int, r int, c int, b int, pd bool) *Cell {
	candidate := make(map[int]int)
	for i := 0; i < 9; i++ {
		candidate[1+i] = 1 + i
	}
	return &Cell{
		value:      v,
		preDefined: pd,
		candidate:  candidate,
		row:        r,
		col:        c,
		block:      b,
	}
}
