package modle

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfoMenu struct {
	Introduce string
}

type UserHistory struct {
	MovieNum int
	Comment  string
	Label    string
}
