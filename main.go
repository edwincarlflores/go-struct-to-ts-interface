package main

import (
	"ts-utils/go-struct-to-ts-interface/utils"
)

// SECTION: Struct Types Declaration
// Declare your struct types here
type Demo struct {
	Foo string `json:"foo,omitempty"`
	Bar uint32 `json:"bar"`
	Baz Baz    `json:"baz"`
}

type Baz struct {
	FirstField  bool    `json:"first_field"`
	SecondField *string `json:"second_field"`
}

func main() {
	// SECTION: Struct Types to Process
	// List all the struct types you want to convert into TS interface here
	interfaceList := []interface{}{
		Demo{},
		Baz{},
	}

	utils.ConvertToTSInterface(interfaceList)
}
