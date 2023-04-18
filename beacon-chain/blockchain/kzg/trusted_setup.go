package kzg

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed trusted_setup.json
	embeddedTrustedSetup []byte // 1.2Mb
)

func main() {
	fmt.Println("vim-go")
}
