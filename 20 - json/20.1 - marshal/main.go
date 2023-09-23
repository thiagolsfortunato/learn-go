package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Breed string `json:"breed"`
	Age uint `json:"age"`
}

func main() {
	// json.Marshal() to json
	// json.Unmarshal() to struct or map

	c := dog{"Brisa", "Vira-Lata", 6}
	cJSON, error := json.Marshal(c)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(cJSON)
	fmt.Println(bytes.NewBuffer(cJSON))

	c2 := map[string]string {
		"name": "Luna",
		"breed": "shitzu",
		"age": "7",
	}

	c2JSON, error := json.Marshal(c2)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(c2JSON)
	fmt.Println(bytes.NewBuffer(c2JSON))
}