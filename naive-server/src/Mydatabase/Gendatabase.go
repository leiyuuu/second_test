package Mydatabase

import (
	"github.com/leiyuuu/second-test/naive-server/src/loggenerator"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Username     string `gorm:"default:noname"`
	Password     string `gorm:"default:nopassword"`
	Access_token string `gorm:"default:notoken"`
	Has_checkin  bool   `gorm:"default:false"`
	Has_signin   bool   `gorm:"default:false"`
	Can_delete   bool   `gorm:"default:true"`
} //注意必须要大写！！！！！！！！！！！！！！！！！！！

var DB *gorm.DB
var Auto_sign_in bool

func Gendatabase() { //生成 迁移 数据库
	db, err := gorm.Open(sqlite.Open("people.db"), &gorm.Config{})
	if err != nil {
		loggenerator.Fatal("Generate Database Error:can't open or generate the database")
	}
	loggenerator.Trace("gen database succeed")
	db.AutoMigrate(&Person{})
	loggenerator.Trace("migrate database succeed")
	DB = db
	// db.Where("username = ?", "ZLY").Delete(&Person{})
}

func Set_auto_sign_in(bol bool) {
	Auto_sign_in = bol
}
