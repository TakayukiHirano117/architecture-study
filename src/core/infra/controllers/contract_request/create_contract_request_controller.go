package contractrequest_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/TakayukiHirano117/architecture-study/src/core/app/contractrequestapp"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/contract_requestdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdbimpl"
)

type CreateContractRequestController struct {
	contractRequestRepo contract_requestdm.ContractRequestRepository
}

func NewCreateContractRequestController() *CreateContractRequestController {
	return &CreateContractRequestController{
		contractRequestRepo: rdbimpl.NewContractRequestRepositoryImpl(),
	}
}

func (c *CreateContractRequestController) Exec(ctx *gin.Context) {
	var in contractrequestapp.CreateContractRequestRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	if err := contractrequestapp.NewCreateContractRequestAppService(NewContractRequestRepositoryImpl()).Exec(ctx.Request.Context(), &in); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Contract request created successfully"})
}
