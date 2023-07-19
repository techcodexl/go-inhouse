package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	First_name string
	Last_name  string
	Email      string
	Password   string
	Age        uint8
	Phone_no   uint16
}