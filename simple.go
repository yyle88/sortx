// Package sortx: Flexible slice sorting utilities with custom comparison functions
// Provides generic sorting functions using Go's sort.Interface for type-safe operations
// Supports both index-based and value-based comparison strategies with stable sorting options
//
// sortx: 灵活的切片排序工具，支持自定义比较函数
// 使用 Go 的 sort.Interface 提供类型安全的泛型排序函数
// 支持基于索引和基于值的比较策略，提供稳定排序选项
package sortx

import "sort"

// SortByIndex sorts the slice using the index-based comparison function
// Provides flexible sorting where comparison is done using element indices
//
// SortByIndex 使用基于索引的比较函数对切片进行排序
// 提供灵活的排序方式，通过元素索引进行比较
func SortByIndex[V any](a []V, iLess func(i, j int) bool) {
	// Sort the slice using the index comparison function.
	// 使用索引比较函数对切片进行排序。
	sort.Sort(NewSortByIndex(a, iLess))
}

// SortByValue sorts the slice using the value-based comparison function
// Provides flexible sorting where comparison is done using element values
//
// SortByValue 使用基于值的比较函数对切片进行排序
// 提供灵活的排序方式，通过元素值进行比较
func SortByValue[V any](a []V, vLess func(a, b V) bool) {
	// Sort the slice using the value comparison function.
	// 使用值比较函数对切片进行排序。
	sort.Sort(NewSortByValue(a, vLess))
}

// SortIStable performs stable sorting using index-based comparison function
// Preserves the original order of equal elements during sorting operation
//
// SortIStable 使用基于索引的比较函数执行稳定排序
// 在排序过程中保持相等元素的原始顺序
func SortIStable[V any](a []V, iLess func(i, j int) bool) {
	// Perform a stable sort using the index comparison function.
	// 使用索引比较函数执行稳定排序。
	sort.Stable(NewSortByIndex(a, iLess))
}

// SortVStable performs stable sorting using value-based comparison function
// Preserves the original order of equal elements during sorting operation
//
// SortVStable 使用基于值的比较函数执行稳定排序
// 在排序过程中保持相等元素的原始顺序
func SortVStable[V any](a []V, vLess func(a, b V) bool) {
	// Perform a stable sort using the value comparison function.
	// 使用值比较函数执行稳定排序。
	sort.Stable(NewSortByValue(a, vLess))
}
