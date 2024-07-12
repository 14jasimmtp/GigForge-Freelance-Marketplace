package broker

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func ConnectAMQP() *amqp.Connection{
	fmt.Println(viper.GetString("AmqpUrl"),"url")
	conn, err := amqp.Dial(viper.GetString("AmqpUrl"))
	if err != nil {
		fmt.Println("error", err)
	}
	return conn
}