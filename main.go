package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	addr := "https://api.quotable.io/random"

	res, err := http.Get(addr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	decoded := &Quote{}

	err = json.NewDecoder(res.Body).Decode(&decoded)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf(decoded.Content + "\n~ " + decoded.Author)
}

type Quote struct {
	ID      string   `json:"_id"`
	Tags    []string `json:"tags"`
	Content string   `json:"content"`
	Author  string   `json:"author"`
	Length  int      `json:"length"`
}
