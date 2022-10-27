package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
	"webservice/controller"

	"github.com/julienschmidt/httprouter"
)

func indexForm(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var filepath = path.Join("static", "view.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if r.Method == "GET" {
		var tmplt = template.Must(template.New("form").ParseFiles("index.html"))
		var err = tmplt.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
	http.Error(w, "", http.StatusBadRequest)

}

func main() {

	// ctx := context.Background()
	// repository.Insert(ctx)

	fmt.Println("SERVER RUNINGGG")
	router := httprouter.New()
	router.GET("/", indexForm)
	router.POST("/user", controller.CreateUserWeb)
	router.PUT("/user/:userId", controller.UpdateUserWeb)
	router.DELETE("/user/:userId", controller.DeleteUserWeb)
	router.GET("/users", controller.FindAllUserWeb)
	router.GET("/users/:userId", controller.FindByIdUserWeb)
	router.POST("/user/login", controller.SignInWeb)

	router.POST("/message/", controller.InsertMassageWeb)

	// http.HandleFunc("/index", controller.Create)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err.Error())
	}

	// db := config.GetConnection()
	// defer db.Close()

	// ctx := context.Background()

	// result, err := db.ExecContext(ctx, "INSERT INTO customer(id,name) VALUES ('gilang1','budi');")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(result)
	// fmt.Println("success insert")

	// type Customer struct {
	// 	id         string
	// 	name       string
	// 	email      sql.NullString
	// 	balance    int
	// 	rating     float32
	// 	created_at time.Time
	// 	birth_date time.Time
	// 	merried    bool
	// }

	// query := "SELECT * FROM customer"
	// rows, err := db.QueryContext(ctx, query)
	// helper.ErrHelper(err)

	// var eko Customer
	// for rows.Next() {

	// 	err := rows.Scan(&eko.id, &eko.name, &eko.email, &eko.balance, &eko.rating, &eko.created_at, &eko.birth_date, &eko.merried)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("id: ", eko.email)
	// 	fmt.Println("name: ", eko.name)
	// }
	// defer rows.Close()

}
