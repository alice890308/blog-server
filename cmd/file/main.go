package file

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alice890308/blog-server/modules/file/middleware"
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
	return &cobra.Command{
		Use:   "file",
		Short: "start file's service",
		RunE:  runFile,
	}
}

func runFile(_ *cobra.Command, _ []string) error {
	ctx := context.Background()
	var args APIArgs
	if _, err := flags.NewParser(&args, flags.Default).Parse(); err != nil {
		log.Fatal("failed to parse flag", err.Error())
	}

	logger := logkit.NewLogger(&args.LoggerConfig)

	ctx = logger.WithContext(ctx)
	jwtManager := authkit.NewJWTManager(ctx, &args.JWTConfig)
	svc := service.NewService(jwtManager)

	router := gin.Default()

	router.Use(middleware.CORS())

	router.StaticFS("/static", gin.Dir("/static", false))

	router.GET("/", func(c *gin.Context) {
		svc.Status(c)
	})

	router.POST("/upload", func(c *gin.Context) {
		svc.Upload(c)
	})

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = logger.Sync()
	}()

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	<-shutdownCh

	return nil
}
