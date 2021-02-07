package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// LoginParam は /yamabiko が受けとるJSONパラメータを定義します。
type LoginParam struct {
	Adress   string `json:"mail"`
	Password string `json:"pass"`
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
				fmt.Println("パスワードが違います", param)
				return c.JSON(http.StatusOK, map[string]interface{}{"reslut": "false"})
			}
		} else {
			fmt.Println("これ", param)
			return c.JSON(http.StatusOK, map[string]interface{}{"reslut": "false"})
		}
		// auth.getToken()
		fmt.Println(param)
		token := getToken(user.name, strconv.Itoa(user.id))
		return c.JSON(http.StatusOK, map[string]interface{}{"reslut": param.Adress, "JWT": token})
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

func GetHomeAPI() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
	}
}
