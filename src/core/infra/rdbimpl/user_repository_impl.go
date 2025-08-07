package rdbimpl

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/TakayukiHirano117/architecture-study/config"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
	"github.com/TakayukiHirano117/architecture-study/src/core/infra/models"
)

type UserRepositoryImpl struct {
	Connect *sqlx.DB
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	dbConfig := config.NewDBConfig()
	db, err := dbConfig.Connect()

	if err != nil {
		panic(errors.New("failed to connect to database: " + err.Error()))
	}

	return &UserRepositoryImpl{Connect: db}
}

func (r *UserRepositoryImpl) FindByName(ctx context.Context, name userdm.UserName) (*userdm.User, error) {
	query := `
		SELECT * FROM users WHERE name = :name
	`
	rows, err := r.Connect.NamedQueryContext(ctx, query, map[string]interface{}{"name": name})
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		return nil, errors.New("user name already exists")
	}

	return nil, nil
}

func (r *UserRepositoryImpl) Store(ctx context.Context, user *userdm.User) error {
	skills := make([]models.SkillModel, len(user.Skills()))
	for i, skill := range user.Skills() {
		skills[i] = models.SkillModel{
			TagId:             skill.TagId().String(),
			Evaluation:        skill.Evaluation(),
			YearsOfExperience: skill.YearsOfExperience(),
		}
	}

	careers := make([]models.CareerModel, len(user.Careers()))
	for i, career := range user.Careers() {
		careers[i] = models.CareerModel{
			Detail:    career.Detail().String(),
			StartYear: career.StartYear().Int(),
			EndYear:   career.EndYear().Int(),
		}
	}

	userModel := &models.UserModel{
		ID:               user.Id().String(),
		Name:             user.Name().String(),
		Email:            user.Email().String(),
		Password:         user.Password().String(),
		Skills:           skills,
		Careers:          careers,
		SelfIntroduction: user.SelfIntroduction().String(),
		CreatedAt:        user.CreatedAt(),
		UpdatedAt:        user.UpdatedAt(),
	}

	query := `
		INSERT INTO users (id, name, email, password, self_introduction, created_at, updated_at)
		VALUES (:id, :name, :email, :password, :self_introduction, NOW(), NOW())
	`
	// TODO: skillsとcareersをinsertする
	_, err := r.Connect.NamedExecContext(ctx, query, userModel)
	if err != nil {
		return err
	}
	return nil
}
