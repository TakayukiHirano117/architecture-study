package userapp

import (
	"context"
	"errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/domain_service"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

type CreateUserAppService struct {
	userRepo          userdm.UserRepository
	userDomainService domain_service.UserDomainService
	tagDomainService  domain_service.TagDomainService
}

func NewCreateUserAppService(userRepo userdm.UserRepository, userDomainService domain_service.UserDomainService, tagDomainService domain_service.TagDomainService) *CreateUserAppService {
	return &CreateUserAppService{
		userRepo:          userRepo,
		userDomainService: userDomainService,
		tagDomainService:  tagDomainService,
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
	TagId             string
	Evaluation        int
	YearsOfExperience int
}

type CreateCareerRequest struct {
	Detail    string
	StartYear int
	EndYear   int
}

func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	userName, err := userdm.NewUserName(req.Name)
	if err != nil {
		return err
	}
	b, err := app.userDomainService.IsExistByUserName(ctx, *userName)
	if err != nil {
		return err
	}
	if b {
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
	for i, reqSkill := range req.Skills {
		// ここでtagIdの存在チェック
		tagId, err := tagdm.NewTagId(reqSkill.TagId)
		if err != nil {
			return err
		}

		b, err := app.tagDomainService.IsExistByTagId(ctx, tagId)
		if err != nil {
			return err
		}
		if !b {
			return errors.New("tag id not found")
		}

		skills[i] = userdm.SkillParamIfCreate{
			TagId:             reqSkill.TagId,
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

	return app.userRepo.Store(ctx, user)
}
