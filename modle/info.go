package modle

type MovieInfo struct {
	Num int //编号

	Name      string
	OtherName string
	Score     string //评分
	Area      string
	Year      int    //出版年份
	Time      string //时长
	Starring  string //主演
	Director  string //导演
	Types     string //类型

	Introduce  string //简介
	CommentNum int
	Writer     string //编剧
	Language   string
	Cast       string //演员名单

	FUsername   string
	FTime       string
	FilmCritics string //影评
	EUsername   string
	ETime       string
	Essay       string //短评
	WantSee     int
	Seen        int //看过
}
