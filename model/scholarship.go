package model

import (
	"Reward/common"
	"sync"
	"time"

	"gorm.io/gorm"
)

type Scholarship struct {
	Id        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"` // 奖学金id
	Name      string    `gorm:"column:name;NOT NULL" json:"name"`               // 奖学金名称
	CollegeId int64     `gorm:"column:college_id;NOT NULL" json:"college_id"`   // 学院id
	StartTime time.Time `gorm:"column:start_time;NOT NULL" json:"start_time"`   // 开始时间
	EndTime   time.Time `gorm:"column:end_time;NOT NULL" json:"end_time"`       // 结束时间
}

func (Scholarship) TableName() string {
	return "`scholarship`"
}

type ScholarshipDao struct{}

var scholarshipDaoOnce sync.Once
var scholarshipDao *ScholarshipDao

func GetScholarshipDao() *ScholarshipDao {
	scholarshipDaoOnce.Do(func() {
		scholarshipDao = &ScholarshipDao{}
	})
	return scholarshipDao
}

func (*ScholarshipDao) Create(db *gorm.DB, s *Scholarship) (int64, error) {
	err := db.Model(&Scholarship{}).Create(s).Error
	if err != nil {
		return 0, err
	}
	return s.Id, nil
}

func (*ScholarshipDao) DeleteByID(db *gorm.DB, id int64) error {
	err := db.Delete(&Scholarship{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (*ScholarshipDao) GetList(db *gorm.DB, condi map[string]interface{}) ([]*Scholarship, error) {
	scholarships := make([]*Scholarship, 0)
	limit := condi[common.CondiLimit].(int)
	page := condi[common.CondiPage].(int)
	collegeId := condi[common.CondiCollegeId].(int)
	if err := db.Where("college_id = ?", collegeId).Order("end_time desc").Offset((page - 1) * limit).Limit(limit).Find(&scholarships).Error; err != nil {
		return nil, err
	}

	return scholarships, nil
}

func (*ScholarshipDao) GetCountByCollegeId(db *gorm.DB, collegeId int) (int64, error) {
	var total int64
	if err := db.Model(&Scholarship{}).Where("college_id = ?", collegeId).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
