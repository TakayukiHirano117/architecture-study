package rdbimpl

import (
	"context"
	"errors"
	"time"

	"strconv"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
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
		SELECT id FROM users WHERE name = $1
	`
	rows, err := conn.QueryContext(ctx, query, name.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		return &userdm.User{}, nil
	}

	return nil, nil
}

// IDに基づいてユーザーを取得する
func (r *UserRepositoryImpl) FindByID(ctx context.Context, id userdm.UserID) (*userdm.User, error) {
	// 接続取得
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// usersからskills, careers以外のデータをまず取得する。
	getUserQuery := `
		SELECT id, name, email, password, self_introduction, created_at, updated_at FROM users WHERE id = $1
	`

	userRows, err := conn.QueryContext(ctx, getUserQuery, id.String())
	if err != nil {
		return nil, err
	}
	defer userRows.Close()

	getSkillsQuery := `
		SELECT id, tag_id, evaluation, years_of_experience, created_at, updated_at FROM skills WHERE user_id = $1
	`

	skillRows, err := conn.QueryContext(ctx, getSkillsQuery, id.String())
	if err != nil {
		return nil, err
	}
	defer skillRows.Close()

	skills := make([]userdm.Skill, 0)
	for skillRows.Next() {
		var id, tagId, evaluation, yearsOfExperience string
		var createdAt, updatedAt time.Time

		err = skillRows.Scan(&id, &tagId, &evaluation, &yearsOfExperience, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		skillID, err := userdm.NewSkillIDByVal(id)
		if err != nil {
			return nil, err
		}

		tagID, err := tagdm.NewTagIDByVal(tagId)
		if err != nil {
			return nil, err
		}

		evaluationInt, err := strconv.Atoi(evaluation)
		if err != nil {
			return nil, err
		}

		yearsOfExperienceInt, err := strconv.Atoi(yearsOfExperience)
		if err != nil {
			return nil, err
		}

		// evaluation, err := userdm.NewEvaluationByVal(evaluation)
		// if err != nil {
		// 	return nil, err
		// }

		// yearsOfExperience, err := userdm.NewYearsOfExperienceByVal(yearsOfExperience)
		// if err != nil {
		// 	return nil, err
		// }

		skill, err := userdm.NewSkillByVal(skillID, tagID, evaluationInt, yearsOfExperienceInt)
		if err != nil {
			return nil, err
		}
		skills = append(skills, *skill)
	}

	getCareersQuery := `
		SELECT id, detail, start_year, end_year, created_at, updated_at FROM careers WHERE user_id = $1
	`
	careerRows, err := conn.QueryContext(ctx, getCareersQuery, id.String())
	if err != nil {
		return nil, err
	}
	defer careerRows.Close()

	careers := make([]userdm.Career, 0)
	for careerRows.Next() {
		var id, detail, startYear, endYear string
		var createdAt, updatedAt time.Time

		err = careerRows.Scan(&id, &detail, &startYear, &endYear, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		careerID, err := userdm.NewCareerIDByVal(id)
		if err != nil {
			return nil, err
		}

		careerDetail, err := userdm.NewCareerDetailByVal(detail)
		if err != nil {
			return nil, err
		}

		startYearInt, err := strconv.Atoi(startYear)
		if err != nil {
			return nil, err
		}
		careerStartYear, err := userdm.NewCareerStartYearByVal(startYearInt)
		if err != nil {
			return nil, err
		}

		endYearInt, err := strconv.Atoi(endYear)
		if err != nil {
			return nil, err
		}
		careerEndYear, err := userdm.NewCareerEndYearByVal(endYearInt)
		if err != nil {
			return nil, err
		}

		career, err := userdm.NewCareerByVal(careerID, careerDetail, careerStartYear, careerEndYear)
		if err != nil {
			return nil, err
		}

		careers = append(careers, *career)
	}

	if !userRows.Next() {
		return nil, errors.New("user not found")
	}

	var userID, name, email, password, selfIntroduction string
	var createdAt, updatedAt time.Time

	err = userRows.Scan(&userID, &name, &email, &password, &selfIntroduction, &createdAt, &updatedAt)

	userIDByVal, err := userdm.NewUserIDByVal(userID)
	if err != nil {
		return nil, err
	}
	userNameByVal, err := userdm.NewUserNameByVal(name)
	if err != nil {
		return nil, err
	}

	passwordByVal, err := userdm.NewPasswordByVal(password)
	if err != nil {
		return nil, err
	}

	emailByVal, err := userdm.NewEmailByVal(email)
	if err != nil {
		return nil, err
	}

	selfIntroductionByVal, err := userdm.NewSelfIntroductionByVal(selfIntroduction)
	if err != nil {
		return nil, err
	}

	user, err := userdm.NewUserByVal(userIDByVal, userNameByVal, passwordByVal, emailByVal, skills, careers, &selfIntroductionByVal)
	if err != nil {
		return nil, err
	}

	return user, nil
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

	// 既存のスキルを更新するだけでなく、tag_idとuser_idの組み合わせがなかったら追加にしなくてはいけない。
	if len(user.Skills()) > 0 {
		skillQuery := `
			UPDATE skills SET tag_id = $1, evaluation = $2, years_of_experience = $3, updated_at = NOW() WHERE user_id = $4
		`
		for _, skill := range user.Skills() {
			_, err := conn.ExecContext(ctx, skillQuery,
				user.ID().String(),
				skill.TagID().String(),
				skill.Evaluation(),
				skill.YearsOfExperience(),
			)
			if err != nil {
				return err
			}
		}
	}

	// こちらも追加・更新を2つやる必要あり。
	if len(user.Careers()) > 0 {
		careerQuery := `
			UPDATE careers SET detail = $1, start_year = $2, end_year = $3, updated_at = NOW() WHERE user_id = $4
		`
		for _, career := range user.Careers() {
			_, err := conn.ExecContext(ctx, careerQuery,
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

	return nil
}
