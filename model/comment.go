package model

type UserComment struct {
	Topic    string
	MovieNum int
	Username string
	Txt      string
	Time     string
	LikeNum  int
	Url      string
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
