# sortslice

`sortslice` 是一个 Go 包，它提供了一种简单灵活的方式来使用自定义比较函数对切片进行排序。它利用了 Go 的泛型和 `sort.Interface`，避免了为不同类型重复实现排序逻辑。

## 英文文档

[English README](README.md)

## 安装

要安装 `sortslice` 包，可以使用以下命令：

```bash
go get github.com/yyle88/sortslice
```

## 使用

该包提供了几种排序函数，支持不同的比较策略。以下是可用的主要函数：

### `SortByIndex`

使用基于索引的比较函数 `iLess` 对切片 `a` 进行排序。

```go
sortslice.SortByIndex(a []V, iLess func(i, j int) bool)
```

- `a`: 要排序的切片。
- `iLess`: 比较切片中两个元素索引的函数。
- 使用提供的基于索引的比较函数对切片进行排序。

### `SortByValue`

使用基于值的比较函数 `vLess` 对切片 `a` 进行排序。

```go
sortslice.SortByValue(a []V, vLess func(a, b V) bool)
```

- `a`: 要排序的切片。
- `vLess`: 比较切片中两个元素值的函数。
- 使用提供的基于值的比较函数对切片进行排序。

### `SortIStable`

使用基于索引的比较函数 `iLess` 对切片 `a` 进行排序，并保持相等元素的原始顺序（稳定排序）。

```go
sortslice.SortIStable(a []V, iLess func(i, j int) bool)
```

- `a`: 要排序的切片。
- `iLess`: 比较切片中两个元素索引的函数。
- 使用基于索引的比较函数对切片进行排序，同时保持相等元素的原始顺序（稳定排序）。

### `SortVStable`

使用基于值的比较函数 `vLess` 对切片 `a` 进行排序，并保持相等元素的原始顺序（稳定排序）。

```go
sortslice.SortVStable(a []V, vLess func(a, b V) bool)
```

- `a`: 要排序的切片。
- `vLess`: 比较切片中两个元素值的函数。
- 使用基于值的比较函数对切片进行排序，同时保持相等元素的原始顺序（稳定排序）。

## 示例

以下是如何使用 `SortByIndex` 和 `SortByValue` 的基本示例：

```go
package main

import (
	"fmt"
	"github.com/yyle88/sortslice"
)

func main() {
	// 示例 1：按索引排序
	numbers := []int{5, 3, 8, 1, 4}
	sortslice.SortByIndex(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j] // 按索引处的值进行比较
	})
	fmt.Println("按索引排序:", numbers)

	// 示例 2：按值排序
	strings := []string{"apple", "banana", "cherry", "date"}
	sortslice.SortByValue(strings, func(a, b string) bool {
		return a < b // 按字符串值进行比较
	})
	fmt.Println("按值排序:", strings)
}
```

---

## 许可证类型

项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE)。

---

## 贡献新代码

非常欢迎贡献代码！贡献流程：

1. 在 GitHub 上 Fork 仓库 （通过网页界面操作）。
2. 克隆Forked项目 (`git clone https://github.com/yourname/repo-name.git`)。
3. 在克隆的项目里 (`cd repo-name`)
4. 创建功能分支（`git checkout -b feature/xxx`）。
5. 添加代码 (`git add .`)。
6. 提交更改（`git commit -m "添加功能 xxx"`）。
7. 推送分支（`git push origin feature/xxx`）。
8. 发起 Pull Request （通过网页界面操作）。

请确保测试通过并更新相关文档。

---

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
