package file

import (
	"context"
	"log"

	"github.com/alice890308/blog-server/modules/file/service"
	"github.com/alice890308/blog-server/pkg/authkit"
	"github.com/alice890308/blog-server/pkg/logkit"
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"github.com/spf13/cobra"
)

type APIArgs struct {
	logkit.LoggerConfig `group:"logger" namespace:"logger" env-namespace:"LOGGER"`
	authkit.JWTConfig   `group:"jwt" namespace:"jwt" env-namespace:"JWT"`
}

func NewFileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "file [service]",
		Short: "start file's service",
	}

	ctx := context.Background()
	var args APIArgs
	if _, err := flags.NewParser(&args, flags.Default).Parse(); err != nil {
		log.Fatal("failed to parse flag", err.Error())
	}

	logger := logkit.NewLogger(&args.LoggerConfig)
	defer func() {
		_ = logger.Sync()
	}()

	ctx = logger.WithContext(ctx)
	jwtManager := authkit.NewJWTManager(ctx, &args.JWTConfig)
	svc := service.NewService(jwtManager)

	router := gin.Default()
	router.StaticFS("/static", gin.Dir("/static", false))

	router.GET("/status", func(c *gin.Context) {
		svc.Status(c)
	})

	router.POST("/upload", func(c *gin.Context) {
		svc.Upload(c)
	})

	router.Run(":8080")

	return cmd
}
