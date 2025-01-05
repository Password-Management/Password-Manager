package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetMasterDetails(masterId string) {
	log.Println("insdie the client")
	client := &http.Client{

	}
	log.Println("After client defination")

	// Make a GET request
	resp, err := client.Get(`http://localhost:8001/customer?masterId=` + masterId)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	// Print the response
	fmt.Println("Response Body:", string(body))
}
