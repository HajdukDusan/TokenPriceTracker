package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SendHTTPGetRequest sends a get request to the specified url and returns the responce string
func SendHTTPGetRequest(url string) (string, error) {

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	response.Body.Close()

	//Convert the body to type string
	return string(body), nil
}

// parseJSONResponse parses a string http response to json format
func ParseJSONResponse(res string) (map[string]interface{}, error) {

	resBytes := []byte(res) 
	// declaring a map for key names as string and values as interface{}
	var jsonResponse map[string]interface{}       
	err := json.Unmarshal(resBytes, &jsonResponse)

	if err != nil {
		fmt.Println(err)
	}

	return jsonResponse, nil
}