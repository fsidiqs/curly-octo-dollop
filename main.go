package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"technical_test/soal_0/palindrome"
	"technical_test/soal_1/book"
	"technical_test/soal_2/missingnumber"
	"text/template"
)

// successResponse define the response when the service run successfully
var successResponse struct {
	Status int
	Input  string
	Result interface{}
}

// errorResponse define the response when the service returns an error
var errorResponse struct {
	Status  int
	Message string
}

type serviceable interface {
	RunService(string) (interface{}, error)
}

func init() {
	errorResponse.Status = 500
	successResponse.Status = 200
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test/", serviceHandler)
	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func serviceHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	// fetchin the input query param, and making sure it is sent using GET method
	if _, ok := query["input"]; ok && r.Method == "GET" {
		input := query["input"][0]
		// run the service
		data, err := serviceImpl(r.URL.Path, input)

		// when error occurs
		if err != nil {
			errorResponse.Message = err.Error()
			json, _ := json.Marshal(errorResponse)
			w.Write(json)
			return
		}

		successResponse.Input = input
		successResponse.Result = data
		// return the data as successResponse
		json, _ := json.Marshal(successResponse)
		w.Write(json)
		return
	}

	errorResponse.Message = "invalid_request"
	json, _ := json.Marshal(errorResponse)
	w.Write(json)
	return
}

func serviceImpl(urlPath string, input string) (data interface{}, err error) {
	var service serviceable
	// each service has method named "RunService"
	if urlPath == "/test/soal_0" {
		service = palindrome.Service{}
	} else if urlPath == "/test/soal_1" {
		service = book.Service{}
	} else if urlPath == "/test/soal_2" {
		service = missingnumber.Service{}
	}
	// run the service
	data, err = service.RunService(input)
	return
}
