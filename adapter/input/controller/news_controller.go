package controller

import (
	"fmt"
	"net/http"

	"github.com/Jeielsantosdev/arch_hexagonal/adapter/input/model/request"
	"github.com/Jeielsantosdev/arch_hexagonal/application/domain"
	"github.com/Jeielsantosdev/arch_hexagonal/application/port/input"
	"github.com/Jeielsantosdev/arch_hexagonal/configuration/logger"
	"github.com/Jeielsantosdev/arch_hexagonal/configuration/validation"
	"github.com/gin-gonic/gin"
)

type newsController struct{
	newsUseCase input.NewsuserCase
}


func NewNewsController(newsUsCase input.NewsuserCase,) *newsController {

	return &newsController{newsUsCase}
}


func (nc * newsController)GetNews(c * gin.Context){
	logger.Info("Init GetNews controller api")
	newsRequest := request.NewsRequest{}

	if err := c.ShouldBindQuery(&newsRequest); err != nil {
		logger.Error("Erro trying  to validate fields from request", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}
	fmt.Println(newsRequest.From.Format("2006-01-02 "))
	newsDomain := domain.NewsReqDomain{
		Subject: newsRequest.Subject,
		From: newsRequest.From.Format("2006-01-02 "),
	}
	newsResponseDomain, err := nc.newsUseCase.GetNewsService(newsDomain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, newsResponseDomain)
}