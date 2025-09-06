package sortx

import "sort"

// SortByIndex sorts the slice `a` using the index-based comparison function `iLess`.
// SortByIndex 使用基于索引的比较函数 `iLess` 对切片 `a` 进行排序。
func SortByIndex[V any](a []V, iLess func(i, j int) bool) {
	// Sort the slice using the index comparison function.
	// 使用索引比较函数对切片进行排序。
	sort.Sort(NewSortByIndex(a, iLess))
}

// SortByValue sorts the slice `a` using the value-based comparison function `vLess`.
// SortByValue 使用基于值的比较函数 `vLess` 对切片 `a` 进行排序。
func SortByValue[V any](a []V, vLess func(a, b V) bool) {
	// Sort the slice using the value comparison function.
	// 使用值比较函数对切片进行排序。
	sort.Sort(NewSortByValue(a, vLess))
}

// SortIStable sorts the slice `a` using the index-based comparison function `iLess`
// and preserves the original order of equal elements (stable sort).
// SortIStable 使用基于索引的比较函数 `iLess` 对切片 `a` 进行排序，并保持相等元素的原始顺序（稳定排序）。
func SortIStable[V any](a []V, iLess func(i, j int) bool) {
	// Perform a stable sort using the index comparison function.
	// 使用索引比较函数执行稳定排序。
	sort.Stable(NewSortByIndex(a, iLess))
}

// SortVStable sorts the slice `a` using the value-based comparison function `vLess`
// and preserves the original order of equal elements (stable sort).
// SortVStable 使用基于值的比较函数 `vLess` 对切片 `a` 进行排序，并保持相等元素的原始顺序（稳定排序）。
func SortVStable[V any](a []V, vLess func(a, b V) bool) {
	// Perform a stable sort using the value comparison function.
	// 使用值比较函数执行稳定排序。
	sort.Stable(NewSortByValue(a, vLess))
}
