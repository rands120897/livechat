package helper

import (
	"database/sql"
	"fmt"
)

func CommitOrRollback(tx *sql.Tx, err error) {
	if err != nil {
		errorRollback := tx.Rollback()
		fmt.Println("error Rollback")
		ErrHelper(errorRollback)

	} else {
		errorCommit := tx.Commit()
		fmt.Println("error commit")
		ErrHelper(errorCommit)
	}
}
