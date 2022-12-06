package controller

import (
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB //this struct is required to make method for the established db connection
}

func Service(db *gorm.DB) *Repo {
	return &Repo{db} // function to pass the db connection
}
