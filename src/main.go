package main

import (
	"mock"
	"real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string 
}
type Poster interface {
	Post(url string,form map[string]string)string
}
/*
func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func upload(p Poster) {
	p.Post("http://www.imooc.com", map[string]string{"name":"commer", "course":"golang"})
}
*/
type RetrieverPoster interface {
	Retriever
	Poster
}

const abc  = "http://www.imooc.com"

func seession(s RetrieverPoster) string {
	s.Post(abc,map[string]string{"contents":"another fake imooc.com"})
	return s.Get(abc)
}

func main(){
	var r Retriever
	retriever:= &mock.Retriever{"this is a fake imooc.com"}

	r=retriever
	inspect(r)
	r= &real.Retriever{UserAgent:"hello", TimeOut:time.Minute*10}
	inspect(r)

	//type assertion 类型断言
	if realRetriever,ok:=r.(*real.Retriever);ok{
		fmt.Println(realRetriever.TimeOut)
	}else {
		fmt.Println("not a real Retriener")
	}
	fmt.Println("try a session")
	fmt.Println(seession(retriever))

}
func inspect(r Retriever){
	fmt.Printf("%T %v\n",r,r)

	switch v:=r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:",v.Contents)
	case *real.Retriever:{
		fmt.Println("Useragent:",v.UserAgent)
	}
	}
}
