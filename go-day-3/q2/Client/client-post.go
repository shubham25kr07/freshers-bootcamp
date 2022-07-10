package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var reqBody, err = json.Marshal(map[string]interface{}{
		"Name":     "hello",
		"LastName": "kumar",
		"Dob":      "05-02-1999",
		"Subject":  "DSA",
		"Address":  "jhvhvh",
		"Marks":    80,
	})
	if err != nil {
		print(err)
	}
	resp, err := http.Post("http://localhost:8080/student-api/student",
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))
}
