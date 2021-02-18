package db

import (
	"context"
	"testing"

	"github.com/thoohv5/gf/db/standard"
)

func TestGetConnect(t *testing.T) {
	wrap, _ := GetConnect(Gorm).Connect(&standard.Config{
		Driver:          "mysql",
		Source:          "",
		ConnMaxLifeTime: 10,
		MaxIdleConns:    10,
		MaxOpenConns:    10,
		Slave:           nil,
		Log: standard.Log{
			Mode: 2,
			Cat:  "sql",
		},
	})

	ret := make([]*struct {
		Id          int    `db:"id" gorm:"id"`
		SectionName string `db:"section_name" gorm:"section_name"`
		SectionId   int    `db:"section_id" gorm:"section_id"`
		SectionType string `db:"section_type" gorm:"section_type"`
		CreatedAt   int    `db:"created_at" gorm:"created_at"`
		UpdatedAt   int    `db:"updated_at" gorm:"updated_at"`
	}, 0)
	err := wrap.Read().Query(context.TODO(), &ret, "select * from agent_section limit 1")
	t.Log(err, ret)
}
