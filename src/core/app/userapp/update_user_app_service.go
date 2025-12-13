package userapp

import (
	"context"

	"github.com/cockroachdb/errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

type UpdateUserAppService struct {
	userRepo                     userdm.UserRepository
	IsExistByUserNameExcludeSelf userdm.IsExistByUserNameExcludeSelfDomainService
	BuildTags                    tagdm.BuildTagsDomainService
}

func NewUpdateUserAppService(
	userRepo userdm.UserRepository,
	IsExistByUserNameExcludeSelf userdm.IsExistByUserNameExcludeSelfDomainService,
	BuildTags tagdm.BuildTagsDomainService,
) *UpdateUserAppService {
	return &UpdateUserAppService{
		userRepo:                     userRepo,
		IsExistByUserNameExcludeSelf: IsExistByUserNameExcludeSelf,
		BuildTags:                    BuildTags,
	}
}

type UpdateUserRequest struct {
	ID               string                `json:"id"`
	Name             string                `json:"name"`
	Email            string                `json:"email"`
	SelfIntroduction string                `json:"self_introduction"`
	Skills           []UpdateSkillRequest  `json:"skills"`
	Careers          []UpdateCareerRequest `json:"careers"`
}

type UpdateSkillRequest struct {
	ID                string           `json:"id"`
	Tag               UpdateTagRequest `json:"tag"`
	Evaluation        uint8            `json:"evaluation"`
	YearsOfExperience uint8            `json:"years_of_experience"`
}

type UpdateTagRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateCareerRequest struct {
	ID        string `json:"id"`
	Detail    string `json:"detail"`
	StartYear uint16 `json:"start_year"`
	EndYear   uint16 `json:"end_year"`
}

func (app *UpdateUserAppService) Exec(ctx context.Context, req *UpdateUserRequest) error {
	userName, err := userdm.NewUserName(req.Name)
	if err != nil {
		return err
	}

	userID, err := shared.NewUUIDByVal(req.ID)
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

	// タグリクエストを作成
	tagRequests := make([]tagdm.TagRequest, len(req.Skills))
	for i, reqSkill := range req.Skills {
		tagRequests[i] = tagdm.TagRequest{
			ID:   reqSkill.Tag.ID,
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

	skills := make([]userdm.SkillParamIfUpdate, len(req.Skills))
	for i, reqSkill := range req.Skills {
		var skillID *string
		if reqSkill.ID != "" {
			skillID = &reqSkill.ID
		}

		tag, ok := tagMap[reqSkill.Tag.Name]
		if !ok {
			return errors.Newf("tag with name %s not found", reqSkill.Tag.Name)
		}

		tagID := tag.ID().String()
		skills[i] = userdm.SkillParamIfUpdate{
			ID: skillID,
			Tag: userdm.TagParamIfUpdate{
				ID:   &tagID,
				Name: reqSkill.Tag.Name,
			},
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

	return app.userRepo.Update(ctx, user)
}
