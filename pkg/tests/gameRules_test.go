package tests

import (
	"reflect"
	"testing"

	"github.com/Dysax/GameOfLife/pkg/board"
)

// test to check if dead cells with no live neighbors stay dead
func TestDeadStayDead(t *testing.T) {
	init_state1 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	expected_next_state1 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	actual_next_state1 := board.CalculateNewBoard(init_state1)
	if !reflect.DeepEqual(expected_next_state1, actual_next_state1) {
		t.Errorf("Expected board state did not match actual.\nExpected: %v\nActual: %v\n", expected_next_state1, actual_next_state1)
	}
}

// should make the target cell alive
func TestComeAlive(t *testing.T) {
	init_state2 := [][]int{
		{0, 0, 1},
		{0, 1, 1},
		{0, 0, 0},
	}

	expected_next_state2 := [][]int{
		{0, 1, 1},
		{0, 1, 1},
		{0, 0, 0},
	}
	actual_next_state2 := board.CalculateNewBoard(init_state2)
	if !reflect.DeepEqual(expected_next_state2, actual_next_state2) {
		t.Errorf("Expected board state did not match actual.\nExpected: %v\nActual: %v\n", expected_next_state2, actual_next_state2)
	}
}

// should kill the target cell
func TestDieProperly(t *testing.T) {
	init_state3 := [][]int{
		{0, 0, 1},
		{0, 1, 1},
		{0, 1, 1},
	}

	expected_next_state3 := [][]int{
		{0, 1, 1},
		{0, 0, 0},
		{0, 1, 1},
	}

	actual_next_state3 := board.CalculateNewBoard(init_state3)
	if !reflect.DeepEqual(expected_next_state3, actual_next_state3) {
		t.Errorf("Expected board state did not match actual.\nExpected: %v\nActual: %v\n", expected_next_state3, actual_next_state3)
	}
}
