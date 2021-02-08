package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"

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

//TalkContent 受信の時
type TalkContent struct {
	Content string `json:"Content"`
}

//LiveTalkContent 配信するとき
type LiveTalkContent struct {
	Type    string `json:"Type"`
	ID      int    `json:"ID"`
	Name    string `json:"name"`
	Content string `json:"Content"`
	Time    string `json:"Time"`
	UserID  int    `json:"UserID"`
}

// WebSocket 更新用
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//CustomMiddleware WebScoket時にトークンを検証する。
func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if CheckToken(c.Request().Header["Sec-Websocket-Protocol"][0]) {
			return next(c)
		}
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
}

//CheckHeader リクエストヘッダの確認用
func CheckHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Request().Header)
		return next(c)
	}
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
				return c.JSON(http.StatusOK, map[string]interface{}{"reslut": "false"})
			}
		} else {
			return c.JSON(http.StatusOK, map[string]interface{}{"reslut": "false"})
		}
		// auth.getToken()
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

var client = make(map[*websocket.Conn]bool)

//TalkAPI タイムライン投稿API
func TalkAPI(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		content := new(TalkContent)
		if err := c.Bind(content); err != nil {
			return err
		}
		_, err := db.Exec("insert into timeline(name,created_time,content,userid) values(?,default,?,?)", claims["name"], content.Content, claims["sub"])
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{"result": "no"})
		}
		// // Write
		liveTalkContent := new(LiveTalkContent)
		liveTalkContent.Type = "push"
		liveTalkContent.ID = 1
		liveTalkContent.Name = claims["name"].(string)
		liveTalkContent.Content = content.Content
		liveTalkContent.Time = claims["iat"].(string)
		liveTalkContent.UserID, _ = strconv.Atoi(claims["sub"].(string))
		jsonString, _ := json.Marshal(liveTalkContent)
		for cl := range client {
			err := cl.WriteMessage(websocket.TextMessage, []byte(jsonString))
			if err != nil {
				c.Logger().Error(err)
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
	}
}

func GetTimeLine(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		fmt.Println(id)
		rows, err := db.Query("select * FROM timeline limit ?;", id)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		var liveTalkContents []LiveTalkContent
		for rows.Next() {
			var id int
			var name string
			var content string
			var time string
			var userid int
			if err := rows.Scan(&id, &name, &time, &content, &userid); err != nil {
				log.Fatal(err)
			}
			liveTalkContent := LiveTalkContent{"push", id, name, content, time, userid}
			liveTalkContents = append(liveTalkContents, liveTalkContent)
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"result": liveTalkContents})
	}
}

//WebsocketServer タイムライン配信用
func WebsocketServer(c echo.Context) error {
	c.Response().Header().Set("Sec-Websocket-Protocol", c.Request().Header["Sec-Websocket-Protocol"][0])
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())

	if err != nil {
		return err
	}
	defer ws.Close()
	client[ws] = true

	for {
		// Read
		_, _, err := ws.ReadMessage()
		if err != nil {
			delete(client, ws)
			return err
		}
	}

}
