package getInitialData

import "fmt"

var Name = "getInitialData"

type Params struct {
	Name string `json:"name"`
}

type Response struct {
	FuckName string `json:"fuckName"`
}

func Do(p Params) *Response {

	fmt.Println(p)

	return &Response{p.Name + "-gay"}
}