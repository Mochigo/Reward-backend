package handler

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"Reward/common/errno"
	"Reward/common/response"
	"Reward/common/utils"
	"Reward/log"
)

type UploadFileResponse struct {
	Urls []string `json:"urls"`
}

func UploadFile(c *gin.Context) {
	log.Info("UploadFile function called.",
		zap.String("X-Request-Id", utils.GetReqID(c)))

	// file, err := c.FormFile("file")
	form, err := c.MultipartForm()
	if err != nil {
		response.SendBadRequest(c, errno.ErrFileNotFound, nil, err.Error(), utils.GetUpFuncInfo(2))
		return
	}
	resp := UploadFileResponse{make([]string, 0)}
	files := form.File["files"]
	for _, file := range files {
		filename := filepath.Join(viper.GetString("file_storage"), file.Filename)
		err := c.SaveUploadedFile(file, filename)
		if err != nil {
			response.SendInternalServerError(c, errno.ErrUploadFailed, nil, err.Error(), utils.GetUpFuncInfo(2))
			return
		}
		resp.Urls = append(resp.Urls, viper.GetString("url")+"/"+filename)
	}

	response.SendResponse(c, errno.OK, resp)
}
