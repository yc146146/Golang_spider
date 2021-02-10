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

//爬取单个页面的函数
func SpiderPage(i int, page chan int){
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn="+strconv.Itoa((i-1)*50)

	result, err := HttpGet(url)

	if err != nil {
		fmt.Println("HttpGet err:", err)
		return
	}

	//fmt.Println("result=", result)

	//将读到的整网页数据保存成一个文件
	f,err := os.Create("第 "+strconv.Itoa(i)+" 页"+".html")
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	f.WriteString(result)
	//保存好一个文件，关闭一个文件
	f.Close()

	//与主构成完成同步
	page <- i
}


func working2(start, end int){
	fmt.Printf("正在爬取 %d 页到 %d 页....\n", start, end)

	page := make(chan int)

	//循环爬取每一页数据
	for i:=start;i<=end;i++{
		go SpiderPage(i, page)

	}

	for i:=start;i<=end;i++ {
		fmt.Printf("第 %d 个网页爬取完成\n",<-page)
	}

}



func main() {
	//指定爬取起始、终止页面
	var start,end int
	fmt.Print("请输入起始页（》=1）：")
	fmt.Scan(&start)
	fmt.Print("请输入爬虫终止页（》=start）:")
	fmt.Scan(&end)

	working2(start, end)
}
