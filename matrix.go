package matrix

import (
	"fmt"
	"strings"
)

type Matrix struct {
	objs [][]interface{}
}


func NewEmptyMatrix(row int, col int) *Matrix {
	return NewMatrix(row, col, new(interface{}))
}

func NewMatrix(row int, col int, obj interface{}) *Matrix {
	if row <= 0 || col <= 0 {
		panic("row or col of the matrix can't be negative")
	}

	m := new(Matrix)

	for i := 0; i < row; i++ {
		var column []interface{}
		for j := 0; j < col; j++ {
			column = append(column, obj)
		}
		m.objs = append(m.objs, column)
	}

	return m
}

func (m *Matrix) GetRowsNumber() int {
	return len(m.objs)
}

func (m *Matrix) GetColsNumber() int {
	return len(m.objs[0])
}

func (m *Matrix) Get(row int, col int) interface{} {
	return m.objs[row][col]
}

func (m *Matrix) Set(row int, col int, i interface{}) {
	m.objs[row][col] = i
}

func (m *Matrix) GetCol(col int) []interface{} {
	var res []interface{}
	for i := 0; i < m.GetRowsNumber(); i++ {
		res = append(res, m.objs[i][col])
	}
	return res
}

func (m *Matrix) SetCol(col int, obj interface{}) {
	for i := 0; i < m.GetRowsNumber(); i++ {
		m.objs[i][col] = obj
	}
}

func (m *Matrix) GetRow(row int) []interface{} {
	return m.objs[row]
}

func (m *Matrix) SetRow(row int, obj interface{}) {
	for i := 0; i < m.GetColsNumber(); i++ {
		m.objs[row][i] = obj
	}
}

func (m *Matrix) ToString() string {
	var b strings.Builder
	for i := 0; i < m.GetRowsNumber(); i++ {
		for j := 0; j < m.GetColsNumber(); j++ {
			b.WriteString("[")
			b.WriteString(fmt.Sprintf("%v", m.objs[i][j]))
			b.WriteString("]")
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (m *Matrix) Fill(obj interface{}) {
	for i := 0; i < m.GetRowsNumber(); i++ {
		for j := 0; j < m.GetColsNumber(); j++ {
			m.objs[i][j] = obj
		}
	}
}