package main

import (
	"apiserver/server/api/getInitialData"
	"encoding/json"
	"errors"
	"frontserver/dbpool"
	"frontserver/proto"
	"github.com/golang/protobuf/proto"
	"gorpc/rpc"
	"log"
	"os"
)

type Response struct {
	Response  interface{} `json:"response"`
	ApiServer string      `json:"apiServer"`
}

var rpcClient *rpc.Client

var invalidParams = errors.New("invalid params")

var serverName = getServerName()

func main() {
	dbpool.InitDb()

	addr := os.Getenv("RPC_ADDR")

	if addr == "" {
		addr = "localhost:9999"
	}

	rpcClient = rpc.NewClient(addr, func(apiName string, body []byte) ([]byte, error) {
		switch apiName {
		case "apiCall":
			var apiCallStruct pb.ApiCall
			err := proto.Unmarshal(body, &apiCallStruct)

			if err != nil {
				return nil, err
			}

			user, err := getUser(apiCallStruct.UserId)

			if err != nil {
				return nil, err
			}

			var objMap []*json.RawMessage
			err = json.Unmarshal(apiCallStruct.Params, &objMap)

			if err != nil {
				return nil, err
			}

			var params getInitialData.Params

			err = json.Unmarshal(*objMap[1], &params)

			if err != nil {
				return nil, err
			}

			response := getInitialData.Do(user, params)

			jsonResponse, err := json.Marshal(&Response{response, serverName})

			if err != nil {
				return nil, err
			}

			return jsonResponse, nil
		}

		log.Printf("Unknown apiMethod for parse %s\n", apiName)
		return nil, rpc.ApiNotFound
	})

	err := rpcClient.Connect()

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Server is started [%s].\n", serverName)

	addSigTermHandler()
}

func getServerName() string {
	name := os.Getenv("SERVER_NAME")

	if name != "" {
		return name
	}

	return "Unknown server"
}
