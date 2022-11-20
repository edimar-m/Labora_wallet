package utils

import (
	"fmt"
	"net/http"
)

func UserMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wellcome, you have got your virtual wallet!")
}
