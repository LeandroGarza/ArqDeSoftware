package main

type Credentials struct {
	user     string `json: "user"`
	password string `json: "password"`
}

type token struct {
	Token string `json: "token"`
}
