# sortx

`sortx` 是一个 Go 包，它提供了一种简单灵活的方式来使用自定义比较函数对切片进行排序。它利用了 Go 的泛型和 `sort.Interface`，避免了为不同类型重复实现排序逻辑。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 安装

要安装 `sortx` 包，可以使用以下命令：

```bash
go get github.com/yyle88/sortx
```

## 使用

该包提供了几种排序函数，支持不同的比较策略。以下是可用的主要函数：

### `SortByIndex`

使用基于索引的比较函数 `iLess` 对切片 `a` 进行排序。

```go
sortx.SortByIndex(a []V, iLess func(i, j int) bool)
```

- `a`: 要排序的切片。
- `iLess`: 比较切片中两个元素索引的函数。
- 使用提供的基于索引的比较函数对切片进行排序。

### `SortByValue`

使用基于值的比较函数 `vLess` 对切片 `a` 进行排序。

```go
sortx.SortByValue(a []V, vLess func(a, b V) bool)
```

- `a`: 要排序的切片。
- `vLess`: 比较切片中两个元素值的函数。
- 使用提供的基于值的比较函数对切片进行排序。

### `SortIStable`

使用基于索引的比较函数 `iLess` 对切片 `a` 进行排序，并保持相等元素的原始顺序（稳定排序）。

```go
sortx.SortIStable(a []V, iLess func(i, j int) bool)
```

- `a`: 要排序的切片。
- `iLess`: 比较切片中两个元素索引的函数。
- 使用基于索引的比较函数对切片进行排序，同时保持相等元素的原始顺序（稳定排序）。

### `SortVStable`

使用基于值的比较函数 `vLess` 对切片 `a` 进行排序，并保持相等元素的原始顺序（稳定排序）。

```go
sortx.SortVStable(a []V, vLess func(a, b V) bool)
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
	"github.com/yyle88/sortx"
)

func main() {
	// 示例 1：按索引排序
	numbers := []int{5, 3, 8, 1, 4}
	sortx.SortByIndex(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j] // 按索引处的值进行比较
	})
	fmt.Println("按索引排序:", numbers)

	// 示例 2：按值排序
	strings := []string{"apple", "banana", "cherry", "date"}
	sortx.SortByValue(strings, func(a, b string) bool {
		return a < b // 按字符串值进行比较
	})
	fmt.Println("按值排序:", strings)
}
```

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-06 04:53:24.895249 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **意见反馈？** 欢迎所有建议和宝贵意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Pull Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Pull Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**使用这个包快乐编程！** 🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->
