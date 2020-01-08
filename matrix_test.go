package matrix

import (
	"fmt"
	"strings"
	"testing"
)

func TestEmptyMatrix(t *testing.T) {
	m := NewEmptyMatrix(5, 5)
	if m.Get(0, 0) == nil {
		t.Errorf("Empty matrix not well initialized, got: nil, want: something")
	}
}

func TestMatrix_Get_and_Set(t *testing.T) {
	m := NewMatrix(5, 6, 0)
	if m.Get(1, 1) != 0 {
		t.Errorf("Matrix not well initialized, got: %v, want: %d", m.Get(1, 1), 0)
	}
	if m.GetRowsNumber() != 5 {
		t.Errorf("Error with row number, got: %d, want: %d", m.GetRowsNumber(), 5)
	}
	if m.GetColsNumber() != 6 {
		t.Errorf("Error with col number, got; %d, want: %d", m.GetColsNumber(), 6)
	}
	m.Set(1, 1, 8)
	if m.Get(1, 1) != 8 {
		t.Errorf("Cell not well set, got: %d, want: %d", m.Get(1, 1), 8)
	}
	m.Set(1, 1, "X")
	if m.Get(1, 1) != "X" {
		t.Errorf("Cell not well updated, got: %v, want: %s", m.Get(1, 1), "X")
	}
	m.Fill("F")
	if m.Get(1, 5) != "F" {
		t.Errorf("Bad matrix filling, got: %v, want: %s", m.Get(1, 5), "F")
	}
	m.SetCol(2, "C")
	if fmt.Sprintf("%v", m.GetRow(3)) != "[F F C F F F]" {
		t.Errorf("Col not well set, got: %s, want: %s", fmt.Sprintf("%v", m.GetRow(3)), "[F F C F F F]")
	}
	m.SetRow(0, "R")
	if fmt.Sprintf("%v", m.GetCol(2)) != "[R C C C C]" {
		t.Errorf("Row not well set, got: %s, want: %s", fmt.Sprintf("%v", m.GetCol(2)), "[R C C C C]")
	}
	if strings.Compare(m.ToString(), "[R][R][R][R][R][R]\n[F][F][C][F][F][F]\n[F][F][C][F][F][F]\n[F][F][C][F][F][F]\n[F][F][C][F][F][F]\n") != 0 {
		t.Errorf("no got: %s, want %s", m.ToString(), "[R][R][R][R][R][R]\n[F][F][C][F][F][F]\n[F][F][C][F][F][F]\n[F][F][C][F][F][F]\n[F][F][C][F][F][F]\n")
	}
}

func TestMatrix_bad_init(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but got nothing")
		}
	}()
	NewMatrix(0, 5, 0)
}