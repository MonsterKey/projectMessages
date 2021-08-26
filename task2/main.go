//+build ignore

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//新建一个Go项目，使用go mod进行包。完成一个命令行程序，接受
	//h,type参数。输入h时候，命令行显示使用帮助；type值为1时输出一句
	//唐诗，为2时输出一句宋词。
	var h string
	var tp int
	for {
		fmt.Scanf("%s",&h)
		if h=="h" {
			fmt.Println("输入1===》输出一句唐诗\n输入2===》输出一句宋词")
			continue
		}else if h=="1" || h=="2"{
			i,_ := strconv.Atoi(h)
			tp = i
			if tp==1 {
				fmt.Println("白日依山尽，黄河入海流。")
				break
			} else if tp==2{
				fmt.Println("明月几时有？把酒问青天。")
				break
			}
		}else {
			fmt.Println("请输入h查看提示！")
		}
	}

}
