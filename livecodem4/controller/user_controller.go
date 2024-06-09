package controller

import (
	"bookapi/entity"
	"bookapi/service"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "select users",
		"users": service.GetAllUsers(),
	})
}

func Register(c *gin.Context) {
	var user entity.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400,gin.H{
			"message":"error",
			"error": err.Error(),
		})
		return
	}
	user = service.Register(user)
	c.JSON(200, gin.H{
		"message": "Insert user",
		"user" : user,
	})
}

func Profile(c *gin.Context) {
	identifiant, error := strconv.ParseUint(c.Param("id"), 10,64)
	if error != nil {
		c.JSON(400,gin.H{
			"message":"error",
			"error": error.Error(),
		})
		return
	}
	user, err := service.Profile(identifiant)
	if err != nil {
		c.JSON(404,gin.H{
			"message":"error",
			"error": err.Error(),
		})
		return 
	}
	c.JSON(200, gin.H{
		"message": "select user searched using ID",
		"user": user,
	})
}

func UpdateProfile(c *gin.Context) {
	var user entity.User
	c.ShouldBind(&user)
	identifiant, _ := strconv.ParseUint(c.Param("id"), 10,64)
	
	err := service.UpdateProfile(user,identifiant)
	if err != nil {
		c.JSON(404,gin.H{
			"message":"User doesn't exist",
		})
	}else {
		c.JSON(200, gin.H{
		"message": "user updated using id" + c.Param("id"),
		})
	}
}

func DeleteAccount(c *gin.Context) {
	identifiant, error := strconv.ParseUint(c.Param("id"), 10,64)
	if error != nil {
		c.JSON(400,gin.H{
			"message":"error",
			"error": error.Error(),
		})
		return
	}
	err := service.DeleteAccount(identifiant)
	if err != nil{
		c.JSON(404,gin.H{
			"message":"User doesn't exist",
		})
	}else{	
		c.JSON(200, gin.H{
		"message": "Delete user using id" + c.Param("id"),
		})
	}
}