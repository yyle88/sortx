[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/sortx/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/sortx/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/sortx)](https://pkg.go.dev/github.com/yyle88/sortx)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/sortx/master.svg)](https://coveralls.io/github/yyle88/sortx?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/sortx.svg)](https://github.com/yyle88/sortx/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/sortx)](https://goreportcard.com/report/github.com/yyle88/sortx)

# sortx

`sortx` is a Go package that provides a simple and flexible way to sort slices using custom comparison functions. It leverages Go's generics and the `sort.Interface` to avoid repeating the implementation of sorting logic for different types.

## CHINESE README

[中文说明](README.zh.md)

## Installation

To install the `sortx` package, you can use the following command:

```bash
go get github.com/yyle88/sortx
```

## Usage

The package offers several functions for sorting slices with different comparison strategies. Below are the key functions available:

### `SortByIndex`

Sorts the slice `a` using an index-based comparison function `iLess`.

```go
sortx.SortByIndex(a []V, iLess func(i, j int) bool)
```

- `a`: The slice to be sorted.
- `iLess`: The function that compares the indices of two elements in the slice.
- Sorts the slice in place using the provided index-based comparison function.

### `SortByValue`

Sorts the slice `a` using a value-based comparison function `vLess`.

```go
sortx.SortByValue(a []V, vLess func(a, b V) bool)
```

- `a`: The slice to be sorted.
- `vLess`: The function that compares the values of two elements in the slice.
- Sorts the slice in place using the provided value-based comparison function.

### `SortIStable`

Sorts the slice `a` using an index-based comparison function `iLess` and preserves the order of equal elements (stable sort).

```go
sortx.SortIStable(a []V, iLess func(i, j int) bool)
```

- `a`: The slice to be sorted.
- `iLess`: The function that compares the indices of two elements in the slice.
- Sorts the slice in place while maintaining the original order of equal elements (stable sort).

### `SortVStable`

Sorts the slice `a` using a value-based comparison function `vLess` and preserves the order of equal elements (stable sort).

```go
sortx.SortVStable(a []V, vLess func(a, b V) bool)
```

- `a`: The slice to be sorted.
- `vLess`: The function that compares the values of two elements in the slice.
- Sorts the slice in place while maintaining the original order of equal elements (stable sort).

## Example

Here's a basic example of how to use `SortByIndex` and `SortByValue`:

```go
package main

import (
	"fmt"
	"github.com/yyle88/sortx"
)

func main() {
	// Example 1: Sorting by index
	numbers := []int{5, 3, 8, 1, 4}
	sortx.SortByIndex(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j] // Compare by values at indices
	})
	fmt.Println("Sorted by index:", numbers)

	// Example 2: Sorting by value
	strings := []string{"apple", "banana", "cherry", "date"}
	sortx.SortByValue(strings, func(a, b string) bool {
		return a < b // Compare by string values
	})
	fmt.Println("Sorted by value:", strings)
}
```

---

## License

MIT License. See [LICENSE](LICENSE).

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repo on GitHub (using the webpage interface).
2. Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. Navigate to the cloned project (`cd repo-name`)
4. Create a feature branch (`git checkout -b feature/xxx`).
5. Stage changes (`git add .`)
6. Commit changes (`git commit -m "Add feature xxx"`).
7. Push to the branch (`git push origin feature/xxx`).
8. Open a pull request on GitHub (on the GitHub webpage).

Please ensure tests pass and include relevant documentation updates.

---

## Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

If you find this package valuable, give me some stars on GitHub! Thank you!!!

**Thank you for your support!**

**Happy Coding with this package!** 🎉

Give me stars. Thank you!!!

---

## GitHub Stars

[![starring](https://starchart.cc/yyle88/sortx.svg?variant=adaptive)](https://starchart.cc/yyle88/sortx)
