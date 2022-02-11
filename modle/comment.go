package modle

type UserComment struct {
	Topic     string
	MovieNum  int
	MovieName string
	Username  string
	Txt       string
	Time      string
	LikeNum   int
	Url       string
}

type CommentArea struct {
	Num        int
	MovieNum   int
	Url        string
	CommentNum int
	Topic      string
	Username   string
	Time       string
	Comment    string
	LikeNum    int
}
