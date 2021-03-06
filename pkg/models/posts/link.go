package posts

import (
	"strings"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/qorpress/qorpress/core/validations"
)

//go:generate gp-extender -structs Link -output link-funcs.go
type Link struct {
	ID       uint `gorm:"primary_key" json:"id"`
	URL      string
	Name     string `gorm:"index:name"`
	Title    string `gorm:"type:mediumtext"`
	ImageUrl string
	PostID   uint
}

func (l Link) Validate(db *gorm.DB) {
	if strings.TrimSpace(l.URL) == "" {
		db.AddError(validations.NewError(l, "URL", "URL can not be empty"))
	}
}

func (l *Link) BeforeCreate() (err error) {
	log.Printf("======> New link: %#v\n", l.URL)
	return nil
}
