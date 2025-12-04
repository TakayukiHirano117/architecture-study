package rdbimpl

import (
	"context"
	"errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/models"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

type UserRepositoryImpl struct {}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) FindByName(ctx context.Context, name userdm.UserName) (*userdm.User, error) {
	conn, err := rdb.ExecFromCtx(ctx)

	if err != nil {
		return nil, err
	}

	query := `
		SELECT
				u.id AS user_id,
				u.name,
				u.email,
				u.password,
				u.self_introduction,
				s.id AS skill_id,
				s.tag_id AS skill_tag_id,
				s.evaluation AS skill_evaluation,
				s.years_of_experience AS skill_years_of_experience,
				c.id AS career_id,
				c.detail AS career_detail,
				c.start_year AS career_start_year,
				c.end_year AS career_end_year
		FROM users u
		LEFT JOIN skills s ON s.user_id = u.id
		LEFT JOIN careers c ON c.user_id = u.id
		WHERE u.name = $1;
	`
	rows, err := conn.QueryxContext(ctx, query, name.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userDetailRows := []models.UserDetailModel{}
	for rows.Next() {
		var r models.UserDetailModel
		if err := rows.StructScan(&r); err != nil {
			return nil, err
		}
		userDetailRows = append(userDetailRows, r)
	}

	if len(userDetailRows) == 0 {
		return nil, nil
	}

	skillModels := []models.SkillModel{}
	careerModels := []models.CareerModel{}

	for _, row := range userDetailRows {
		if row.SkillID.Valid {
			skillModels = append(skillModels, models.SkillModel{
				SkillID:           row.SkillID.String,
				TagID:             row.SkillTagID.String,
				Evaluation:        uint8(row.SkillEvaluation.Int64),
				YearsOfExperience: uint8(row.SkillYearsOfExperience.Int64),
			})
		}

		if row.CareerID.Valid {
			careerModels = append(careerModels, models.CareerModel{
				CareerID:  row.CareerID.String,
				Detail:    row.CareerDetail.String,
				StartYear: int(row.CareerStart.Int64),
				EndYear:   uint16(row.CareerEnd.Int64),
			})
		}
	}

	u := userDetailRows[0]

	userID, err := userdm.NewUserIDByVal(u.UserID)
	if err != nil {
		return nil, err
	}

	userName, err := userdm.NewUserNameByVal(u.UserName)
	if err != nil {
		return nil, err
	}

	email, err := userdm.NewEmailByVal(u.Email)
	if err != nil {
		return nil, err
	}
	password, err := userdm.NewPasswordByVal(u.Password)
	if err != nil {
		return nil, err
	}

	selfIntroductionStr := ""
	if u.SelfIntroduction.Valid {
		selfIntroductionStr = u.SelfIntroduction.String
	}
	selfIntroduction, err := userdm.NewSelfIntroductionByVal(selfIntroductionStr)
	if err != nil {
		return nil, err
	}

	skills := []userdm.Skill{}
	for _, s := range skillModels {
		tagID, err := tagdm.NewTagIDByVal(s.TagID)
		if err != nil {
			return nil, err
		}
		tagName, err := tagdm.NewTagNameByVal(s.TagName)
		if err != nil {
			return nil, err
		}
		tag, err := tagdm.NewTagByVal(tagID, tagName)
		if err != nil {
			return nil, err
		}
		ev, err := userdm.NewEvaluationByVal(s.Evaluation)
		if err != nil {
			return nil, err
		}
		yoe, err := userdm.NewYearsOfExperienceByVal(s.YearsOfExperience)
		if err != nil {
			return nil, err
		}
		skill, err := userdm.NewSkillByVal(userdm.NewSkillID(), tag, ev, yoe)
		if err != nil {
			return nil, err
		}
		skills = append(skills, *skill)
	}

	careers := []userdm.Career{}
	for _, c := range careerModels {
		idVo, err := userdm.NewCareerIDByVal(c.CareerID)
		if err != nil {
			return nil, err
		}
		detailVo, err := userdm.NewCareerDetailByVal(c.Detail)
		if err != nil {
			return nil, err
		}
		startVo, err := userdm.NewCareerStartYearByVal(c.StartYear)
		if err != nil {
			return nil, err
		}
		endVo, err := userdm.NewCareerEndYearByVal(c.EndYear)
		if err != nil {
			return nil, err
		}
		career, err := userdm.NewCareerByVal(idVo, detailVo, startVo, endVo)
		if err != nil {
			return nil, err
		}
		careers = append(careers, *career)
	}
	return userdm.NewUserByVal(
		userID,
		userName,
		password,
		email,
		skills,
		careers,
		&selfIntroduction,
	)
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, id userdm.UserID) (*userdm.User, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT
				u.id AS user_id,
				u.name,
				u.email,
				u.password,
				u.self_introduction,
				s.id AS skill_id,
				s.tag_id AS skill_tag_id,
				t.name AS skill_tag_name,
				s.evaluation AS skill_evaluation,
				s.years_of_experience AS skill_years_of_experience,
				c.id AS career_id,
				c.detail AS career_detail,
				c.start_year AS career_start_year,
				c.end_year AS career_end_year
		FROM users u
		LEFT JOIN skills s ON s.user_id = u.id
		LEFT JOIN careers c ON c.user_id = u.id
		LEFT JOIN tags t ON t.id = s.tag_id
		WHERE u.id = $1;
		`

	rows, err := conn.QueryxContext(ctx, query, id.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userDetailRows := []models.UserDetailModel{}
	for rows.Next() {
		var r models.UserDetailModel
		if err := rows.StructScan(&r); err != nil {
			return nil, err
		}
		userDetailRows = append(userDetailRows, r)
	}

	if len(userDetailRows) == 0 {
		return nil, errors.New("user not found")
	}

	skillModels := []models.SkillModel{}
	careerModels := []models.CareerModel{}

	for _, row := range userDetailRows {
		if row.SkillID.Valid {
			skillModels = append(skillModels, models.SkillModel{
				SkillID:           row.SkillID.String,
				TagID:             row.SkillTagID.String,
				TagName:           row.SkillTagName.String,
				Evaluation:        uint8(row.SkillEvaluation.Int64),
				YearsOfExperience: uint8(row.SkillYearsOfExperience.Int64),
			})
		}

		if row.CareerID.Valid {
			careerModels = append(careerModels, models.CareerModel{
				CareerID:  row.CareerID.String,
				Detail:    row.CareerDetail.String,
				StartYear: int(row.CareerStart.Int64),
				EndYear:   uint16(row.CareerEnd.Int64),
			})
		}
	}

	u := userDetailRows[0]

	userID, err := userdm.NewUserIDByVal(u.UserID)
	if err != nil {
		return nil, err
	}

	userName, err := userdm.NewUserNameByVal(u.UserName)
	if err != nil {
		return nil, err
	}

	email, err := userdm.NewEmailByVal(u.Email)
	if err != nil {
		return nil, err
	}
	password, err := userdm.NewPasswordByVal(u.Password)
	if err != nil {
		return nil, err
	}

	selfIntroductionStr := ""
	if u.SelfIntroduction.Valid {
		selfIntroductionStr = u.SelfIntroduction.String
	}
	selfIntroduction, err := userdm.NewSelfIntroductionByVal(selfIntroductionStr)
	if err != nil {
		return nil, err
	}

	skills := []userdm.Skill{}
	for _, s := range skillModels {
		tagID, err := tagdm.NewTagIDByVal(s.TagID)
		if err != nil {
			return nil, err
		}
		tagName, err := tagdm.NewTagNameByVal(s.TagName)
		if err != nil {
			return nil, err
		}
		tag, err := tagdm.NewTagByVal(tagID, tagName)
		if err != nil {
			return nil, err
		}
		ev, err := userdm.NewEvaluationByVal(s.Evaluation)
		if err != nil {
			return nil, err
		}
		yoe, err := userdm.NewYearsOfExperienceByVal(s.YearsOfExperience)
		if err != nil {
			return nil, err
		}
		skill, err := userdm.NewSkillByVal(userdm.NewSkillID(), tag, ev, yoe)
		if err != nil {
			return nil, err
		}
		skills = append(skills, *skill)
	}

	careers := []userdm.Career{}
	for _, c := range careerModels {
		idVo, err := userdm.NewCareerIDByVal(c.CareerID)
		if err != nil {
			return nil, err
		}
		detailVo, err := userdm.NewCareerDetailByVal(c.Detail)
		if err != nil {
			return nil, err
		}
		startVo, err := userdm.NewCareerStartYearByVal(c.StartYear)
		if err != nil {
			return nil, err
		}
		endVo, err := userdm.NewCareerEndYearByVal(c.EndYear)
		if err != nil {
			return nil, err
		}
		career, err := userdm.NewCareerByVal(idVo, detailVo, startVo, endVo)
		if err != nil {
			return nil, err
		}
		careers = append(careers, *career)
	}
	return userdm.NewUserByVal(
		userID,
		userName,
		password,
		email,
		skills,
		careers,
		&selfIntroduction,
	)
}

func (r *UserRepositoryImpl) Store(ctx context.Context, user *userdm.User) error {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return errors.New("transaction not found")
	}

	userQuery := `
		INSERT INTO users (id, name, email, password, self_introduction, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
	`
	_, err = conn.ExecContext(ctx, userQuery,
		user.ID().String(),
		user.Name().String(),
		user.Email().String(),
		user.Password().String(),
		user.SelfIntroduction().String(),
	)
	if err != nil {
		return err
	}

	if len(user.Careers()) > 0 {
		careerQuery := `
			INSERT INTO careers (id, user_id, detail, start_year, end_year, created_at, updated_at)
			VALUES (gen_random_uuid(), $1, $2, $3, $4, NOW(), NOW())
		`
		for _, career := range user.Careers() {
			_, err = conn.ExecContext(ctx, careerQuery,
				user.ID().String(),
				career.Detail().String(),
				career.StartYear().Int(),
				career.EndYear().Uint16(),
			)
			if err != nil {
				return err
			}
		}
	}

	if len(user.Skills()) > 0 {
		skillQuery := `
			INSERT INTO skills (id, user_id, tag_id, created_at, updated_at)
			VALUES (gen_random_uuid(), $1, $2, NOW(), NOW())
		`
		for _, skill := range user.Skills() {
			_, err = conn.ExecContext(ctx, skillQuery,
				user.ID().String(),
				skill.TagID().String(),
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *userdm.User) error {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return errors.New("transaction not found")
	}

	query := `
		UPDATE users SET name = $1, email = $2, password = $3, self_introduction = $4, updated_at = NOW() WHERE id = $5
	`

	_, err = conn.ExecContext(ctx, query,
		user.Name().String(),
		user.Email().String(),
		user.Password().String(),
		user.SelfIntroduction().String(),
		user.ID().String(),
	)
	if err != nil {
		return err
	}

	if len(user.Skills()) > 0 {
		skillQuery := `
			UPDATE skills SET tag_id = $1, evaluation = $2, years_of_experience = $3, updated_at = NOW() WHERE id = $4
		`
		for _, skill := range user.Skills() {
			_, err := conn.ExecContext(ctx, skillQuery,
				skill.TagID().String(),
				skill.Evaluation().Uint8(),
				skill.YearsOfExperience().Uint8(),
				skill.ID().String(),
			)
			if err != nil {
				return err
			}
		}
	}

	if len(user.Careers()) > 0 {
		careerQuery := `
			UPDATE careers SET detail = $1, start_year = $2, end_year = $3, updated_at = NOW() WHERE id = $4
		`
		for _, career := range user.Careers() {
			_, err := conn.ExecContext(ctx, careerQuery,
				career.Detail().String(),
				career.StartYear().Int(),
				career.EndYear().Uint16(),
				career.ID().String(),
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
