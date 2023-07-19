package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keshav1411/crud-with-go/controller"
	"github.com/keshav1411/crud-with-go/initialzers"
)
func init(){
	initialzers.LoadEnvVariable()
	initialzers.ConnectToDB()
}
func main() {
	r:= gin.Default()

	r.POST("/posts",controller.PostsCreate)
	r.GET("/",controller.PostsIndex)
	r.GET("/post/:id",controller.PostGetOne)
	r.POST("/post/:id",controller.UpdateOne)
	r.DELETE("/post/:id",controller.DeleteOne)


	r.POST("/createUser",controller.CreateUser)
	r.GET("/getUser",controller.GetUser)
	r.GET("/getUser/:id",controller.GetOneUser)
	r.PUT("/updateUser/:id",controller.UpdateOneUser)
	r.DELETE("/deleteUser/:id",controller.Deletedata)


	r.Run()
}