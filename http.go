package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getResponse(url string, out interface{}) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(responseData, &out)
}
