# GoUtils

## 简介 (Introduction)

**GoUtils** 是一个高性能、模块化的 Go 语言工具包，旨在提供常用的字符串、切片、并发和通用工具方法。它注重性能优化，使用如 `strings.Builder`、`sync.Pool` 和 `unsafe` 等技术，适合需要高效处理数据的场景。工具包按功能模块化设计，易于扩展和维护，包含单元测试。

## 功能 (Features)

- **字符串工具 (String Utilities)**：

  - 高效字符串拼接 (`Concat`) 和零拷贝转换 (`FastStringToBytes`, `FastBytesToString`)。
  - 首尾空白裁剪 (`TrimWhitespace`)、首字母大写 (`Capitalize`)、子字符串检查 (`ContainsAny`) 和按长度分割 (`SplitByLength`)。

- **切片工具 (Slice Utilities)**：

  - 元素删除 (`Remove`)、去重 (`Deduplicate`)、过滤 (`Filter`)、映射 (`Map`) 和反转 (`Reverse`)。

- **并发工具 (Concurrent Utilities)**：

  - 并发安全的字符串拼接 (`ConcurrentConcat`) 和线程安全计数器 (`SafeCounter`)。

- **通用工具 (Common Utilities)**：

  - 字符串判空 (`IsEmpty`)、反转 (`Reverse`)、随机字符串生成 (`RandomString`) 和数值范围限制 (`Clamp`)。

- **性能优化 (Performance Optimization)**：

  - 使用 `strings.Builder` 和预分配内存减少内存分配。
  - 使用 `sync.Pool` 管理对象，降低 GC 压力。
  - 提供 `unsafe` 方法用于极致性能场景。

- **模块化设计 (Modular Design)**：

  - 按功能分为 `stringutils`、`sliceutils`、`concurrentutils` 和 `commonutils` 包。
  - 每个包包含单元测试，确保可靠性。

## 安装 (Installation)

1. 确保你的 Go 版本 &gt;= 1.21。
2. 使用 `go get` 安装工具包：

```bash
go get github.com/yourusername/goutils
```

3. 在代码中导入所需包：

```go
import (
    "github.com/yourusername/goutils/stringutils"
    "github.com/yourusername/goutils/sliceutils"
    // 其他包
)
```

## 使用示例 (Usage Examples)

以下是一些常见用法的示例：

### 字符串拼接 (String Concatenation)

```go
package main

import (
    "fmt"
    "github.com/yourusername/goutils/stringutils"
)

func main() {
    su := stringutils.NewStringUtils()
    result := su.Concat("Hello", " ", "World", "!")
    fmt.Println(result) // 输出: Hello World!
}
```

### 切片去重 (Slice Deduplication)

```go
package main

import (
    "fmt"
    "github.com/yourusername/goutils/sliceutils"
)

func main() {
    su := sliceutils.NewSliceUtils()
    input := []string{"a", "b", "a", "c", "b"}
    result := su.Deduplicate(input)
    fmt.Println(result) // 输出: [a b c]
}
```

### 并发安全计数器 (Thread-Safe Counter)

```go
package main

import (
    "fmt"
    "sync"
    "github.com/yourusername/goutils/concurrentutils"
)

func main() {
    sc := concurrentutils.NewSafeCounter()
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            sc.Increment()
        }()
    }
    wg.Wait()
    fmt.Println(sc.Value()) // 输出: 100
}
```

### 随机字符串生成 (Random String Generation)

```go
package main

import (
    "fmt"
    "github.com/yourusername/goutils/commonutils"
)

func main() {
    cu := commonutils.NewCommonUtils()
    result := cu.RandomString(10)
    fmt.Println(result) // 输出: 随机字符串，如 "aB7kP9mN2x"
}
```

## 目录结构 (Directory Structure)

```
goutils/
├── stringutils/
│   ├── stringutils.go      // 字符串工具实现
│   ├── stringutils_test.go // 字符串工具测试
├── sliceutils/
│   ├── sliceutils.go       // 切片工具实现
│   ├── sliceutils_test.go  // 切片工具测试
├── concurrentutils/
│   ├── concurrentutils.go  // 并发工具实现
│   ├── concurrentutils_test.go // 并发工具测试
├── commonutils/
│   ├── commonutils.go      // 通用工具实现
│   ├── commonutils_test.go // 通用工具测试
├── go.mod                  // Go 模块定义
├── README.md               // 仓库说明
```

## 开发与贡献 (Development & Contribution)

1. 克隆仓库：

```bash
git clone https://github.com/yourusername/goutils.git
```

2. 运行测试：

```bash
go test ./...
```

3. 提交 Pull Request：
   - 欢迎贡献新功能、优化性能或修复 Bug。
   - 请确保代码通过测试并遵循 Go 编码规范。

## 许可证 (License)

本项目采用 MIT 许可证。详情见 LICENSE 文件。

## 联系 (Contact)

如有问题或建议，请在 GitHub Issues 中提出，或联系：3174285493@qq.com。
