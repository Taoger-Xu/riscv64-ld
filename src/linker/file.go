package linker

import (
	"bytes"
	"os"

	"github.com/Taoger-Xu/riscv64-ld/src/utils"
)

type File struct {
	Name    string
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
