package linker

import "debug/elf"

// .o文件
type ObjectFile struct {
	// 继承, 可以直接通过.访问InputFile中的成员
	InputFile
	// .symtab section - 符号表(symbol table), 每一个表项记录一个符号的信息
	SymbolTableSection *Shdr
}

func NewObjectFile(file *File) *ObjectFile {
	obj := &ObjectFile{
		InputFile: NewInputFile(file),
	}
	return obj
}

// 完成objectFile文件解析
func (o *ObjectFile) ParseFile() {
	o.SymbolTableSection = o.FindSectionByType(uint32(elf.SHT_SYMTAB))

	if o.SymbolTableSection != nil {
		o.FirstGlobalSymbolIdx = int64(o.SymbolTableSection.Info)
		o.InitSymbolTable(o.SymbolTableSection)
		// o.SymbolTableSection.Link用来存储symbol table所使用的string table的索引
		o.SymStrTab = o.GetSectionBytesByIdx(uint64(o.SymbolTableSection.Link))

	}

	// 符号解析
}
