package main

import (
	"apiserver/server/structs"
	"errors"
	"frontserver/dbpool"
)

var NotFound = errors.New("not found")

func getUser(id uint64) (*structs.User, error) {
	con, err := dbpool.GetConnection()

	if err != nil {
		return nil, err
	}

	rows, err := con.Query("SELECT name FROM users WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		var name string
		rows.Scan(&name)
		return &structs.User{id, name}, nil
	}

	return nil, NotFound
}
