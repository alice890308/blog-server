package service

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/alice890308/blog-server/pkg/authkit"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service struct {
	authkit.JWT
}

func NewService(jwtManager authkit.JWT) *Service {
	return &Service{jwtManager}
}

func (s *Service) Status(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": "ok",
	})
}

func (s *Service) Upload(c *gin.Context) {
	userID := s.getUserID(c.Request.Header["Authorization"][0])
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid token",
		})
		return
	}

	checkDir(userID)

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "upload file error",
		})
		return
	}

	buffer := make([]byte, 512)
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "file header open error",
		})
		return
	}

	file.Read(buffer)
	contentType := http.DetectContentType(buffer)
	file.Seek(0, io.SeekStart)
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "file read error",
		})
		return
	}

	filePath := "/static/" + userID + "/" + uuid.New().String()
	switch contentType {
	case "image/jpg":
		filePath = filePath + ".jpg"
	case "image/png":
		filePath = filePath + ".png"
	case "image/jpeg":
		filePath = filePath + ".jpeg"
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "file type error",
		})
		return
	}

	ioutil.WriteFile(filePath, fileBytes, 0644)
	c.JSON(http.StatusAccepted, gin.H{
		"message":  "success",
		"filepath": filePath,
	})
}

func (s *Service) getUserID(accessToken string) string {

	accessToken = string(accessToken[7:])
	log.Println(accessToken)

	payload, err := s.JWT.Verify(accessToken)
	if err != nil {
		return ""
	}

	return payload.UserID
}

func checkDir(userID string) {
	path := "/static/" + userID
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, 0666)
	}
}
