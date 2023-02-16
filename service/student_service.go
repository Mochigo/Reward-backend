package service

import (
	"Reward/common"
	"Reward/common/token"
	"Reward/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type StudentService struct {
	ctx        *gin.Context
	studentDao *model.StudentDao
}

func NewStudentService(ctx *gin.Context) *StudentService {
	return &StudentService{
		ctx:        ctx,
		studentDao: model.GetStudentDao(),
	}
}

// 判断当前uid是否存在，且密码与之对应
func (s *StudentService) VerifyStudent(uid, password string) (bool, error) {
	stu, err := s.studentDao.GetStudentByUID(model.DB.Self, uid)
	if err != nil {
		return false, err
	}

	if stu.Password != password {
		return false, err
	}

	return true, nil
}

func (s *StudentService) Sign(uid string) (string, error) {
	stu, err := s.studentDao.GetStudentByUID(model.DB.Self, uid)
	if err != nil {
		return common.StringEmpry, err
	}

	t, err := token.GenerateToken(&token.PayLoad{
		UserID:    int(stu.Id),
		CollegeId: int(stu.CollegeId),
		Expired:   time.Duration(viper.GetInt("token_expired") * int(time.Second)),
	})
	if err != nil {
		return common.StringEmpry, err
	}

	return t, nil
}
