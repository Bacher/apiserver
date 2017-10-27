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
)

var rpcClient *rpc.Client

var invalidParams = errors.New("invalid params")

func main() {
	dbpool.InitDb()

	rpcClient = rpc.NewClient(func(apiName string, body []byte) ([]byte, error) {
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

			jsonResponse, err := json.Marshal(response)

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

	log.Println("Server is started.")

	addSigTermHandler()
}
