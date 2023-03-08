package common

import "errors"

const StringEmpty = ""

const (
	StatusPROCESS  = "PROCESS"
	StatusAPPROVE  = "APPROVE"
	StatusREJECTED = "REJECTED"
)

const (
	LevelSchool   = "01"
	LevelProvince = "02"
	LevelCountry  = "03"
)

const (
	OperationApprove = "APPROVE"
	OperationReject  = "REJECT"
)

const (
	CondiPage      = "page"
	CondiLimit     = "limit"
	CondiCollegeId = "college_id"
	CondiStudentId = "student_id"
)

const (
	TokenUserID    = "userID"
	TokenCollegeID = "collegeID"
)

var (
	ErrTimeParse = errors.New("fail to parse time")
)
