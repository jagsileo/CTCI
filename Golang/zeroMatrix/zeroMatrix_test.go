package zeroMatrix

import (
	"reflect"
	"testing"
)

func TestSetZeroMatrix(t *testing.T) {
	matrix := [][]int{{1, 2, 3, 4}, {5, 0, 7, 8}, {9, 10, 0, 12}}

	newMatrix := [][]int{{1, 0, 0, 4}, {0, 0, 0, 0}, {0, 0, 0, 0}}

	if !reflect.DeepEqual(newMatrix, setZeroMatrix(matrix)) {

		t.Error("Test Failed for #1")
	}

	matrix = [][]int{{1, 2, 3, 0}, {5, 0, 7, 8}, {9, 10, 0, 12}}

	newMatrix = [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}

	if !reflect.DeepEqual(newMatrix, setZeroMatrix(matrix)) {

		t.Error("Test Failed for #2")
	}
}
