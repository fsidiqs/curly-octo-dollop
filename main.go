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

var successResponse struct {
	Status int
	Input  string
	Result interface{}
}

var errorResponse struct {
	Status  int
	Message string
}

func init() {
	errorResponse.Status = 500
	successResponse.Status = 200
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test/", serviceHandler)
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

	if _, ok := query["input"]; ok && r.Method == "GET" {
		var err error
		var data interface{}
		input := query["input"][0]
		if r.URL.Path == "/test/soal_0" {

			data, err = palindrome.CountPalindromePossibilities(input)
		} else if r.URL.Path == "/test/soal_1" {
			data, err = book.ParseAndGetSortedBooks(input)
		} else if r.URL.Path == "/test/soal_2" {
			data, err = missingnumber.FindMissingNumber(input)
		}

		if err != nil {
			errorResponse.Message = err.Error()
			json, _ := json.Marshal(errorResponse)
			w.Write(json)
			return
		}

		successResponse.Input = input
		successResponse.Result = data
		json, _ := json.Marshal(successResponse)
		w.Write(json)
		return
	}

	errorResponse.Message = "invalid_request"
	json, _ := json.Marshal(errorResponse)
	w.Write(json)
	return
}
