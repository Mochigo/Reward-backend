package model

import (
	"errors"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"Reward/log"
)

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

type TeacherDao struct{}

var teacherDaoOnce sync.Once
var teacherDao *TeacherDao

func GetTeacherDao() *TeacherDao {
	teacherDaoOnce.Do(func() {
		teacherDao = &TeacherDao{}
	})

	return teacherDao
}

func (*TeacherDao) GetTeacherByUID(db *gorm.DB, uid string) (*Teacher, error) {
	t := &Teacher{}
	err := db.Model(&Teacher{}).Where("uid = ?", uid).Find(t).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("[GetTeacherByUID] record not found",
			zap.String("uid", uid))
		return nil, gorm.ErrRecordNotFound
	}

	if err != nil {
		log.Error("[GetTeacherByUID] failed to get",
			zap.String("uid", uid),
			zap.String("err", err.Error()))
		return nil, err
	}

	return t, nil
}
