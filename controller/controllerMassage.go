package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webservice/models"
	"webservice/service"

	"github.com/julienschmidt/httprouter"
)

func InsertMassageWeb(wr http.ResponseWriter, r *http.Request, params httprouter.Params) {

	decoder := json.NewDecoder(r.Body)
	insertMessage := models.WebRequestInsertMessage{}
	err := decoder.Decode(&insertMessage)
	if err != nil {
		fmt.Fprintf(wr, "Decode gagal")
	}

	MessageService := service.CreateMessage(r.Context(), insertMessage)

	webResponResult := models.WebResponWeb{
		Code:   200,
		Status: "OK",
		Data:   MessageService,
	}
	wr.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(wr)
	encoder.Encode(webResponResult)

}
