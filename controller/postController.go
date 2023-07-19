package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/keshav1411/crud-with-go/initialzers"
	"github.com/keshav1411/crud-with-go/models"
)

func PostsCreate(c *gin.Context) {


	var body struct{
		Title string
		Body string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title,Body: body.Body}

	result := initialzers.DB.Create(&post)

	if result.Error !=nil{
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}


func PostsIndex(c *gin.Context){
	var posts []models.Post
	initialzers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"post": posts,
	})
}
func PostGetOne(c *gin.Context){

	id:=c.Param("id")

	var post models.Post
	initialzers.DB.First(&post,id)

	c.JSON(200, gin.H{
		"post": post,
	})

}
func UpdateOne(c *gin.Context)  {
	id:= c.Param("id")

	var body struct{
		Title string
		Body string
	}

	c.Bind(&body)

	var post models.Post
	initialzers.DB.First(&post,id)
	
	initialzers.DB.Model(&post).Updates(models.Post{
		Title:body.Title ,
		Body: body.Body})

		c.JSON(200, gin.H{
			"post": post,
		})
	
}

func DeleteOne(c *gin.Context){
	id := c.Param("id")

	initialzers.DB.Delete(&models.Post{},id)
	c.JSON(200, gin.H{
		"message": "Data has been deleted",
	})

}