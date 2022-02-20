package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestJsonStreamDecoderFromFile(t *testing.T) {

	reader, err := os.Open("person.json")

	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(reader)

	for {
		var c Customer
		err := decoder.Decode(&c)

		if err != nil {
			break
		}

		fmt.Println(c)
	}

}

func TestJsonStreamDecoderFromInputStream(t *testing.T) {

	jsonStream := `
				{
					"firstName":"Danu",
					"middleName":"Budi",
					"lastName":"Raharjo",
					"age":25,
					"isMarried":false
				}
				{
					"firstName":"John",
					"middleName":"Doe",
					"lastName":"Smith",
					"age":25,
					"isMarried":false
				}
				{
					"firstName":"Jane",
					"middleName":"Doe",
					"lastName":"Smith",
					"age":25,
					"isMarried":false
				}`

	decoder := json.NewDecoder(strings.NewReader(jsonStream))

	for {
		var c Customer
		err := decoder.Decode(&c)

		if err != nil {
			break
		}

		fmt.Printf("%s %s %s\n", c.FirstName, c.MiddleName, c.LastName)
	}

}
