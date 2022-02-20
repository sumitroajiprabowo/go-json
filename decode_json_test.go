package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {

	jsonRequest := `{"FirstName":"Danu","MiddleName":"Budi","LastName":"Raharjo","Age":25,"Married":false}` // json string

	var c Customer // json object

	err := json.Unmarshal([]byte(jsonRequest), &c) // json decode

	if err != nil {
		panic(err)
	}

	fmt.Println(c) // print json object

	fmt.Println(c.FirstName) // print json object property value by property name (FirstName)

	fmt.Println(c.MiddleName) // print json object property value by property name (MiddleName)

	fmt.Println(c.LastName) // print json object property value by property name (LastName)

	fmt.Println(c.Age) // print json object property value by property name (Age)

	fmt.Println(c.Married) // print json object property value by property name (Married)

}

func TestDecodeJsonExample(t *testing.T) {

	jsonString := `{"FirstName":"Danu","MiddleName":"Budi","LastName":"Raharjo","Age":25,"Married":false}` // json string

	jsonBytes := []byte(jsonString) // convert json string to json bytes

	customer := Customer{} // create customer object

	err := json.Unmarshal(jsonBytes, &customer) // json decode

	if err != nil {
		panic(err)
	}

	// print json object
	fmt.Println(customer)

	// print json object property value by property name (FirstName)
	fmt.Println(customer.FirstName)

	// print json object property value by property name (MiddleName)
	fmt.Println(customer.MiddleName)

	// print json object property value by property name (LastName)
	fmt.Println(customer.LastName)

	// print json object property value by property name (Age)
	fmt.Println(customer.Age)

	// print json object property value by property name (Married)
	fmt.Println(customer.Married)

}
