package contractcontroller

import "github.com/gin-gonic/gin"

type ContractController struct {
	createContractController *CreateContractController
}

func NewContractController() *ContractController {
	createContractController := NewCreateContractController()

	return &ContractController{
		createContractController: createContractController,
	}
}

func (cc *ContractController) SetupRoutes(router gin.IRouter) {
	router.POST("/contracts", cc.createContractController.Exec)
}
