package main

import (
	"fmt"
	"github.com/funnylookinhat/easyjson/lib"
	"log"
)

func main() {

	str := `
{
	"f": 123,
	"ff": 123.456,
	"s": "Some string",
	"m": {
		"a": "AAA",
		"b": "BBB",
		"dc": "DCDCDC",
		"dd": {
			"dda": "DDADDADDA",
			"ddb": "DDBDDBDDB",
			"ddc": "DDCDDCDDC"
		}
	},
	"b": false,
	"bb": true,
	"n": [
		1,2,3
	]
}`

	data, err := easyjson.DecodeJson([]byte(str))

	if err != nil {
		log.Fatal(err)
	}

	f1, _ := easyjson.GetFloat(data, "f")
	f2, _ := easyjson.GetFloat(data, "ff")

	fmt.Printf("Two floats: %f , %f \n", f1, f2)

	s1, _ := easyjson.GetString(data, "s")

	fmt.Printf("A string: %s \n", s1)

	m, _ := easyjson.GetMap(data, "m", "dd")

	fmt.Printf("Mapped values: \n")

	for k, v := range m {
		fmt.Printf("  %s : %s \n", k, v)
	}

	b1, _ := easyjson.GetBool(data, "b")
	b2, _ := easyjson.GetBool(data, "bb")

	fmt.Printf("Booleans: %t , %t \n", b1, b2)

	n, _ := easyjson.GetSlice(data, "n")

	fmt.Printf("Slice of floats: %f, %f, %f \n", n[0], n[1], n[2])
}
