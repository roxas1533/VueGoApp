package handler

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"

	"github.com/labstack/echo"
)

// LoginParam は /yamabiko が受けとるJSONパラメータを定義します。
type LoginParam struct {
	Adress   string `json:"mail"`
	Password string `json:"pass"`
}

//RegisterParam は登録時のパラメーターを定義します
type RegisterParam struct {
	Name     string `json:"name"`
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

type UpdateUserContent struct {
	UserName     string `json:"UserName"`
	ProfileImage string `json:"ProfileImage"`
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

// LoginAPI はログイン時のPost時のJSONデータ生成処理を行います。
func LoginAPI(db *sql.DB) echo.HandlerFunc {
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
			if param.Password != user.password {
				return c.JSON(http.StatusOK, map[string]interface{}{"reslut": "false"})
			}
		} else {
			return c.JSON(http.StatusOK, map[string]interface{}{"reslut": "false"})
		}
		// auth.getToken()
		token := getToken(user.name, strconv.Itoa(user.id))
		return c.JSON(http.StatusOK, map[string]interface{}{"reslut": "ok", "JWT": token, "userName": user.name, "userID": user.id})
	}
}

//RegisterAPI 登録時にDBサーバに問い合わせます
func RegisterAPI(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(RegisterParam)

		if err := c.Bind(param); err != nil {
			return err
		}
		// var user UserInfo
		var count int
		db.QueryRow("select COUNT(*) FROM users where mail_adress=?;", param.Adress).Scan(&count)
		if count != 0 {
			return c.JSON(http.StatusOK, map[string]interface{}{"result": "already"})
		}

		var id int
		var address string
		var password string
		var name string
		var createdTime string
		var updatedTime string
		db.Exec("insert into users(mail_adress,password,name,created_time,updated_time) values(?,?,?,default,default)", param.Adress, param.Password, param.Name)
		db.QueryRow("select * FROM users order by id desc limit 1;").Scan(&id, &address, &password, &name, &createdTime, &updatedTime)
		if count != 0 {
			return c.JSON(http.StatusOK, map[string]interface{}{"result": "already"})
		}
		token := getToken(name, strconv.Itoa(id))
		return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok", "JWT": token})
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
		db.Exec(`insert into timeline(created_time,content,userid) values(default,?,?)`, content.Content, claims["sub"])
		row := db.QueryRow(`select timeline.id,timeline.created_time,content,name 
		from timeline inner join users on timeline.userid=users.id order by id desc limit 1`)
		var id int
		var name string
		var time string
		var talkContent string
		if err := row.Scan(&id, &time, &talkContent, &name); err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusOK, map[string]interface{}{"result": "no"})
		}
		// // Write
		liveTalkContent := new(LiveTalkContent)
		liveTalkContent.Type = "push"
		liveTalkContent.ID = 1
		liveTalkContent.Name = name
		liveTalkContent.Content = content.Content
		liveTalkContent.Time = time
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

//GetTimeLine タイムライン取得
func GetTimeLine(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		from := c.Param("from")
		var rows *sql.Rows
		var err error
		if from == "0" {
			rows, err = db.Query(`select timeline.id,timeline.created_time,content,timeline.userid,name 
			from timeline inner join users on timeline.userid=users.id order by id desc limit ?`, id)
		} else {
			rows, err = db.Query(`select timeline.id,timeline.created_time,content,timeline.userid,name 
			from timeline inner join users on timeline.userid=users.id where timeline.id<? order by id desc limit ?;`, from, id)
		}
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		return returnTalkContents(c, rows)
	}
}

func GetTimeLineUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userid := c.Param("userid")
		id := c.Param("id")
		from := c.Param("from")
		var rows *sql.Rows
		var err error
		if from == "0" {
			rows, err = db.Query(`select timeline.id,timeline.created_time,content,timeline.userid,name 
			from timeline inner join users on timeline.userid=users.id where timeline.userid=? order by id desc limit ?`, userid, id)
		} else {
			rows, err = db.Query(`select timeline.id,timeline.created_time,content,timeline.userid,name 
			from timeline inner join users on timeline.userid=users.id where timeline.id<? and timeline.userid=? order by id desc limit ?;`, from, userid, id)
		}
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		return returnTalkContents(c, rows)
	}
}

func returnTalkContents(c echo.Context, rows *sql.Rows) error {
	var liveTalkContents []LiveTalkContent
	for rows.Next() {
		var id int
		var name string
		var content string
		var time string
		var userid int
		if err := rows.Scan(&id, &time, &content, &userid, &name); err != nil {
			log.Fatal(err)
		}
		liveTalkContent := LiveTalkContent{"push", id, name, content, time, userid}
		liveTalkContents = append(liveTalkContents, liveTalkContent)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"result": liveTalkContents})
}

//UpdateUserInfo はユーザーデータをアップデートします。
func UpdateUserInfo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ucontent := new(UpdateUserContent)
		if err := c.Bind(ucontent); err != nil {
			return err
		}
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		splited := strings.Split(ucontent.ProfileImage, ",")
		if len(splited) > 1 {
			Imagebyte, err := base64.StdEncoding.DecodeString(splited[1])
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"result": "faild"})
			}
			w, err := os.Create(claims["sub"].(string) + ".png")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"result": "failed"})
			}
			defer w.Close()
			w.Write(Imagebyte)
		}
		fmt.Println(ucontent.UserName, claims["sub"])
		r, _ := db.Exec(`update users set name=? where id=?;`, ucontent.UserName, claims["sub"])
		fmt.Println(r.RowsAffected())
		return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})

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

//GetUserProfileImg はユーザアイコンを取得します
func GetUserProfileImg() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		_, err := os.Stat(id)
		fmt.Println(id)
		if err == nil {
			return c.File(id)
		}
		return c.File("default_user.png")
	}
}
