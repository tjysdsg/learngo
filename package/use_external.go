/*
https://go.dev/doc/tutorial/getting-started
https://go.dev/blog/using-go-modules
*/
package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
}
