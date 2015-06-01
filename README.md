# easyjson

**Quickly consume JSON in Go.**

[![Build Status](https://travis-ci.org/funnylookinhat/easyjson.svg?branch=master)](https://travis-ci.org/funnylookinhat/easyjson)

While defining structs is a great means to generate a JSON API, it's a bit 
annoying when you want to consume some arbitrary JSON from another server.

This makes that easier.

```
someJson := `
{
	"some": {
		"nested": {
			"path": "a string!"
		}
	}
}`

data, err := easyjson.DecodeJson([]byte(someJson))

if err != nil {
	log.Fatal(err)
}

someString, err := easyjson.GetString(data, "some, "nested", "path")

// Prints "a string!"
fmt.Printf("%s", someString)
```