package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Obteniendo solicitud

func SearchResult() {

	url := "https://api.checks.truora.com/v1/checks/CHK9ab2d610ff6c497d53e0cfe82764f666"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Add("Truora-API-Key", os.Getenv("api-key-truora"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
