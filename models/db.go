package models

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
	//tables = []interface{}{
	//	(*User)(nil),
	//	(*Issue)(nil),
	//}
)

func Connect(ctx context.Context, dbURL string) (err error) {
	db, err = sqlx.ConnectContext(ctx, "mysql", dbURL)
	if err != nil {
		return
	}
	return nil
}
