package main

import (
	"fmt"
	"os"
)

func main() {
	byte, err := os.ReadFile("examples/custom-binary-le.bin")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(byte)
	}
}
