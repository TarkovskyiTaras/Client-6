package main

import (
	"bytes"
	"net/http"
)

func main() {
	client := &http.Client{}
	url := "http://localhost:8080/person/3"

	jsonMessage := `{
						"first_name" : "fsafasf",
						"last_name" : "fsafsaf",
						"dob" : "1990-12-28T00:00:00Z",
						"home_address" : "fsafdf",
						"cellphone" : "ffsdsa"
					}`

	body := bytes.NewBuffer([]byte(jsonMessage))

	request, _ := http.NewRequest("PUT", url, body)
	client.Do(request)
}
