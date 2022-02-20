package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {

	bytes, err := json.Marshal(data) // json encode

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

}

func TestEncodeJson(t *testing.T) {

	logJson("hello")                                // json string
	logJson(123)                                    // json number
	logJson(123.456)                                // json number
	logJson(true)                                   // json boolean
	logJson(nil)                                    // json null
	logJson([]int{1, 2, 3})                         // json array
	logJson(map[string]int{"a": 1, "b": 2})         // json object
	logJson(map[string]interface{}{"a": 1, "b": 2}) // json object

}
