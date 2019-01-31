package models

import (
	"database/sql"
)

type Contents struct {
	Cid          uint
	Title        string
	Slug         string
	Created      int64
	Modified     int64
	Text         string `db:"size:9999999"`
	Order        uint
	AuthorId     uint
	Template     string
	Type         string
	Status       string
	Password     sql.NullString
	CommentsNum  uint
	AllowComment []byte
	AllowPing    []byte
	AllowFeed    []byte
	Parent       uint
}
