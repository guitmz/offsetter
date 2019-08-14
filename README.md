# offsetter
[![Go Report Card](https://goreportcard.com/badge/github.com/guitmz/offsetter)](https://goreportcard.com/report/github.com/guitmz/offsetter)

Small package to convert between file offsets and virtual addresses (partialy inspired by some code of [pwnlib](https://github.com/Gallopsled/pwntools/tree/dev/pwnlib)). For the moment, only `ELF` files are supported.

# Install
Install with `go get -u github.com/guitmz/offsetter`

# Example
```
$ cat example.go
```

 ```go
 package main

import (
	"debug/elf"
	"fmt"

	"github.com/guitmz/offsetter"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	targetFile, err := elf.Open("/bin/ls")
	check(err)

	loadAddress, err := offsetter.GetLoadAddress(targetFile)
	check(err)

	hexOffset, uintOffset, err := offsetter.VaddrToOffset(targetFile, loadAddress, targetFile.Entry)
	check(err)
	fmt.Printf("Offset: %s (%d)\n", hexOffset, uintOffset)

	hexVaddr, uintVaddr, err := offsetter.OffsetToVaddr(targetFile, uintOffset)
	check(err)
	fmt.Printf("Vaddr: %s (%d)\n", hexVaddr, uintVaddr)
}
```

```
$ go run example.go
Offset: 0x42d4 (17108)
Vaddr: 0x4042d4 (4211412)
```
