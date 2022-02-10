package modle

type User struct {
	NickName  string `json:"nick_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Introduce string
}

type UserHistory struct {
	MovieNum int
	Comment  string
	Label    string
}
