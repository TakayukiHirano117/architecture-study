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
	TagName           string
	Evaluation        int
	YearsOfExperience int
}

type UpdateCareerRequest struct {
	ID        string
	Detail    string
	StartYear int
	EndYear   int
}

func (app *UpdateUserAppService) Exec(ctx context.Context, req *UpdateUserRequest) error {
	// username重複チェック
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

	userID, err := userdm.NewUserIDByVal(req.ID)
	if err != nil {
		return err
	}

	user, err := app.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	user.UpdateProfile(userName, email, skills, careers, selfIntroduction)
	if err != nil {
		return err
	}

	return app.userRepo.Update(ctx, user)
	return nil
}
