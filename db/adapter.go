package db

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/thoohv5/gf/db/impl"
	"github.com/thoohv5/gf/db/standard"
)

type typeName string

const (
	Database typeName = "database"
	Gorm     typeName = "gorm"
)

// 数据库连接适配
func GetConnect(typeName typeName) standard.Connecter {
	var connect standard.Connecter
	switch typeName {
	case Database:
		connect = impl.NewSDb()
	case Gorm:
		connect = impl.NewGDb()
	default:
		connect = impl.NewSDb()
	}
	return connect
}
