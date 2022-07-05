package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	row, column int
	Grid        [][]int
}

func (m1 *Matrix) getRow() int {
	return m1.row
}
func (m1 *Matrix) getColumn() int {
	return m1.column
}
func (m *Matrix) setElemet(i, j, val int) {
	if i < 0 || i >= m.getRow() {
		fmt.Println("Cannot set the element")
		return
	}
	if j < 0 || j >= m.getColumn() {
		fmt.Println("Cannot set the element")
		return
	}
	m.Grid[i][j] = val
}

func (m1 *Matrix) addMatrix(m2 Matrix) [][]int {

	addMatrix := make([][]int, m1.getRow())

	for i := range addMatrix {
		addMatrix[i] = make([]int, m1.getColumn())
	}

	for i := 0; i < m1.getRow(); i++ {
		for j := 0; j < m1.getColumn(); j++ {
			addMatrix[i][j] = m1.Grid[i][j] + m2.Grid[i][j]
		}
	}
	return addMatrix
}

func main() {
	n := 5
	mm := 4

	m1 := Matrix{row: n, column: mm, Grid: make([][]int, n)}
	for i := range m1.Grid {
		m1.Grid[i] = make([]int, m1.getColumn())
	}

	m1.setElemet(0, 0, 4)
	m1.setElemet(1, 0, 4)

	m2 := Matrix{row: n, column: mm, Grid: make([][]int, n)}
	for i := range m2.Grid {
		m2.Grid[i] = make([]int, m2.getColumn())
	}

	m2.setElemet(1, 1, 6)

	fmt.Println("Fisrt Matrix", m1.Grid)
	fmt.Println("Second Matrix", m2.Grid)

	addMatrix := m1.addMatrix(m2)

	fmt.Println("Add Matrix", addMatrix)

	data, _ := json.Marshal(addMatrix)
	fmt.Printf("%s\n", data)

}
