package service

import (
	"Reward/common"
	"Reward/common/utils"
	"Reward/log"
	"Reward/model"
	"Reward/service/entity"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApplicationService struct {
	Ctx                *gin.Context
	applicationDao     *model.ApplicationDao
	scholarshipItemDao *model.ScholarshipItemDao
	scholarshipDao     *model.ScholarshipDao
}

func NewApplicationService(ctx *gin.Context) *ApplicationService {
	return &ApplicationService{
		Ctx:                ctx,
		applicationDao:     model.GetApplicationDao(),
		scholarshipItemDao: model.GetScholarshipItemDao(),
		scholarshipDao:     model.GetScholarshipDao(),
	}
}

func (s *ApplicationService) CreateApplication(req *entity.CreateApplicationEntity) error {
	deadline, err := utils.GetDateTime(req.Deadline)
	if err != nil {
		log.Error("[CreateApplication] fail to parse time",
			zap.String("time", req.Deadline))
		return common.ErrTimeParse
	}

	return s.applicationDao.Create(model.DB.Self, &model.Application{
		ScholarshipItemId: req.ScholarshipItemId,
		ScholarshipId:     req.ScholarshipId,
		StudentId:         int64(s.Ctx.GetInt(common.TokenUserID)),
		Status:            common.StatusPROCESS,
		Deadline:          deadline,
	})
}

func (s *ApplicationService) GetUserApplication(req *entity.GetUserApplicationEntity) ([]*entity.ApplicationEntity, int64, error) {
	condi := make(map[string]interface{})
	condi[common.CondiPage] = req.Page
	condi[common.CondiLimit] = req.Limit
	condi[common.CondiStudentId] = s.Ctx.GetInt(common.TokenUserID)
	// 查询申请
	al, err := s.applicationDao.GetList(model.DB.Self, condi)
	if err != nil {
		return nil, 0, err
	}

	scholarship := make(map[int64]*model.Scholarship)
	scholarshipIds := make([]int64, 0, len(al))

	scholarshipItem := make(map[int64]*model.ScholarshipItem)
	scholarshipItemIds := make([]int64, 0, len(al))

	for _, a := range al {
		scholarshipItemIds = append(scholarshipItemIds, a.ScholarshipItemId)
		if _, ok := scholarship[a.ScholarshipId]; !ok {
			scholarshipIds = append(scholarshipIds, a.ScholarshipId)
			scholarship[a.ScholarshipId] = &model.Scholarship{}
		}
	}

	scholarshipItems, err := s.scholarshipItemDao.BatchGetByIds(model.DB.Self, scholarshipItemIds)
	if err != nil {
		return nil, 0, err
	}
	for _, item := range scholarshipItems {
		scholarshipItem[item.Id] = item
	}

	scholarships, err := s.scholarshipDao.BatchGetByIds(model.DB.Self, scholarshipIds)
	if err != nil {
		return nil, 0, err
	}
	for _, s := range scholarships {
		scholarship[s.Id].Name = s.Name
		scholarship[s.Id].CollegeId = s.CollegeId
		scholarship[s.Id].EndTime = s.EndTime
		scholarship[s.Id].StartTime = s.StartTime
	}

	// 返回拼接申请内容
	applications := make([]*entity.ApplicationEntity, 0, len(al))
	for _, a := range al {
		if _, ok := scholarshipItem[a.ScholarshipItemId]; !ok {
			continue
		}
		if _, ok := scholarship[a.ScholarshipId]; !ok {
			continue
		}
		tmp := &entity.ApplicationEntity{
			Id:                  a.Id,
			ScholarshipItemId:   a.ScholarshipItemId,
			ScholarshipItemName: scholarshipItem[a.ScholarshipItemId].Name,
			ScholarshipId:       a.ScholarshipId,
			ScholarshipName:     scholarship[a.ScholarshipId].Name,
			StudentId:           a.StudentId,
			Status:              a.Status,
			Deadline:            a.Deadline.Format(utils.LayoutDateTime),
		}
		applications = append(applications, tmp)
	}

	total, err := s.applicationDao.GetCountByStudentId(model.DB.Self, s.Ctx.GetInt(common.TokenUserID))
	if err != nil {
		return nil, 0, err
	}

	return applications, total, nil
}
