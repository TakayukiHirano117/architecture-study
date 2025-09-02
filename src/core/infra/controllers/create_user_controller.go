package controllers

import (
	"net/http"

	"github.com/TakayukiHirano117/architecture-study/src/core/app/userapp"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdbimpl"
	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	userRepo          userdm.UserRepository
	IsExistByUserName userdm.IsExistByUserNameDomainService
	IsExistByTagID    tagdm.IsExistByTagIDDomainService
	FindIDByTagName   tagdm.FindIDByTagNameDomainService
}

func NewCreateUserController() *CreateUserController {
	return &CreateUserController{
		userRepo:          rdbimpl.NewUserRepositoryImpl(),
		IsExistByUserName: userdm.NewIsExistByUserNameDomainService(rdbimpl.NewUserRepositoryImpl()),
		IsExistByTagID:    tagdm.NewIsExistByTagIDDomainService(rdbimpl.NewTagRepositoryImpl()),
		FindIDByTagName:   tagdm.NewFindIDByTagNameDomainService(rdbimpl.NewTagRepositoryImpl()),
	}
}

func (c *CreateUserController) Exec(ctx *gin.Context) {
	var in userapp.CreateUserRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	if err := userapp.NewCreateUserAppService(c.userRepo, c.IsExistByUserName, c.IsExistByTagID, c.FindIDByTagName).Exec(ctx.Request.Context(), &in); err != nil {
		ctx.Error(err)
		return
	}
}
