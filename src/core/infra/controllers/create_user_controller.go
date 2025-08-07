package controllers

import (
	"net/http"

	"github.com/TakayukiHirano117/architecture-study/src/core/app/userapp"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/domain_service"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdbimpl"
	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	userRepo userdm.UserRepository
	userDomainService domain_service.UserDomainService
}

func NewCreateUserController() *CreateUserController {
	return &CreateUserController{
		userRepo: rdbimpl.NewUserRepositoryImpl(),
		userDomainService: domain_service.NewUserDomainService(rdbimpl.NewUserRepositoryImpl()),
	}
}

func (c *CreateUserController) Exec(ctx *gin.Context) {
	var in userapp.CreateUserRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	if err := userapp.NewCreateUserAppService(c.userRepo, c.userDomainService).Exec(ctx, &in); err != nil {
		ctx.Error(err)
		return
	}
}
