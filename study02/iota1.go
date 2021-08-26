// +build ignore

package main
import "fmt"

func foo1(i, j int) (r1,r2 int) {
	fmt.Println("r1", r1)
	fmt.Println("r2", r2)

	r1 = 100
	r2 = 200
	return
}

func main() {
		//fmt.Println(foo1(1, 2))
	println(foo1(1, 2))
}
