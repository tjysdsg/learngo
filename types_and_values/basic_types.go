package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var speed1 int = 100
	var speed2 = 100
	speed3 := 100

	var (
		fp1 float32 = 3.1415
		fp2         = 3.1415
	)

	var s1 = "a string"

	var _, _ bool
	var b1, b2 = true, false

	fmt.Println("Integer values:", speed1, speed2, speed3)
	fmt.Println("Default int size:", unsafe.Sizeof(speed1), "bytes") // However: https://pkg.go.dev/builtin#int
	fmt.Println("Float values:", fp1)
	fmt.Println("Default float type:", reflect.TypeOf(fp2))

	fmt.Printf("s1 = %q\n", s1)

	fmt.Println("Bool value b1 and b2:", b1, b2)
	// swap values
	b1, b2 = b2, b1
	fmt.Println("Swapped b1 and b2:", b1, b2)

	{
		var s1 = true
		fmt.Println("Redeclaration of s1 shadows previous declaration:", s1)
	}
}
