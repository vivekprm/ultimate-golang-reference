package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	ID   int
	Name string
}

func main() {
	u, err := retrieveUser(1234)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", *u)
}

func retrieveUser(id int) (*user, error) {
	r, err := getUser(id)
	if err != nil {
		return nil, err
	}
	// not very readable, instead of pointer use normal variable
	var u *user
	err = json.Unmarshal([]byte(r), &u)
	return u, err
}

func getUser(id int) (string, error) {
	response := fmt.Sprintf(`{"id": %d, "name": "sally"}`, id)
	return response, nil
}
