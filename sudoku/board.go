package sudoku

import (
	"errors"
	"fmt"
	"log"
)

type Board struct {
	Size       uint8
	Difficulty uint8
	Cells      []Cell `json:"squares"`
}

type Cell struct {
	X   uint8 `json:"x"`
	Y   uint8 `json:"y"`
	Val uint8 `json:"value"`
}

func Create(size, difficulty uint8) *Board {
	cells := []Cell{}

	body := fetch(size, difficulty)
	dto := new(Response)
	must(dto.fromJson(body))

	var i uint8 = 0
	var j uint8 = 0
	for ; i < size; i++ {
		for ; j < size; j++ {
			cells = append(cells, Cell{X: i, Y: j, Val: 0})
		}
		j = 0
	}

	for i := range dto.Squares {
		curr := dto.Squares[i]
		index, err := curr.GetIndex(size)
		must(err)

		cells[index].Val = curr.Val
	}

	return &Board{Size: size, Difficulty: difficulty, Cells: cells}
}

func (b *Board) PrintGraph() {
	fmt.Printf("Size: %v\nDifficulty: %v\n\n", b.Size, b.Difficulty)
	for i := range b.Cells {
		c := b.Cells[i]
		fmt.Printf("{ X: %v, Y: %v, Value: %v }\n", c.X, c.Y, c.Val)
	}
}

func (b *Board) Print() {
	fmt.Printf("Size: %v Difficulty: %v\n\n", b.Size, b.Difficulty)
	s := int(b.Size)
	for i := 0; i < s; i++ {
		fmt.Printf("[%v]", i)
		for j := 0; j < s; j++ {
			c := b.Cells[i*s+j].Val
			fmt.Printf(" %v", c)
		}
		fmt.Printf("\n")
	}
}

func (c *Cell) GetIndex(size uint8) (uint8, error) {
	result := (c.X * size) + c.Y
	if !(result < (size * size)) {

		message := fmt.Sprintf("index %v out of range", result)
		return result, errors.New(message)
	}
	return result, nil
}

func must(err error) {
	if err != nil {
		log.Fatalf("Error, %v\n", err)
	}
}