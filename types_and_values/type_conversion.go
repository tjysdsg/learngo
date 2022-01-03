package main

import (
	"fmt"
	"strconv"
)

func main() {
	speed := 100 // int
	force := 2.5 // float64

	// golang doesn't support implicit type conversion
	// speed = speed * force

	speed = speed * int(force)
	fmt.Println(speed)

	var f32 float32 = 1
	var f64 float64 = 2
	fmt.Println(float64(f32) * f64)

	// ascii code to string
	num := 65
	fmt.Println("ASCII code 65 is:", string(num))
	fmt.Println("int to string:", strconv.Itoa(num))

	// bytes to string
	fmt.Println(string([]byte{104, 105}))
}
