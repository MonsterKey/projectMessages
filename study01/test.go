
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replaceSpace( s string ) string {
	// write code here
	return strings.Replace(s," ","%20",-1)
}

func replaceSpace1(s string) string {
	ss:=[]byte(s)
	//fmt.Println(len(ss))
	for i:=0;i<len(ss);i++{
		if ss[i]==' '{
			n:=len(ss)
			ss=append(ss,ss[n-2:]...)
			for k:=n-1;k>i+2;k--{
				ss[k]=ss[k-2]
			}
			ss[i]='%'
			ss[i+1]='2'
			ss[i+2]='0'
		}
	}
	return string(ss)
}

//func replaceSpace2(s string) string {
//	str:= make([]byte,len(s))
//	for i:=0;i<len(s);i++{
//		if s[i]==' ' {
//			str = append([]byte,"%20")
//		}else{
//
//		}
//	}
//}

func main() {
	in := bufio.NewReader(os.Stdin)
	s, err := in.ReadString('\n')
	if err!=nil {
		fmt.Println("无法输入！")
	}
	fmt.Println(replaceSpace1(s))
}