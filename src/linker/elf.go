package linker

import (
	"bytes"
	"unsafe"
)

const EhdrSize = int(unsafe.Sizeof(Ehdr{}))
const ShdrSize = int(unsafe.Sizeof(Shdr{}))
const SymEntrySize = int(unsafe.Sizeof(SymTabEntry{}))

// elf header的抽象
type Ehdr struct {
	Ident     [16]uint8
	Type      uint16
	Machine   uint16
	Version   uint32
	Entry     uint64
	PhOff     uint64
	ShOff     uint64 //  section header table’s file offset in bytes.
	Flags     uint32
	EhSize    uint16
	PhEntSize uint16
	PhNum     uint16
	ShEntSize uint16
	ShNum     uint16 // section header table中entry的数量，每一个entry描述每个section的基本信息
	// 即Section Header String Table在Section Header Table中的索引
	// Suppose the value of e_shstrndx is 5, meaning that the 5th entry in the section header table is the Section Header String Table.
	// Other sections will then use their sh_name field as an offset into this string table to determine their names
	ShStrndx uint16
}

// Section Header，是section header table的每一个entry
type Shdr struct {
	Name   uint32 //在section header string table section中的索引
	Type   uint32
	Flags  uint64
	Addr   uint64
	Offset uint64 // section在file中起始的byte的索引
	Size   uint64 // section的byte数量
	Link   uint32
	Info   uint32 // 含有extra information, 不同section含义不同，对于symtab section，记录第一个非local的变量
	// 所有索引小于 sh_info 的符号都是局部符号，而索引大于或等于 sh_info 的符号都是全局符号
	AddrAlign uint64
	EntSize   uint64
}

// Symbol Table Entry
type SymTabEntry struct {
	// C代码中的全局变量名/函数名, 作为符号的键值
	// 在ELF格式中, 同样是字符串在.strtab节中的位置
	// 局部的自动变量名不是符号
	// 例如，如果symbol的st_name的值为32，
	// 这个值是Symbol Table header中sh_link字段指定的section中的偏移量，这个section也是string table，用来确定符号的名称
	Name  uint32
	Info  uint8
	Other uint8
	Shndx uint16 // 符号在程序中的地址
	Val   uint64 // 变量/函数的长度
	Size  uint64
}

// 通过sh_name的offset得到section的name
func GetSectionNameByOffset(strTab []byte, sh_name uint32) string {
	// 找到下一个字符串`0`的index, Index([]str1, []str2)，找到str2字符串在str1中第一次出现的位置
	next := uint32(bytes.Index(strTab[sh_name:], []byte{0}))
	return string(strTab[sh_name : sh_name+next])
}
