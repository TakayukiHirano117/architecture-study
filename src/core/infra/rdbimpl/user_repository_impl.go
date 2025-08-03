package rdbimpl

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/TakayukiHirano117/architecture-study/config"
	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
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

type userDTO struct {
	ID               string    `db:"id"`
	Name             string    `db:"name"`
	Email            string    `db:"email"`
	Password         string    `db:"password"`
	SelfIntroduction string    `db:"self_introduction"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}

func (dto *userDTO) fromDomain(user *userdm.User) {
	dto.ID = user.Id().String()
	dto.Name = string(user.Name())
	dto.Email = string(user.Email())
	dto.Password = string(user.Password())
	dto.SelfIntroduction = string(user.SelfIntroduction())
	dto.CreatedAt = user.CreatedAt()
	dto.UpdatedAt = user.UpdatedAt()
}

var (
	users []*userdm.User
)

func (r *UserRepositoryImpl) FindByName(ctx context.Context, name userdm.UserName) (*userdm.User, error) {
	query := `
		SELECT * FROM users WHERE name = :name
	`
	rows, err := r.Connect.NamedQueryContext(ctx, query, map[string]interface{}{"name": name})
	if err != nil {
		return nil, err
	}
	// DBからとってきた値でreconstruct

	for _, user := range users {
		if user.Name().Equal(name) {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *UserRepositoryImpl) Store(ctx context.Context, user *userdm.User) error {
	// users = append(users, user)
	dto := &userDTO{}
	dto.fromDomain(user)
	query := `
		INSERT INTO users (id, name, email, password, self_introduction, created_at, updated_at)
		VALUES (:id, :name, :email, :password, :self_introduction, NOW(), NOW())
	`
	_, err := r.Connect.NamedExecContext(ctx, query, dto)
	if err != nil {
		return err
	}
	return nil
}
