package planapp

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/plandm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/support/customerr"
)

type CreatePlanAppService struct {
	isExistByUserIDDomainService     userdm.IsExistByUserIDDomainService
	isExistByCategoryIDDomainService categorydm.IsExistByCategoryIDDomainService
	planRepo                         plandm.PlanRepository
	buildTags                        tagdm.BuildTagsDomainService
}

type CreatePlanRequest struct {
	UserID           string                 `json:"user_id"`
	Title            string                 `json:"title"`
	CategoryID       string                 `json:"category_id"`
	Tags             []CreatePlanTagRequest `json:"tags"`
	Content          string                 `json:"content"`
	Status           string                 `json:"status"`
	ConsultationType string                 `json:"consultation_type"`
	Price            uint32                 `json:"price"`
}

type CreatePlanTagRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCreatePlanAppService(
	isExistByUserIDDomainService userdm.IsExistByUserIDDomainService,
	isExistByCategoryIDDomainService categorydm.IsExistByCategoryIDDomainService,
	planRepo plandm.PlanRepository,
	buildTags tagdm.BuildTagsDomainService,
) *CreatePlanAppService {
	return &CreatePlanAppService{
		isExistByUserIDDomainService:     isExistByUserIDDomainService,
		isExistByCategoryIDDomainService: isExistByCategoryIDDomainService,
		planRepo:                         planRepo,
		buildTags:                        buildTags,
	}
}

func (app *CreatePlanAppService) Exec(ctx context.Context, req *CreatePlanRequest) error {
	userID, err := shared.NewUUIDByVal(req.UserID)
	if err != nil {
		return customerr.BadRequestWrapf(err, "%s", err.Error())
	}

	isExistUser, err := app.isExistByUserIDDomainService.Exec(ctx, userID)
	if err != nil {
		return customerr.InternalWrapf(err, "failed to check if user exists: %s", err.Error())
	}

	if !isExistUser {
		return customerr.NotFound("user not found")
	}

	categoryID, err := categorydm.NewCategoryIDByVal(req.CategoryID)
	if err != nil {
		return customerr.BadRequestWrapf(err, "%s", err.Error())
	}

	isExistCategory, err := app.isExistByCategoryIDDomainService.Exec(ctx, categoryID)
	if err != nil {
		return customerr.InternalWrapf(err, "failed to check if category exists: %s", err.Error())
	}

	if !isExistCategory {
		return customerr.NotFound("category not found")
	}

	tagRequests := make([]tagdm.TagRequest, len(req.Tags))
	for i, t := range req.Tags {
		tagRequests[i] = tagdm.TagRequest{ID: t.ID, Name: t.Name}
	}

	tags, err := app.buildTags.Exec(ctx, tagRequests)
	if err != nil {
		return customerr.InternalWrapf(err, "failed to build tags: %s", err.Error())
	}

	tagIDs := make([]shared.UUID, len(tags))
	for i, t := range tags {
		tagIDs[i] = t.ID()
	}

	status, err := plandm.NewStatus(req.Status)
	if err != nil {
		return customerr.BadRequestWrapf(err, "%s", err.Error())
	}

	consultationType, err := plandm.NewConsultationType(req.ConsultationType)
	if err != nil {
		return customerr.BadRequestWrapf(err, "%s", err.Error())
	}

	mentorID, err := userdm.NewUserIDByVal(req.UserID)
	if err != nil {
		return customerr.BadRequestWrapf(err, "%s", err.Error())
	}

	plan, err := plandm.NewPlan(
		shared.NewUUID(),
		mentorID,
		req.Title,
		categoryID,
		tagIDs,
		req.Content,
		status,
		&consultationType,
		req.Price,
	)
	if err != nil {
		return customerr.InternalWrapf(err, "%s", err.Error())
	}

	return app.planRepo.Store(ctx, plan)
}
