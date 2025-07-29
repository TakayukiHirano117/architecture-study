package controllers

import "github.com/gin-gonic/gin"

func NewUserController(g *gin.Engine) {
	g.POST("/users", func(ctx *gin.Context) {
		NewCreateUserController().Exec(ctx)
	})
}
