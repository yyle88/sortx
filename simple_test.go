package sortx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortByIndex(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	SortByIndex(a, func(i, j int) bool { return a[i] < a[j] })
	t.Log(a)
	require.Equal(t, a, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestSortByValue(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	SortByValue(a, func(a, b int) bool { return a < b })
	t.Log(a)
	require.Equal(t, a, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestSortIStable(t *testing.T) {
	// Test stable sort with index comparison
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 30},
		{"David", 25},
	}

	// Sort by age, preserving original order for equal ages
	SortIStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	expected := []Person{
		{"Bob", 25},
		{"David", 25},
		{"Alice", 30},
		{"Charlie", 30},
	}

	require.Equal(t, expected, people)
}

func TestSortVStable(t *testing.T) {
	// Test stable sort with value comparison
	type Item struct {
		Value    int
		Priority int
	}

	items := []Item{
		{1, 5},
		{2, 3},
		{3, 5},
		{4, 3},
	}

	// Sort by priority, preserving original order for equal priorities
	SortVStable(items, func(a, b Item) bool {
		return a.Priority < b.Priority
	})

	expected := []Item{
		{2, 3},
		{4, 3},
		{1, 5},
		{3, 5},
	}

	require.Equal(t, expected, items)
}

func TestSortEmptySlice(t *testing.T) {
	// Test sorting empty slices
	var empty []int

	SortByIndex(empty, func(i, j int) bool { return empty[i] < empty[j] })
	require.Empty(t, empty)

	SortByValue(empty, func(a, b int) bool { return a < b })
	require.Empty(t, empty)

	SortIStable(empty, func(i, j int) bool { return empty[i] < empty[j] })
	require.Empty(t, empty)

	SortVStable(empty, func(a, b int) bool { return a < b })
	require.Empty(t, empty)
}

func TestSortSingleElement(t *testing.T) {
	// Test sorting single element slices
	single := []string{"hello"}

	SortByIndex(single, func(i, j int) bool { return single[i] < single[j] })
	require.Equal(t, []string{"hello"}, single)

	SortByValue(single, func(a, b string) bool { return a < b })
	require.Equal(t, []string{"hello"}, single)
}
