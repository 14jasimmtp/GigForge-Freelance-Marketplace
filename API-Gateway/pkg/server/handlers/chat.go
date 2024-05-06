package handler

import (
	"fmt"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/chat"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatHandler struct {
	chat chat.ChatServiceClient
}

func NewChatHandler(chat chat.ChatServiceClient) ChatHandler {
	return ChatHandler{chat: chat}

}

type Client struct {
	Conn   *websocket.Conn
	ChatId primitive.ObjectID
	UserId uint
}

var (
	connections = make(map[*websocket.Conn]*Client)
	users       = make(map[uint]*websocket.Conn)
)

func (h *ChatHandler) Chat(c *websocket.Conn) {
	userId, ok := 1, true
	if !ok {
		errRes := MakeResponse(fiber.StatusUnauthorized, "unauthorised", nil, "error in retrieving user id")
		c.WriteJSON(errRes)
		c.Close()
		return
	}

	chatId, err := primitive.ObjectIDFromHex(c.Params("chatid"))
	if err != nil {
		errRes := MakeResponse(fiber.StatusBadRequest, "string conversion failed", nil, err.Error())
		c.WriteJSON(errRes)
		c.Close()
		return
	}

	client := &Client{Conn: c, ChatId: chatId, UserId: uint(userId)}
	connections[c] = client
	users[uint(userId)] = c

	go handleConnection(client)
}

func handleConnection(client *Client) {
	defer func() {
		client.Conn.Close()
		delete(connections, client.Conn)
		delete(users, client.UserId)
	}()

	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			break
		}

		// Save message to the database
		_, err = SaveMessage(client.ChatId, client.UserId, string(msg))
		if err != nil {
			fmt.Println("error in saving message")
			break
		}

		// Broadcast the message to the recipient
		client.Conn.WriteMessage(websocket.TextMessage, msg)

		recipient, err := FetchRecipient(client.ChatId, client.UserId)
		if err != nil {
			fmt.Println("error in fetching recipient id")
			break
		}

		if recipientConn, ok := users[recipient]; ok {
			err = recipientConn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				delete(connections, recipientConn)
				delete(users, recipient)
			}
		}
	}
}

func MakeResponse(status int, message string, data interface{}, error string) fiber.Map {
	return fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
		"error":   error,
	}
}

// SaveMessage is a placeholder function to save the message to the database
func SaveMessage(chatID primitive.ObjectID, userID uint, message string) (interface{}, error) {
	// Implement your logic to save the message to the database
	fmt.Printf("Saving message: %s to chat %s from user %d\n", message, chatID.Hex(), userID)
	return nil, nil
}

// FetchRecipient is a placeholder function to fetch the recipient's ID
func FetchRecipient(chatID primitive.ObjectID, userID uint) (uint, error) {
	// Implement your logic to fetch the recipient's ID from the database
	fmt.Printf("Fetching recipient for chat %s and user %d\n", chatID.Hex(), userID)
	recipientID := uint(0) // Replace with the actual recipient ID
	return recipientID, nil
}
