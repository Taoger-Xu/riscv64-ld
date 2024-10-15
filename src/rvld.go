package main

import (
	"fmt"
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

	obj_file := linker.NewObjectFile(file)

	obj_file.ParseFile()

	utils.Assert(len(obj_file.SectionHeaderTable) == 12)
	utils.Assert(obj_file.FirstGlobalSymbolIdx == 12)
	utils.Assert(len(obj_file.SymbolTable) == 14)

	println(len(obj_file.SectionHeaderTable))

	fmt.Println(len(obj_file.SymbolTable))
	// fmt.Printf("%d\n", obj_file.FirstGlobalSymbolIdx)
	// for _, section := range file_with_resolved.ElfSections {
	// 	// fmt.Printf("Section %d:\n", i)
	// 	fmt.Printf("  Name:      %s\n", linker.GetSectionNameByOffset(file_with_resolved.ShStrTab, section.Name))
	// 	// fmt.Printf("  Type:      %d\n", section.Type)
	// 	// fmt.Printf("  Flags:     %d\n", section.Flags)
	// 	// fmt.Printf("  Addr:      %d\n", section.Addr)
	// 	// fmt.Printf("  Offset:    0x%X\n", section.Offset)
	// 	// fmt.Printf("  Size:      %d\n", section.Size)
	// 	// fmt.Printf("  Link:      %d\n", section.Link)
	// 	// fmt.Printf("  Info:      %d\n", section.Info)
	// 	// fmt.Printf("  AddrAlign: %d\n", section.AddrAlign)
	// 	// fmt.Printf("  EntSize:   %d\n", section.EntSize)
	// }

	for _, symbol := range obj_file.SymbolTable {
		fmt.Printf("symbol name is    %s\n", linker.GetSectionNameByOffset(obj_file.SymStrTab, symbol.Name))
	}
}
