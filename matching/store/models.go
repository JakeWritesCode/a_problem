package store

import (
	"a_problem/lines/store"
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type Participant struct {
	store.GormUUIDModel
	StudiesStarted              pq.StringArray `gorm:"type:text[]"`
	StudiesCompleted            pq.StringArray `gorm:"type:text[]"`
	StudiesApproved             pq.StringArray `gorm:"type:text[]"`
	StudiesRejected             pq.StringArray `gorm:"type:text[]"`
	ParticipantGroupMemberships pq.StringArray `gorm:"type:text[]"`
	Status                      string         // Banned etc
	FilterResponses             datatypes.JSON
}

// These models wouldn't be in the matching service, but I'm putting them here so
// we can use them for queries with sensible values

// Study represents a study, we only care about the ID really
type Study struct {
	store.GormUUIDModel
}

// ParticipantGroup represents a ParticipantGroup, we only care about the ID really
type ParticipantGroup struct {
	store.GormUUIDModel
}

type MultiSelectQuestion struct {
	store.GormUUIDModel
	FilterId          string         `gorm:"unique"`
	PossibleResponses pq.StringArray `gorm:"type:text[]"`
}
