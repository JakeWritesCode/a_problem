package store

func (s *PostgresStore) CreateParticipant(participant *Participant) error {
	return s.Conn.Create(participant).Error
}

func (s *PostgresStore) UpdateParticipant(participant *Participant) error {
	return s.Conn.Save(participant).Error
}

func (s *PostgresStore) GetParticipantByID(id string) (*Participant, error) {
	var participant Participant
	err := s.Conn.First(&participant, "id = ?", id).Error
	return &participant, err
}

func (s *PostgresStore) CreateStudy(study *Study) error {
	return s.Conn.Create(study).Error
}

func (s *PostgresStore) UpdateStudy(study *Study) error {
	return s.Conn.Save(study).Error
}

func (s *PostgresStore) GetStudyByID(id string) (*Study, error) {
	var study Study
	err := s.Conn.First(&study, "id = ?", id).Error
	return &study, err
}

func (s *PostgresStore) GetAllStudies() ([]Study, error) {
	var studies []Study
	err := s.Conn.Find(&studies).Error
	return studies, err
}

func (s *PostgresStore) CreateParticipantGroup(participantGroup *ParticipantGroup) error {
	return s.Conn.Create(participantGroup).Error
}

func (s *PostgresStore) UpdateParticipantGroup(participantGroup *ParticipantGroup) error {
	return s.Conn.Save(participantGroup).Error
}

func (s *PostgresStore) GetParticipantGroupByID(id string) (*ParticipantGroup, error) {
	var participantGroup ParticipantGroup
	err := s.Conn.First(&participantGroup, "id = ?", id).Error
	return &participantGroup, err
}

func (s *PostgresStore) GetAllParticipantGroups() ([]ParticipantGroup, error) {
	var participantGroups []ParticipantGroup
	err := s.Conn.Find(&participantGroups).Error
	return participantGroups, err
}

func (s *PostgresStore) CreateMultiSelectQuestion(multiSelectQuestion *MultiSelectQuestion) error {
	return s.Conn.Create(multiSelectQuestion).Error
}

func (s *PostgresStore) UpdateMultiSelectQuestion(multiSelectQuestion *MultiSelectQuestion) error {
	return s.Conn.Save(multiSelectQuestion).Error
}

func (s *PostgresStore) GetMultiSelectQuestionByID(id string) (*MultiSelectQuestion, error) {
	var multiSelectQuestion MultiSelectQuestion
	err := s.Conn.First(&multiSelectQuestion, "id = ?", id).Error
	return &multiSelectQuestion, err
}

func (s *PostgresStore) GetAllMultiSelectQuestions() ([]MultiSelectQuestion, error) {
	var multiSelectQuestions []MultiSelectQuestion
	err := s.Conn.Find(&multiSelectQuestions).Error
	return multiSelectQuestions, err
}
