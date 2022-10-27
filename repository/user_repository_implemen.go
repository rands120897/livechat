package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"webservice/helper"
	"webservice/models"
)

func Insert(ctx context.Context, tx *sql.Tx, user models.User) models.User {

	sql := "INSERT INTO user(name,email,password,created_at) VALUES (?,?,?,?)"

	result, err := tx.ExecContext(ctx, sql, user.Name, user.Email, user.Password, time.Now())
	if err != nil {
		panic(err.Error())

	}
	fmt.Println("Repository")

	id, err := result.LastInsertId()
	user.Id = int(id)
	return user

}

func Update(ctx context.Context, tx *sql.Tx, user models.User) models.User {
	sql := "update user set name = ?,email=?,password=? where id=?"
	_, err := tx.ExecContext(ctx, sql, user.Name, user.Email, user.Password, user.Id)
	helper.ErrHelper(err)

	return user

}

func Delete(ctx context.Context, tx *sql.Tx, Id int) error {
	sql := "delete from user where id =?"
	_, err := tx.ExecContext(ctx, sql, Id)
	if err != nil {
		fmt.Println("repo Delete error")
		fmt.Println(err.Error())
	}

	return err

}

func SelectAllUser(ctx context.Context, tx *sql.Tx) ([]models.User, error) {
	sql := "select * from user"
	rows, err := tx.QueryContext(ctx, sql)
	if err != nil {
		fmt.Println("repo Delete error")
		fmt.Println(err.Error())
	}

	users := []models.User{}
	user := models.User{}
	for rows.Next() {

		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at)
		if err != nil {
			fmt.Println(err.Error())
		}
		users = append(users, user)

	}

	return users, err

}

func FindById(ctx context.Context, tx *sql.Tx, id int) (models.User, error) {
	sql := "select * from user where id=?"
	rows, err := tx.QueryContext(ctx, sql, id)
	helper.ErrHelper(err)

	user := models.User{}
	if rows.Next() {
		errScan := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at)
		return user, errScan
	} else {

		return user, fmt.Errorf("data tidak ditemukan")

	}

}

func LoginUserRepo(ctx context.Context, tx *sql.Tx, model models.User) (models.User, error) {
	sql := "select * from user where email=? AND password=?"

	rows, err := tx.QueryContext(ctx, sql, model.Email, model.Password)
	helper.ErrHelper(err)

	fmt.Println(rows.Err())
	user := models.User{}
	if rows.Next() {
		errScan := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at)

		return user, errScan
	} else {
		return user, fmt.Errorf("data tidak ditemukan")
	}

}
