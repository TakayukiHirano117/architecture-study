package controllers

import (
	"net/http"

	"github.com/TakayukiHirano117/architecture-study/src/core/app/userapp"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdbimpl"
	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	userRepo userdm.UserRepository
}

func NewUpdateUserController() *UpdateUserController {
	return &UpdateUserController{
		userRepo: rdbimpl.NewUserRepositoryImpl(),
	}
}

func (c *UpdateUserController) Exec(ctx *gin.Context) {
	var in userapp.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	if err := userapp.NewUpdateUserAppService(c.userRepo).Exec(ctx, &in); err != nil {
		ctx.Error(err)
		return
	}
}