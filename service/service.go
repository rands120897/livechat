package service

import (
	"context"
	"database/sql"
	"fmt"
	"webservice/config"
	"webservice/helper"
	"webservice/models"
	"webservice/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	DB *sql.DB
}

func (service *UserServiceImpl) CreateUser(ctx context.Context, request models.WebRequestData) models.WebResponse {
	db := config.GetConnection()

	tx, errDB := db.Begin()
	fmt.Println("service impl")

	defer helper.CommitOrRollback(tx, errDB)
	defer db.Close()

	fmt.Println(request)
	validate := validator.New()
	errValidate := validate.Struct(&request)
	if errValidate != nil {
		panic(errValidate.Error())
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	user = repository.Insert(ctx, tx, user)
	result := helper.ConvertToRespon(user)
	return result
}

func UpdateUser(ctx context.Context, request models.WebRequestUpdate) models.WebResponse {
	db := config.GetConnection()

	tx, errDB := db.Begin()
	fmt.Println("service impl update")

	defer helper.CommitOrRollback(tx, errDB)
	defer db.Close()
	user := models.User{
		Id:       request.Id,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	fmt.Println(user)
	user = repository.Update(ctx, tx, user)
	result := helper.ConvertToRespon(user)
	return result

}

func DeleteUser(ctx context.Context, idUser int) error {
	db := config.GetConnection()

	tx, errDB := db.Begin()
	fmt.Println("service impl Delete")

	defer helper.CommitOrRollback(tx, errDB)
	defer db.Close()

	err := repository.Delete(ctx, tx, idUser)
	return err

}

func FindAllUser(ctx context.Context) []models.WebResponseAllUser {
	db := config.GetConnection()

	tx, errDB := db.Begin()
	fmt.Println("service impl FindAllUser")

	defer helper.CommitOrRollback(tx, errDB)
	defer db.Close()

	users, err := repository.SelectAllUser(ctx, tx)
	if err != nil {
		fmt.Println("error Service FindAllUser")
	}
	result := helper.ConvertToResponMany(users)

	return result

}

func FindByIdUser(ctx context.Context, id int) (models.WebResponseAllUser, error) {
	db := config.GetConnection()

	tx, errDB := db.Begin()
	fmt.Println("service impl update")

	defer helper.CommitOrRollback(tx, errDB)
	defer db.Close()

	user, errorRepo := repository.FindById(ctx, tx, id)
	fmt.Println(user)
	result := helper.ConvertToResponOne(user)
	return result, errorRepo

}

func SignInUser(ctx context.Context, data models.WebRequestLogin) error {
	db := config.GetConnection()

	tx, errDB := db.Begin()

	defer helper.CommitOrRollback(tx, errDB)

	dataModels := models.User{
		Email:    data.Email,
		Password: data.Password,
	}

	_, errRepo := repository.LoginUserRepo(ctx, tx, dataModels)
	defer db.Close()

	return errRepo

}
