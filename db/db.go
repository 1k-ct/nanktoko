package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Toko dbの投稿内容 今から作成するもの
type Toko struct {
	ID   int
	Name string
	Text string
	// CreatedAd string or time.Time
}

//ConnectGorm connect db
func ConnectGorm() *gorm.DB { // 下のところは自分のものに変更してください
	db, err := gorm.Open("mysql", "user_name:password@/database_name")
	if err != nil {
		panic(err)
	}
	return db
}

// AddNewInDB DBに新しく追加する
func AddNewInDB(text string, name string) { //, createdAt string
	db := ConnectGorm()
	db.Create(&Toko{Text: text, Name: name}) //, CreatedAt: createdAt
	defer db.Close()
}

// GetDBContents DBの全ての投稿を取得する
func GetDBContents() []Toko {
	db := ConnectGorm()
	var tokos []Toko
	db.Find(&tokos)
	db.Close()
	return tokos
}

// DeleteDB 選択したidをDBから削除
func DeleteDB(id int) {
	db := ConnectGorm()
	var toko Toko
	db.First(&toko, id)
	db.Delete(&toko)
	db.Close()
}
