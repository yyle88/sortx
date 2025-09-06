package sortx

import (
	"sort"

	"github.com/pkg/errors"
)

// Slice represents a slice that can be sorted using either an index-based or value-based comparison function.
// Slice 结构体表示一个可以使用索引或值比较函数排序的切片。
type Slice[V any] struct {
	slice []V                 // The underlying slice to be sorted // 要排序的基础切片
	iLess func(i, j int) bool // Index-based comparison function // 基于索引的比较函数
	vLess func(a, b V) bool   // Value-based comparison function // 基于值的比较函数
}

// NewSortByIndex creates a new sort.Interface for sorting by index comparison.
// NewSortByIndex 根据索引比较函数创建一个新的 sort.Interface。
func NewSortByIndex[V any](a []V, iLess func(i, j int) bool) sort.Interface {
	return &Slice[V]{slice: a, iLess: iLess}
}

// NewSortByValue creates a new sort.Interface for sorting by value comparison.
// NewSortByValue 根据值比较函数创建一个新的 sort.Interface。
func NewSortByValue[V any](a []V, vLess func(a, b V) bool) sort.Interface {
	return &Slice[V]{slice: a, vLess: vLess}
}

// Len returns the number of elements in the slice.
// Len 返回切片中的元素数量。
func (T *Slice[V]) Len() int {
	return len(T.slice)
}

// Less compares two elements at indexes i and j for sorting.
// Less 比较索引 i 和 j 处的两个元素，用于排序。
func (T *Slice[V]) Less(i, j int) bool {
	switch {
	// If the index-based comparison function is provided, use it.
	// 如果提供了基于索引的比较函数，则使用它。
	case T.iLess != nil:
		return T.iLess(i, j)
	// If the value-based comparison function is provided, use it.
	// 如果提供了基于值的比较函数，则使用它。
	case T.vLess != nil:
		return T.vLess(T.slice[i], T.slice[j])
	// If neither comparison function is provided, panic with an error.
	// 如果没有提供比较函数，则引发错误。
	default:
		panic(errors.New("missing less function"))
	}
}

// Swap swaps the elements with indexes i and j.
// Swap 交换索引 i 和 j 处的两个元素。
func (T *Slice[V]) Swap(i, j int) {
	T.slice[i], T.slice[j] = T.slice[j], T.slice[i]
}
