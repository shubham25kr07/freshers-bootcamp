package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// FOR GETTING ALL USERS

	resp, err := http.Get("http://localhost:8080/student-api/student")

	//  FOR GETTING SPECIFIC  USER   (WE WILL GIVE THE ID OF USERS)

	// resp, err := http.Get("http://localhost:8080/student-api/student/2")

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}
