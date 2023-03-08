package model

import (
	"Reward/common"
	"sync"

	"gorm.io/gorm"
)

type Application struct {
	Id                int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`   // 申请id
	ScholarshipItemId int64  `gorm:"column:scholarship_item_id;NOT NULL"`    // 奖学金子项id
	ScholarshipId     int64  `gorm:"column:scholarship_id;NOT NULL"`         // 奖学金id
	StudentId         int64  `gorm:"column:student_id;NOT NULL"`             // 申请学生id
	Status            string `gorm:"column:status;default:PROCESS;NOT NULL"` // 申请状态，APPROVE-通过|PROCESS-处理中|FAILURE-驳回
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
	return db.Model(&Application{}).Create(a).Error
}

func (*ApplicationDao) GetList(db *gorm.DB, condi map[string]interface{}) ([]*Application, error) {
	applications := make([]*Application, 0)
	limit := condi[common.CondiLimit].(int)
	page := condi[common.CondiPage].(int)
	studentId := condi[common.CondiStudentId].(int)
	if err := db.Where("student_id = ?", studentId).Offset((page - 1) * limit).Limit(limit).Find(&applications).Error; err != nil {
		return nil, err
	}

	return applications, nil
}

func (*ApplicationDao) GetCountByStudentId(db *gorm.DB, studentId int) (int64, error) {
	var total int64
	if err := db.Model(&Application{}).Where("student_id = ?", studentId).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}
