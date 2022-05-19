package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(BadRequest, gin.H{
			"error": "request contains no file",
		})

		return
	}

	uuid := uuid.New().String()
	if err := c.SaveUploadedFile(file, "/static/"+uuid); err != nil {
		c.JSON(InternalServerError, gin.H{
			"error": "internal server error",
		})

		return
	}

	c.JSON(SuccessRequest, gin.H{
		"message": "upload file successfully",
	})
}

// func (s *Service) DownloadFile(c *gin.Context) {
// 	filePath := c.Query("file_path")

// }
