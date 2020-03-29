package controllers

import (
	"bufio"
	"encoding/json"
	"net/http"
)

// GetService http request method get
func GetService(url string) (response map[string]interface{}, err error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan(); i++ {
		jsonData := []byte(scanner.Text())
		json.Unmarshal(jsonData, &response)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return response, err
}
