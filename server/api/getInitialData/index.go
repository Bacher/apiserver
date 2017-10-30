package getInitialData

import (
	"apiserver/server/structs"
	"log"
	"time"
)

var Name = "getInitialData"

type Params struct {
	Name string `json:"name"`
}

type Response struct {
	FuckName string `json:"fuckName"`
}

func Do(u *structs.User, p Params) *Response {
	log.Println("api call start")
	time.Sleep(5 * time.Second)
	return &Response{p.Name + "-gay " + u.Name}
}
