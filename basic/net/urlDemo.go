package main

import (
	"net/url"
	"fmt"
)

/**
 * url包解析URL并实现了查询的逸码
 */
func main() {
	testParseQuery()
	testParseRequestQuery()
}


func testParseQuery() {
	s := "keyword=开倒车&a=x&b=b&c=c&gourl=http://search.bilibili.com/all?keyword=开倒车"
	val,err:=url.ParseQuery(s)
	if err!=nil {

	}
	fmt.Println(val.Get("keyword"))
	for i,v:=range val {
		fmt.Println(i,v)
	}
}
func testParseRequestQuery() {
	s:="http://search.bilibili.com/all?keyword=开倒车"
	url,err:=url.ParseRequestURI(s)
	if err != nil {

	}
	fmt.Println(url.RawQuery)
}