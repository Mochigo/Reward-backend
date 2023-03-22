package entity

type StudentEntity struct {
	Id      int64   `json:"id"`      // 学生id
	Score   float64 `json:"score"`   // 学生学分绩
	Uid     string  `json:"uid"`     // 学号
	College string  `json:"college"` // 学院名称
}

type CreateStudentEntity struct {
	Score     float64 `json:"score"`      // 学生学分绩
	Uid       string  `json:"uid"`        // 学号
	CollegeId int     `json:"college_id"` // 学院id
	Password  string  `json:"passrod"`    // 密码
}

type UploadScoreEntity struct {
	SuccessfulList []string `json:"successful_list"`
	FailedList     []string `json:"failed_list"`
}

type GetUserApplicationEntity struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
