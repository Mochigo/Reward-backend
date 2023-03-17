package parse

import (
	"Reward/common"
	"Reward/log"
	"Reward/service/entity"
	"mime/multipart"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

type XlsxParseStrategy struct {
}

func (*XlsxParseStrategy) Parse(ctx *gin.Context, file *multipart.FileHeader) ([]*entity.CreateStudentEntity, error) {
	// 打开xlsx文件
	fileName := filepath.Join(viper.GetString("file_storage"), file.Filename)
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Error("[Parse] fail to open file",
			zap.String("filename", fileName),
			zap.String("err", err.Error()))
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Error("[Parse] fail to close file",
				zap.String("filename", fileName),
				zap.String("err", err.Error()))
		}
	}()

	// 获取sheet1信息
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Error("[Parse] fail to get data in Sheet1",
			zap.String("filename", fileName),
			zap.String("err", err.Error()))
	}

	// 根据模板解析内容
	students := make([]*entity.CreateStudentEntity, 0)
	for _, row := range rows[1:] {
		score, err := strconv.ParseFloat(row[1], 64)
		uid := row[0]
		if err != nil {
			log.Error("[Parse] fail to parse score",
				zap.String("uid", uid),
				zap.String("score_str", row[1]),
				zap.String("err", err.Error()))
			continue
		}
		students = append(students, &entity.CreateStudentEntity{
			Score:     score,
			Uid:       uid,
			CollegeId: ctx.GetInt(common.TokenCollegeID),
			Password:  common.DefaultPassword,
		})
	}

	return students, nil
}
