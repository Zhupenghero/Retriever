package main

import "fmt"
import "mock"

func main()  {
	var a="lilei"
	funName(a)
	fmt.Println(funName(a))
}
func funName(a interface{}) string {
	return a.(string)
}

func (r Retriever) Post(url string,form map[string]string)string{
	r.Contents=form["contents"]
	fmt.Println("aaa:",r.Contents)
	return "ok"
}