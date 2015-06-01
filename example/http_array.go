package main

import (
	"fmt"
	"github.com/funnylookinhat/easyjson"
	"io/ioutil"
	"log"
	"net/http"
)

// Demo application where the root of the returned JSON is an array.
func main() {
	url := "https://api.github.com/users/funnylookinhat/repos?per_page=10"

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

	repositories, err := easyjson.GetSlice(data)

	for _, v := range repositories {
		fmt.Printf("Name: %s \n", v.(map[string]interface{})["name"])
	}
}
