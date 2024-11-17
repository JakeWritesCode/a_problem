package store

import (
	"a_problem/lines/store"
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"time"
)

type Participant struct {
	store.GormUUIDModel
	StudiesStarted              pq.StringArray `gorm:"type:text[]"`
	StudiesCompleted            pq.StringArray `gorm:"type:text[]"`
	StudiesApproved             pq.StringArray `gorm:"type:text[]"`
	StudiesRejected             pq.StringArray `gorm:"type:text[]"`
	ParticipantGroupMemberships pq.StringArray `gorm:"type:text[]"`
	LastActiveAt                time.Time      `gorm:"index"`
	Status                      string         // Banned etc
	DataA                       datatypes.JSON `gorm:"index,type:gin"`
	DataB                       datatypes.JSON `gorm:"index,type:gin"`
	DataC                       datatypes.JSON `gorm:"index,type:gin"`
	DataD                       datatypes.JSON `gorm:"index,type:gin"`
	DataE                       datatypes.JSON `gorm:"index,type:gin"`
	DataF                       datatypes.JSON `gorm:"index,type:gin"`
	DataG                       datatypes.JSON `gorm:"index,type:gin"`
	DataH                       datatypes.JSON `gorm:"index,type:gin"`
	DataI                       datatypes.JSON `gorm:"index,type:gin"`
	DataJ                       datatypes.JSON `gorm:"index,type:gin"`
	DataK                       datatypes.JSON `gorm:"index,type:gin"`
	DataL                       datatypes.JSON `gorm:"index,type:gin"`
	DataM                       datatypes.JSON `gorm:"index,type:gin"`
	DataN                       datatypes.JSON `gorm:"index,type:gin"`
	DataO                       datatypes.JSON `gorm:"index,type:gin"`
	DataP                       datatypes.JSON `gorm:"index,type:gin"`
	DataQ                       datatypes.JSON `gorm:"index,type:gin"`
	DataR                       datatypes.JSON `gorm:"index,type:gin"`
	DataS                       datatypes.JSON `gorm:"index,type:gin"`
	DataT                       datatypes.JSON `gorm:"index,type:gin"`
	DataU                       datatypes.JSON `gorm:"index,type:gin"`
	DataV                       datatypes.JSON `gorm:"index,type:gin"`
	DataW                       datatypes.JSON `gorm:"index,type:gin"`
	DataX                       datatypes.JSON `gorm:"index,type:gin"`
	DataY                       datatypes.JSON `gorm:"index,type:gin"`
	DataZ                       datatypes.JSON `gorm:"index,type:gin"`
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
