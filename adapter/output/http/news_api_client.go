package http

import (
	"fmt"

	"github.com/Jeielsantosdev/arch_hexagonal/adapter/output/models/response"
	"github.com/Jeielsantosdev/arch_hexagonal/application/domain"
	"github.com/Jeielsantosdev/arch_hexagonal/configuration/env"
	"github.com/Jeielsantosdev/arch_hexagonal/configuration/rest_err"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
)

type newsClient struct {}

func NewNewsclient() *newsClient {
	client = resty.New().SetBaseURL("https://newsapi.org/v2")
	return &newsClient{}
}

var (
	client *resty.Client
)

func (nc *newsClient) GetNewsPorts(newsDomain domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr){
	newsResponse := response.NewsClientResponse{}

	_, err := client.R().
		SetQueryParams(map[string]string{
			"q":      newsDomain.Subject,
			"from":   newsDomain.From,
			"apiKey": env.GetNewsTokenAPI(),
		}).
		SetResult(newsResponse).
		Get("/everything")
	fmt.Println(err)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro trying to call NewsAPI with params")
	}

	newResponseDomain := &domain.NewsDomain{}
	copier.Copy(newResponseDomain , newsResponse)
	return  newResponseDomain, nil
}