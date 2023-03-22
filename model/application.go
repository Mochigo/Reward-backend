package model

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"Reward/common"
	"Reward/log"
)

type Application struct {
	Id                int64     `gorm:"column:id"`                       // 申请id
	ScholarshipItemId int64     `gorm:"column:scholarship_item_id"`      // 奖学金子项id
	ScholarshipId     int64     `gorm:"column:scholarship_id"`           // 奖学金id
	StudentId         int64     `gorm:"column:student_id"`               // 申请学生id
	Status            string    `gorm:"column:status"`                   // 申请状态，APPROVE-通过|PROCESS-处理中|FAILURE-驳回
	Deadline          time.Time `gorm:"column:deadline" json:"deadline"` // 截止时间
}

func (m *Application) TableName() string {
	return "`application`"
}

type ApplicationDao struct{}

var applicationDaoOnce sync.Once
var applicationDao *ApplicationDao

func GetApplicationDao() *ApplicationDao {
	applicationDaoOnce.Do(func() {
		applicationDao = &ApplicationDao{}
	})
	return applicationDao
}

func (*ApplicationDao) Create(db *gorm.DB, a *Application) error {
	if err := db.Model(&Application{}).Create(a).Error; err != nil {
		log.Error("[ApplicationDao][Create] fail to create",
			zap.String("err", err.Error()),
			zap.String("value", fmt.Sprintf("%+v", *a)))
		return err
	}
	return nil
}

func (*ApplicationDao) GetList(db *gorm.DB, condi map[string]interface{}) ([]*Application, error) {
	applications := make([]*Application, 0)
	limit := condi[common.CondiLimit].(int)
	page := condi[common.CondiPage].(int)
	studentId := condi[common.CondiStudentId].(int)
	if err := db.Where("student_id = ?", studentId).Offset((page - 1) * limit).Limit(limit).Order("deadline DESC").Find(&applications).Error; err != nil {
		return nil, err
	}

	return applications, nil
}

func (*ApplicationDao) GetCountByStudentId(db *gorm.DB, studentId int) (int64, error) {
	var total int64
	if err := db.Model(&Application{}).Where("student_id = ?", studentId).Count(&total).Error; err != nil {
		log.Error("[ApplicationDao][GetCountByStudentId] fail to get count",
			zap.Int("student_id", studentId))
		return 0, err
	}

	return total, nil
}

func (*ApplicationDao) GetCountByScholarshipItemId(db *gorm.DB, scholarshipItemId int64) (int64, error) {
	var total int64
	if err := db.Model(&Application{}).Where("scholarship_item_id = ?", scholarshipItemId).Count(&total).Error; err != nil {
		log.Error("[ApplicationDao][GetCountByScholarshipItemId] fail to get count",
			zap.Int64("scholarship_item_id", scholarshipItemId))
		return 0, err
	}

	return total, nil
}

func (*ApplicationDao) DeleteByScholarshipItemId(db *gorm.DB, scholarshipItemId int) error {
	return db.Where("scholarship_item_id = ?", scholarshipItemId).Delete(&Application{}).Error
}

func (*ApplicationDao) GetListByScholarshipItemId(db *gorm.DB, scholarshipItemId int64, page, limit int) ([]*Application, error) {
	applications := make([]*Application, 0)
	if err := db.Where("scholarship_item_id = ?", scholarshipItemId).Offset((page - 1) * limit).Limit(limit).Find(&applications).Error; err != nil {
		log.Error("[ApplicationDao][GetListByScholarshipItemId] fail to get",
			zap.Int64("scholarship_item_id", scholarshipItemId))
		return nil, err
	}
	return applications, nil
}

func (*ApplicationDao) Update(db *gorm.DB, id int64, updates map[string]interface{}) error {
	if err := db.Table("application").Where("id = ?", id).Updates(updates).Error; err != nil {
		log.Error("[ApplicationDao][Update] fail to update",
			zap.String("err", err.Error()),
			zap.String("updates", fmt.Sprintf("%+v", updates)))
		return err
	}
	return nil
}
