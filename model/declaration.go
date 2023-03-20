package model

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"Reward/common"
	"Reward/log"
)

// 申报表
type Declaration struct {
	Id             int64  `gorm:"column:id" json:"id"`
	ApplicationId  int64  `gorm:"column:application_id" json:"application_id"`
	StudentId      int64  `gorm:"column:student_id" json:"student_id"`
	Name           string `gorm:"column:name" json:"name"`
	Level          string `gorm:"column:level" json:"level"`
	Status         string `gorm:"column:status" json:"status"`
	RejectedReason string `gorm:"column:rejected_reason" json:"rejected_reason"`
	Url            string `gorm:"column:url" json:"url"`
}

func (*Declaration) TableName() string {
	return "declaration"
}

type DeclarationDao struct{}

var declarationDaoOnce sync.Once
var declarationDao *DeclarationDao

func GetDeclarationDao() *DeclarationDao {
	declarationDaoOnce.Do(func() {
		declarationDao = &DeclarationDao{}
	})

	return declarationDao
}

func (*DeclarationDao) Create(db *gorm.DB, d *Declaration) error {
	if err := db.Model(&Declaration{}).Create(d).Error; err != nil {
		log.Error("[DeclarationDao][Create] fail to create",
			zap.String("err", err.Error()),
			zap.String("value", fmt.Sprintf("%+v", *d)))
		return err
	}
	return nil
}

func (*DeclarationDao) Delete(db *gorm.DB, id int64) error {
	if err := db.Delete(&Declaration{}, id).Error; err != nil {
		log.Error("[DeclarationDao][Delete] fail to delete",
			zap.String("err", err.Error()),
			zap.Int64("id", id))
		return err
	}
	return nil
}

func (*DeclarationDao) Update(db *gorm.DB, id int64, updates map[string]interface{}) error {
	if err := db.Table("declaration").Where("id = ?", id).Updates(updates).Error; err != nil {
		log.Error("[DeclarationDao][Update] fail to update",
			zap.String("err", err.Error()),
			zap.String("updates", fmt.Sprintf("%+v", updates)))
		return err
	}
	return db.Table("declaration").Where("id = ?", id).Updates(updates).Error
}

func (*DeclarationDao) GetDeclarationsByApplicationIdAndStatus(db *gorm.DB, applicationId int64, status string) ([]*Declaration, error) {
	db = db.Model(&Declaration{})
	if status != common.StringEmpty {
		db = db.Where("status = ?", status)
	}
	dl := make([]*Declaration, 0)
	if err := db.Where("application_id = ?", applicationId).Find(&dl).Error; err != nil {
		log.Error("[DeclarationDao][GetDeclarationsByApplicationIdAndStatus] fail to get",
			zap.String("err", err.Error()),
			zap.Int64("application_id", applicationId),
			zap.String("status", status),
		)
		return nil, err
	}

	return dl, nil
}

func (*DeclarationDao) GetDeclarationsByStudentIdsAndStatus(db *gorm.DB, stuIds []int64, status string, page int, limit int) ([]*Declaration, error) {
	db = db.Model(&Declaration{})
	if status != common.StringEmpty {
		db = db.Where("status = ?", status)
	}

	dl := make([]*Declaration, 0)
	if err := db.Where("student_id in ?", stuIds).Offset((page - 1) * limit).Limit(limit).Find(&dl).Error; err != nil {
		log.Error("[DeclarationDao][GetDeclarationsByStudentIdAndStatus] fail to get",
			zap.String("err", err.Error()),
			zap.String("stuIds", fmt.Sprintf("%+v", stuIds)),
			zap.String("status", status),
		)
		return nil, err
	}

	return dl, nil
}

func (*DeclarationDao) GetCountByStudentIdAndStatus(db *gorm.DB, stuIds []int64, status string) (int64, error) {
	var total int64

	db = db.Model(&Declaration{})
	if status != common.StringEmpty {
		db = db.Where("status = ?", status)
	}

	if err := db.Where("student_id in ?", stuIds).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}
