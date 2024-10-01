package controller;

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func GetUser(c *gin.Context){

	var newUser = User{
		Name:"Harsh",
		Age:24,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user": newUser,
	});

}