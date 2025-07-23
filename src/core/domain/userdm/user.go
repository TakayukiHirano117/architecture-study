package userdm

import (
	"time"

	"github.com/cockroachdb/errors"
)

type User struct {
<<<<<<< HEAD
<<<<<<< HEAD
  id UserID
=======
  id UserId
>>>>>>> d99dc49 (fix conflicts)
  name string
  password Password
  
=======
	id               UserId
	name             UserName
	password         Password
	skills           []Skill
	careers          []*Career
	email            Email
<<<<<<< HEAD
	selfIntroduction SelfIntroduction
<<<<<<< HEAD
>>>>>>> dfd255d (バリューオブジェクト実装途中、一旦コミット)
=======
=======
	selfIntroduction *SelfIntroduction
>>>>>>> 6b0b5cf (エンティティと値オブジェクトの実装)
	createdAt        time.Time
	updatedAt        time.Time
>>>>>>> 68d42e9 (値オブジェクトやエンティティなど実装、一旦コミット)
}

func NewUser(name UserName, password Password, skills []Skill, careers []*Career, email Email, selfIntroduction *SelfIntroduction) (*User, error) {
	if len(skills) <= 0 {
		return nil, errors.New("skills must be at least 1")
	}

	return &User{
		name:             name,
		password:         password,
		email:            email,
		selfIntroduction: selfIntroduction,
		skills:           skills,
		careers:          careers,
		createdAt:        time.Now(),
		updatedAt:        time.Now(),
	}, nil
}
