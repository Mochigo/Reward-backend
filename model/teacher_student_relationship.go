package model

import (
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"Reward/log"
)

type TeacherStudentRelationship struct {
	Id        int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	TeacherId int64 `gorm:"column:teacher_id;NOT NULL"` // 教师id
	StudentId int64 `gorm:"column:student_id;NOT NULL"` // 学生id
}

func (*TeacherStudentRelationship) TableName() string {
	return "`teacher_student_relationship`"
}

type TeacherStudentRelationshipDao struct{}

var teacherStudentRelationshipDao *TeacherStudentRelationshipDao
var teacherStudentRelationshipDaoOnce sync.Once

func GetTeacherStudentRelationshipDao() *TeacherStudentRelationshipDao {
	teacherStudentRelationshipDaoOnce.Do(func() {
		teacherStudentRelationshipDao = &TeacherStudentRelationshipDao{}
	})

	return teacherStudentRelationshipDao
}

func (*TeacherStudentRelationshipDao) GetStudentsByTeacherId(db *gorm.DB, teacherId int64) ([]*TeacherStudentRelationship, error) {
	relationships := make([]*TeacherStudentRelationship, 0)
	if err := db.Model(&TeacherStudentRelationship{}).Where("teacher_id = ?", teacherId).Find(&relationships).Error; err != nil {
		log.Error("[TeacherStudentRelationshipDao][GetStudentsByTeacherId] failed to get",
			zap.Int64("teacher_id ", teacherId),
			zap.String("err", err.Error()))
		return nil, err
	}
	return relationships, nil
}
