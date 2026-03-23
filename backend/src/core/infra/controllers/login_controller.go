package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TakayukiHirano117/architecture-study/src/core/app/userapp"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdbimpl"
)

type LoginController struct {
	loginAppService *userapp.LoginAppService
}

func NewLoginController() *LoginController {
	return &LoginController{
		loginAppService: userapp.NewLoginAppService(rdbimpl.NewUserRepositoryImpl()),
	}
}

func (c *LoginController) Exec(ctx *gin.Context) {
	var in userapp.LoginRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	response, err := c.loginAppService.Exec(ctx.Request.Context(), &in)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  response.Token,
	})
}
