package model

type User struct {
	Username string `gorm:"primaryKey;not null;unique;type:varchar(20)"`
	Password string `gorm:"type:varchar(100)"`
	Nickname string `gorm:"type:varchar(20)"`
}

type UserMenu struct {
	Username  string `gorm:"primaryKey;not null;ForeignKey:UserUsername;AssociationForeignKey:Username"`
	Introduce string `gorm:"default:''"`
	Avatar    string `gorm:"default:'http://49.235.99.195:8080/pictures/1644517805test.png'"`
}

type UserEncrypted struct {
	uID      uint   `gorm:"primaryKey;AUTO_INCREMENT=1;not null"`
	Username string `gorm:"ForeignKey:UserUsername;AssociationForeignKey:Username"`
	Question string
	Answer   string
}

type UserWantSee struct {
	uID      int    `gorm:"primaryKey;AUTO_INCREMENT=1;not null"`
	Username string `gorm:"ForeignKey:UserUsername;AssociationForeignKey:Username"`
	Num      int    `gorm:"ForeignKey:MovieInfoNum;AssociationForeignKey:Num"`
	Txt      string
	Label    string
}

type UserSeen struct {
	uID      int    `gorm:"primaryKey;AUTO_INCREMENT=1;not null"`
	Username string `gorm:"ForeignKey:UserUsername;AssociationForeignKey:Username"`
	Num      int    `gorm:"ForeignKey:MovieInfoNum;AssociationForeignKey:Num"`
	Txt      string
	Label    string
}
