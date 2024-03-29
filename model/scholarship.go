package model

import (
	"Reward/common"
	"Reward/log"
	"errors"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
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

func (*ScholarshipDao) GetScholarshipById(db *gorm.DB, id int64) (*Scholarship, error) {
	s := &Scholarship{}
	err := db.Model(&Scholarship{}).Where("id = ?", id).First(s).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("[ScholarshipDao][GetScholarshipById] record not found",
			zap.Int64("id", id))
		return nil, gorm.ErrRecordNotFound
	}

	if err != nil {
		log.Error("[ScholarshipDao][GetScholarshipById] failed to get",
			zap.Int64("id", id),
			zap.String("err", err.Error()))
		return nil, err
	}

	return s, nil
}

func (*ScholarshipDao) BatchGetByIds(db *gorm.DB, Ids []int64) ([]*Scholarship, error) {
	scholarships := make([]*Scholarship, 0)
	if err := db.Model(&Scholarship{}).Where("id IN ?", Ids).Find(&scholarships).Error; err != nil {
		log.Error("[ScholarshipDao][BatchGetByIds] fail to get",
			zap.String("ids", fmt.Sprintf("%+v", Ids)),
			zap.String("err", err.Error()))
		return nil, err
	}
	return scholarships, nil
}
