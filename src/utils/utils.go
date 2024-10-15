package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"runtime/debug"
)

const (
	Red   = "\033[0;31m"
	Reset = "\033[0m"
)

func Fatal(v any) {
	fmt.Printf("%srvld: fatal:%s %v\n", Red, Reset, v)
	debug.PrintStack()
	os.Exit(1)
}

func MustNo(err error) {
	if err != nil {
		Fatal(err)
	}
}

/*
*

	从 data 字节切片中读取一个类型为 T 的值，读取的字节数取决于类型 T 的大小
	[T any] 表示函数 Read 是一个泛型函数
	val 是返回值的名称。在函数体内，你可以通过 val 来引用返回值

	[T any]：它定义了一个类型参数 T。T 是一个类型占位符，可以是任何类型。关键字 any 表示 T 可以是任意类型，相当于 interface{}

*
*/
func Read[T any](data []byte) (val T) {
	reader := bytes.NewReader(data)                       // 创建一个字节读取器，从 data 字节切片中读取数据
	err := binary.Read(reader, binary.LittleEndian, &val) // 从 reader 中读取数据，并将其解码为小端序的 T 类型值
	MustNo(err)
	return val
}

func Assert(condition bool) {
	if !condition {
		Fatal("Assert Failed")
	}
}
