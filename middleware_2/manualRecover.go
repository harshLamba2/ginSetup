package main;

import (
	"github.com/gin-gonic/gin"
)

func main(){
	notDefaultPanicProtectedRouter:= gin.New();

	notDefaultPanicProtectedRouter.Use(gin.Recovery()); // by default return status code 500

	notDefaultPanicProtectedRouter.GET("/", func(c * gin.Context){

		panic("THIS IS SOME CODE WHICH SHOWS PANIC");// this will stop the program because of it's unability to deal with the panic created unless a manual panic Recovery middleware is not added

		c.JSON(200, gin.H{
			"message":"Hello There",
		})

	});


	notDefaultPanicProtectedRouter.GET("/something", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"Server is wokring Fine",
			"success": true,
		})
	})

	notDefaultPanicProtectedRouter.Run(":1000");


}