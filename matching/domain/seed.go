package domain

import (
	"a_problem/lines/utils"
	"a_problem/matching/store"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/cespare/xxhash/v2"
	"log"
	"math/rand"
	"strings"
	"time"
)

func (d *Domain) CreateStudy() error {
	study := &store.Study{}
	return d.store.CreateStudy(study)
}

func (d *Domain) CreateParticipantGroup() error {
	participantGroup := &store.ParticipantGroup{}
	return d.store.CreateParticipantGroup(participantGroup)
}

func randBetween(min, max int) int {
	return rand.Intn(max-min) + min
}

func (d *Domain) CreateMultiSelectQuestion() error {
	choices := []string{}
	for i := 0; i < randBetween(10, 100); i++ {
		choices = append(choices, faker.Word())
	}
	multiSelectQuestion := &store.MultiSelectQuestion{
		FilterId:          strings.ToLower(fmt.Sprintf("%v-%v-%v", faker.Word(), faker.Word(), faker.Word())),
		PossibleResponses: choices,
	}
	return d.store.CreateMultiSelectQuestion(multiSelectQuestion)
}

// SeedCoreData seeds the database with core data, ready to create participants.
func (d *Domain) SeedCoreData() {
	// Create 100000 studies
	log.Println("Creating studies")
	studyTasks := make(chan func(), 100000)
	for i := 0; i < 100000; i++ {
		studyTasks <- func() {
			d.CreateStudy()
		}
	}
	close(studyTasks)
	utils.RunInWorkerPool(studyTasks, 50)

	// Create 100000 participant groups
	log.Println("Creating participant groups")
	pgTasks := make(chan func(), 100000)
	for i := 0; i < 100000; i++ {
		pgTasks <- func() {
			d.CreateParticipantGroup()
		}
	}
	close(pgTasks)
	utils.RunInWorkerPool(pgTasks, 50)

	// Create 500 multi select questions
	log.Println("Creating multi select questions")
	questionTasks := make(chan func(), 500)
	for i := 0; i < 500; i++ {
		questionTasks <- func() {
			d.CreateMultiSelectQuestion()
		}
	}
	close(questionTasks)
	utils.RunInWorkerPool(questionTasks, 50)
}

func RandomSelections[T any](slice []T, numChoices int) []T {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Shuffle the slice
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	// Adjust numChoices if it is greater than the length of the slice
	if numChoices > len(slice) {
		numChoices = len(slice)
	}

	// Select the first `numChoices` elements
	return slice[:numChoices]
}

func StudyInStudySlice(study *store.Study, studies []*store.Study) bool {
	for _, s := range studies {
		if s.ID == study.ID {
			return true
		}
	}
	return false
}

func GenerateQuestionResponse(question *store.MultiSelectQuestion) (filterId string, selectedValues []string) {
	filterId = question.FilterId
	selectedValues = RandomSelections(question.PossibleResponses, randBetween(1, 3))
	return
}

func GenerateQuestionsAnswered(questions []*store.MultiSelectQuestion) map[string][]string {
	questionsAnswered := make(map[string][]string, len(questions))
	for _, q := range questions {
		filterId, selectedValues := GenerateQuestionResponse(q)
		questionsAnswered[filterId] = selectedValues
	}
	return questionsAnswered
}

func generate32BitHash(input string) string {
	data := []byte(input)
	hash := xxhash.Sum64(data)
	// Take the lower 32 bits of the hash
	hash32bit := uint32(hash & 0xFFFFFFFF)
	return fmt.Sprintf("%08x", hash32bit)
}

func (d *Domain) GenerateAndSaveParticipant(
	studies []*store.Study,
	pgs []*store.ParticipantGroup,
	questions []*store.MultiSelectQuestion,
	returnChan chan store.Participant,
) error {

	// Let's assume an average participant has taken 100ish studies with a 97% completion rate
	// They're in 10 participant groups
	// And they've answered at least 100 multi select questions

	studiesStarted := RandomSelections(studies, randBetween(90, 110))
	studyIdsStarted := make([]string, len(studiesStarted))
	for i, study := range studiesStarted {
		studyIdsStarted[i] = study.ID.String()
	}
	studiesCompleted := RandomSelections(studiesStarted, randBetween(len(studiesStarted)-10, len(studiesStarted)-5))
	studyIdsCompleted := make([]string, len(studiesCompleted))
	for i, study := range studiesCompleted {
		studyIdsCompleted[i] = study.ID.String()
	}
	var studiesRejected []*store.Study
	for _, study := range studiesStarted {
		if !StudyInStudySlice(study, studiesCompleted) {
			studiesRejected = append(studiesRejected, study)
		}
	}
	studyIdsRejected := make([]string, len(studiesRejected))
	for i, study := range studiesRejected {
		studyIdsRejected[i] = study.ID.String()
	}
	pgMemberships := RandomSelections(pgs, randBetween(8, 12))
	pgIds := make([]string, len(pgMemberships))
	for i, pg := range pgMemberships {
		pgIds[i] = pg.ID.String()
	}
	questionsSelected := RandomSelections(questions, randBetween(100, 200))

	participant := store.Participant{
		StudiesStarted:              studyIdsStarted,
		StudiesCompleted:            studyIdsCompleted,
		StudiesRejected:             studyIdsRejected,
		ParticipantGroupMemberships: pgIds,
		Status:                      "Active",
	}

	err := d.store.CreateParticipant(&participant)
	if err != nil {
		return err
	}
	questionData := make(map[string][]string, len(questionsSelected))
	for _, q := range questionsSelected {
		filterId, selectedValues := GenerateQuestionResponse(q)
		questionData[filterId] = selectedValues
	}
	for filterId, selectedValues := range questionData {
		switch filterId[0] {
		case 'a':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesA{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'b':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesB{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'c':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesC{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'd':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesD{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'e':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesE{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'f':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesF{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'g':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesG{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'h':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesH{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'i':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesI{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'j':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesJ{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'k':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesK{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'l':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesL{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'm':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesM{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'n':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesN{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'o':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesO{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'p':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesP{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'q':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesQ{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'r':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesR{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 's':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesS{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 't':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesT{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'u':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesU{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'v':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesV{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'w':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesW{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'x':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesX{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'y':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesY{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		case 'z':
			for _, value := range selectedValues {
				attr := &store.ParticipantAttributesZ{
					ParticipantID: participant.ID.String(),
					FilterID:      filterId,
					Response:      value,
				}
				err := d.store.GetStore().Create(attr).Error
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (d *Domain) CreateParticipant(participant *store.Participant) error {
	return d.store.CreateParticipant(participant)
}

func sliceToPointerSlice[T any](slice []T) []*T {
	pointerSlice := make([]*T, len(slice))
	for i, item := range slice {
		pointerSlice[i] = &item
	}
	return pointerSlice
}

func (d *Domain) SeedParticipants() {
	studies, err := d.store.GetAllStudies()
	if err != nil {
		panic(err)
	}
	studiesPointers := sliceToPointerSlice(studies)

	pgs, err := d.store.GetAllParticipantGroups()
	if err != nil {
		panic(err)
	}
	pgPointers := sliceToPointerSlice(pgs)

	questions, err := d.store.GetAllMultiSelectQuestions()
	if err != nil {
		panic(err)
	}
	questionPointers := sliceToPointerSlice(questions)

	// Create 1.8M participants
	log.Println("Generating participants")
	for i := 0; i < 1000; i++ {
		participantTasks := make(chan func(), 1800)
		for j := 0; j < 1800; j++ {
			participantTasks <- func() {
				d.GenerateAndSaveParticipant(studiesPointers, pgPointers, questionPointers, nil)
			}
		}
		close(participantTasks)
		// Limit workers here to avoid overloading the database connection pool
		utils.RunInWorkerPool(participantTasks, 50)
	}
}

func (d *Domain) Seed() {
	d.SeedCoreData()
	d.SeedParticipants()
}

func (d *Domain) CreateLastActive() error {
	qtyLastActive90Days, err := d.store.GetNumberOfParticipantsActiveInLast90Days()
	if err != nil {
		panic(err)
	}
	if qtyLastActive90Days < 250000 {
		log.Println(fmt.Sprintf("There are %v participants active in the last 90 days, marking 1000 as active", qtyLastActive90Days))
		err = d.store.Mark1000ParticipantsActive()
		if err != nil {
			panic(err)
		}
	}
	log.Println("Last active participants updated")
	return nil
}
