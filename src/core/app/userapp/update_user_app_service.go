package userapp

import (
	"context"

	"github.com/cockroachdb/errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

type UpdateUserAppService struct {
	userRepo                     userdm.UserRepository
	tagRepo                      tagdm.TagRepository
	IsExistByUserNameExcludeSelf userdm.IsExistByUserNameExcludeSelfDomainService
	IsExistByTagID               tagdm.IsExistByTagIDDomainService
	FindIDByTagName              tagdm.FindIDByTagNameDomainService
}

func NewUpdateUserAppService(
	userRepo userdm.UserRepository,
	tagRepo tagdm.TagRepository,
	IsExistByUserNameExcludeSelf userdm.IsExistByUserNameExcludeSelfDomainService,
	IsExistByTagID tagdm.IsExistByTagIDDomainService,
	FindIDByTagName tagdm.FindIDByTagNameDomainService,
) *UpdateUserAppService {
	return &UpdateUserAppService{
		userRepo:                     userRepo,
		tagRepo:                      tagRepo,
		IsExistByUserNameExcludeSelf: IsExistByUserNameExcludeSelf,
		IsExistByTagID:               IsExistByTagID,
		FindIDByTagName:              FindIDByTagName,
	}
}

type UpdateUserRequest struct {
	ID               string
	Name             string
	Email            string
	Skills           []UpdateSkillRequest
	Careers          []UpdateCareerRequest
	SelfIntroduction string
}

type UpdateSkillRequest struct {
	ID                string
	Tag               UpdateTagRequest
	Evaluation        uint8
	YearsOfExperience uint8
}

type UpdateTagRequest struct {
	ID   string
	Name string
}

type UpdateCareerRequest struct {
	ID        string
	Detail    string
	StartYear int
	EndYear   int
}

func (app *UpdateUserAppService) Exec(ctx context.Context, req *UpdateUserRequest) error {
	userName, err := userdm.NewUserName(req.Name)
	if err != nil {
		return err
	}

	userID, err := userdm.NewUserIDByVal(req.ID)
	if err != nil {
		return err
	}

	isExistByUserNameExcludeSelf, err := app.IsExistByUserNameExcludeSelf.Exec(ctx, *userName, userID)
	if err != nil {
		return err
	}

	if isExistByUserNameExcludeSelf {
		return errors.New("user name already exists")
	}

	user, err := app.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	careers := make([]userdm.CareerParamIfUpdate, len(req.Careers))

	for i, reqCareer := range req.Careers {
		var careerID *string

		if reqCareer.ID != "" {
			careerID = &reqCareer.ID
		}
		careers[i] = userdm.CareerParamIfUpdate{
			ID:        careerID,
			Detail:    reqCareer.Detail,
			StartYear: reqCareer.StartYear,
			EndYear:   reqCareer.EndYear,
		}
	}

	skills := make([]userdm.SkillParamIfUpdate, len(req.Skills))
	newTagNames := []string{}

	for i, reqSkill := range req.Skills {
		var skillID *string

		if reqSkill.ID != "" {
			skillID = &reqSkill.ID
		}

		var tagParam userdm.TagParamIfUpdate

		if reqSkill.Tag.ID != "" {
			isExist, err := app.IsExistByTagID.Exec(ctx, reqSkill.Tag.ID)
			if err != nil {
				return err
			}

			if !isExist {
				return errors.Newf("tag with ID %s not found", reqSkill.Tag.ID)
			}

			tagParam = userdm.TagParamIfUpdate{
				ID:   &reqSkill.Tag.ID,
				Name: reqSkill.Tag.Name,
			}
		} else {
			newTagNames = append(newTagNames, reqSkill.Tag.Name)
			tagParam = userdm.TagParamIfUpdate{
				ID:   nil,
				Name: reqSkill.Tag.Name,
			}
		}

		skills[i] = userdm.SkillParamIfUpdate{
			ID:                skillID,
			Tag:               tagParam,
			Evaluation:        reqSkill.Evaluation,
			YearsOfExperience: reqSkill.YearsOfExperience,
		}
	}

	if err := user.UpdateProfile(
		req.Name,
		req.Email,
		skills,
		careers,
		req.SelfIntroduction,
	); err != nil {
		return err
	}

	if len(newTagNames) > 0 {
		newTags := []tagdm.Tag{}
		for _, skill := range user.Skills() {
			tag := skill.Tag()
			for _, name := range newTagNames {
				if tag.Name().String() == name {
					newTags = append(newTags, *tag)
					break
				}
			}
		}

		if len(newTags) > 0 {
			if err := app.tagRepo.BulkInsert(ctx, newTags); err != nil {
				return err
			}
		}
	}

	return app.userRepo.Update(ctx, user)
}
