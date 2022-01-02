package main

import "fmt"

type Client struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewClient(username, password *string) *Client {
	c := Client{
		*username,
		*password,
	}

	return &c
}

func main() {
	client := NewClient(nil, nil) //=> panic: runtime error: invalid memory address or nil pointer dereference

	fmt.Println(client)
}
