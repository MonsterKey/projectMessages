//+build ignore

package main

import "fmt"

func main() {
	s := "1234"
	ss := []byte(s)
	fmt.Println(len(ss))
}