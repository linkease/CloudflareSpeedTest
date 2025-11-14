package color

// 极简本地 color 库：只保留与 utils/color.go 兼容的 API，
// 不做真实着色，只在开启时输出原始文本；关闭时静默。

import (
	"fmt"
	"io"
	"os"
)

type Attribute int

// 保留常量以兼容原调用（实际不使用）。
const (
	FgRed Attribute = iota
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgHiCyan
	FgWhite
	Bold
)

// 全局开关：true 打印；false 不打印（静默）。
var enabled = true

func SetEnabled(b bool) { enabled = b }

type Color struct{}

// New 兼容原 fatih/color.New(...) 签名。
func New(attrs ...Attribute) *Color { return &Color{} }

func (c *Color) Printf(format string, a ...interface{}) (int, error) {
	if !enabled {
		return 0, nil
	}
	return fmt.Printf(format, a...)
}

func (c *Color) Println(a ...interface{}) (n int, err error) {
	if !enabled {
		return 0, nil
	}
	return fmt.Println(a...)
}

func (c *Color) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	if !enabled {
		return 0, nil
	}
	return fmt.Fprintln(w, a...)
}

// 兼容全局变量名。
var (
	Red     = New(FgRed)
	Green   = New(FgGreen)
	Yellow  = New(FgYellow)
	Blue    = New(FgBlue, Bold)
	Magenta = New(FgMagenta)
	Cyan    = New(FgHiCyan, Bold)
	White   = New(FgWhite)
)

var Stdout = os.Stdout
