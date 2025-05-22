package ws

import (
	//"encoding/json"
	"net/http"
	//"mind-set/internal/model"
	"sync"

	"github.com/gorilla/websocket"
)

type User struct {
	Conn        *websocket.Conn
	Name        string
	Token       string
	AccessToken string
	Id          int
	//Cid         int
	//Avatar      string
	//Role        Role
	Mux sync.Mutex
}

var ClientList = make(map[string]*User)
var upgrader = websocket.Upgrader{}
var Mux sync.RWMutex

func init() {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

// 添加用户
func AddClientToList(user *User) {
	oldUser, ok := ClientList[user.AccessToken]
	if oldUser != nil || ok {
		oldUser.Conn.Close()
		delete(ClientList, user.AccessToken)
	}
	ClientList[user.AccessToken] = user
}

/*
// 发送消息
func SendMessage(AccessToken string, msg *model.Chat) {

	str, _ := json.Marshal(msg)
	user, ok := ClientList[AccessToken]
	if !ok || user == nil || user.Conn == nil {
		return
	}
	user.Mux.Lock()
	defer user.Mux.Unlock()
	user.Conn.WriteMessage(websocket.TextMessage, str)
}
*/
