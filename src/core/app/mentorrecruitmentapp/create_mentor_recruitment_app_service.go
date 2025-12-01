package mentorrecruitmentapp

import (
	"context"

	"github.com/cockroachdb/errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/categorydm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/mentor_recruitmentdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

type CreateMentorRecruitmentAppService struct {
	isExistByUserIDDomainService userdm.IsExistByUserIDDomainService
	mentorRecruitmentRepo mentor_recruitmentdm.MentorRecruitmentRepository
	tagRepo               tagdm.TagRepository
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
		isExistByUserIDDomainService: isExistByUserIDDomainService,
		isExistByCategoryIDDomainService: isExistByCategoryIDDomainService,
		mentorRecruitmentRepo: mentorRecruitmentRepo,
		tagRepo:               tagRepo,
	}
}

func (app *CreateMentorRecruitmentAppService) Exec(ctx context.Context, req *CreateMentorRecruitmentRequest) error {
	userID, err := userdm.NewUserIDByVal(req.UserID.String())
	if err != nil {
		return err
	}

	isExistUser, err := app.isExistByUserIDDomainService.Exec(ctx, userID)
	if err != nil {
		return err
	}
	if !isExistUser {
		return errors.New("user not found")
	}

	categoryID, err := categorydm.NewCategoryIDByVal(req.CategoryID.String())
	if err != nil {
		return err
	}

	isExistCategory, err := app.isExistByCategoryIDDomainService.Exec(ctx, categoryID)
	if err != nil {
		return err
	}
	if !isExistCategory {
		return errors.New("category not found")
	}

	status := mentor_recruitmentdm.Published

	var existingTagIDs []tagdm.TagID
	var newTags []tagdm.Tag

	for _, reqTag := range req.Tags {
		if reqTag.ID == "" {
			tagName, err := tagdm.NewTagNameByVal(reqTag.Name)
			if err != nil {
				return err
			}
			newTag, err := tagdm.NewTag(tagdm.NewTagID(), tagName)
			if err != nil {
				return err
			}
			newTags = append(newTags, *newTag)
		} else {
			tagID, err := tagdm.NewTagIDByVal(reqTag.ID)
			if err != nil {
				return err
			}
			existingTagIDs = append(existingTagIDs, tagID)
		}
	}

	if len(newTags) > 0 {
		if err := app.tagRepo.BulkInsert(ctx, newTags); err != nil {
			return err
		}
	}

	var existingTags []tagdm.Tag
	if len(existingTagIDs) > 0 {
		var err error
		existingTags, err = app.tagRepo.FindByIDs(ctx, existingTagIDs)
		if err != nil {
			return err
		}

		if len(existingTags) != len(existingTagIDs) {
			return errors.Newf("some tags not found: requested %d, found %d", len(existingTagIDs), len(existingTags))
		}
	}

	tags := append(existingTags, newTags...)

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
		return err
	}

	return app.mentorRecruitmentRepo.Store(ctx, mentorRecruitment)
}
