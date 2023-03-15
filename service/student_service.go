package service

import (
	"Reward/common"
	"Reward/common/token"
	"Reward/model"
	"Reward/service/entity"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type StudentService struct {
	ctx        *gin.Context
	studentDao *model.StudentDao
	collegeDao *model.CollegeDao
}

func NewStudentService(ctx *gin.Context) *StudentService {
	return &StudentService{
		ctx:        ctx,
		studentDao: model.GetStudentDao(),
		collegeDao: model.GetCollegeDao(),
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
		return common.StringEmpty, err
	}

	t, err := token.GenerateToken(&token.PayLoad{
		UserID:    int(stu.Id),
		CollegeId: int(stu.CollegeId),
		Expired:   time.Duration(viper.GetInt("token_expired") * int(time.Second)),
	})
	if err != nil {
		return common.StringEmpty, err
	}

	return t, nil
}

func (s *StudentService) GetStudentInfo(sid int64) (*entity.StudentEntity, error) {
	stu, err := s.studentDao.GetStudentById(model.DB.Self, sid)
	if err != nil {
		return nil, err
	}

	college, err := s.collegeDao.GetCollegeById(model.DB.Self, stu.CollegeId)
	if err != nil {
		return nil, err
	}
	return &entity.StudentEntity{
		Id:      stu.Id,
		Score:   stu.Score,
		Uid:     stu.Uid,
		College: college.Name,
	}, nil
}
