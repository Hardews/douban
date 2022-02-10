package modle

type UserComment struct {
	Topic     string
	MovieNum  int
	MovieName string
	Username  string
	Txt       string
	Time      string
	LikeNum   int
}

type CommentArea struct {
	Num        int
	MovieNum   int
	CommentNum int
	Topic      string
	Username   string
	Time       string
	Comment    string
	LikeNum    int
}
