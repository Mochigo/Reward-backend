package model

type Student struct {
	Id        int64   `gorm:"column:id;primary_key;AUTO_INCREMENT"` // 学生id
	Score     float64 `gorm:"column:score;default:0;NOT NULL"`      // 学生学分绩
	Uid       string  `gorm:"column:uid;NOT NULL"`                  // 学号
	Password  string  `gorm:"column:password;NOT NULL"`             // 账号密码，默认为学生学号
	CollegeId int64   `gorm:"column:college_id;NOT NULL"`           // 学院id
}

func (m *Student) TableName() string {
	return "`student`"
}
