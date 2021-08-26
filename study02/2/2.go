package main

import "fmt"

func arrayT(array []int)  {
	for _, i2 := range array {
		fmt.Println("value=",i2)
	}
	array[0] = 100
}

func main() {
	myArray := []int{1,2,3,4}
	for i := 0; i < len(myArray); i++ {
		fmt.Printf("myArray[%d]: %d\n", i, myArray[i])
	}
	arrayT(myArray)
	fmt.Println("=====")
	for i := 0; i < len(myArray); i++ {
		fmt.Printf("myArray[%d]: %d\n", i, myArray[i])
	}
}
