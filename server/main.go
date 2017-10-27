package main

import (
	"apiserver/server/api/getInitialData"
	"encoding/json"
	"fmt"
	"gorpc/rpc"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := rpc.NewClient(func(apiName string, body []byte) ([]byte, error) {
		switch apiName {
		case "apiCall":
			log.Println(string(body))

			return []byte(`{"result":"OK"}`), nil
		}

		log.Printf("Unknown apiMethod for parse %s\n", apiName)
		return nil, rpc.ApiNotFound
	})

	err := client.Connect()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Server is started.")

	addSigTermHandler()
}

func handler(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fail(w, 400, "Bad Request")
		return
	}

	var objMap []*json.RawMessage

	err = json.Unmarshal(jsonBytes, &objMap)

	if err != nil {
		fmt.Println("Invalid JSON:", jsonBytes)
		fail(w, 400, "Bad Request")
		return
	}

	fmt.Println(objMap)

	var token string
	json.Unmarshal(*objMap[0], &token)

	fmt.Printf("token %s\n", token)

	var params getInitialData.Params

	err = json.Unmarshal(*objMap[1], &params)

	if err != nil {
		fmt.Println("Invalid JSON", err)
		fail(w, 400, "Bad Request")
		return
	}

	response := getInitialData.Do(params)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		fmt.Println(err)
		fail(w, 500, "Internal Server Error")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func fail(w http.ResponseWriter, code int, message string) {
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(code)
	w.Write([]byte(message))
}
