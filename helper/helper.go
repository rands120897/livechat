package helper

import (
	"fmt"
	"net/http"
)

func ErrHelper(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func ErrWeb(w http.ResponseWriter, err error, text string, kode int) {
	if err != nil {
		http.Error(w, text, kode)
	}
}
