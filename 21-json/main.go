package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var text = []byte(`{
	"first_name": "San",
	"last_name": "Zhang",
	"age": 10,
	"admin": false
}`)

type user struct {
	FirstName string `json:"first_name"`
	LastName  string
	Age       uint

	IsAdmin bool `json:"admin,omitempty"`
}

func parse() {
	u := user{}
	if err := json.Unmarshal(text, &u); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("user: %#v\n", u)
}

func dump(u user) {

	if bs, err := json.MarshalIndent(u, "", "  "); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("json: %s\n", bs)
	}
}

func main() {
	u := user{
		FirstName: "Si",
		LastName:  "Li",
		Age:       30,
		IsAdmin:   true,
	}

	parse()
	dump(u)
	dump(user{})
}
