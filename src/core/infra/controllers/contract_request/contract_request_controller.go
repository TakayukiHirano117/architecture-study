package contractrequest_controllers

import "github.com/gin-gonic/gin"

type ContractRequestController struct {
	createContractRequestController *CreateContractRequestController
}

func NewContractRequestController() *ContractRequestController {
	createContractRequestController := NewCreateContractRequestController()

	return &ContractRequestController{
		createContractRequestController: createContractRequestController,
	}
}

func (c *ContractRequestController) SetupRoutes(router *gin.Engine) {
	router.POST("/contract-requests", c.createContractRequestController.Exec)
}
