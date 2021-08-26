package main

import "fmt"

func swap(a *int, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}

func main() {
	var (
		a = 1
		b = 2
	)
	swap(&a,&b)
	fmt.Println("a=",a)
	fmt.Println("b=",b)

	var p *int
	p = &a

	var pp **int
	pp = &p

	var x *int = &b
	fmt.Println("x=", x)
	fmt.Println("p=",p)
	fmt.Println("pp=",pp)


}