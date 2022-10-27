package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webservice/auth"
	"webservice/models"
	"webservice/service"

	"github.com/julienschmidt/httprouter"
)

func CreateUserWeb(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	webRequest := models.WebRequestData{}
	err := decoder.Decode(&webRequest)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("controller")
	response := service.UserServiceImpl{}
	ResponCreate := response.CreateUser(request.Context(), webRequest)
	webResponResult := models.WebResponWeb{
		Code:   200,
		Status: "OK",
		Data:   ResponCreate,
	}
	// fmt.Println("ResponCreate ", webResponResult)

	writer.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writer)
	encoder.Encode(webResponResult)

}

func UpdateUserWeb(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	webRequest := models.WebRequestUpdate{}
	err := decoder.Decode(&webRequest)
	if err != nil {
		panic(err.Error())
	}
	paramsId := Params.ByName("userId")
	fmt.Println("params id", paramsId)
	convertId, err := strconv.Atoi(paramsId)
	webRequest.Id = convertId
	fmt.Println("web request Update :", webRequest)

	ResponUpdate := service.UpdateUser(r.Context(), webRequest)
	webResponResult := models.WebResponWeb{
		Code:   200,
		Status: "OK",
		Data:   ResponUpdate,
	}
	w.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(webResponResult)

}

func DeleteUserWeb(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	paramsId := params.ByName("userId")
	fmt.Println("params id", paramsId)
	convertId, err := strconv.Atoi(paramsId)
	if err != nil {
		fmt.Fprintf(w, "data tidak sesuai")
	}
	fmt.Println("web request Delete :", convertId)

	ResponUpdate := service.DeleteUser(r.Context(), convertId)
	if ResponUpdate != nil {
		webResponResult := models.WebResponWeb{
			Code:   400,
			Status: "FAIL",
			Data:   ResponUpdate.Error(),
		}
		w.Header().Add("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		encoder.Encode(webResponResult)

	} else {
		webResponResult := models.WebResponWeb{
			Code:   200,
			Status: "OK",
			Data:   "User Berhasil di hapus",
		}
		w.Header().Add("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		encoder.Encode(webResponResult)
	}

}

func FindAllUserWeb(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {

	ResponGetAllUser := service.FindAllUser(r.Context())
	webResponResult := models.WebResponWeb{
		Code:   200,
		Status: "OK",
		Data:   ResponGetAllUser,
	}
	w.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(webResponResult)

}

func FindByIdUserWeb(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {

	cookie, err := r.Cookie("AUTHENTICATION")
	token := cookie.Value

	validation := auth.CheckValidTokenJwt(token)

	if validation == true {
		paramsId := Params.ByName("userId")
		fmt.Println("params id", paramsId)
		convertId, _ := strconv.Atoi(paramsId)

		fmt.Println("web request FindById :", convertId)

		ResponFindById, err := service.FindByIdUser(r.Context(), convertId)
		if err != nil {
			webResponResult := models.WebResponWeb{
				Code:   400,
				Status: "FAIL",
				Data:   err.Error(),
			}
			w.Header().Add("Content-Type", "application/json")

			encoder := json.NewEncoder(w)
			encoder.Encode(webResponResult)

		} else {
			webResponResult := models.WebResponWeb{
				Code:   200,
				Status: "OK",
				Data:   ResponFindById,
			}
			w.Header().Add("Content-Type", "application/json")

			encoder := json.NewEncoder(w)
			encoder.Encode(webResponResult)
		}

	} else {
		http.Error(w, err.Error(), 400)
	}

}

func SignInWeb(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	webRequestLogin := models.WebRequestLogin{}
	errDecode := decoder.Decode(&webRequestLogin)
	if errDecode != nil {
		fmt.Println()
	}

	errorData := service.SignInUser(request.Context(), webRequestLogin)

	if errorData != nil {
		responseLogin := models.WebResponWeb{
			Code:   400,
			Status: "FAIL",
			Data:   "Email atau password Tidak sesuai",
		}

		writer.Header().Add("Content-Type", "application/json")

		encoder := json.NewEncoder(writer)

		encoder.Encode(responseLogin)

	} else {
		token := auth.GetTokenJWT(webRequestLogin.Email)
		fmt.Println("Fungsi Berjalan")
		cookie := new(http.Cookie)
		cookie.Name = "AUTHENTICATION"
		cookie.Value = token
		cookie.Path = "/"
		fmt.Println("Fungsi Berjalan sass" + cookie.Value)
		http.SetCookie(writer, cookie)

		responseLogin := models.WebResponWeb{
			Code:   200,
			Status: "OK",
			Data:   "Login berhasil",
		}
		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		encoder.Encode(responseLogin)

	}

}
