package model

import (
	"errors"
	"fmt"
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
	err := db.Model(&Student{}).Where("uid = ?", uid).First(s).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Warn("[StudentDao][GetStudentByUID] record not found",
			zap.String("uid", uid))
		return nil, gorm.ErrRecordNotFound
	}

	if err != nil {
		log.Error("[StudentDao][GetStudentByUID] failed to get",
			zap.String("uid", uid),
			zap.String("err", err.Error()))
		return nil, err
	}

	return s, nil
}

func (*StudentDao) GetStudentById(db *gorm.DB, id int64) (*Student, error) {
	s := &Student{}
	err := db.Model(&Student{}).Where("id = ?", id).First(s).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("[StudentDao][GetStudentByUID] record not found",
			zap.Int64("id", id))
		return nil, gorm.ErrRecordNotFound
	}

	if err != nil {
		log.Error("[StudentDao][GetStudentByUID] failed to get",
			zap.Int64("id", id),
			zap.String("err", err.Error()))
		return nil, err
	}

	return s, nil
}

func (*StudentDao) Save(db *gorm.DB, s *Student) error {
	if err := db.Save(s).Error; err != nil {
		log.Error("[StudentDao][Save] fail to save",
			zap.String("student", fmt.Sprintf("%+v", s)))
		return err
	}
	return nil
}

func (*StudentDao) Create(db *gorm.DB, s *Student) error {
	if err := db.Model(&Student{}).Create(s).Error; err != nil {
		log.Error("[StudentDao][Create] fail to create",
			zap.String("student", fmt.Sprintf("%+v", s)))
		return err
	}
	return nil
}
