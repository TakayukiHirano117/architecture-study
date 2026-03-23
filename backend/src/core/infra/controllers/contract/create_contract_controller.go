package contractcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TakayukiHirano117/architecture-study/src/core/app/contractapp"
	"github.com/TakayukiHirano117/architecture-study/src/core/app/query_service"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/contractdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdbimpl"
)

type CreateContractController struct {
	contractRepo                                        contractdm.ContractRepository
	isExistByUserIDDomainService                        userdm.IsExistByUserIDDomainService
	isExistByPlanIDDomainService                        plandm.IsExistByPlanIDDomainService
	isExistPlanUserCombinationFromContractsQueryService query_service.IsExistPlanUserCombinationFromContractsQueryService
}

func NewCreateContractController() *CreateContractController {
	contractRepo := rdbimpl.NewContractRepositoryImpl()
	return &CreateContractController{
		contractRepo:                                        contractRepo,
		isExistByUserIDDomainService:                        userdm.NewIsExistByUserIDDomainService(rdbimpl.NewUserRepositoryImpl()),
		isExistByPlanIDDomainService:                        plandm.NewIsExistByPlanIDDomainService(rdbimpl.NewPlanRepositoryImpl()),
		isExistPlanUserCombinationFromContractsQueryService: query_service.NewIsExistPlanUserCombinationFromContractsQueryService(contractRepo),
	}
}

func (c *CreateContractController) Exec(ctx *gin.Context) {
	var in contractapp.CreateContractRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	if err := contractapp.NewCreateContractAppService(c.contractRepo, c.isExistByUserIDDomainService, c.isExistByPlanIDDomainService, c.isExistPlanUserCombinationFromContractsQueryService).Exec(ctx.Request.Context(), &in); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Contract created successfully"})
}
