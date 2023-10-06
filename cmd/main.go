package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func main() {
	str := "㎡"
	nfc := norm.NFC.String(str)   // NFC 形式
	nfkc := norm.NFKC.String(str) // NFKC 形式

	fmt.Printf("NFC String RESULT IS: % s\n", nfc)
	fmt.Printf("NFKC String RESULT IS: % s\n", nfkc)
	fmt.Printf("NFC Bytes RESULT IS: % x\n", nfc)
	fmt.Printf("NFKC Bytes RESULT IS: % x\n", nfkc)
	fmt.Printf("NFC Rune RESULT IS: % +q\n", nfc)
	fmt.Printf("NFKC Rune RESULT IS: % +q\n", nfkc)

	// 检查是否相等
	fmt.Println(nfc == nfkc) // 输出：false
}
