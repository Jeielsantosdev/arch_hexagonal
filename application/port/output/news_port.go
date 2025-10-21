package output

import (
	"github.com/Jeielsantosdev/arch_hexagonal/application/domain"
	"github.com/Jeielsantosdev/arch_hexagonal/configuration/rest_err"
)


type NewsPort interface {
	GetNewsPorts(domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr)
}