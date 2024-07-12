package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/notification-svc/pb"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/notification-svc/pkg/repository"
	"github.com/rabbitmq/amqp091-go"
)

type Service struct {
	pb.NotificationServiceServer
	Repo *repository.Repo
	AmqpConn *amqp091.Connection
}

func NewNotificationService(repo *repository.Repo, amqp *amqp091.Connection) *Service {
	return &Service{Repo: repo}
}

func (s *Service) ChatReciever() {
	ch, err := s.AmqpConn.Channel()
	if err != nil {
		fmt.Println("error", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"notification", // name
		false,
		false,
		false,
		false,
		nil,
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
			err := s.Repo.CreateNewNotification(d.Body)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// func (s *Service) CreateNotification(ctx context.Context, req *pb.GNReq) (*pb.GNRes, error) {
// 	err := s.Repo.CreateNewNotification(req)
// 	if err != nil {
// 		return &pb.GNRes{Error: err.Error()}, nil
// 	}
// 	return &pb.GNRes{N}
// }

func (s *Service) GetNotification(ctx context.Context, req *pb.GNReq) (*pb.GNRes, error) {
	notifications, err := s.Repo.GetNotificationsForUser(req.UserId)
	if err != nil {
		return &pb.GNRes{Error: err.Error(),Status: http.StatusBadRequest}, nil
	}
	return &pb.GNRes{Notification: notifications,Status: http.StatusOK},nil
}
