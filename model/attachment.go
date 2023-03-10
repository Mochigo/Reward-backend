package model

import (
	"sync"

	"gorm.io/gorm"
)

type Attachment struct {
	Id            int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"` // 附件子项id
	ScholarshipId int64  `gorm:"column:scholarship_id;NOT NULL"`       // 奖学金id
	Url           string `gorm:"column:url;NOT NULL"`                  // 地址
}

func (m *Attachment) TableName() string {
	return "`attachment`"
}

type AttachmentDao struct{}

var attachmentDaoOnce sync.Once
var attachmentDao *AttachmentDao

func GetAttachmentDao() *AttachmentDao {
	attachmentDaoOnce.Do(func() {
		attachmentDao = &AttachmentDao{}
	})
	return attachmentDao
}

func (*AttachmentDao) Create(db *gorm.DB, a *Attachment) error {
	return db.Model(&Attachment{}).Create(a).Error
}

func (*AttachmentDao) GetListByScholarshipId(db *gorm.DB, scholarshipId int64) ([]*Attachment, error) {
	attachments := make([]*Attachment, 0)
	if err := db.Model(&Attachment{}).Where("scholarship_id = ?", scholarshipId).Find(&attachments).Error; err != nil {
		return nil, err
	}

	return attachments, nil
}
