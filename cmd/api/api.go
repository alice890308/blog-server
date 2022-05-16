package api

import (
	"context"
	"log"
	"net"

	"github.com/alice890308/blog-server/modules/api/dao"
	"github.com/alice890308/blog-server/modules/api/pb"
	"github.com/alice890308/blog-server/modules/api/service"
	"github.com/alice890308/blog-server/pkg/authkit"
	"github.com/alice890308/blog-server/pkg/logkit"
	"github.com/alice890308/blog-server/pkg/mongokit"
	"github.com/alice890308/blog-server/pkg/runkit"
	flags "github.com/jessevdk/go-flags"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func newAPICommand() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "start api server",
		RunE:  runAPI,
	}
}

type APIArgs struct {
	GRPCAddr              string `long:"grpc_addr" env:"RCPC_ADDR" default:":8081"`
	runkit.GracefulConfig `group:"graceful" namespace:"graceful" env-namespace:"GRACEFUL"`
	logkit.LoggerConfig   `group:"logger" namespace:"logger" env-namespace:"LOGGER"`
	mongokit.MongoConfig  `group:"mongo" namespace:"mongo" env-namespace:"MONGO"`
	authkit.JWTConfig     `group:"jwt" namespace:"jwt" env-namespace:"JWT"`
}

func runAPI(_ *cobra.Command, _ []string) error {
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

	mongoClient := mongokit.NewMongoClient(ctx, &args.MongoConfig)
	defer func() {
		if err := mongoClient.Close(); err != nil {
			logger.Fatal("failed to close mongo client", zap.Error(err))
		}
	}()

	postDAO := dao.NewMongoPostDAO(mongoClient.Database().Collection("posts"))
	userDAO := dao.NewMongoUserDAO(mongoClient.Database().Collection("users"))
	jwtManager := authkit.NewJWTManager(ctx, &args.JWTConfig)
	svc := service.NewService(postDAO, userDAO, jwtManager)

	logger.Info("listen to gRPC addr", zap.String("grpc_addr", args.GRPCAddr))
	lis, err := net.Listen("tcp", args.GRPCAddr)
	if err != nil {
		logger.Fatal("failed to listen gRPC addr", zap.Error(err))
	}
	defer func() {
		if err := lis.Close(); err != nil {
			logger.Fatal("failed to close gRPC listner", zap.Error(err))
		}
	}()

	auth := authkit.NewAuthService(jwtManager)

	return runkit.GracefulRun(serveGRPC(lis, svc, logger, grpc.UnaryInterceptor(auth.UnaryServerInterceptor())), &args.GracefulConfig)
}

func serveGRPC(lis net.Listener, svc *service.Service, logger *logkit.Logger, opt ...grpc.ServerOption) runkit.GracefulRunFunc {
	grpcServer := grpc.NewServer(opt...)
	pb.RegisterPostServer(grpcServer, svc)
	pb.RegisterUserServer(grpcServer, svc)
	pb.RegisterSessionServer(grpcServer, svc)

	return func(ctx context.Context) error {
		go func() {
			if err := grpcServer.Serve(lis); err != nil {
				logger.Error("failed to run gRPC srever", zap.Error(err))
			}
		}()

		<-ctx.Done()

		grpcServer.GracefulStop()

		return nil
	}
}
