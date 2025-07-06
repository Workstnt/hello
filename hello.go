package main

import (
	"fmt"
	"os"
	"runtime"
)

const (
	goRoot = "/path/to/your/go"     // 替换为你的 GOROOT 路径
	goPath = "/path/to/your/gopath" // 替换为你的 GOPATH 路径
	goPath2
)
const (
	a1 = iota
	a2
	a3
	a4
	a5
)
const (
	b1 = iota
	b2
)

const (
	c1 = iota + 1 // 从 1 开始
	c2 = 100
	c3 = iota
	c4
	c5
)

const (
	d1, d2 = iota + 1, iota + 1 // 从 1 开始
	d3, d4 = iota + 1, iota + 1 // 从 1 开始
	d5, d6 = iota + 1, iota + 1 // 从 1 开始
)

const (
	// display constants for << operator
	e1 = 1 << iota // 1 << 0 = 1
	e2             // 1 << 1 = 2
	e3             // 1 << 2 = 4
	e4             // 1 << 3 = 8
	e5             // 1 << 4 = 16
)

func main() {

	fmt.Println("便携Go环境运行成功！")
	fmt.Println("Hello Copilot")
	fmt.Println("GOROOT:", os.Getenv("GOROOT"))
	fmt.Println("GOPATH:", os.Getenv("GOPATH"))
	fmt.Println("GO版本:", runtime.Version())
	fmt.Println("GOOS:", runtime.GOOS)
	fmt.Println("GOARCH:", runtime.GOARCH)
	fmt.Println(goRoot)
	fmt.Println(goPath)
	fmt.Println(goPath2)
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("a4:", a4)
	fmt.Println("a5:", a5)
	fmt.Println("a1 + a2:", a1+a2)
	fmt.Println("a2 + a3:", a2+a3)
	fmt.Println("a3 + a4:", a3+a4)
	fmt.Println("a4 + a5:", a4+a5)
	fmt.Println("a1 + a2 + a3:", a1+a2+a3)
	fmt.Println("a2 + a3 + a4:", a2+a3+a4)
	fmt.Println("a3 + a4 + a5:", a3+a4+a5)
	fmt.Println("a1 + a2 + a3 + a4:", a1+a2+a3+a4)
	fmt.Println("a2 + a3 + a4 + a5:", a2+a3+a4+a5)
	fmt.Println("a1 + a2 + a3 + a4 + a5:", a1+a2+a3+a4+a5)
	fmt.Println("a1 * a2:", a1*a2)
	fmt.Println("a2 * a3:", a2*a3)
	fmt.Println("a3 * a4:", a3*a4)
	fmt.Println("a4 * a5:", a4*a5)
	fmt.Println("a1 * a2 * a3:", a1*a2*a3)
	fmt.Println("a2 * a3 * a4:", a2*a3*a4)
	fmt.Println("a3 * a4 * a5:", a3*a4*a5)
	fmt.Println("a1 * a2 * a3 * a4:", a1*a2*a3*a4)
	fmt.Println("a2 * a3 * a4 * a5:", a2*a3*a4*a5)
	fmt.Println("a1 * a2 * a3 * a4 * a5:", a1*a2*a3*a4*a5)
	fmt.Println("a1 / a2:", a1/(a2+1))                                     // 避免除以0为除数
	fmt.Println("a2 / a3:", a2/(a3+1))                                     // 避免除以0为除数
	fmt.Println("a3 / a4:", a3/(a4+1))                                     // 避免除以0为除数
	fmt.Println("a4 / a5:", a4/(a5+1))                                     // 避免除以0为除数
	fmt.Println("a1 / a2 / a3:", a1/(a2+1)/(a3+1))                         // 避免除以0为除数
	fmt.Println("a2 / a3 / a4:", a2/(a3+1)/(a4+1))                         // 避免除以0为除数
	fmt.Println("a3 / a4 / a5:", a3/(a4+1)/(a5+1))                         // 避免除以0为除数
	fmt.Println("a1 / a2 / a3 / a4:", a1/(a2+1)/(a3+1)/(a4+1))             // 避免除以0为除数
	fmt.Println("a2 / a3 / a4 / a5:", a2/(a3+1)/(a4+1)/(a5+1))             // 避免除以0为除数
	fmt.Println("a1 / a2 / a3 / a4 / a5:", a1/(a2+1)/(a3+1)/(a4+1)/(a5+1)) // 避免除以0为除数
	fmt.Println("a1 % a2:", a1%(a2+1))                                     // 避免除以0为除数
	fmt.Println("a2 % a3:", a2%(a3+1))                                     // 避免  除以0为除数
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)
	fmt.Println("c5:", c5)
	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)
	fmt.Println("d5:", d5)
	fmt.Println("d6:", d6)
	fmt.Println("e1:", e1)         // BIN: 1
	fmt.Println("e2:", e2)         // BIN: 10
	fmt.Println("e3:", e3)         // BIN: 100
	fmt.Println("e4:", e4)         // BIN: 1000
	fmt.Println("e5:", e5)         // BIN: 10000
	fmt.Println("e1 >> 1:", e1>>1) // 右移操作
	fmt.Println("e2 >> 1:", e2>>1) // 右移操作
	fmt.Println("e3 >> 1:", e3>>1) // 右移操作
	fmt.Println("e4 >> 1:", e4>>1) // 右移操作
	fmt.Println("e5 >> 1:", e5>>1) // 右移操作

}

// 运行命令：go run hello.go
// 输出：
// 便携Go环境运行成功！
// Hello Copilot
// GOROOT: /path/to/your/go
// GOPATH: /path/to/your/gopath
// GO版本: go1.20.3
// GOOS: linux
// GOARCH: amd64
// 注意：请根据实际情况修改 GOROOT 和 GOPATH 的路径
// 你可以在终端中运行 `go run hello.go` 来测试这个程序
// 确保你已经安装了 Go 语言环境，并且设置了 GOROOT 和 GOPATH 环境变量
// 如果你使用的是 Windows 系统，请确保路径格式正确，例如：`C:\path\to\your\go` 和 `C:\path\to\your\gopath`
