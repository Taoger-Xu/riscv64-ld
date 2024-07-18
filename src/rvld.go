package main

import (
	"os"

	"github.com/Taoger-Xu/riscv64-ld/src/linker"
	"github.com/Taoger-Xu/riscv64-ld/src/utils"
)

func main() {
	// 读取elf文件
	if len(os.Args) < 2 {
		utils.Fatal("wrong args")
	}

	filename := os.Args[1]
	file := linker.MustNewFile(filename)

	if !linker.CheckMagic(file.Content) {
		utils.Fatal("not an ELF file")
	}

	println(len(file.Content))
}
