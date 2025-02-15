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
	//GetAllParticipants() ([]Participant, error)
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
	GetNumberOfParticipantsActiveInLast90Days() (int, error)
	Mark1000ParticipantsActive() error
	GetStore() *gorm.DB
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
		ParticipantAttributesA{},
		ParticipantAttributesB{},
		ParticipantAttributesC{},
		ParticipantAttributesD{},
		ParticipantAttributesE{},
		ParticipantAttributesF{},
		ParticipantAttributesG{},
		ParticipantAttributesH{},
		ParticipantAttributesI{},
		ParticipantAttributesJ{},
		ParticipantAttributesK{},
		ParticipantAttributesL{},
		ParticipantAttributesM{},
		ParticipantAttributesN{},
		ParticipantAttributesO{},
		ParticipantAttributesP{},
		ParticipantAttributesQ{},
		ParticipantAttributesR{},
		ParticipantAttributesS{},
		ParticipantAttributesT{},
		ParticipantAttributesU{},
		ParticipantAttributesV{},
		ParticipantAttributesW{},
		ParticipantAttributesX{},
		ParticipantAttributesY{},
		ParticipantAttributesZ{},
	)
}

func NewPostgresStore() (*PostgresStore, error) {
	connectionUrl := "postgres://postgres:JakesP0Str3sP@ssw0rd@192.168.1.200:5432/matching_test_three"
	store := &PostgresStore{ConnectionURL: connectionUrl}
	err := store.Initialize()
	if err != nil {
		return nil, err
	}
	return store, nil
}
