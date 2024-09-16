package event

import (
	"JakesBettingAlgorithmV2/lines/logging"
	"errors"
	"fmt"
	"github.com/wagslane/go-rabbitmq"
)

type EventReceiver func(d rabbitmq.Delivery) rabbitmq.Action

type EventConsumer interface {
	Run(receiver EventReceiver) error
}

type EventHandler interface {
	NewConsumer() *EventConsumer
	Close()
}

type RabbitMQEventHandlerConfig struct {
	URL          string
	Logger       logging.Logger
	ExchangeName string
	AppName      string
	Concurrency  int
}

type RabbitEventHandlerInterface interface {
	Connect() error
	Close()
	NewConsumer(config ConsumerConfig) error
}

type RabbitEventHandler struct {
	Config RabbitMQEventHandlerConfig
	conn   *rabbitmq.Conn
}

func (r *RabbitEventHandler) Close() {
	r.Config.Logger.Info("event", "Close", fmt.Sprintf("Closing RabbitMQ consumer connection for app %v", r.Config.AppName))
	r.conn.Close()
}

func (r *RabbitEventHandler) Connect() error {
	conn, err := rabbitmq.NewConn(
		r.Config.URL,
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		r.Config.Logger.Error("event", "Connect", err.Error())
		return errors.New("failed to create connection")
	}
	r.conn = conn
	return nil
}

type RabbitMQConsumer interface {
	Run(handler rabbitmq.Handler) error
}

type ConsumerConfig struct {
	RoutingKey   string
	QueueName    string
	ConsumerFunc rabbitmq.Handler
}

func (r *RabbitEventHandler) NewConsumer(config ConsumerConfig) error {
	consumer, err := rabbitmq.NewConsumer(
		r.conn,
		config.QueueName,
		rabbitmq.WithConsumerOptionsRoutingKey(config.RoutingKey),
		rabbitmq.WithConsumerOptionsExchangeName(r.Config.ExchangeName),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsConcurrency(r.Config.Concurrency),
	)
	if err != nil {
		r.Config.Logger.Error("event", "NewConsumer", err.Error())
		return errors.New("failed to create consumer")
	}
	r.Config.Logger.Info("event", "NewConsumer", fmt.Sprintf("Created consumer for routing key %v, queue name: %v, exchange_name: %v", config.RoutingKey, config.QueueName, r.Config.ExchangeName))

	go consumer.Run(config.ConsumerFunc)
	return nil
}
