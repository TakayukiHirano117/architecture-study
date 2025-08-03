package controllers

import (
	"github.com/gin-gonic/gin"
)

type FindUserByNameController struct{}

func NewFindUserByNameController() *FindUserByNameController {
	return &FindUserByNameController{}
}

func (c *FindUserByNameController) Exec(ctx *gin.Context) {

}
