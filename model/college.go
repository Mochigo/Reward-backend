package model

import (
	"Reward/log"
	"errors"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type College struct {
	Id   int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"column:name;NOT NULL"` // 学院名称
}

func (m *College) TableName() string {
	return "`college`"
}

type CollegeDao struct{}

var collegeDaoOnce sync.Once
var collegeDao *CollegeDao

func GetCollegeDao() *CollegeDao {
	collegeDaoOnce.Do(func() {
		collegeDao = &CollegeDao{}
	})
	return collegeDao
}

func (*CollegeDao) GetCollegeById(db *gorm.DB, id int64) (*College, error) {
	college := &College{}
	err := db.Model(&College{}).Where("id = ?", id).Find(college).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("[GetCollegeById] record not found",
			zap.Int64("id", id))
		return nil, gorm.ErrRecordNotFound
	}

	if err != nil {
		log.Error("[GetCollegeById] failed to get",
			zap.Int64("id", id),
			zap.String("err", err.Error()))
		return nil, err
	}

	return college, nil
}
