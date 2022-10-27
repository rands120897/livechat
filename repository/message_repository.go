package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"webservice/models"
)

func InsertMessageRepo(ctx context.Context, tx *sql.Tx, message models.Message) models.Message {

	sql := "INSERT INTO message(pesan,penerima,pengirim,status,created_at) VALUES (?,?,?,?,?)"
	message.Status = "TERKIRIM"
	message.Created_at = time.Now()
	result, err := tx.ExecContext(ctx, sql, message.Pesan, message.Pengirim, message.Penerima, message.Status, message.Created_at)
	if err != nil {
		panic(err.Error())

	}

	fmt.Println("Repository", message.Status, message.Penerima)

	id, err := result.LastInsertId()
	message.Id = int(id)
	return message

}
