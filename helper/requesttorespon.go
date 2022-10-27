package helper

import (
	"strconv"
	"webservice/models"
)

func ConvertToRespon(req models.User) models.WebResponse {

	result := models.WebResponse{
		Name:  req.Name,
		Email: req.Email,
	}

	return result

}

func ConvertToResponMany(req []models.User) []models.WebResponseAllUser {

	users := []models.WebResponseAllUser{}
	for _, user := range req {
		userId := strconv.Itoa(user.Id)
		result := models.WebResponseAllUser{

			Id:         userId,
			Name:       user.Name,
			Email:      user.Email,
			Created_at: user.Created_at,
		}
		users = append(users, result)
	}

	return users

}

func ConvertToResponOne(req models.User) models.WebResponseAllUser {

	userId := strconv.Itoa(req.Id)
	result := models.WebResponseAllUser{

		Id:         userId,
		Name:       req.Name,
		Email:      req.Email,
		Created_at: req.Created_at,
	}

	return result

}

func ConvertToResponMessage(req models.Message) models.WebResponInsertMessage {
	result := models.WebResponInsertMessage{
		Pesan:      req.Pesan,
		Penerima:   req.Penerima,
		Pengirim:   req.Pengirim,
		Status:     req.Status,
		Created_at: req.Created_at,
	}

	return result
}
