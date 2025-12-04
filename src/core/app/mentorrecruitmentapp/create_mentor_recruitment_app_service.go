package mentorrecruitmentapp

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/support/customerr"
)

type CreateMentorRecruitmentAppService struct {
	isExistByUserIDDomainService     userdm.IsExistByUserIDDomainService
	mentorRecruitmentRepo            mentor_recruitmentdm.MentorRecruitmentRepository
	tagRepo                          tagdm.TagRepository
	isExistByCategoryIDDomainService categorydm.IsExistByCategoryIDDomainService
}

type CreateMentorRecruitmentRequest struct {
	UserID             userdm.UserID                           `json:"user_id"`
	Title              string                                  `json:"title"`
	CategoryID         categorydm.CategoryID                   `json:"category_id"`
	ConsultationType   mentor_recruitmentdm.ConsultationType   `json:"consultation_type"`
	ConsultationMethod mentor_recruitmentdm.ConsultationMethod `json:"consultation_method"`
	Description        string                                  `json:"description"`
	BudgetFrom         uint32                                  `json:"budget_from"`
	BudgetTo           uint32                                  `json:"budget_to"`
	Tags               []CreateMentorRecruitmentTagRequest     `json:"tags"`
}

// TODO: IDのみを受け取り、NameはDBから取得する様に書き換える
type CreateMentorRecruitmentTagRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCreateMentorRecruitmentAppService(
	isExistByUserIDDomainService userdm.IsExistByUserIDDomainService,
	isExistByCategoryIDDomainService categorydm.IsExistByCategoryIDDomainService,
	mentorRecruitmentRepo mentor_recruitmentdm.MentorRecruitmentRepository,
	tagRepo tagdm.TagRepository,
) *CreateMentorRecruitmentAppService {
	return &CreateMentorRecruitmentAppService{
		isExistByUserIDDomainService:     isExistByUserIDDomainService,
		isExistByCategoryIDDomainService: isExistByCategoryIDDomainService,
		mentorRecruitmentRepo:            mentorRecruitmentRepo,
		tagRepo:                          tagRepo,
	}
}

func (app *CreateMentorRecruitmentAppService) Exec(ctx context.Context, req *CreateMentorRecruitmentRequest) error {
	userID, err := userdm.NewUserIDByVal(req.UserID.String())
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

	categoryID, err := categorydm.NewCategoryIDByVal(req.CategoryID.String())
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

	status := mentor_recruitmentdm.Published

	tagRequests := make([]tagdm.TagRequest, len(req.Tags))
	for i, t := range req.Tags {
		tagRequests[i] = tagdm.TagRequest{ID: t.ID, Name: t.Name}
	}

	tags, err := tagdm.NewBuildTagsDomainService(app.tagRepo).Exec(ctx, tagRequests)
	if err != nil {
		return customerr.InternalWrapf(err, "failed to build tags: %s", err.Error())
	}

	mentorRecruitment, err := mentor_recruitmentdm.NewMentorRecruitment(
		mentor_recruitmentdm.NewMentorRecruitmentID(),
		userID,
		req.Title,
		req.Description,
		categoryID,
		req.ConsultationType,
		req.ConsultationMethod,
		req.BudgetFrom,
		req.BudgetTo,
		mentor_recruitmentdm.NewApplicationPeriod(),
		status,
		tags,
	)
	if err != nil {
		return customerr.InternalWrapf(err, "%s", err.Error())
	}

	return app.mentorRecruitmentRepo.Store(ctx, mentorRecruitment)
}
