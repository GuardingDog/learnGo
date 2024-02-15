package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())
	y := v.Elem()
	fmt.Println("type of y:", y.Type())
	fmt.Println("settability of y:", y.CanSet())

	y.SetFloat(7.1)
	fmt.Println("修改后的 x 的值:", x)
}
