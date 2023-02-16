package service

import (
	"Reward/common"
	"Reward/common/token"
	"Reward/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type TeacherService struct {
	ctx        *gin.Context
	teacherDao *model.TeacherDao
}

func NewTeacherService(ctx *gin.Context) *TeacherService {
	return &TeacherService{
		ctx:        ctx,
		teacherDao: model.GetTeacherDao(),
	}
}

// 判断当前uid是否存在，且密码与之对应
func (s *TeacherService) VerifyTeacher(uid, password string) (bool, error) {
	teacher, err := s.teacherDao.GetTeacherByUID(model.DB.Self, uid)
	if err != nil {
		return false, err
	}

	if teacher.Password != password {
		return false, err
	}

	return true, nil
}

func (s *TeacherService) Sign(uid string) (string, error) {
	teacher, err := s.teacherDao.GetTeacherByUID(model.DB.Self, uid)
	if err != nil {
		return common.StringEmpry, err
	}

	t, err := token.GenerateToken(&token.PayLoad{
		UserID:    int(teacher.Id),
		CollegeId: int(teacher.CollegeId),
		Expired:   time.Duration(viper.GetInt("token_expired") * int(time.Second)),
	})
	if err != nil {
		return common.StringEmpry, err
	}

	return t, nil
}
