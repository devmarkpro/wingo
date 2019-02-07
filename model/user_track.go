package model

import (
	"database/sql"
)

type UserTrack struct {
	Base
	ContestID             uint
	QuestionID            uint
	QuestionIndex         uint
	CanPlay               bool
	CanUseCorrector       bool
	IsSelectCorrectAnswer sql.NullBool
	State                 TrackStateEnum
	BeforeTickets         uint
	BeforeCorrectors      uint
	AfterTickets          uint
	AfterCorrectors       uint
	MetaData              sql.NullString
}