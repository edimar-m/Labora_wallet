package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Generando solicitud

func SearchRequest(w http.ResponseWriter, r *http.Request) {
	url := "https://api.checks.truora.com/v1/checks"
	method := "POST"

	payload := strings.NewReader("national_id=42388604&country=PE&type=person&user_authorized=true&force_creation=true")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	req.Header.Add("Truora-API-Key", os.Getenv("api-key-truora"))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Add("Truora-API-Key", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiIiwiYWRkaXRpb25hbF9kYXRhIjoie30iLCJjbGllbnRfaWQiOiJUQ0kxOTE4YzU5N2Y3YjdkNjEzYjQwNzBlZWNiYWNjZTc0MCIsImV4cCI6MzI0NTQ0MzcwNSwiZ3JhbnQiOiIiLCJpYXQiOjE2Njg2NDM3MDUsImlzcyI6Imh0dHBzOi8vY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb20vdXMtZWFzdC0xX1FOMVZxVWFneSIsImp0aSI6IjVmZGM4YmJjLTlmNzktNDM1YS1hODc3LTM1OGUxOTJhNzI4MCIsImtleV9uYW1lIjoibGFib3JhX2FwaSIsImtleV90eXBlIjoiYmFja2VuZCIsInVzZXJuYW1lIjoicGFydGljdWxhci1sYWJvcmFfYXBpIn0.NKEpa-oP3tvstL_TNU0ObXdp69TLOW7RUKioy8_Dr9E")

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
	fmt.Println(res)
	fmt.Println(string(body))
}
