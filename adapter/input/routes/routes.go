package routes

import (
	"github.com/Jeielsantosdev/arch_hexagonal/adapter/input/controller"
	"github.com/Jeielsantosdev/arch_hexagonal/adapter/output/http"
	"github.com/Jeielsantosdev/arch_hexagonal/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes( r *gin.Engine){
	newsClient := http.NewNewsclient()
	newsService := service.NewNewsService(newsClient)
	newsController := controller.NewNewsController(newsService)

	r.GET("/new", newsController.GetNews)

}