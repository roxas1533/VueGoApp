package main

import (
	"database/sql"

	"work/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(vuegosample_sqldb_1:3306)/sample_db")
	if err != nil {
		panic(nil)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.POST("/login", handler.YamabikoAPI(db))
	e.POST("/register", handler.RegisterAPI(db))
	r := e.Group("/home")
	r.Use(middleware.JWT([]byte(handler.Secret)))
	r.POST("", handler.GetHomeAPI())

	// サーバー起動
	e.Start(":8000")
}
