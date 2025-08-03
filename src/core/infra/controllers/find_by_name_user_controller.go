package controllers

import (
	"net/http"

	"github.com/TakayukiHirano117/architecture-study/src/core/app/userapp"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdbimpl"
	"github.com/gin-gonic/gin"
)

type FindByNameUserController struct{}

func NewFindByNameUserController() *FindByNameUserController {
	return &FindByNameUserController{}
}

func (c *FindByNameUserController) Exec(ctx *gin.Context) {
	name := ctx.Param("name")

	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	userRepoImpl := rdbimpl.NewUserRepositoryImpl()

	user, err := userapp.NewFindByNameUserAppService(userRepoImpl).Exec(ctx, name)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
