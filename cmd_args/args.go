package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args) // https://pkg.go.dev/fmt
	fmt.Println("Path:", os.Args[0])
	fmt.Println("Number of items inside os.Args:", len(os.Args))
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%d-th argument: %s\n", i, os.Args[i])
	}
}
