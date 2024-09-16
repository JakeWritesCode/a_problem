package store

import (
	"a_problem/lines/logging"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/wagslane/go-rabbitmq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// PostgresDBConfig is the configuration for a PostgresDB.
type PostgresDBConfig struct {
	TestRunner       bool
	Logger           logging.Logger
	ConnectionString string
	AppName          string
}

type PostgresStoreInterface interface {
	BeginTransaction() error
	RollbackTransaction() error
	Models() []PostgresModel
	RecordNotFound(err error) bool
	Close()
}

// PostgresStore is a struct that contains an initialized PostgresDB instance.
type PostgresStore struct {
	Postgres GormInstanceInterface
	Config   PostgresDBConfig
}

func (s *PostgresStore) BeginTransaction() error {
	s.Postgres = s.Postgres.Begin()
	return nil
}

func (s *PostgresStore) RollbackTransaction() error {
	s.Postgres = s.Postgres.Rollback()
	return nil
}

func (s *PostgresStore) Models() []PostgresModel {
	return []PostgresModel{}
}

func (s *PostgresStore) RecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (s *PostgresStore) Close() {
	s.Config.Logger.Info(s.Config.AppName, "PostgresStore.Close", fmt.Sprintf("Closing Postgres connection for app: %v", s.Config.AppName))
}

type ModelValidationError struct {
	Field   string
	Message string
}

type PostgresModel interface {
	Validate() []ModelValidationError
}

type IntegrationTestStore interface {
	BeginTransaction() error
	RollbackTransaction() error
}

type GormInstanceInterface interface {
	Create(value interface{}) (tx *gorm.DB)
	Begin(opts ...*sql.TxOptions) *gorm.DB
	Rollback() *gorm.DB
	AutoMigrate(dst ...interface{}) error
	Where(query interface{}, args ...interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Model(value interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Clauses(conds ...clause.Expression) *gorm.DB
	Preload(query string, args ...interface{}) (tx *gorm.DB)
}

type RabbitMQConfig struct {
	URL          string
	Logger       logging.Logger
	AppName      string
	ExchangeName string
	TestRunner   bool
}

type RabbitMQStore struct {
	Config    RabbitMQConfig
	conn      *rabbitmq.Conn
	publisher *rabbitmq.Publisher
}

func (s *RabbitMQStore) Connect() {
	conn, err := rabbitmq.NewConn(
		s.Config.URL,
		rabbitmq.WithConnectionOptionsLogging,
	)

	if err != nil {
		s.Config.Logger.Fatal(s.Config.AppName, "RabbitStore.Connect", fmt.Sprintf("Failed to connect to RabbitMQ: %v", err))
	}

	s.Config.Logger.Info(s.Config.AppName, "RabbitStore.Connect", fmt.Sprintf("Connected to RabbitMQ for app: %v", s.Config.AppName))

	// Create a publisher
	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName(s.Config.ExchangeName),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)

	if err != nil {
		s.Config.Logger.Fatal(s.Config.AppName, "RabbitStore.Connect", fmt.Sprintf("Failed to create publisher: %v", err))
	}
	s.Config.Logger.Info(s.Config.AppName, "RabbitStore.Connect", fmt.Sprintf("Creating publisher for app: %v on exchange: %v", s.Config.AppName, s.Config.ExchangeName))

	s.publisher = publisher
	s.conn = conn
}

func (s *RabbitMQStore) PublishMessage(message []byte, routingKeys []string) error {
	return s.publisher.Publish(
		message,
		routingKeys,
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange(s.Config.ExchangeName),
	)
}

func (s *RabbitMQStore) Close() {
	s.Config.Logger.Info(s.Config.AppName, "RabbitStore.Close", fmt.Sprintf("Closing RabbitMQ connection for app: %v", s.Config.AppName))
	s.conn.Close()
}

type GormUUIDModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RabbitMQStoreInterface interface {
	PublishMessage(message []byte, routingKeys []string) error
}
