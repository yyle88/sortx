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

func TestSliceMethods(t *testing.T) {
	// Test Slice struct methods directly
	data := []int{3, 1, 4, 1, 5}

	// Test with index-based comparison
	slice1 := &Slice[int]{
		slice: data,
		iLess: func(i, j int) bool { return data[i] < data[j] },
	}

	require.Equal(t, 5, slice1.Len())
	require.True(t, slice1.Less(1, 0))  // data[1]=1 < data[0]=3
	require.False(t, slice1.Less(0, 1)) // data[0]=3 > data[1]=1

	// Test Swap method
	slice1.Swap(0, 1)
	require.Equal(t, []int{1, 3, 4, 1, 5}, data)

	// Test with value-based comparison
	data2 := []string{"banana", "apple", "cherry"}
	slice2 := &Slice[string]{
		slice: data2,
		vLess: func(a, b string) bool { return a < b },
	}

	require.Equal(t, 3, slice2.Len())
	require.True(t, slice2.Less(1, 0))  // "apple" < "banana"
	require.False(t, slice2.Less(0, 1)) // "banana" > "apple"

	slice2.Swap(0, 1)
	require.Equal(t, []string{"apple", "banana", "cherry"}, data2)
}

func TestSlicePanic(t *testing.T) {
	// Test panic when no comparison function is provided
	slice := &Slice[int]{
		slice: []int{1, 2, 3},
		// Neither iLess nor vLess is set
	}

	require.Panics(t, func() {
		slice.Less(0, 1)
	}, "Should panic when no comparison function is provided")
}

func TestSliceWithBothComparisons(t *testing.T) {
	// Test that iLess takes priority over vLess
	data := []int{3, 1, 4}
	slice := &Slice[int]{
		slice: data,
		iLess: func(i, j int) bool { return data[i] < data[j] }, // Should be used
		vLess: func(a, b int) bool { return a > b },             // Should be ignored
	}

	// iLess should be used (ascending), not vLess (descending)
	require.True(t, slice.Less(1, 0))  // data[1]=1 < data[0]=3
	require.False(t, slice.Less(0, 1)) // data[0]=3 > data[1]=1
}

func TestSortInterface(t *testing.T) {
	// Test that our types implement sort.Interface correctly
	data := []float64{3.14, 2.71, 1.41, 1.73}

	iface := NewSortByValue(data, func(a, b float64) bool { return a < b })

	// Verify it implements sort.Interface
	var _ sort.Interface = iface

	// Test the interface methods
	require.Equal(t, 4, iface.Len())
	require.True(t, iface.Less(2, 0)) // 1.41 < 3.14

	// Test full sort
	sort.Sort(iface)
	expected := []float64{1.41, 1.73, 2.71, 3.14}
	require.Equal(t, expected, data)
}
