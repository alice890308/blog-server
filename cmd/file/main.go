package file

import (
	"log"
	"net/http"

	"github.com/alice890308/blog-server/modules/file/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func NewFileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "file [service]",
		Short: "start file's service",
	}

	svc := service.NewService()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.POST("/file", func(c *gin.Context) {
		svc.UploadFile(c)
	})

	router.Run(":8080")

	fs := http.FileServer(http.Dir("/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

	return cmd
}
