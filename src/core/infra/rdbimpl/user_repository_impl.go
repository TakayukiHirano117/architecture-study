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

// TODO: これをデータモデルに置き換える
// type userDTO struct {
// 	ID               string      `db:"id"`
// 	Name             string      `db:"name"`
// 	Password         string      `db:"password"`
// 	Email            string      `db:"email"`
// 	Skills           []skillDTO  `db:"skills"`
// 	Careers          []careerDTO `db:"careers"`
// 	SelfIntroduction string      `db:"self_introduction"`
// 	CreatedAt        time.Time   `db:"created_at"`
// 	UpdatedAt        time.Time   `db:"updated_at"`
// }

// type skillDTO struct {
// 	TagId             string `db:"tag_id"`
// 	Evaluation        int    `db:"evaluation"`
// 	YearsOfExperience int    `db:"years_of_experience"`
// }

// type careerDTO struct {
// 	Detail    string `db:"detail"`
// 	StartYear int    `db:"start_year"`
// 	EndYear   int    `db:"end_year"`
// }

// func (dto *userDTO) fromDomain(user *userdm.User) {
// 	dto.ID = user.Id().String()
// 	dto.Name = string(user.Name())
// 	dto.Email = string(user.Email())
// 	dto.Password = string(user.Password())
// 	dto.SelfIntroduction = string(user.SelfIntroduction())
// 	dto.CreatedAt = user.CreatedAt()
// 	dto.UpdatedAt = user.UpdatedAt()
// }

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
	_, err := r.Connect.NamedExecContext(ctx, query, userModel)
	if err != nil {
		return err
	}
	return nil
}
