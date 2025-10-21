package service

import (
	"fmt"

	"github.com/Jeielsantosdev/arch_hexagonal/application/domain"
	"github.com/Jeielsantosdev/arch_hexagonal/application/port/output"
	"github.com/Jeielsantosdev/arch_hexagonal/configuration/logger"
	"github.com/Jeielsantosdev/arch_hexagonal/configuration/rest_err"
)


type newsService struct {
	newsPort output.NewsPort
}

func NewNewsService(newsPort output.NewsPort) *newsService {
	return &newsService{newsPort}
}

func (ns *newsService) GetNewsService(newsDomain domain.NewsReqDomain,)(*domain.NewsDomain, *rest_err.RestErr){
	logger.Info(fmt.Sprintf("Init getNewsService function, subject=%s, from=%s", newsDomain.Subject, newsDomain.From))
	newsaDomainResponse, err := ns.newsPort.GetNewsPorts(newsDomain)
	return newsaDomainResponse, err
}