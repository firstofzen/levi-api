package user

import (
	"context"
	"levi-api/db"
	"levi-api/db/statement"
	"levi-api/db/statement/account"
	error2 "levi-api/model/error"
)

type AccountOps struct{}

func (a AccountOps) AddAccount(email string, name string, avatarUrl string) error {
	var dbErr error2.DBError
	if pool, err := db.Connect(); err != nil {
		return dbErr.Connect(err.Error())
	} else {
		defer pool.Close()
		var a account.AccountStmt
		if _, err := pool.Exec(context.Background(), statement.Transaction(a.Init(email, name, avatarUrl))); err != nil {
			return dbErr.TransError(err.Error())
		} else {
			return nil
		}
	}
}
