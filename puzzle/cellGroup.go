package puzzle

type CellGroup struct {
	cells []*Cell
}

func (cg CellGroup) Add(c *Cell, pos int) {
	cg.cells[pos] = c
}

func NewCellGroup() *CellGroup {
	cells := make([]*Cell, 9)
	return &CellGroup{cells}
}
