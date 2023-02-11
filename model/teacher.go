package model

type Teacher struct {
	Id        int64   `gorm:"column:id;primary_key;AUTO_INCREMENT"` // 学生id
	Score     float64 `gorm:"column:score;default:0;NOT NULL"`      // 学生学分绩
	Uid       string  `gorm:"column:uid;NOT NULL"`                  // 教师编号
	Password  string  `gorm:"column:password;NOT NULL"`             // 账号密码，默认为教师编号
	CollegeId int64   `gorm:"column:college_id;NOT NULL"`           // 学院id
	Role      string  `gorm:"column:role;NOT NULL"`                 // 角色
}

func (m *Teacher) TableName() string {
	return "`teacher`"
}
