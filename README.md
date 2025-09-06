[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/sortx/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/sortx/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/sortx)](https://pkg.go.dev/github.com/yyle88/sortx)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/sortx/master.svg)](https://coveralls.io/github/yyle88/sortx?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/sortx.svg)](https://github.com/yyle88/sortx/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/sortx)](https://goreportcard.com/report/github.com/yyle88/sortx)

# sortx

`sortx` is a Go package that provides a simple and flexible way to sort slices using custom comparison functions. It leverages Go's generics and the `sort.Interface` to avoid repeating the implementation of sorting logic for different types.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

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

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-06 04:53:24.895249 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a bug?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize via reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Common feedback?** Each suggestion and comment is welcome

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a pull request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Happy Coding with this package!** 🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yyle88/sortx.svg?variant=adaptive)](https://starchart.cc/yyle88/sortx)
