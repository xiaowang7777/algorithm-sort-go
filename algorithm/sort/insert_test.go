package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertSort(t *testing.T) {
	var arr []int = []int{4, 3, 5, 4, 2, 1}
	var res []int = []int{1, 2, 3, 4, 4, 5}
	InsertSort(arr)

	assert.Equal(t, arr, res)
}
