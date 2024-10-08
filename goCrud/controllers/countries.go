package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func AddCounty(c *gin.Context){

	fmt.Println("add Country hit");

	c.JSON(200, gin.H{
		"success": true,
	})

}