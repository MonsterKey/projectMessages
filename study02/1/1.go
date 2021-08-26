package main

import (
	"fmt"
	"strconv"
)


func f(n int)(r1 int){
	var s string
	for i:=0;i<n;i++ {
		s+="9"
	}
	res,_:=strconv.Atoi(s)
	return res
}
func main(){
	//var n int
	//fmt.Scanf("%d",&n)
	//if res,error := f(n);error == nil {
	//	fmt.Println(res)
	//}
	//fmt.Println(f(n))
	//myArray := [4]int{1,2,3,4}
	//for _,i2 := range myArray {
	//	fmt.Printf("value: %d\n",i2)
	//}
	//var array [6]int
	//array[0] = 1
	//array[1] = 2
	//array[2] = 3
	//array[3] = 4
	//array[4] = 5
	//fmt.Printf("array.length: %d\n", len(array))
	//array[5] = 6
	//fmt.Printf("array.length: %d\n", len(array))
	var n int
	fmt.Scanf("%d",&n)
	array01 := make([]int, n)
	for i := 0; i < n; i++ {
		array01[i] = i
		fmt.Printf("i: %d\n", array01[i])
	}
	for i, i2 := range array01 {
		fmt.Printf("index: %d, value: %d\n",i,i2)
	}
}
