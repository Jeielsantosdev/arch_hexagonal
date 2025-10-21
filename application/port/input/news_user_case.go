package input

import (
	"github.com/Jeielsantosdev/arch_hexagonal/application/domain"
	"github.com/Jeielsantosdev/arch_hexagonal/configuration/rest_err"
)

type NewsuserCase interface {
	GetNewsService(
		domain.NewsReqDomain,
	)(*domain.NewsDomain, *rest_err.RestErr)
}