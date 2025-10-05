package grpc_server

import (
	"context"
	"database/sql"
	"github.com/medabbassi/go_server/cmd/grpc_server/servers"
	"log"
	"net"

	"github.com/medabbassi/go_server/pkg/config"
	pb "github.com/medabbassi/go_server/pkg/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedErrorServiceServer
}

func (s *server) LogError(ctx context.Context, in *pb.ErrorLogRequest) (*pb.ErrorLogResponse, error) {
	log.Printf("Received error log - ID: %s, Message: %s", in.Id, in.Message)
	return &pb.ErrorLogResponse{Status: "received"}, nil
}

func Start(db *sql.DB) {
	lis, err := net.Listen("tcp", ":"+config.GetConfig().GRPCPort)
	if err != nil {
		log.Printf("⚠️ Failed to start gRPC server on port %s: %v", config.GetConfig().GRPCPort, err)
		log.Printf("gRPC server will not be available, but REST API will continue running")
		return
	}

	s := grpc.NewServer()

	// ✅ Register the DB-backed implementation
	pb.RegisterErrorServiceServer(s, servers.NewErrorServiceServer(db))

	log.Printf("✅ gRPC server started on port %s", config.GetConfig().GRPCPort)

	if err := s.Serve(lis); err != nil {
		log.Printf("⚠️ gRPC server stopped: %v", err)
	}
}
