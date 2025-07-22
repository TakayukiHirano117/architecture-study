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
	email            Email
	selfIntroduction SelfIntroduction
>>>>>>> dfd255d (バリューオブジェクト実装途中、一旦コミット)
}
