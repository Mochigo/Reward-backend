package handler

import (
	"github.com/gin-gonic/gin"
)

type CreateScholarshipRequest struct {
	Attachments []string `json:"attachments"`
	StartTime   string   `json:"start_time"`
	EndTime     string   `json:"end_time"`
}

func CreateScholarship(c *gin.Context) {

}
