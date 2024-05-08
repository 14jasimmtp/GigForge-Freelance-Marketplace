package handler

import (
	"encoding/json"
	"fmt"

	req "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/req_models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatHandler struct {
}

func NewChatHandler() ChatHandler {
	return ChatHandler{}

}

type Client struct {
	Conn   *websocket.Conn
	ChatId primitive.ObjectID
	UserId uint
}

var (
	// connections = make(map[*websocket.Conn]*Client)
	users= make(map[string]*websocket.Conn)
)

func (h *ChatHandler) Chat(c *websocket.Conn) {
	defer delete(users, c.Locals("User_id").(string))
	defer c.Close()

	users[c.Locals("User_id").(string)] = c

	for {
		fmt.Println("loop starts", c.Locals("User_id"), users)

		_, msg, err := c.ReadMessage()
		if err != nil {
			c.WriteJSON(fiber.Map{"Error":err.Error()})
		}

		SendMessageToUser(users, msg, c.Locals("User_id").(string))
	}
}

func  SendMessageToUser(User map[string]*websocket.Conn, msg []byte, userID string) {
	senderConn, ok := User[userID]

	var message req.Message
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		if ok {
			senderConn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		}
		return
	}
	fmt.Println("==", message)
	

	message.Status = "send"
	message.SenderID = userID

	recipientConn, ok := User[message.RecipientID]
	if !ok {
		message.Status = "pending"
		delete(User, message.RecipientID)

		// err := r.KafkaProducer(message)
		// if err != nil {
		// 	senderConn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		// }
		return
	}

	// err := r.KafkaProducer(message)
	// if err != nil {sugano
	// 	senderConn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
	// }

	err := recipientConn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		senderConn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		delete(User, message.RecipientID)
	}
}


