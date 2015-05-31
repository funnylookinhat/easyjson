package main

import (
	"fmt"
	"github.com/funnylookinhat/_easyjson/lib"
	"log"
)

func main() {

	str := `
[
	{
		"a": "a1",
		"b": "b1",
		"c": "c1"
	},
	{
		"a": "a2",
		"b": "b2",
		"c": "c2"
	},
	{
		"a": "a3",
		"b": "b3",
		"c": "c3"
	}
]
`

	data, err := easyjson.DecodeJson([]byte(str))

	if err != nil {
		log.Fatal(err)
	}

	m, _ := easyjson.GetMap(data, 0)

	fmt.Printf("Mapped values: \n")

	for k, v := range m {
		fmt.Printf("  %s : %s \n", k, v)
	}

	s1, _ := easyjson.GetString(data, 0, "a")

	fmt.Printf("A string: %s \n", s1)
}
