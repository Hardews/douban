package modle

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfoMenu struct {
	Introduce   string
	FilmCritics string
	WantSee     string
	Seen        string
}
