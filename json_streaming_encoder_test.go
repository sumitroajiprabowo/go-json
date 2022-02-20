package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestJsonStreammingEncoderToFile(t *testing.T) {

	writer, err := os.Create("output.json") // create a file

	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(writer) // create a encoder

	c := Customer{"Danu", "Budi", "Raharjo", 25, false} // create a customer

	err = encoder.Encode(c) // encode the customer

	if err != nil {
		panic(err)
	}

	writer.Close()

}
