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
		token := CheckToken(c.Request().Header["Sec-Websocket-Protocol"][0])
		if token != nil {
			c.Set("user", token)
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

var (
	globalClient = make(map[*websocket.Conn]bool)
	homeClient   = make(map[*websocket.Conn]map[int]bool)
)

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
		for cl := range globalClient {
			err := cl.WriteMessage(websocket.TextMessage, []byte(jsonString))
			if err != nil {
				c.Logger().Error(err)
			}
		}
		for cl, v := range homeClient {
			for id := range v {
				if strconv.Itoa(id) == claims["sub"].(string) {
					err := cl.WriteMessage(websocket.TextMessage, []byte(jsonString))
					if err != nil {
						c.Logger().Error(err)
					}
				}
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
	}
}

func getFavoritList(db *sql.DB, id string, num string, from string) []int {
	var favoritelist []int
	rows, err := db.Query(`select contentid from favorites where userid=? and contentid<? order by contentid desc limit ?`, id, from, num)
	if err != nil {
		return favoritelist
	}
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		favoritelist = append(favoritelist, id)
	}
	return favoritelist

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
			from = "1000000"
		} else {
			rows, err = db.Query(`select timeline.id,timeline.created_time,content,timeline.userid,name 
			from timeline inner join users on timeline.userid=users.id where timeline.id<? order by id desc limit ?;`, from, id)
		}
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		fmt.Println()
		return returnTalkContents(c, rows, db, id, from)
	}
}

//GetUsersTimeLine 認証済みユーザーのタイムライン取得
func GetUsersTimeLine(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		from := c.Param("from")
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		var rows *sql.Rows
		var err error
		if from == "0" {
			rows, err = db.Query(`select timeline.id,timeline.created_time,content,timeline.userid,name from follows join timeline on 
			timeline.userid=follows.followid and follows.userid=? 
			join users on timeline.userid=users.id order by timeline.id desc limit ?;`, claims["sub"], id)
			from = "1000000"
		} else {
			rows, err = db.Query(`select timeline.id,timeline.created_time,content,timeline.userid,name from follows join timeline on 
			timeline.userid=follows.followid and follows.userid=? 
			join users on timeline.userid=users.id where timeline.id<? order by timeline.id desc limit ?;`, claims["sub"], from, id)
		}
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		return returnTalkContents(c, rows, db, id, from)
	}
}

//GetTimeLineUser 指定したユーザーのツイート取得
func GetTimeLineUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userid := c.Param("userid")
		id := c.Param("id")
		from := c.Param("from")

		var rows *sql.Rows
		var err error
		if from == "0" {
			rows, err = db.Query(`select timeline.id,timeline.created_time,content,timeline.userid,name 
			from timeline inner join users on timeline.userid=users.id and timeline.userid=? order by id desc limit ?`, userid, id)
			from = "1000000"
		} else {
			rows, err = db.Query(`select timeline.id,timeline.created_time,content,timeline.userid,name 
			from timeline inner join users on timeline.userid=users.id 
			where timeline.id<? and timeline.userid=? order by id desc limit ?;`, from, userid, id)
		}
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		return returnTalkContents(c, rows, db, id, from)
	}
}

func returnTalkContents(c echo.Context, rows *sql.Rows, db *sql.DB, id string, from string) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
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
	return c.JSON(http.StatusOK, map[string]interface{}{
		"result":   liveTalkContents,
		"favolist": getFavoritList(db, claims["sub"].(string), id, from),
	})
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
		r, _ := db.Exec(`update users set name=? where id=?;`, ucontent.UserName, claims["sub"])
		fmt.Println(r.RowsAffected())
		return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})

	}
}

//WebsocketGlobalServer タイムライン配信用
func WebsocketGlobalServer(c echo.Context) error {
	c.Response().Header().Set("Sec-Websocket-Protocol", c.Request().Header["Sec-Websocket-Protocol"][0])
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())

	if err != nil {
		return err
	}
	defer ws.Close()
	globalClient[ws] = true
	// claims := user.Claims.(jwt.MapClaims)
	for {
		// Read
		_, _, err := ws.ReadMessage()
		if err != nil {
			delete(globalClient, ws)
			return err
		}
	}
}

//WebsocketHomeServer タイムライン配信用
func WebsocketHomeServer(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Sec-Websocket-Protocol", c.Request().Header["Sec-Websocket-Protocol"][0])
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())

		if err != nil {
			return err
		}
		defer ws.Close()
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		rows, err := db.Query(`select followid from follows where userid=?;`, claims["sub"])
		if len(homeClient) == 0 {
			homeClient[ws] = map[int]bool{}
		} else {
			homeClient[ws] = make(map[int]bool)
		}
		for rows.Next() {
			var id int
			if err := rows.Scan(&id); err != nil {
				log.Fatal(err)
			}
			homeClient[ws][id] = true
		}
		for {
			_, _, err := ws.ReadMessage()
			if err != nil {
				delete(homeClient, ws)
				return err
			}
		}
	}
}

//GetUserProfileImg はユーザアイコンを取得します
func GetUserProfileImg() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		_, err := os.Stat(id)
		if err == nil {
			return c.File(id)
		}
		return c.File("default_user.png")
	}
}

//IsFollow は認証ユーザーが指定したユーザーをフォローしてるか確認します。
func IsFollow(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var count int
		id := c.Param("id")
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		db.QueryRow("select COUNT(*) FROM follows where userid=? and followid=?;", claims["sub"], id).Scan(&count)
		if count == 0 {
			return c.JSON(http.StatusOK, map[string]interface{}{"result": false})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"result": true})
	}
}

//Follow は指定したユーザをフォローします。
func Follow(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var count int
		id := c.Param("id")
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		if id == claims["sub"] {
			return c.JSON(http.StatusOK, map[string]interface{}{"result": false})
		}
		db.QueryRow("select COUNT(*) FROM follows where userid=? and followid=?;", claims["sub"], id).Scan(&count)
		if count != 0 {
			return c.JSON(http.StatusOK, map[string]interface{}{"result": false})
		}
		_, error := db.Exec("insert into follows(userid,followid,created_time) value(?,?,default)", claims["sub"], id)
		if error != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{"result": false})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"result": true})

	}
}

//UnFollow は指定したユーザのフォローを解除します。
func UnFollow(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var count int
		id := c.Param("id")
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		db.QueryRow("select COUNT(*) FROM follows where userid=? and followid=?;", claims["sub"], id).Scan(&count)
		if count == 0 {
			return c.JSON(http.StatusOK, map[string]interface{}{"result": false})
		}
		_, error := db.Exec("delete from follows where userid=? and followid=?;", claims["sub"], id)
		if error != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{"result": false})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"result": true})

	}
}

//TweetCount は指定したユーザのツイート数を取得します。
func TweetCount(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var count int
		id := c.Param("id")
		db.QueryRow("select COUNT(*) FROM timeline where userid=?", id).Scan(&count)
		return c.JSON(http.StatusOK, map[string]interface{}{"result": true, "count": count})
	}
}

//GetFollowNumber は指定したユーザのフォロー人数を取得します。
func GetFollowNumber(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var count int
		id := c.Param("id")
		db.QueryRow("select COUNT(*) FROM follows where userid=?", id).Scan(&count)
		return c.JSON(http.StatusOK, map[string]interface{}{"result": true, "count": count})
	}
}

//GetFollowerNumber は指定したユーザのフォロワー人数を取得します。
func GetFollowerNumber(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var count int
		id := c.Param("id")
		db.QueryRow("select COUNT(*) FROM follows where followid=?", id).Scan(&count)
		return c.JSON(http.StatusOK, map[string]interface{}{"result": true, "count": count})
	}
}

//FavoritTalk は指定したtツイートにfavoriteします、すでにされていた場合取り消します。
func FavoritTalk(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		var count int
		db.QueryRow("select COUNT(*) from favorites where userid=? and contentid=?", claims["sub"], id).Scan(&count)
		fmt.Println(count)
		if count == 0 {
			db.Exec("insert into favorites(userid,contentid,created_time) values(?,?,default)", claims["sub"], id)
		} else {
			db.Exec("delete from favorites where userid=? and contentid=?", claims["sub"], id)
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"result": true})
	}
}
