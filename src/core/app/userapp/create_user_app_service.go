package userapp

import (
	"context"
	"errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

type CreateUserAppService struct {
	userRepo          userdm.UserRepository
	tagRepo           tagdm.TagRepository
	IsExistByUserName userdm.IsExistByUserNameDomainService
	IsExistByTagID    tagdm.IsExistByTagIDDomainService
}

func NewCreateUserAppService(
	userRepo userdm.UserRepository,
	tagRepo tagdm.TagRepository,
	isExistByUserNameDomainService userdm.IsExistByUserNameDomainService,
	isExistByTagIDDomainService tagdm.IsExistByTagIDDomainService,
) *CreateUserAppService {
	return &CreateUserAppService{
		userRepo:          userRepo,
		tagRepo:           tagRepo,
		IsExistByUserName: isExistByUserNameDomainService,
		IsExistByTagID:    isExistByTagIDDomainService,
	}
}

type CreateUserRequest struct {
	Name             string
	Email            string
	Password         string
	Skills           []CreateSkillRequest
	Careers          []CreateCareerRequest
	SelfIntroduction string
}

type CreateSkillRequest struct {
	ID                *string
	Tag               TagParamRequest
	Evaluation        uint8
	YearsOfExperience uint8
}

type CreateCareerRequest struct {
	Detail    string
	StartYear int
	EndYear   int
}

type TagParamRequest struct {
	ID   *string
	Name string
}

func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	userName, err := userdm.NewUserName(req.Name)
	if err != nil {
		return err
	}

	isExistByUserName, err := app.IsExistByUserName.Exec(ctx, *userName)
	if err != nil {
		return err
	}
	if isExistByUserName {
		return errors.New("user name already exists")
	}

	email, err := userdm.NewEmail(req.Email)
	if err != nil {
		return err
	}

	password, err := userdm.NewPassword(req.Password)
	if err != nil {
		return err
	}

	skills := make([]userdm.SkillParamIfCreate, len(req.Skills))
	newTagNames := []string{}

	for i, reqSkill := range req.Skills {
		var tagParam userdm.TagParamIfCreate

		if reqSkill.Tag.ID != nil && *reqSkill.Tag.ID != "" {
			isExist, err := app.IsExistByTagID.Exec(ctx, *reqSkill.Tag.ID)
			if err != nil {
				return err
			}

			if !isExist {
				return errors.New("tag with ID " + *reqSkill.Tag.ID + " not found")
			}

			tagParam = userdm.TagParamIfCreate{
				ID:   reqSkill.Tag.ID,
				Name: reqSkill.Tag.Name,
			}
		} else {
			newTagNames = append(newTagNames, reqSkill.Tag.Name)
			tagParam = userdm.TagParamIfCreate{
				ID:   nil,
				Name: reqSkill.Tag.Name,
			}
		}

		skills[i] = userdm.SkillParamIfCreate{
			ID:                reqSkill.ID,
			Tag:               tagParam,
			Evaluation:        reqSkill.Evaluation,
			YearsOfExperience: reqSkill.YearsOfExperience,
		}
	}

	careers := make([]userdm.CareerParamIfCreate, len(req.Careers))
	for i, reqCareer := range req.Careers {
		careers[i] = userdm.CareerParamIfCreate{
			Detail:          reqCareer.Detail,
			CareerStartYear: reqCareer.StartYear,
			CareerEndYear:   reqCareer.EndYear,
		}
	}

	selfIntroduction, err := userdm.NewSelfIntroduction(req.SelfIntroduction)
	if err != nil {
		return err
	}

	user, err := userdm.GenIfCreate(*userName, *email, *password, careers, skills, selfIntroduction)

	if err != nil {
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

	return app.userRepo.Store(ctx, user)
}
