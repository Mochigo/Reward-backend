package model

type TeacherStudentRelationship struct {
	Id        int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	TeacherId int64 `gorm:"column:teacher_id;NOT NULL"` // 教师id
	StudentId int64 `gorm:"column:student_id;NOT NULL"` // 学生id
}

func (m *TeacherStudentRelationship) TableName() string {
	return "`teacher_student_relationship`"
}
