package model

import (
	"errors"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"Reward/log"
)

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

type StudentDao struct{}

var studentDaoOnce sync.Once
var studentDao *StudentDao

func GetStudentDao() *StudentDao {
	studentDaoOnce.Do(func() {
		studentDao = &StudentDao{}
	})

	return studentDao
}

func (*StudentDao) GetStudentByUID(db *gorm.DB, uid string) (*Student, error) {
	s := &Student{}
	err := db.Model(&Student{}).Where("uid = ?", uid).Find(s).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("[GetStudentByUID] record not found",
			zap.String("uid", uid))
		return nil, gorm.ErrRecordNotFound
	}

	if err != nil {
		log.Error("[GetStudentByUID] failed to get",
			zap.String("uid", uid),
			zap.String("err", err.Error()))
		return nil, err
	}

	return s, nil
}
