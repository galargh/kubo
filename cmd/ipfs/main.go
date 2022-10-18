// cmd/ipfs implements the primary CLI binary for ipfs
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Fprintln(os.Stderr, "Error: Kubo is not distributed through Chocolatey anymore (https://github.com/ipfs/kubo/issues/9341).")
	} else {
		fmt.Fprintln(os.Stderr, "Error: Kubo is not distributed through Snap anymore (https://github.com/ipfs/kubo/issues/8688).")
	}
	fmt.Fprintln(os.Stderr, "Error: Please download Kubo from https://dist.ipfs.tech/#kubo.")
	os.Exit(1)
}
