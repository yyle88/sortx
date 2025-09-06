package sortx

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSortByIndex(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	sort.Sort(NewSortByIndex(a, func(i, j int) bool { return a[i] < a[j] }))
	t.Log(a)
	require.Equal(t, a, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestNewSortByValue(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	sort.Sort(NewSortByValue(a, func(a, b int) bool { return a < b }))
	t.Log(a)
	require.Equal(t, a, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}
