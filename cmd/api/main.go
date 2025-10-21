package main

import (
	"github.com/Jeielsantosdev/arch_hexagonal/adapter/input/routes"
	"github.com/Jeielsantosdev/arch_hexagonal/configuration/logger"
	"github.com/gin-gonic/gin"
)


func main(){
	logger.Info("About to start application")
	router := gin.Default()


	routes.InitRoutes(router)
	if err := router.Run(":8080"); err != nil{
		logger.Error("Erro trying to init application on port 8080", err)
		return
	}
}