package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
Create a struct that will be encoded to JSON and decoded back to a struct.
Json Primitive Types: https://golang.org/pkg/encoding/json/#Marshal
*/
type Person struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hoobies    []string
}

func TestJsonArrayEncode(t *testing.T) {

	p := Person{
		FirstName:  "Danu",
		MiddleName: "Budi",
		LastName:   "Raharjo",
		Age:        25,
		Married:    false,
		Hoobies:    []string{"sport", "music", "movie"},
	}

	bytes, err := json.Marshal(p)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

}

func TestJsonArrayDecode(t *testing.T) {

	jsonString := `{"FirstName":"Danu","MiddleName":"Budi","LastName":"Raharjo","Age":25,"Married":false,"Hoobies":["sport","music","movie"]}`

	jsonBytes := []byte(jsonString)

	p := Person{}

	err := json.Unmarshal(jsonBytes, &p)

	if err != nil {
		panic(err)
	}

	fmt.Println(p)
	fmt.Println(p.FirstName)
	fmt.Println(p.MiddleName)
	fmt.Println(p.LastName)
	fmt.Println(p.Age)
	fmt.Println(p.Married)
	fmt.Println(p.Hoobies)

}

/*
Create a struct that will be encoded to JSON and decoded back to a struct.
Json Object Types: https://golang.org/pkg/encoding/json/#Marshal
Example: https://golang.org/pkg/encoding/json/#Marshal
Example using json complex
*/

type Address struct {
	Street  string
	City    string
	State   string
	Zip     string
	Country string
}

type Pelanggan struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hoobies    []string
	Addresses  []Address
}

func TestJsonArrayComplexEncode(t *testing.T) {

	p := Pelanggan{
		FirstName:  "Danu",
		MiddleName: "Budi",
		LastName:   "Raharjo",
		Age:        25,
		Married:    false,
		Hoobies:    []string{"sport", "music", "movie"},
		Addresses: []Address{
			{
				Street:  "Jl. Raya",
				City:    "Bandung",
				State:   "Jawa Barat",
				Zip:     "40132",
				Country: "Indonesia",
			},
			{
				Street:  "Jl. Raya",
				City:    "Bandung",
				State:   "Jawa Barat",
				Zip:     "40132",
				Country: "Indonesia",
			},
		},
	}

	bytes, err := json.Marshal(p)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

}

func TestJsonArrayComplexDecode(t *testing.T) {

	jsonString := `{
		"FirstName":"Danu",
		"MiddleName":"Budi",
		"LastName":"Raharjo",
		"Age":25,
		"Married":false,
		"Hoobies":["sport","music","movie"],
		"Addresses":[
			{
				"Street":"Jl. Raya",
				"City":"Bandung",
				"State":"Jawa Barat",
				"Zip":"40132",
				"Country":"Indonesia"
			},
			{
				"Street":"Jl. Raya",
				"City":"Bandung",
				"State":"Jawa Barat",
				"Zip":"40132",
				"Country":"Indonesia"
			}
		]
	}`

	jsonBytes := []byte(jsonString)

	p := Pelanggan{}

	err := json.Unmarshal(jsonBytes, &p)

	if err != nil {
		panic(err)
	}

	fmt.Println(p)                      // print struct
	fmt.Println(p.FirstName)            // print struct field FirstName
	fmt.Println(p.MiddleName)           // print struct field MiddleName
	fmt.Println(p.LastName)             // print struct field LastName
	fmt.Println(p.Age)                  // print struct field Age
	fmt.Println(p.Married)              // print struct field Married
	fmt.Println(p.Hoobies)              // print struct field Hoobies
	fmt.Println(p.Addresses)            // print struct field Addresses
	fmt.Println(p.Addresses[0].Street)  // print struct field Addresses[0].Street
	fmt.Println(p.Addresses[0].City)    // print struct field Addresses[0].City
	fmt.Println(p.Addresses[0].State)   // print struct field Addresses[0].State
	fmt.Println(p.Addresses[0].Zip)     // print struct field Addresses[0].Zip
	fmt.Println(p.Addresses[0].Country) // print struct field Addresses[0].Country
	fmt.Println(p.Addresses[1].Street)  // print struct field Addresses[1].Street
	fmt.Println(p.Addresses[1].City)    // print struct field Addresses[1].City
	fmt.Println(p.Addresses[1].State)   // print struct field Addresses[1].State
	fmt.Println(p.Addresses[1].Zip)     // print struct field Addresses[1].Zip
	fmt.Println(p.Addresses[1].Country) // print struct field Addresses[1].Country

}

func TestOnlyJsonComplexDecode(t *testing.T) {

	jsonString := `[
			{
				"Street":"Jl. Raya",
				"City":"Bandung",
				"State":"Jawa Barat",
				"Zip":"40132",
				"Country":"Indonesia"
			},
			{
				"Street":"Jl. Raya",
				"City":"Bandung",
				"State":"Jawa Barat",
				"Zip":"40132",
				"Country":"Indonesia"
			}
	]`

	jsonBytes := []byte(jsonString)

	var addresses []Address

	err := json.Unmarshal(jsonBytes, &addresses)

	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)            // print struct
	fmt.Println(addresses[0].Street)  // print struct field Addresses[0].Street
	fmt.Println(addresses[0].City)    // print struct field Addresses[0].City
	fmt.Println(addresses[0].State)   // print struct field Addresses[0].State
	fmt.Println(addresses[0].Zip)     // print struct field Addresses[0].Zip
	fmt.Println(addresses[0].Country) // print struct field Addresses[0].Country
	fmt.Println(addresses[1].Street)  // print struct field Addresses[1].Street
	fmt.Println(addresses[1].City)    // print struct field Addresses[1].City
	fmt.Println(addresses[1].State)   // print struct field Addresses[1].State
	fmt.Println(addresses[1].Zip)     // print struct field Addresses[1].Zip
	fmt.Println(addresses[1].Country) // print struct field Addresses[1].Country

}

func TestOnlyJsonComplexEncode(t *testing.T) {

	addresses := []Address{
		{
			Street:  "Jl. Raya",
			City:    "Bandung",
			State:   "Jawa Barat",
			Zip:     "40132",
			Country: "Indonesia",
		},
		{
			Street:  "Jl. Raya",
			City:    "Bandung",
			State:   "Jawa Barat",
			Zip:     "40132",
			Country: "Indonesia",
		},
	}

	jsonBytes, err := json.Marshal(addresses)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))

}
