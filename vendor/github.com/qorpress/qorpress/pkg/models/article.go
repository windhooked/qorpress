package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qorpress/publish2"
)

//go:generate gp-extender -structs Article -output article-funcs.go
type Article struct {
	gorm.Model
	Author   User
	AuthorID uint
	Title    string
	Content  string `gorm:"type:text"`
	publish2.Version
	publish2.Schedule
	publish2.Visible
}
