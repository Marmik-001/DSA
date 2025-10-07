package concepts

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	IsRemote bool   `json:"remote"`
	Address
}

type Address struct {
	City    string `json:"myCity"`
	Country string `json:"myCountry"`
}

func StructDemo() {
	addressEmployee1 := Address{
		City:    "abd",
		Country: "Ind",
	}
	employee1 := Employee{
		Name:     "marmik",
		Age:      20,
		IsRemote: false,
		Address:  addressEmployee1,
	}

	fmt.Println(employee1)

	jsonData , _ := json.Marshal(employee1)
	fmt.Println(string(jsonData))

}
