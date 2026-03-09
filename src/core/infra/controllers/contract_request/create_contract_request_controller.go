package contractrequest_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/TakayukiHirano117/architecture-study/src/core/app/contractrequestapp"
)

type CreateContractRequestController struct {
}

func NewCreateContractRequestController() *CreateContractRequestController {
	return &CreateContractRequestController{}
}

func (c *CreateContractRequestController) Exec(ctx *gin.Context) {
	var in contractrequestapp.CreateContractRequestRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	if err := contractrequestapp.NewCreateContractRequestAppService(c.contractRequestRepo).Exec(ctx.Request.Context(), &in); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Contract request created successfully"})
}
