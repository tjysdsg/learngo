/*
https://go.dev/doc/tutorial/getting-started
https://go.dev/blog/using-go-modules
*/
package main

import "fmt"
import "learngo/mylib" // see go.mod
import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
	fmt.Println(mylib.Exported())
}
