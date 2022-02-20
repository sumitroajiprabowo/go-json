package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	ImageURL string  `json:"image_url"`
}

func TestJsonTagEncode(t *testing.T) {
	p := Product{
		Id:       1,
		Name:     "Product 1",
		Price:    10.99,
		ImageURL: "http://example.com/product1.jpg",
	}
	bytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{
		"id":1,
		"name":"Product 1",
		"price":10.99,
		"image_url":"http://example.com/product1.jpg"
		}`
	jsonBytes := []byte(jsonString)
	p := Product{}
	err := json.Unmarshal(jsonBytes, &p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
	fmt.Println(p.Id)
	fmt.Println(p.Name)
	fmt.Println(p.Price)
	fmt.Println(p.ImageURL)
}
