package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
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
	users = make(map[int]*websocket.Conn)
)

// Chat handles the WebSocket connection for real-time chat messaging.
// @Summary WebSocket Chat
// @Description Establish a WebSocket connection for real-time chat messaging. This endpoint allows users to send and receive messages in real time.
// @security Authorization
// @Tags Chat
// @Param User_id path int true "User ID"
// @Produce json
// @Router /chat [get]
func (h *ChatHandler) Chat(c *websocket.Conn) {
	defer delete(users, c.Locals("User_id").(int))
	defer c.Close()

	users[c.Locals("User_id").(int)] = c

	for {
		fmt.Println("loop starts", c.Locals("User_id"), users)

		_, msg, err := c.ReadMessage()
		if err != nil {
			c.WriteJSON(fiber.Map{"Error": err.Error()})
		}

		h.SendMessageToUser(users, msg, c.Locals("User_id").(int))
	}
}

// func (h *ChatHandler) DeleteMessage(c *fiber.Ctx) error{

// }

func (h *ChatHandler) SendMessageToUser(User map[int]*websocket.Conn, msg []byte, userID int) {
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

	recipientConn, ok := (User[message.RecipientID])
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

// GetMessages retrieves chat messages between the sender and receiver.
// @Summary Get Chat Messages
// @Description Retrieve chat messages between the logged-in user and the specified receiver.
// @security Authorization
// @Tags Chat
// @Produce json
// @Param User_id path string true "Sender User ID"
// @Param receiver_id path string true "Receiver User ID"
// @Success 200 {object} res.CommonRes
// @Failure 400 {object} res.CommonRes
// @Router /chat/messages/{receiver_id} [get]
func (h *ChatHandler) GetMessages(c *fiber.Ctx) error {
	sender_id := c.Locals("User_id").(int)
	receiver_id := c.Params("receiver_id")
	fmt.Println(receiver_id,sender_id)
	res, err := h.chat.GetChats(context.Background(), &chat.GetChatReq{SenderId: strconv.Itoa(sender_id), RecieverId: receiver_id})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
