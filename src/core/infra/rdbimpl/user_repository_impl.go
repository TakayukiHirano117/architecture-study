package rdbimpl

import (
	"context"
	"errors"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/models"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/rdb"
)

type UserRepositoryImpl struct {
}

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
				Evaluation:        int(row.SkillEvaluation.Int64),
				YearsOfExperience: int(row.SkillYearsOfExperience.Int64),
			})
		}

		if row.CareerID.Valid {
			careerModels = append(careerModels, models.CareerModel{
				CareerID:  row.CareerID.String,
				Detail:    row.CareerDetail.String,
				StartYear: int(row.CareerStart.Int64),
				EndYear:   int(row.CareerEnd.Int64),
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

	// Skill / Career を VO に詰め替え
	skills := []userdm.Skill{}
	for _, s := range skillModels {
		tagID, _ := tagdm.NewTagIDByVal(s.TagID)
		ev, _ := userdm.NewEvaluationByVal(s.Evaluation)
		yoe, _ := userdm.NewYearsOfExperienceByVal(s.YearsOfExperience)
		skill, _ := userdm.NewSkillByVal(userdm.NewSkillID(), tagID, ev, yoe)
		skills = append(skills, *skill)
	}

	careers := []userdm.Career{}
	for _, c := range careerModels {
		idVo, _ := userdm.NewCareerIDByVal(c.CareerID)
		detailVo, _ := userdm.NewCareerDetailByVal(c.Detail)
		startVo, _ := userdm.NewCareerStartYearByVal(c.StartYear)
		endVo, _ := userdm.NewCareerEndYearByVal(c.EndYear)
		career, _ := userdm.NewCareerByVal(idVo, detailVo, startVo, endVo)
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
				s.evaluation AS skill_evaluation,
				s.years_of_experience AS skill_years_of_experience,
				c.id AS career_id,
				c.detail AS career_detail,
				c.start_year AS career_start_year,
				c.end_year AS career_end_year
		FROM users u
		LEFT JOIN skills s ON s.user_id = u.id
		LEFT JOIN careers c ON c.user_id = u.id
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
				Evaluation:        int(row.SkillEvaluation.Int64),
				YearsOfExperience: int(row.SkillYearsOfExperience.Int64),
			})
		}

		if row.CareerID.Valid {
			careerModels = append(careerModels, models.CareerModel{
				CareerID:  row.CareerID.String,
				Detail:    row.CareerDetail.String,
				StartYear: int(row.CareerStart.Int64),
				EndYear:   int(row.CareerEnd.Int64),
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

	// Skill / Career を VO に詰め替え
	skills := []userdm.Skill{}
	for _, s := range skillModels {
		tagID, _ := tagdm.NewTagIDByVal(s.TagID)
		ev, _ := userdm.NewEvaluationByVal(s.Evaluation)
		yoe, _ := userdm.NewYearsOfExperienceByVal(s.YearsOfExperience)
		skill, _ := userdm.NewSkillByVal(userdm.NewSkillID(), tagID, ev, yoe)
		skills = append(skills, *skill)
	}

	careers := []userdm.Career{}
	for _, c := range careerModels {
		idVo, _ := userdm.NewCareerIDByVal(c.CareerID)
		detailVo, _ := userdm.NewCareerDetailByVal(c.Detail)
		startVo, _ := userdm.NewCareerStartYearByVal(c.StartYear)
		endVo, _ := userdm.NewCareerEndYearByVal(c.EndYear)
		career, _ := userdm.NewCareerByVal(idVo, detailVo, startVo, endVo)
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
				career.EndYear().Int(),
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
				skill.Evaluation().Int(),
				skill.YearsOfExperience().Int(),
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
				career.EndYear().Int(),
				career.ID().String(),
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
