package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Dob      string `json:"dob"`
	Subject  string `json:"subject"`
	Address  string `json:"address"`
	Marks    int    `json:"marks"`
}

func main() {

	todo := Todo{
		Name:     "hola",
		LastName: "kumar",
		Dob:      "05-02-1999",
		Subject:  "DSA",
		Address:  "jhvhvh",
		Marks:    20,
	}

	reqBody, err := json.Marshal(todo)
	if err != nil {
		print(err)
	}
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/student-api/student/2", bytes.NewBuffer(reqBody))

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
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
