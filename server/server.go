package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"net"
	"os"
	"strings"

	zk "github.com/fbac/zkproof-grpc/internal/zk"
	"github.com/fbac/zkproof-grpc/pkg/check"
	pb "github.com/fbac/zkproof-grpc/protobuf"
	"google.golang.org/grpc"
)

// The AuthServer interface forces us to embed UnimplementedAuthServer.
type grpcServer struct {
	pb.UnimplementedAuthServer
	zk.ZKVerifier
}

// userData holds the exchanged Y's
type userData struct {
	Y1 int64
	Y2 int64
}

var (
	users                = make(map[string]userData)
	challenge      int64 = 0
	hostName, port string
)

// grpcServer has to implement AuthServer.
// this line won't compile if there is not interface compliance.
var _ pb.AuthServer = (*grpcServer)(nil)

func (s *grpcServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Basic sanity check
	// userName has to contain at least one char (not white spaces)
	if !check.IsValidString(in.User) {
		slog.WarnContext(ctx, "Invalid", slog.String("user", in.User))
		return &pb.RegisterResponse{}, fmt.Errorf("userName is invalid")
	}

	// Save the password if the user doesn 't exist.
	// Skip saving password if it's already registered.
	if _, exists := users[in.User]; exists {
		slog.InfoContext(ctx, "User already registered")
		return &pb.RegisterResponse{}, nil
	}

	users[in.User] = userData{
		Y1: in.Y1,
		Y2: in.Y2,
	}
	slog.InfoContext(ctx, "Registered new", slog.String("user", in.User))

	return &pb.RegisterResponse{}, nil
}

func (s *grpcServer) CreateAuthenticationChallenge(ctx context.Context, in *pb.AuthenticationChallengeRequest) (*pb.AuthenticationChallengeResponse, error) {
	userName := strings.TrimSpace(in.User)

	// If the user doesn't exist, don't create a challenge.
	_, ok := users[userName]
	if !ok {
		slog.WarnContext(ctx, "Do not create challenge for non existing", slog.String("user", userName))
		return &pb.AuthenticationChallengeResponse{}, fmt.Errorf("user already exists")
	}

	// Update challenge
	slog.InfoContext(ctx, "Creating challenge for", slog.String("user", userName))
	challenge = s.Challenge()

	// Return the proper challenge for this user.
	return &pb.AuthenticationChallengeResponse{
		AuthId: userName,
		C:      challenge}, nil
}

func (s *grpcServer) VerifyAuthentication(ctx context.Context, in *pb.AuthenticationAnswerRequest) (*pb.AuthenticationAnswerResponse, error) {
	userName := strings.TrimSpace(in.AuthId)

	// Calculations sourced from https://crypto.stackexchange.com/questions/99262/chaum-pedersen-protocol
	// If the user doesn't exist, return immediately.
	user, ok := users[userName]
	if !ok {
		slog.WarnContext(ctx, "Not existing", slog.String("user", userName))
		return &pb.AuthenticationAnswerResponse{}, fmt.Errorf("user doesn't exist")
	}

	// Authenticate the user if the answer to the challenge is correct
	if s.Verify(big.NewInt(user.Y1), big.NewInt(user.Y2), big.NewInt(in.S), big.NewInt(challenge)) {
		slog.InfoContext(ctx, "Authentication verified for", slog.String("user", userName))
		return &pb.AuthenticationAnswerResponse{SessionId: "ValidSession"}, nil
	}

	// If not verified, return an invalid session.
	slog.WarnContext(ctx, "Authentication failed for", slog.String("user", userName))
	return nil, fmt.Errorf("authentication failed for user %s", userName)
}

func main() {
	flag.StringVar(&hostName, "host", "localhost", "connect to hostname")
	flag.StringVar(&port, "port", "50051", "TCP port")
	flag.Parse()

	if !check.IsValidPort(port) {
		log.Printf("port %s is invalid", port)
		os.Exit(1)
	}

	// Hostname basic sanity check
	if !check.IsValidString(hostName) {
		log.Printf("hostName %s is invalid", hostName)
		os.Exit(1)
	}

	// Grab a tcp socket to listen on
	host := fmt.Sprintf("%s:%s", hostName, port)
	l, err := net.Listen("tcp", host)
	if err != nil {
		log.Printf("failed to listen on %s: %v", host, err)
		os.Exit(1)
	}

	// Initialize the grpc server
	s := grpc.NewServer()

	// Register the AuthServer methods to this server
	pb.RegisterAuthServer(s, &grpcServer{
		ZKVerifier: &zk.ZKServer{},
	})

	// Serve on the listener
	log.Printf("gRPC server listening at %v", l.Addr())
	if err := s.Serve(l); err != nil {
		log.Printf("error listening at %v", l.Addr())
		os.Exit(1)
	}
}
