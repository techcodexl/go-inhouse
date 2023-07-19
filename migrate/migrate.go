package main

import (
	"github.com/keshav1411/crud-with-go/initialzers"
	"github.com/keshav1411/crud-with-go/models"
)

func init() {
	initialzers.LoadEnvVariable()
	initialzers.ConnectToDB()
}

func main(){
	initialzers.DB.AutoMigrate(&models.Post{})
	initialzers.DB.AutoMigrate(&models.User{})
}