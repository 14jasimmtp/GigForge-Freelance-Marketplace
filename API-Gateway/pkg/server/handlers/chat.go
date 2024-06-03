package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/chat"
	req "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/req_models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ChatHandler struct {
	chat chat.ChatServiceClient
}

func NewChatHandler(chat chat.ChatServiceClient) ChatHandler {
	return ChatHandler{chat: chat}

}

var (
	users = make(map[string]*websocket.Conn)
)

func (h *ChatHandler) Chat(c *websocket.Conn) {
	defer delete(users, c.Locals("User_id").(string))
	defer c.Close()

	users[c.Locals("User_id").(string)] = c

	for {
		fmt.Println("loop starts", c.Locals("User_id"), users)

		_, msg, err := c.ReadMessage()
		if err != nil {
			c.WriteJSON(fiber.Map{"Error": err.Error()})
		}

		h.SendMessageToUser(users, msg, c.Locals("User_id").(string))
	}
}

// func (h *ChatHandler) DeleteMessage(c *fiber.Ctx) error{

// }

func (h *ChatHandler) SendMessageToUser(User map[string]*websocket.Conn, msg []byte, userID string) {
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

		err := h.RabbitmqSender(message)
		if err != nil {
			senderConn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		}
		return
	}

	err := h.RabbitmqSender(message)
	if err != nil {
		senderConn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
	}

	err = recipientConn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		senderConn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		delete(User, message.RecipientID)
	}
}

func (h *ChatHandler) RabbitmqSender(msg req.Message) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"message", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msgbyte, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = ch.PublishWithContext(ctx,
		"",    
		q.Name,
		false, 
		false, 
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msgbyte,
		})
	if err != nil {
		return err
	}
	return nil
}

func (h *ChatHandler) GetMessages(c *fiber.Ctx) error {
	sender_id := c.Locals("User_id").(string)
	receiver_id := c.Params("receiver_id")
	res, err := h.chat.GetChats(context.Background(), &chat.GetChatReq{SenderId: sender_id, RecieverId: receiver_id})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
