package linker

import (
	"bytes"
	"os"

	"github.com/Taoger-Xu/riscv64-ld/src/utils"
)

// 对磁盘文件的抽象，是interface
type File struct {
	// 名称
	Name string
	// 字节流
	Content []byte
}

func MustNewFile(filename string) *File {
	contents, err := os.ReadFile(filename)
	utils.MustNo(err)

	return &File{
		Name:    filename,
		Content: contents,
	}
}

func CheckMagic(content []byte) bool {
	return bytes.HasPrefix(content, []byte("\177ELF"))
}
