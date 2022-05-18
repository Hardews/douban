package model

type MovieInfo struct {
	Num       int    `gorm:"primaryKey;AUTO_INCREMENT=1;not null"`
	Name      string `gorm:"type:varchar(20);not null"`
	Introduce string `gorm:"default:''"`
	Score     string `gorm:"default:0"`
	Area      string
	Year      string
	Types     string `gorm:"comments:类型"`
	Img       string
}

type MovieExtra struct {
	Num        int `gorm:"primaryKey;ForeignKey:MovieInfoNum;AssociationForeignKey:Num"`
	Time       int
	CommentNum int `gorm:"default:0"`
	WantSeeNum int `gorm:"default:0"`
	SeenNum    int `gorm:"default:0"`
}
