package main

import (
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

	dJSON := `{"name":"Brisa","breed":"Vira-Lata","age":6}`
	
	var d dog;

	if error := json.Unmarshal([]byte(dJSON), &d) ; error != nil {
		log.Fatal(error)
	}

	fmt.Println(d)
	fmt.Println(d.Name)

	d2JSON := `{"age":"7","breed":"shitzu","name":"Luna"}`

	d2 := make(map[string]string)

	if error := json.Unmarshal([]byte(d2JSON), &d2) ; error != nil {
		log.Fatal(error)
	}

	fmt.Println(d2)
	fmt.Println(d2["name"])
}
