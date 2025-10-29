package userapp

import (
	"context"

	"github.com/cockroachdb/errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

type UpdateUserAppService struct {
	userRepo          userdm.UserRepository
	IsExistByUserName userdm.IsExistByUserNameDomainService
	IsExistByTagID    tagdm.IsExistByTagIDDomainService
	FindIDByTagName   tagdm.FindIDByTagNameDomainService
}

func NewUpdateUserAppService(userRepo userdm.UserRepository, IsExistByUserName userdm.IsExistByUserNameDomainService, IsExistByTagID tagdm.IsExistByTagIDDomainService, FindIDByTagName tagdm.FindIDByTagNameDomainService) *UpdateUserAppService {
	return &UpdateUserAppService{
		userRepo:          userRepo,
		IsExistByUserName: IsExistByUserName,
		IsExistByTagID:    IsExistByTagID,
		FindIDByTagName:   FindIDByTagName,
	}
}

type UpdateUserRequest struct {
	ID               string
	Name             string
	Email            string
	Password         string
	Skills           []UpdateSkillRequest
	Careers          []UpdateCareerRequest
	SelfIntroduction string
}

type UpdateSkillRequest struct {
	ID                string
	TagID             string
	Evaluation        int
	YearsOfExperience int
}

type UpdateCareerRequest struct {
	ID        string
	Detail    string
	StartYear int
	EndYear   int
}

// パスワード更新は別のユースケースになると思うので今回は考慮しない
func (app *UpdateUserAppService) Exec(ctx context.Context, req *UpdateUserRequest) error {
	userName, err := userdm.NewUserName(req.Name)
	if err != nil {
		return err
	}

	b, err := app.IsExistByUserName.Exec(ctx, *userName)
	// モダンなフロントエンドなら入力中にすでに存在します的なメッセージ返すと思うが今は考慮しない
	if err != nil {
		return err
	}
	if b {
		return errors.New("user name already exists")
	}

	userID, err := userdm.NewUserIDByVal(req.ID)
	if err != nil {
		return err
	}

	user, err := app.userRepo.FindByID(ctx, userID)

	if err != nil {
		return err
	}

	skills := make([]userdm.SkillParamIfUpdate, len(req.Skills))
	careers := make([]userdm.CareerParamIfUpdate, len(req.Careers))

	for i, reqCareer := range req.Careers {
		careers[i] = userdm.CareerParamIfUpdate{
			ID:     &reqCareer.ID,
			Detail: reqCareer.Detail,
		}
	}

	for i, reqSkill := range req.Skills {
		skills[i] = userdm.SkillParamIfUpdate{
			ID:    &reqSkill.ID,
			TagID: reqSkill.TagID,
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

	return app.userRepo.Update(ctx, user)
}
