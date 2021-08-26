package main

import (
	"fmt"
	"os"
)

func main() {
	for _, a := range os.Args[1:] {
		//fmt.Printf("Argument is %s\n", a)
		if a=="h" {
			fmt.Println("输入1===》输出一句唐诗\t输入2===》输出一句宋词")
		}
		if a=="2" {
			fmt.Println("明月几时有？把酒问青天。")
		}
		if a=="1" {
			fmt.Println("白日依山尽，黄河入海流。")
		}
	}

}