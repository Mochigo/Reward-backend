package model

import (
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

func (*ScholarshipDao) Create(db *gorm.DB, s *Scholarship) error {
	err := db.Model(&Scholarship{}).Create(s).Error
	if err != nil {
		return err
	}
	return nil
}

func (*ScholarshipDao) DeleteByID(db *gorm.DB, id int64) error {
	err := db.Delete(&Scholarship{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (*ScholarshipDao) GetList(db *gorm.DB, condi map[string]interface{}) ([]*Scholarship, error) {
	var scholarships []*Scholarship
	limit, _ := condi["limit"].(int)
	page, _ := condi["page"].(int)

	if err := db.Model(&Scholarship{}).Where("college_id = ?", condi["college_id"]).Order("end_time desc").Offset((page - 1) * limit).Limit(limit).Find(&scholarships).Error; err != nil {
		return nil, err
	}

	return scholarships, nil
}
