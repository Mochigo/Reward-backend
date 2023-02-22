package model

import (
	"Reward/common"
	"sync"

	"gorm.io/gorm"
)

type Certificate struct {
	Id             int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`   // 荣誉id
	ApplicationId  int64  `gorm:"column:application_id;NOT NULL"`         // 申请id
	Name           string `gorm:"column:name;NOT NULL"`                   // 荣誉名称
	Level          string `gorm:"column:level;NOT NULL"`                  // 荣誉级别, 校级-01|省级-02|国家级|-03
	Status         string `gorm:"column:status;default:PROCESS;NOT NULL"` // 申请状态，APPROVED-通过|PROCESS-待处理|REJECTED-驳回
	RejectedReason string `gorm:"column:rejected_reason"`                 // 驳回理由
	Url            string `gorm:"column:url;NOT NULL"`                    // 证明文件连接
}

func (m *Certificate) TableName() string {
	return "`certificate`"
}

type CertificateDao struct{}

var certificateDaoOnce sync.Once
var certificateDao *CertificateDao

func GetCertificateDao() *CertificateDao {
	certificateDaoOnce.Do(func() {
		certificateDao = &CertificateDao{}
	})

	return certificateDao
}

func (*CertificateDao) Create(db *gorm.DB, c *Certificate) error {
	return db.Model(&Certificate{}).Create(c).Error
}

func (*CertificateDao) Delete(db *gorm.DB, id int64) error {
	return db.Delete(&Certificate{}, id).Error
}

func (*CertificateDao) Update(db *gorm.DB, id int64, updates map[string]interface{}) error {
	return db.Table("certificate").Where("id = ?", id).Updates(updates).Error
}

func (*CertificateDao) GetCertificatesByApplicationIdAndStatus(db *gorm.DB, applicationId int64, status string) ([]*Certificate, error) {
	db = db.Model(&Certificate{})
	if status != common.StringEmpty {
		db = db.Where("status = ?", status)
	}
	cl := make([]*Certificate, 0)
	if err := db.Where("application_id = ?", applicationId).Find(&cl).Error; err != nil {
		return nil, err
	}

	return cl, nil
}
