package entity

type StudentEntity struct {
	Id      int64   `json:"id"`      // 学生id
	Score   float64 `json:"score"`   // 学生学分绩
	Uid     string  `json:"uid"`     // 学号
	College string  `json:"college"` // 学院id
}
