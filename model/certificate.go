package model

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
