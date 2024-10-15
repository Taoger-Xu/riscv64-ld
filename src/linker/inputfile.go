package linker

import (
	"debug/elf"
	"fmt"

	"github.com/Taoger-Xu/riscv64-ld/src/utils"
)

/**
	InputFile是对ELF文件解析内容的抽象，是File的子类
	File ：指向File对象的指针，是从文件中读取的字节流
	ElfSections：从file中解析得到的entries in the section header table数组
	ShStrTab : 字符串表，存储ELF文件所有的字符串
**/

type InputFile struct {
	File *File
	// section header table
	SectionHeaderTable []Shdr
	// 第一个non-local的符号索引
	FirstGlobalSymbolIdx int64

	// SectionHeader String Table
	ShStrTab []byte

	// SymbolTable String Table
	SymStrTab []byte

	//Symbol Table
	SymbolTable []SymTabEntry
}

func NewInputFile(file *File) InputFile {

	f := InputFile{
		File: file,
	}

	// 文件的长度比elf header的长度小
	if len(file.Content) < EhdrSize {
		utils.Fatal("file too small")
	}

	// 检查elf header的magic number
	if !CheckMagic(f.File.Content) {
		utils.Fatal("not an ELF file")
	}

	// 读取ELF Header
	ehdr := utils.Read[Ehdr](f.File.Content)
	contents := file.Content[ehdr.ShOff:]
	// 从section header table中读取第一个section header
	shdr := utils.Read[Shdr](contents)

	// 初始化SectionHeaderTable
	f.InitSectionHeaderTable(&ehdr)

	// 初始化String Table
	// 特殊值 SHN_XINDEX：如果 e_shnum 或 e_shstrndx 的值为 SHN_XINDEX（0xFFFF），
	// 	表示实际的节头部表条目数或字符串表索引需要从节头部表的第一个条目中获取
	sh_str_ndx := uint64(ehdr.ShStrndx)
	if ehdr.ShStrndx == uint16(elf.SHN_XINDEX) {
		sh_str_ndx = uint64(shdr.Link)
	}
	f.ShStrTab = f.GetSectionBytesByIdx(sh_str_ndx)

	fmt.Printf("str tab is %s, index is %d\n", f.ShStrTab, sh_str_ndx)

	return f

}

// 根据section header得到该section对应的byte[]
func (f *InputFile) GetSectionBytesByShdr(s *Shdr) []byte {
	start := s.Offset
	end := s.Offset + s.Size
	if uint64(len(f.File.Content)) < end {
		utils.Fatal(fmt.Sprintf("section header is out of range %d", s.Offset))
	}
	return f.File.Content[start:end]
}

// 根据在section header table中的index得到该section对应的byte[]
func (f *InputFile) GetSectionBytesByIdx(idx uint64) []byte {
	section_header := f.SectionHeaderTable[idx]
	start := section_header.Offset
	end := section_header.Offset + section_header.Size

	if uint64(len(f.File.Content)) < end {
		utils.Fatal(fmt.Sprintf("section header is out of range %d", section_header.Offset))
	}
	return f.File.Content[start:end]
}

// 找到某种类型的section header
func (f *InputFile) FindSectionByType(tp uint32) *Shdr {
	for i := 0; i < len(f.SectionHeaderTable); i++ {
		shdr := &f.SectionHeaderTable[i]
		if shdr.Type == tp {
			return shdr
		}
	}
	return nil
}

// 根据elf header的信息去初始化SectionHeaderTable
func (f *InputFile) InitSectionHeaderTable(ehdr *Ehdr) {
	// 读取ELF Header
	// ehdr := utils.Read[Ehdr](f.File.Content)
	// contents为section header table开始的位置，为切片
	contents := f.File.Content[ehdr.ShOff:]
	// 从section header table中读取第一个section header
	shdr := utils.Read[Shdr](contents)
	//当 e_shnum 为 0 时，表示节头部表条目数的信息可能在 sh_size 字段中
	numSection := int64(ehdr.ShNum)
	if numSection == 0 {
		numSection = int64(shdr.Size)
	}
	// 第一个section
	f.SectionHeaderTable = []Shdr{shdr}
	// 每个entry的大小固定，依次循环读取即可
	for numSection > 1 {
		contents = contents[ShdrSize:]
		f.SectionHeaderTable = append(f.SectionHeaderTable, utils.Read[Shdr](contents))
		numSection--
	}
}

// 根据section header table entry的信息去初始化SymbolTable
func (f *InputFile) InitSymbolTable(s *Shdr) {
	sym_bytes := f.GetSectionBytesByShdr(s)
	entry_num := len(sym_bytes) / SymEntrySize
	f.SymbolTable = make([]SymTabEntry, 0, entry_num)
	for entry_num > 0 {
		f.SymbolTable = append(f.SymbolTable, utils.Read[SymTabEntry](sym_bytes))
		sym_bytes = sym_bytes[SymEntrySize:]
		entry_num--
	}
}
