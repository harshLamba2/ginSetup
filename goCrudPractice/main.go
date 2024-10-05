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

	router.Run(":8080");
}