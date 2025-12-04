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
	BuildTags         tagdm.BuildTagsDomainService
}

func NewCreateUserAppService(
	userRepo userdm.UserRepository,
	isExistByUserNameDomainService userdm.IsExistByUserNameDomainService,
	buildTags tagdm.BuildTagsDomainService,
) *CreateUserAppService {
	return &CreateUserAppService{
		userRepo:          userRepo,
		IsExistByUserName: isExistByUserNameDomainService,
		BuildTags:         buildTags,
	}
}

type CreateUserRequest struct {
	Name             string
	Email            string
	Password         string
	SelfIntroduction string
	Skills           []CreateSkillRequest
	Careers          []CreateCareerRequest
}

type CreateSkillRequest struct {
	ID                *string
	Tag               TagParamRequest
	Evaluation        uint8
	YearsOfExperience uint8
}

type CreateCareerRequest struct {
	Detail    string
	StartYear uint16
	EndYear   uint16
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

	// タグリクエストを作成
	tagRequests := make([]tagdm.TagRequest, len(req.Skills))
	for i, reqSkill := range req.Skills {
		tagID := ""
		if reqSkill.Tag.ID != nil {
			tagID = *reqSkill.Tag.ID
		}
		tagRequests[i] = tagdm.TagRequest{
			ID:   tagID,
			Name: reqSkill.Tag.Name,
		}
	}

	// BuildTagsDomainService でタグを一括取得/作成（新規タグの保存も含む）
	tags, err := app.BuildTags.Exec(ctx, tagRequests)
	if err != nil {
		return err
	}

	// タグ名をキーにしたマップを作成（スキルとタグを紐付けるため）
	tagMap := make(map[string]tagdm.Tag)
	for _, tag := range tags {
		tagMap[tag.Name().String()] = tag
	}

	skills := make([]userdm.SkillParamIfCreate, len(req.Skills))
	for i, reqSkill := range req.Skills {
		tag, ok := tagMap[reqSkill.Tag.Name]
		if !ok {
			return errors.New("tag with name " + reqSkill.Tag.Name + " not found")
		}

		tagID := tag.ID().String()
		skills[i] = userdm.SkillParamIfCreate{
			ID: reqSkill.ID,
			Tag: userdm.TagParamIfCreate{
				ID:   &tagID,
				Name: reqSkill.Tag.Name,
			},
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
