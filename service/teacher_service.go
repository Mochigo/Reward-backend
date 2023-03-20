package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"Reward/common"
	"Reward/common/token"
	"Reward/model"
	"Reward/service/entity"
)

type TeacherService struct {
	ctx                           *gin.Context
	teacherDao                    *model.TeacherDao
	teacherStudentRelationshipDao *model.TeacherStudentRelationshipDao
	declarationDao                *model.DeclarationDao
}

func NewTeacherService(ctx *gin.Context) *TeacherService {
	return &TeacherService{
		ctx:                           ctx,
		teacherDao:                    model.GetTeacherDao(),
		teacherStudentRelationshipDao: model.GetTeacherStudentRelationshipDao(),
		declarationDao:                model.GetDeclarationDao(),
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

func (s *TeacherService) GetTeacherRole(uid string) (string, error) {
	teacher, err := s.teacherDao.GetTeacherByUID(model.DB.Self, uid)
	if err != nil {
		return "", err
	}

	return teacher.Role, nil
}

func (s *TeacherService) Sign(uid string) (string, error) {
	teacher, err := s.teacherDao.GetTeacherByUID(model.DB.Self, uid)
	if err != nil {
		return common.StringEmpty, err
	}

	t, err := token.GenerateToken(&token.PayLoad{
		UserID:    int(teacher.Id),
		CollegeId: int(teacher.CollegeId),
		Expired:   time.Duration(viper.GetInt("token_expired") * int(time.Second)),
	})
	if err != nil {
		return common.StringEmpty, err
	}

	return t, nil
}

func (s *TeacherService) GetStudentsDeclarations(req *entity.GetStudentsDeclarationsEntity) ([]*entity.DeclarationEntity, int64, error) {
	relationships, err := s.teacherStudentRelationshipDao.GetStudentsByTeacherId(model.DB.Self, int64(s.ctx.GetInt(common.TokenUserID)))
	if err != nil {
		return nil, 0, err
	}

	stuIds := make([]int64, 0, len(relationships))
	for _, r := range relationships {
		stuIds = append(stuIds, r.StudentId)
	}

	dl, err := s.declarationDao.GetDeclarationsByStudentIdsAndStatus(model.DB.Self, stuIds, common.StatusPROCESS, req.Page, req.Limit)
	if err != nil {
		return nil, 0, err
	}

	declarations := make([]*entity.DeclarationEntity, 0, len(dl))
	for _, d := range dl {
		declarations = append(declarations, &entity.DeclarationEntity{
			Id:             d.Id,
			ApplicationId:  d.ApplicationId,
			StudentId:      d.StudentId,
			Name:           d.Name,
			Level:          d.Level,
			Status:         d.Status,
			RejectedReason: d.Status,
			Url:            d.Url,
		})
	}

	total, err := s.declarationDao.GetCountByStudentIdAndStatus(model.DB.Self, stuIds, common.StatusPROCESS)
	if err != nil {
		return nil, 0, err
	}

	return declarations, total, nil
}
