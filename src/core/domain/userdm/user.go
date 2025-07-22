package userdm

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
>>>>>>> dfd255d (バリューオブジェクト実装途中、一旦コミット)
}

func NewUser(name UserName, password Password, skills []Skill, careers []Career, email Email, selfIntroduction SelfIntroduction) (*User, error) {
	return &User{
		name:             name,
		password:         password,
		email:            email,
		selfIntroduction: selfIntroduction,
		skills:           skills,
		careers:          careers,
	}, nil
}
