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
	e.POST("/login", handler.LoginAPI(db))
	e.POST("/register", handler.RegisterAPI(db))
	e.GET("/profile/:id", handler.GetUserProfileImg())
	talk := e.Group("")
	// talk.Use(handler.CheckHeader)
	talk.Use(middleware.JWT([]byte(handler.Secret)))
	talk.POST("/talk", handler.TalkAPI(db))
	talk.POST("/get/:id/:from", handler.GetTimeLine(db))
	talk.POST("/getUsersTimeLine/:id/:from", handler.GetUsersTimeLine(db))
	talk.POST("/getusers/:userid/:id/:from", handler.GetTimeLineUser(db))
	talk.POST("/update", handler.UpdateUserInfo(db))
	talk.POST("/follow/:id", handler.Follow(db))
	talk.POST("/unfollow/:id", handler.UnFollow(db))
	talk.POST("/tweetCount/:id", handler.TweetCount(db))
	talk.POST("/isFollow/:id", handler.IsFollow(db))
	talk.POST("/getFollowNumber/:id", handler.GetFollowNumber(db))
	talk.POST("/getFollowerNumber/:id", handler.GetFollowerNumber(db))
	talk.POST("/favorite/:id", handler.FavoritTalk(db))
	// talk.POST("/favoriteList/:id/:from", handler.getFavoritList(db))

	r := e.Group("/home")
	r.Use(handler.CustomMiddleware)
	r.GET("/streamGlobalTimeLine", handler.WebsocketGlobalServer)
	r.GET("/streamHomeTimeLine", handler.WebsocketHomeServer(db))

	// サーバー起動
	e.Logger.Fatal(e.Start(":8000"))
}
