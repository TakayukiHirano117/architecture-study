package controllers

import (
	"github.com/gin-gonic/gin"
)

type FindByNameUserController struct{}

func NewFindByNameUserController() *FindByNameUserController {
	return &FindByNameUserController{}
}

func (c *FindByNameUserController) Exec(ctx *gin.Context) {

}
