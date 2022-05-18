package model

import "gorm.io/gorm"

type MovieReview struct {
	gorm.Model
	Username string `gorm:"ForeignKey:UserUsername;Association_ForeignKey:Username"`
	MovieNum int    `gorm:"column:movie_num;ForeignKey:MovieInfoNum;AssociationForeignKey:Num"`
	Title    string `gorm:"type:varchar(30)"`
	Txt      string
}

type ShortReview struct {
	gorm.Model
	Username string `gorm:"ForeignKey:UserUsername;Association_ForeignKey:Username"`
	MovieNum int    `gorm:"column:movie_num;ForeignKey:MovieInfoNum;AssociationForeignKey:Num"`
	Title    string `gorm:"type:varchar(30)"`
	Txt      string
}

type CommentArea struct {
	gorm.Model
	Username string `gorm:"ForeignKey:UserUsername;Association_ForeignKey:Username"`
	MovieNum int    `gorm:"column:movie_num;ForeignKey:MovieInfoNum;AssociationForeignKey:Num"`
	Topic    string `gorm:"type:varchar(30)"`
	LikeNum  int    `gorm:"column:like_num;default:0"`
}

type Comment struct {
	gorm.Model
	Username  string `gorm:"ForeignKey:UserUsername;Association_ForeignKey:Username"`
	CommentId int    `gorm:"column:comment_id;ForeignKey:CommentAreaID;AssociationForeignKey:ID"`
	Txt       string `gorm:"type:varchar(30)"`
	LikeNum   int    `gorm:"column:like_num;default:0"`
}

type TopicLike struct {
	gorm.Model
	Username string `gorm:"ForeignKey:UserUsername;Association_ForeignKey:Username"`
	TopicId  int    `gorm:"column:topic_id;ForeignKey:CommentAreaId;AssociationForeignKey:Id"`
	MovieNum int    `gorm:"column:movie_num;ForeignKey:MovieInfoNum;AssociationForeignKey:Num"`
}

type CommentLike struct {
	gorm.Model
	Username  string `gorm:"ForeignKey:UserUsername;Association_ForeignKey:Username"`
	CommentId int    `gorm:"column:topic_id;ForeignKey:CommentId;AssociationForeignKey:Id"`
	MovieNum  int    `gorm:"column:movie_num;ForeignKey:MovieInfoNum;AssociationForeignKey:Num"`
}
