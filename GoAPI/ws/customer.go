package ws

import (
	"fmt"
	"log"
	"mind-set/internal/utils"
	"mind-set/internal/utils/gen_token"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func NewStudentServer(c *gin.Context) {
	student, _ := c.Get("user")
	userinfo := student.(gen_token.UserInfo)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	log.Println("user connect: ", userinfo.Name)
	defer conn.Close()
	hashAccessToken := utils.GetMD5Hash(userinfo.Token)
	fmt.Println("NewStudentServer-----", hashAccessToken)
	users := &User{
		Conn: conn,
		Name: userinfo.Name,
		//Avatar:      userinfo.Avatar,
		Id:          userinfo.Id,
		Token:       userinfo.Token,
		AccessToken: hashAccessToken,
		//Role:        RoleStudent,
	}

	AddClientToList(users)

	for {
		var receive []byte
		messageType, receive, err := conn.ReadMessage()
		if err != nil {
			log.Println("ws/student.go ", err)
			log.Println("user close: ", userinfo.Name)
			delete(ClientList, hashAccessToken)
			return
		}
		switch messageType {
		case websocket.TextMessage:
			log.Printf("Received text message: %s", receive)
		case websocket.BinaryMessage:
			log.Printf("Received binary message: %x", receive)
		case websocket.PingMessage:
			log.Printf("Received ping message: %x", receive)
		case websocket.PongMessage:
			log.Printf("Received pong message: %x", receive)
		default:
			log.Printf("Received unknown message type: %d", messageType)
		}
	}
}
