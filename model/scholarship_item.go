package model

import (
	"Reward/log"
	"fmt"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ScholarshipItem struct {
	Id            int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"` // 奖学金子项id
	Name          string `gorm:"column:name;NOT NULL"`                 // 奖学金子项名称
	ScholarshipId int64  `gorm:"column:scholarship_id;NOT NULL"`       // 奖学金id
}

func (m *ScholarshipItem) TableName() string {
	return "`scholarship_item`"
}

type ScholarshipItemDao struct{}

var scholarshipItemDaoOnce sync.Once
var scholarshipItemDao *ScholarshipItemDao

func GetScholarshipItemDao() *ScholarshipItemDao {
	scholarshipItemDaoOnce.Do(func() {
		scholarshipItemDao = &ScholarshipItemDao{}
	})
	return scholarshipItemDao
}

func (*ScholarshipItemDao) Create(db *gorm.DB, item *ScholarshipItem) error {
	return db.Model(&ScholarshipItem{}).Create(item).Error
}

func (*ScholarshipItemDao) DeleteByID(db *gorm.DB, id int64) error {
	return db.Delete(&ScholarshipItem{}, id).Error
}

func (*ScholarshipItemDao) GetList(db *gorm.DB, scholarshipId int64) ([]*ScholarshipItem, error) {
	scholarshipItemList := make([]*ScholarshipItem, 0)
	if err := db.Model(&ScholarshipItem{}).Where("scholarship_id = ?", scholarshipId).Find(&scholarshipItemList).Error; err != nil {
		return nil, err
	}

	return scholarshipItemList, nil
}

func (*ScholarshipItemDao) BatchGetByIds(db *gorm.DB, Ids []int64) ([]*ScholarshipItem, error) {
	scholarshipItemList := make([]*ScholarshipItem, 0)
	if err := db.Model(&ScholarshipItem{}).Where("id IN ?", Ids).Find(&scholarshipItemList).Error; err != nil {
		log.Error("[ScholarshipItemDao][BatchGetByIds] fail to get",
			zap.String("ids", fmt.Sprintf("%+v", Ids)),
			zap.String("err", err.Error()))
		return nil, err
	}
	return scholarshipItemList, nil
}
