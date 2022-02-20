package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJsonObject(t *testing.T) {

	c := Customer{"Danu", "Budi", "Raharjo", 25, false}

	bytes, err := json.Marshal(c)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

}
