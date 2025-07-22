package userdm

import "time"

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
	careers          []Career
	email            Email
	selfIntroduction SelfIntroduction
<<<<<<< HEAD
>>>>>>> dfd255d (バリューオブジェクト実装途中、一旦コミット)
=======
	createdAt        time.Time
	updatedAt        time.Time
>>>>>>> 68d42e9 (値オブジェクトやエンティティなど実装、一旦コミット)
}

func NewUser(name UserName, password Password, skills []Skill, careers []Career, email Email, selfIntroduction SelfIntroduction) (*User, error) {
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
