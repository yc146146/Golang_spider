package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "abc a7c mfc cat 8ca azc cba"

	//``  表示使用原生字符串
	ret := regexp.MustCompile(`a.c`)

	//提取需要信息
	alls := ret.FindAllStringSubmatch(str,-1)
	fmt.Println("alls:",alls)
}
