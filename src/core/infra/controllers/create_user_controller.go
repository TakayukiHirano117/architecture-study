package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TakayukiHirano117/architecture-study/src/core/app/userapp"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdbimpl"
)

type CreateUserController struct {
	userRepo          userdm.UserRepository
	tagRepo           tagdm.TagRepository
	IsExistByUserName userdm.IsExistByUserNameDomainService
	IsExistByTagID    tagdm.IsExistByTagIDDomainService
}

func NewCreateUserController() *CreateUserController {
	return &CreateUserController{
		userRepo:          rdbimpl.NewUserRepositoryImpl(),
		tagRepo:           rdbimpl.NewTagRepositoryImpl(),
		IsExistByUserName: userdm.NewIsExistByUserNameDomainService(rdbimpl.NewUserRepositoryImpl()),
		IsExistByTagID:    tagdm.NewIsExistByTagIDDomainService(),
	}
}

func (c *CreateUserController) Exec(ctx *gin.Context) {
	var in userapp.CreateUserRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	if err := userapp.NewCreateUserAppService(c.userRepo, c.tagRepo, c.IsExistByUserName, c.IsExistByTagID).Exec(ctx.Request.Context(), &in); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "User created successfully"})
}
