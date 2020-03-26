package main

// packages that the program will utilize
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type numbers struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

type numsResponseData struct {
	Add float64 `json:"add"`
}

// this method recieves a numbers struct and returns a numsResponseData struct after doing addition
func process(numsdata numbers) numsResponseData {

	var numsres numsResponseData
	numsres.Add = numsdata.Num1 + numsdata.Num2

	return numsres
}

// this function handles json requests that are recieved through url
func calc(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var numsData numbers
	var numsResData numsResponseData

	decoder.Decode(&numsData)

	numsResData = process(numsData)

	fmt.Println(numsResData)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(numsResData); err != nil {
		panic(err)
	}
}

// main function
func main() {
	http.HandleFunc("/", calc)
	http.ListenAndServe(":8090", nil)
}
