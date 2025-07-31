package userapp

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

type CreateUserAppService struct {
	userRepo userdm.UserRepository
}

func NewCreateUserAppService(userRepo userdm.UserRepository) *CreateUserAppService {
	return &CreateUserAppService{
		userRepo: userRepo,
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
	// ユーザドメイン作成
	// ユーザ名重複チェック
	// ユーザ作成
	userName, err := userdm.NewUserName(req.Name)
	if err != nil {
		return err
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
