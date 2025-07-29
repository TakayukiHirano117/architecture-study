package controllers

import (
	"net/http"

	"github.com/TakayukiHirano117/architecture-study/src/core/app/userapp"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdbimpl"
	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	userAppService userapp.CreateUserAppService
}

func NewCreateUserController() *CreateUserController {
	return &CreateUserController{}
}

func (c *CreateUserController) Exec(ctx *gin.Context) {
	var in userapp.CreateUserRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	userRepoImpl := rdbimpl.NewUserRepositoryImpl()

	if err := userapp.NewCreateUserAppService(userRepoImpl).Exec(ctx, &in); err != nil {
		ctx.Error(err)
		return
	}
}
