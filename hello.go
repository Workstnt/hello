package main

import (
    "fmt"
    "os"
    "runtime"
)

func main() {
    fmt.Println("便携Go环境运行成功！")
    fmt.Println("Hello Copilot")
    fmt.Println("GOROOT:", os.Getenv("GOROOT"))
    fmt.Println("GOPATH:", os.Getenv("GOPATH"))
    fmt.Println("GO版本:", runtime.Version())
    fmt.Println("GOOS:", runtime.GOOS)
    fmt.Println("GOARCH:", runtime.GOARCH)
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