package student

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Birthdate string `json:"birthdate"`
}
