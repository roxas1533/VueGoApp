package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// LoginParam は /yamabiko が受けとるJSONパラメータを定義します。
type LoginParam struct {
	Adress   string `json:"adress"`
	Password string `json:"password"`
}

// UserInfo ログイン時のユーザーデータ
type UserInfo struct {
	id       int
	address  string
	password string
	name     string
	cTime    string
	uTime    string
}

// YamabikoAPI は /api/hello のPost時のJSONデータ生成処理を行います。
func YamabikoAPI(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(LoginParam)

		if err := c.Bind(param); err != nil {
			return err
		}
		var user UserInfo
		var count int
		db.QueryRow("select COUNT(*) FROM users where mail_adress=?;", param.Adress).Scan(&count)
		if count != 0 {
			row := db.QueryRow("SELECT * FROM users where mail_adress=?", param.Adress)
			err := row.Scan(&user.id, &user.address, &user.password, &user.name, &user.cTime, &user.uTime)
			if err != nil {
				panic(err)
			}
			fmt.Println(user.id, user.address, user.password, user.name, user.cTime, user.uTime)
			if param.Password != user.password {
				return c.JSON(http.StatusOK, map[string]interface{}{"reslut": nil})
			}
		} else {
			return c.JSON(http.StatusOK, map[string]interface{}{"reslut": nil})
		}
		fmt.Println("存在確認")
		return c.JSON(http.StatusOK, map[string]interface{}{"reslut": param.Adress})
	}
}

//RegisterAPI 登録時にDBサーバに問い合わせます
func RegisterAPI(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(LoginParam)

		if err := c.Bind(param); err != nil {
			return err
		}
		// var user UserInfo
		var count int
		db.QueryRow("select COUNT(*) FROM users where mail_adress=?;", param.Adress).Scan(&count)
		if count != 0 {
			return c.JSON(http.StatusOK, map[string]interface{}{"result": "already"})
		}
		fmt.Println(param)
		return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
	}
}
