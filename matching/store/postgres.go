package store

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStoreInterface interface {
	Initialize() error
	Close() error
	CreateParticipant(participant *Participant) error
	UpdateParticipant(participant *Participant) error
	GetParticipantByID(id string) (*Participant, error)
	CreateStudy(study *Study) error
	UpdateStudy(study *Study) error
	GetStudyByID(id string) (*Study, error)
	GetAllStudies() ([]Study, error)
	CreateParticipantGroup(participantGroup *ParticipantGroup) error
	UpdateParticipantGroup(participantGroup *ParticipantGroup) error
	GetParticipantGroupByID(id string) (*ParticipantGroup, error)
	GetAllParticipantGroups() ([]ParticipantGroup, error)
	CreateMultiSelectQuestion(multiSelectQuestion *MultiSelectQuestion) error
	UpdateMultiSelectQuestion(multiSelectQuestion *MultiSelectQuestion) error
	GetMultiSelectQuestionByID(id string) (*MultiSelectQuestion, error)
	GetAllMultiSelectQuestions() ([]MultiSelectQuestion, error)
}

type PostgresStore struct {
	ConnectionURL string
	Conn          *gorm.DB
}

func (s *PostgresStore) Initialize() error {
	conn, err := gorm.Open(postgres.Open(s.ConnectionURL), &gorm.Config{})
	if err != nil {
		return err
	}
	s.Conn = conn
	err = s.Migrate()
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) Close() error {
	db, err := s.Conn.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (s *PostgresStore) Migrate() error {
	return s.Conn.AutoMigrate(
		Participant{},
		Study{},
		ParticipantGroup{},
		MultiSelectQuestion{},
	)
}

func NewPostgresStore() (*PostgresStore, error) {
	connectionUrl := "someURL"
	store := &PostgresStore{ConnectionURL: connectionUrl}
	err := store.Initialize()
	if err != nil {
		return nil, err
	}
	return store, nil
}