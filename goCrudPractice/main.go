package main;

import (
	"github.com/gin-gonic/gin"
	"crud/database"
	"crud/routes"
)

func main(){

	dbInstance:=database.InitDb();

	router:=gin.Default();

	routes.CountriesRoute(router, dbInstance);

	// Create a route group for routes that start with /method/
	methodGroup := router.Group("/methods")
	
	// Pass the router group and db instance to GormMethod
	routes.GormMethod(methodGroup, dbInstance)

	// routes.GormMethod(router, dbInstance);

	router.Run(":8080");
}