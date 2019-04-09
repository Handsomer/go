package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)
/*
go语言语法:
初始化	 :=
输出  	 fmt.Println
格式化 	 fmt.Sprintf() //重定向到fmt Println输出
时间用法 time.Now()//当前时间
	 time.Since(//Start对象)
遍历	 for ;;;{} //;;;内容皆可为空
字符串拼接	string.Join(list_obj, " ")

*/

func main(){
	fmt.Println(strings.Join(os.Args[:1]," "))
	join_start := time.Now()
	for index, arg := range os.Args[:]{
		fmt.Println(index, arg)
	}
	//fmt.Println(join_start)
	fmt.Println(time.Since(join_start).Seconds())
//	time.Sleep(time.Duration(2)*time.Second)
	fmt.Println(fmt.Sprintf("%.9fs", time.Since(join_start).Seconds()))
	origin_start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(time.Since(origin_start).Seconds())
	fmt.Println(fmt.Sprintf("%.9fs", time.Since(origin_start).Seconds()))
	fmt.Println(s)

//	fmt.Println(os.Args[:1])
}
