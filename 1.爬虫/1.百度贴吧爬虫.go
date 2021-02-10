package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string) (result string, err error){

	resp, err1 := http.Get(url)

	if err1 != nil {
		//将封装函数内部的错误，传出给调用者
		err = err1
		return
	}

	defer resp.Body.Close()

	buf := make([]byte, 4096)
	//循环读取 网页数据  传出给调用者
	for  {
		n,err2 := resp.Body.Read(buf)
		if n==0{
			fmt.Println("Read web success")
			break
		}
		if err2 != nil && err2 != io.EOF{
			err = err2
			return
		}
		//累加每一次循环读到的buf数据， 存入result一次性返回
		result += string(buf[:n])
	}
	return
}

func working(start, end int){
	fmt.Printf("正在爬取%d页到%d页....\n", start, end)

	//循环爬取每一页数据
	for i:=start;i<=end;i++{
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn="+strconv.Itoa((i-1)*50)

		result, err := HttpGet(url)

		if err != nil {
			fmt.Println("HttpGet err:", err)
			continue
		}

		//fmt.Println("result=", result)

		//将读到的整网页数据保存成一个文件
		f,err := os.Create("第 "+strconv.Itoa(i)+" 页"+".html")
		if err != nil {
			fmt.Println("Create err:", err)
			continue
		}
		f.WriteString(result)
		//保存好一个文件，关闭一个文件
		f.Close()

	}
}



func main() {
	//指定爬取起始、终止页面
	var start,end int
	fmt.Print("请输入起始页（》=1）：")
	fmt.Scan(&start)
	fmt.Print("请输入爬虫终止页（》=start）:")
	fmt.Scan(&end)

	working(start, end)
}
