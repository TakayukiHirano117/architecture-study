package userapp

import (
	"context"
	"errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

type CreateUserAppService struct {
	userRepo          userdm.UserRepository
	IsExistByUserName userdm.IsExistByUserNameDomainService
	IsExistByTagID    tagdm.IsExistByTagIDDomainService
	FindIDByTagName   tagdm.FindIDByTagNameDomainService
}

func NewCreateUserAppService(userRepo userdm.UserRepository, isExistByUserNameDomainService userdm.IsExistByUserNameDomainService, isExistByTagIDDomainService tagdm.IsExistByTagIDDomainService, findIDByTagNameDomainService tagdm.FindIDByTagNameDomainService) *CreateUserAppService {
	return &CreateUserAppService{
		userRepo:          userRepo,
		IsExistByUserName: isExistByUserNameDomainService,
		IsExistByTagID:    isExistByTagIDDomainService,
		FindIDByTagName:   findIDByTagNameDomainService,
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
	TagName           string
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
	b, err := app.IsExistByUserName.Exec(ctx, *userName)
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
		// ここでtagNameの存在チェック
		tagName, err := tagdm.NewTagName(reqSkill.TagName)
		if err != nil {
			return err
		}

		// tagNameがあったらtagIdを取得、DBに保存するのはtagIdなので
		tagId, err := app.FindIDByTagName.Exec(ctx, *tagName)
		if err != nil {
			return err
		}
		if tagId == nil {
			return errors.New("tag name not found")
		}

		skills[i] = userdm.SkillParamIfCreate{
			TagId:             tagId.String(),
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
