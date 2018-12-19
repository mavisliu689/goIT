package main

import (
	"fmt"
	"time"

	"mavis.com/IT-chinese/retriever/mock"
	"mavis.com/IT-chinese/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com",
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetriverPoster interface {
	Retriever
	Poster
	// PosterConnect(host string)
}

const url = "http://www.immocc.com"

func session(s RetriverPoster) string {
	ok := s.Post(url, map[string]string{
		"contents": "another faked immocc.com",
	})
	fmt.Println("in session", ok)
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)

	fmt.Print("type switch")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever
	r = &mock.Retriever{"this is a fake immoc.com"}
	fmt.Printf("%T %v \n", r, r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v \n", r, r)

	retriever := mock.Retriever{"this is a fake immoc.com"}
	inspect(r)

	//Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	// fmt.Println(download(r))
	fmt.Println("try a session")
	fmt.Println(session(&retriever))
}
