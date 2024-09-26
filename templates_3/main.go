package main;

import (
	"github.com/gin-gonic/gin"
)


func main(){
	router:= gin.Default();

	router.LoadHTMLGlob("templates/*.html"); // Load templates from the "templates" directory
	// router.LoadHTMLGlob("templates/subDirName/*.html"); // Load templates from the template's sub directories

	router.Static("/static", "./static"); // Load static resources like css, js, images


	router.GET("/", func(c * gin.Context){

		c.HTML(200, "index.html", gin.H{
			"message":"Success",
			"success":true,
			"page":"Index",
		});

	});

	router.GET("/about", func(c *gin.Context){

		c.HTML(200, "about.html",gin.H{
			"message":"About Content Fetched Successfully",
			"title":"Title From The Backend",
			"discription":"Discription From The Backend",
		})
	})

	router.Run(":8080");


}