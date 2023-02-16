package model

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