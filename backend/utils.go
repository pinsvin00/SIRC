package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func load_request(request * http.Request)map[string]string {
	//fmt.Println("New request...")
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	//log_bytes(body)
	request_dict := make(map[string]string)
	json.Unmarshal(body, &request_dict)
	return request_dict
}