package getInitialData

import "apiserver/server/structs"

var Name = "getInitialData"

type Params struct {
	Name string `json:"name"`
}

type Response struct {
	FuckName string `json:"fuckName"`
}

func Do(u *structs.User, p Params) *Response {
	return &Response{p.Name + "-gay " + u.Name}
}
