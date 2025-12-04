package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TakayukiHirano117/architecture-study/src/core/app/userapp"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdbimpl"
)

type UpdateUserController struct {
	userRepo          userdm.UserRepository
	tagRepo           tagdm.TagRepository
	IsExistByUserName userdm.IsExistByUserNameExcludeSelfDomainService
	IsExistByTagID    tagdm.IsExistByTagIDDomainService
	FindIDByTagName   tagdm.FindIDByTagNameDomainService
}

func NewUpdateUserController() *UpdateUserController {
	return &UpdateUserController{
		userRepo:          rdbimpl.NewUserRepositoryImpl(),
		tagRepo:           rdbimpl.NewTagRepositoryImpl(),
		IsExistByUserName: userdm.NewIsExistByUserNameExcludeSelfDomainService(rdbimpl.NewUserRepositoryImpl()),
		IsExistByTagID:    tagdm.NewIsExistByTagIDDomainService(),
		FindIDByTagName:   tagdm.NewFindIDByTagNameDomainService(rdbimpl.NewTagRepositoryImpl()),
	}
}

func (c *UpdateUserController) Exec(ctx *gin.Context) {
	var in userapp.UpdateUserRequest

	userID := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	in.ID = userID

	if err := userapp.NewUpdateUserAppService(c.userRepo, c.tagRepo, c.IsExistByUserName, c.IsExistByTagID, c.FindIDByTagName).Exec(ctx.Request.Context(), &in); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "User updated successfully"})
}
