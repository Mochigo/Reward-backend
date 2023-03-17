package parse

import (
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"Reward/common"
	"Reward/log"
	"Reward/service/entity"
)

var parseStrategy map[string]ParseStrategy

func init() {
	parseStrategy = make(map[string]ParseStrategy)
	// 注册
	parseStrategy[common.XlsxExt] = &XlsxParseStrategy{}
}

type ParseStrategy interface {
	Parse(ctx *gin.Context, file *multipart.FileHeader) ([]*entity.CreateStudentEntity, error)
}

// 因为策略与状态无关，所以这里选择使用map提前初始化，然后取用
func ParseStrategyFactory(ext string) (ParseStrategy, error) {
	if _, exist := parseStrategy[ext]; !exist {
		log.Error("[ParseStrategyFactory] invaild extension",
			zap.String("ext", ext))
		return nil, common.ErrInvalidFileType
	}

	return parseStrategy[ext], nil
}

type ParseService struct {
	ctx *gin.Context
	// strategy ParseStrategy
}

func NewParseService(c *gin.Context) *ParseService {
	return &ParseService{
		ctx: c,
	}
}

func (s *ParseService) Parse(file *multipart.FileHeader) ([]*entity.CreateStudentEntity, error) {
	// 保存到本地
	if err := saveFile(s.ctx, file); err != nil {
		return nil, err
	}

	strategy, err := ParseStrategyFactory(filepath.Ext(file.Filename))
	if err != nil {
		return nil, err
	}
	return strategy.Parse(s.ctx, file)
}

func saveFile(ctx *gin.Context, file *multipart.FileHeader) error {
	fileName := filepath.Join(viper.GetString("file_storage"), file.Filename)
	err := ctx.SaveUploadedFile(file, fileName)
	if err != nil {
		log.Error("[saveFile] fail to save file",
			zap.String("filename", fileName),
			zap.String("err", err.Error()))
		return err
	}

	return nil
}
