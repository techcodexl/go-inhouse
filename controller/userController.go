package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/keshav1411/crud-with-go/initialzers"
	"github.com/keshav1411/crud-with-go/models"
)

func CreateUser(c *gin.Context) {


	var user struct{
		First_name string
		Last_name  string
		Email      string
		Password   string 
		Age        uint8
		Phone_no   uint16
	}

	c.Bind(&user)


	post := models.User{
		First_name: user.First_name,
		Last_name: user.Last_name,
		Email: user.Email,
		Password: user.Password,
		Age: user.Age,
		Phone_no: user.Phone_no,
	}

	result := initialzers.DB.Create(&post)

	if result.Error !=nil{
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})

}
func GetUser(c *gin.Context){

	var users []models.User
	initialzers.DB.Find(&users)

	c.JSON(200, gin.H{
		"post": users,
	})

}
func GetOneUser(c *gin.Context){

	id:=c.Param("id")

	var users models.User
	initialzers.DB.First(&users,id)

	c.JSON(200, gin.H{
		"post": users,
	})

}
func UpdateOneUser(c *gin.Context)  {
	id:= c.Param("id")

	var body struct{
		First_name string
		Last_name  string
		Email      string
		Password   string 
		Age        uint8
		Phone_no   uint16
	}

	c.Bind(&body)

	var users models.User
	initialzers.DB.First(&users,id)
	
	initialzers.DB.Model(&users).Updates(models.User{
		First_name: body.First_name,
		Last_name: body.Last_name,
		Email: body.Email,
		Password: body.Password,
		Age: body.Age,
		Phone_no: body.Phone_no,
	})

		c.JSON(200, gin.H{
			"post": users,
		})
	
}
func Deletedata(c *gin.Context){
	id := c.Param("id")

	initialzers.DB.Delete(&models.User{},id)
	c.JSON(200, gin.H{
		"message": "Data has been deleted",
	})

}