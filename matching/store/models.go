package store

import (
	"a_problem/lines/store"
	"github.com/lib/pq"
	"time"
)

type Participant struct {
	store.GormUUIDModel
	StudiesStarted              pq.StringArray         `gorm:"type:text[]"`
	StudiesCompleted            pq.StringArray         `gorm:"type:text[]"`
	StudiesApproved             pq.StringArray         `gorm:"type:text[]"`
	StudiesRejected             pq.StringArray         `gorm:"type:text[]"`
	ParticipantGroupMemberships pq.StringArray         `gorm:"type:text[]"`
	LastActiveAt                time.Time              `gorm:"index"`
	Status                      string                 // Banned etc
	AttributesA                 ParticipantAttributesA `gorm:"foreignKey:ParticipantID"`
	AttributesB                 ParticipantAttributesB `gorm:"foreignKey:ParticipantID"`
	AttributesC                 ParticipantAttributesC `gorm:"foreignKey:ParticipantID"`
	AttributesD                 ParticipantAttributesD `gorm:"foreignKey:ParticipantID"`
	AttributesE                 ParticipantAttributesE `gorm:"foreignKey:ParticipantID"`
	AttributesF                 ParticipantAttributesF `gorm:"foreignKey:ParticipantID"`
	AttributesG                 ParticipantAttributesG `gorm:"foreignKey:ParticipantID"`
	AttributesH                 ParticipantAttributesH `gorm:"foreignKey:ParticipantID"`
	AttributesI                 ParticipantAttributesI `gorm:"foreignKey:ParticipantID"`
	AttributesJ                 ParticipantAttributesJ `gorm:"foreignKey:ParticipantID"`
	AttributesK                 ParticipantAttributesK `gorm:"foreignKey:ParticipantID"`
	AttributesL                 ParticipantAttributesL `gorm:"foreignKey:ParticipantID"`
	AttributesM                 ParticipantAttributesM `gorm:"foreignKey:ParticipantID"`
	AttributesN                 ParticipantAttributesN `gorm:"foreignKey:ParticipantID"`
	AttributesO                 ParticipantAttributesO `gorm:"foreignKey:ParticipantID"`
	AttributesP                 ParticipantAttributesP `gorm:"foreignKey:ParticipantID"`
	AttributesQ                 ParticipantAttributesQ `gorm:"foreignKey:ParticipantID"`
	AttributesR                 ParticipantAttributesR `gorm:"foreignKey:ParticipantID"`
	AttributesS                 ParticipantAttributesS `gorm:"foreignKey:ParticipantID"`
	AttributesT                 ParticipantAttributesT `gorm:"foreignKey:ParticipantID"`
	AttributesU                 ParticipantAttributesU `gorm:"foreignKey:ParticipantID"`
	AttributesV                 ParticipantAttributesV `gorm:"foreignKey:ParticipantID"`
	AttributesW                 ParticipantAttributesW `gorm:"foreignKey:ParticipantID"`
	AttributesX                 ParticipantAttributesX `gorm:"foreignKey:ParticipantID"`
	AttributesY                 ParticipantAttributesY `gorm:"foreignKey:ParticipantID"`
	AttributesZ                 ParticipantAttributesZ `gorm:"foreignKey:ParticipantID"`
}

type ParticipantAttributesA struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesB struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesC struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesD struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesE struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesF struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesG struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesH struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesI struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesJ struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesK struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesL struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesM struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesN struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesO struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesP struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesQ struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesR struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesS struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesT struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesU struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesV struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesW struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesX struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesY struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
}

type ParticipantAttributesZ struct {
	store.GormUUIDModel
	ParticipantID string `gorm:"index"`
	FilterID      string `gorm:"index"`
	Response      string `gorm:"index"`
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
