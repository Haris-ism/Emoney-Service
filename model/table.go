package model

import (
	"gorm.io/gorm"
)
type Users struct {
	gorm.Model
	Email    	string `json:"email"`
	Password 	string `json:"password"`
	Name 		string `json:"name"`
}

type Balances struct {
	gorm.Model
	Balance             	int    	
	UsersID          		int
	Users            		*Users 	`json:"Users,omitempty",gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
type History struct {
	gorm.Model
	Price             		int    
	Fee          			int    
	Product            		string 
	Description            	string 
	UsersID          		int
	Users            		*Users 	`json:"Users,omitempty",gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}