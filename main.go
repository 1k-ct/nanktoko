package main

import (
	"strconv"
	"youtubeWeb774inc/db"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	dbdef := db.ConnectGorm()
	defer dbdef.Close()

	router := gin.Default()
	//router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		tokos := db.GetDBContents()
		ctx.HTML(200, "index.html", gin.H{
			"tokos": tokos,
		})
	})
	router.POST("/creat", func(ctx *gin.Context) {
		text := ctx.PostForm("toko_text")
		name := ctx.PostForm("toko_user_name")
		//createdAt = ctx.PostForm("createdAt")
		db.AddNewInDB(text, name) //, createdAt
		ctx.Redirect(302, "/")
	})
	router.POST("/delete", func(ctx *gin.Context) {
		n := ctx.PostForm("db_delete_num")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		db.DeleteDB(id)
		ctx.Redirect(302, "/")
	})
	router.Run()

}
