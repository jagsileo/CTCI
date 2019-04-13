package rotateMatrix

import (
	"log"
	"reflect"
	"testing"
)

func TestRotateMatrix(t *testing.T) {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	newMatrix := rotateMatrix(matrix)
	log.Println(newMatrix)
	if !reflect.DeepEqual(newMatrix, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}) {
		t.Error("Test Failed for 3x3")
	}

	matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

	newMatrix = rotateMatrix(matrix)
	log.Println(newMatrix)
	if !reflect.DeepEqual(newMatrix, [][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}}) {
		t.Error("Test Failed for 4x4")
	}

	matrix = [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}

	newMatrix = rotateMatrix(matrix)
	log.Println(newMatrix)
	if !reflect.DeepEqual(newMatrix, [][]int{{21, 16, 11, 6, 1}, {22, 17, 12, 7, 2}, {23, 18, 13, 8, 3}, {24, 19, 14, 9, 4}, {25, 20, 15, 10, 5}}) {
		t.Error("Test Failed for 5x5")
	}

	matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	newMatrix = [][]int{}
	if !reflect.DeepEqual(newMatrix, rotateMatrix(matrix)) {
		t.Error("Test Failed for 2x4")
	}
}
