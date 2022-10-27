package service

import (
	"context"
	"fmt"
	"webservice/config"
	"webservice/helper"
	"webservice/models"
	"webservice/repository"

	"github.com/go-playground/validator/v10"
)

func CreateMessage(ctx context.Context, request models.WebRequestInsertMessage) models.WebResponInsertMessage {
	db := config.GetConnection()

	tx, errDB := db.Begin()
	fmt.Println("service impl Massage")

	defer helper.CommitOrRollback(tx, errDB)
	defer db.Close()

	fmt.Println(request)
	validate := validator.New()
	errValidate := validate.Struct(&request)
	if errValidate != nil {
		panic(errValidate.Error())
	}

	message := models.Message{
		Pesan:    request.Pesan,
		Penerima: request.Penerima,
		Pengirim: request.Pengirim,
	}
	messageRepo := repository.InsertMessageRepo(ctx, tx, message)
	result := helper.ConvertToResponMessage(messageRepo)
	return result
}
