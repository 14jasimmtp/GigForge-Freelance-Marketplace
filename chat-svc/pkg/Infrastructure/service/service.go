package service

import (
	"context"
	"fmt"
	"log"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pb"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pkg/Infrastructure/repository"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Service struct{
	repo *repository.Repo
	pb.UnimplementedChatServiceServer
}

func NewChatService(repo *repository.Repo) *Service{
	return &Service{repo: repo}
}

func (s *Service) ChatReciever() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("error", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("error", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"message", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		fmt.Println("error", err)
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		fmt.Println("error", err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			err:=s.repo.SaveMessage(d.Body)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func (s *Service) GetChats(ctx context.Context,req *pb.GetChatReq) (*pb.GetChatRes,error){
	chats,err:=s.repo.GetChats(req)
	if err != nil {
		return nil,nil
	}
	
	return &pb.GetChatRes{Chat: chats},nil
}
