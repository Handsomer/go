package main

/*
文件io
生成映射	count := map(map[string]int) 
参数切片 	os.Args[1:]
读取文件	iotile.ReadFile(FileName)
文件切片	strings.Split(string(data), "\n")
字典操作	count += 1
*/
import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main(){
	count := make(map[string]int)
	for _, filename := range os.Args[1:]{
		data, err := ioutil.ReadFile(filename)
		if err != nil{
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n"){
			if count[line] > 0 {
				fmt.Println(line, filename)
			}
			count[line] ++
		}
	}
}
