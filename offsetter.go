package offsetter

import (
	"debug/elf"
	"fmt"
)

// GetLoadAddress returns input ELF file Load Address (in memory)
func GetLoadAddress(f *elf.File) (uint64, error) {
	for _, p := range f.Progs {
		if p.Type == elf.PT_LOAD && p.Flags&elf.PF_X != 0 {
			return p.Vaddr, nil
		}
	}
	return 0, fmt.Errorf("unknown load address")
}

// VaddrToOffset converts a virtual address to a file offset (returns string and uint64 representations of the resulting file offset)
func VaddrToOffset(f *elf.File, load, addr uint64) (string, uint64, error) {
	for _, segment := range f.Progs {
		addr = (addr - load) + load
		offset := addr - segment.Vaddr
		result := segment.Off + offset
		return fmt.Sprintf("%#.4x", result), result, nil
	}
	return "", 0, fmt.Errorf("unable to convert virtual address to offset")
}

// OffsetToVaddr converts a file offset to a virtual address (returns string and uint64 representations of the resulting virtual address)
func OffsetToVaddr(f *elf.File, offset uint64) (string, uint64, error) {
	loadAddressFixup := uint64(0) // (self.address - self.load_addr)

	for _, segment := range f.Progs {
		if segment.Off <= offset && offset <= (segment.Off+segment.Filesz) {
			delta := offset - segment.Off
			result := segment.Vaddr + delta + loadAddressFixup
			return fmt.Sprintf("%#.4x", result), result, nil
		}
	}
	return "", 0, fmt.Errorf("unable to convert offset to virtual address")
}
