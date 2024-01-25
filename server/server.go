package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"math/rand"
	"net"

	zk "github.com/fbac/zkproof-grpc/internal/zk"
	pb "github.com/fbac/zkproof-grpc/protobuf"
	"google.golang.org/grpc"
)

/*
	Calculations and data sourced from https://crypto.stackexchange.com/questions/99262/chaum-pedersen-protocol
*/

// The AuthServer interface forces us to embed UnimplementedAuthServer.
type grpcServer struct {
	pb.UnimplementedAuthServer
}

// UserData holds the exchanged Y's
type UserData struct {
	Y1 int64
	Y2 int64
}

var (
	// Users database
	users           = make(map[string]UserData)
	challenge int64 = 0
)

// grpcServer has to implement AuthServer.
// this line won't compile if there is not interface compliance.
var _ pb.AuthServer = (*grpcServer)(nil)

func (s *grpcServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// If the user already exists it doesn't need to register again.
	if _, ok := users[in.User]; ok {
		slog.WarnContext(ctx, "Existing", slog.String("user", in.User))
		return &pb.RegisterResponse{}, fmt.Errorf("user already exists")
	}

	// Add new user
	users[in.User] = UserData{
		Y1: in.Y1,
		Y2: in.Y2,
	}
	slog.InfoContext(ctx, "Registered", slog.String("user", in.User))

	// This should be invoked as DebugContext properly
	slog.InfoContext(ctx, "Debug: Y's provided", slog.Int64("Y1", users[in.User].Y1), slog.Int64("Y2", users[in.User].Y2))
	return &pb.RegisterResponse{}, nil
}

func (s *grpcServer) CreateAuthenticationChallenge(ctx context.Context, in *pb.AuthenticationChallengeRequest) (*pb.AuthenticationChallengeResponse, error) {
	// If the user doesn't exist, don't create a challenge.
	_, ok := users[in.User]
	if !ok {
		slog.WarnContext(ctx, "Not existing", slog.String("user", in.User))
		return &pb.AuthenticationChallengeResponse{}, fmt.Errorf("user already exists")
	}

	// Update challenge
	slog.InfoContext(ctx, "Creating challenge for", slog.String("user", in.User))
	challenge = rand.Int63n(1000)

	// Return the proper challenge for this user.
	return &pb.AuthenticationChallengeResponse{
		AuthId: in.User,
		C:      challenge}, nil
}

func (s *grpcServer) VerifyAuthentication(ctx context.Context, in *pb.AuthenticationAnswerRequest) (*pb.AuthenticationAnswerResponse, error) {
	// Calculations sourced from https://crypto.stackexchange.com/questions/99262/chaum-pedersen-protocol

	// If the user doesn't exist, return immediately.
	user, ok := users[in.AuthId]
	if !ok {
		slog.WarnContext(ctx, "Not existing", slog.String("user", in.AuthId))
		return &pb.AuthenticationAnswerResponse{}, fmt.Errorf("user doesn't exist")
	}

	// Authenticate the user if the answer to the challenge is correct
	if isVerified := zk.Verify(ctx, big.NewInt(user.Y1), big.NewInt(user.Y2), big.NewInt(in.S), big.NewInt(challenge)); isVerified {
		return &pb.AuthenticationAnswerResponse{SessionId: "ValidSession"}, nil
	}

	// If not verified, return an invalid session.
	return &pb.AuthenticationAnswerResponse{SessionId: "NotValid"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &grpcServer{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
