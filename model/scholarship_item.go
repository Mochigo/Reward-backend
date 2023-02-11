package model

type ScholarshipItem struct {
	Id            int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"` // 奖学金子项id
	Name          string `gorm:"column:name;NOT NULL"`                 // 奖学金子项名称
	ScholarshipId int64  `gorm:"column:scholarship_id;NOT NULL"`       // 奖学金id
}

func (m *ScholarshipItem) TableName() string {
	return "`scholarship_item`"
}
