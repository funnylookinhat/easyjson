package main

import (
	"fmt"
	"github.com/funnylookinhat/easyjson/lib"
	"io/ioutil"
	"log"
	"net/http"
)

// Demo application where the root of the returned JSON is an array.
func main() {
	url := "https://api.github.com/users/funnylookinhat"

	fmt.Printf("Fetching URL %s \n", url)

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	data, err := easyjson.DecodeJson([]byte(body))

	if err != nil {
		log.Fatal(err)
	}

	name, _ := easyjson.GetString(data, "name")
	blog, _ := easyjson.GetString(data, "blog")
	location, _ := easyjson.GetString(data, "location")

	fmt.Printf("Name: %s\nBlog: %s\nLocation: %s\n", name, blog, location)
}
