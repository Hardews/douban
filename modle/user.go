package modle

type User struct {
	NickName  string `form:"nikeName"`
	Username  string `form:"signUsername"`
	Password  string `form:"signPassword"`
	Introduce string
}

type UserHistory struct {
	MovieNum int
	Comment  string
	Label    string
	Url      string
}
