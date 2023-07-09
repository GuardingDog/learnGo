package main

import (
	"fmt"
	"learngo/pkg/data"
)

func main() {
	fmt.Println("GobDemo ...")
	data.GobDemo()

	fmt.Println("BasicDemo ...")
	data.BasicDemo()

	fmt.Println("EncodeDecodeDemo ...")
	data.EncodeDecodeDemo()

	fmt.Println("InterfaceDemo ...")
	data.InterfaceDemo()
}
