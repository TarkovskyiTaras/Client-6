package main

import (
	"bytes"
	"net/http"
)

func main() {
	//jsonMessage := `{
	//					"id" : 1,
	//					"first_name" : "Taras",
	//					"last_name" : "Tarkovskyi",
	//					"dob" : "1992-01-23T00:00:00Z",
	//					"home_address" : "Kyiv",
	//					"cellphone" : "0933115485"
	//				}`

	//jsonMessage := `{
	//					"id" : 2,
	//					"first_name" : "Nikita",
	//					"last_name" : "Tarkovskyi",
	//					"dob" : "1992-01-23T00:00:00Z",
	//					"home_address" : "Kyiv",
	//					"cellphone" : "0933115485"
	//				}`

	jsonMessage := `{
						"id" : 3,
						"first_name" : "Sergey",
						"last_name" : "Onishenko",
						"dob" : "1990-12-28T00:00:00Z",
						"home_address" : "Kyiv",
						"cellphone" : "345345345"
					}`

	client := &http.Client{}
	myUrl := "http://localhost:8080/person"
	body1 := bytes.NewBuffer([]byte(jsonMessage))

	request, _ := http.NewRequest("POST", myUrl, body1)
	client.Do(request)
}
