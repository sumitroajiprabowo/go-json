package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMapEncode(t *testing.T) {

	//create map object for json encode
	data := map[string]interface{}{
		"id":    "1",
		"name":  "Product 3",
		"price": 10.99,
	}

	//json encode
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	// print json string
	fmt.Println(string(bytes))

}

func TestJsonMapDecode(t *testing.T) {

	//example of json map
	jsonString := `{
		"id":"1",
		"name":"Product 2",
		"price":10.99
		}`

	//convert json string to byte array
	jsonBytes := []byte(jsonString)

	//create map object for json decode
	var result map[string]interface{}

	//json decode
	err := json.Unmarshal(jsonBytes, &result)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])

}
