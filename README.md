# go-struct-to-ts-interface

Convert Go lang struct types to TypeScript interfaces

### Running

```
go run .
```

### Example Usage

Go to the following sections in `main.go`

- `Struct Types Declaration` - Declare all the struct types you want to convert into TS Interfaces
- `Struct Types to Process` - Add all the struct types you want to convert into TS interfaces in the `interfaceList` array

Example:

```Go
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
```

Run the app with the command:

```
go run .
```

Example structs `Demo` & `Baz` above should produce the following TS interfaces:

```TypeScript
export interface Demo {
    foo?: string
    bar: number
    baz: Baz
}

export interface Baz {
    first_field: boolean
    second_field: string
}
```

### Built with

- [Go](https://go.dev/)
- [Bel](https://pkg.go.dev/github.com/32leaves/bel)

### Contributors

- Edwin Carl Flores
