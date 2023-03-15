package main

import "fmt"

func main() {
	user := map[string]string{
		"name": "Pedro",
		"surname": "Silva",
	}

	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"first": "João",
			"last": "Pedro",
		},
		"course": {
			"name":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(user2)
	delete(user2, "name")
	fmt.Println(user2)

	user2["signo"] = map[string]string{
		"name": "Thiago",
	}

	fmt.Println(user2)
}
