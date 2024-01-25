package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"math/rand"
	"net"

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

// Users is the users database
var users = make(map[string]UserData)

// Data agreed between server <-> client for zkproof verifications.
// Ideally, this would be moved to internal/zkproof/const.go.
// And the values would be loaded from an encrypted secret.
var (
	q               = big.NewInt(10009)
	g               = big.NewInt(3)
	a               = big.NewInt(10)
	b               = big.NewInt(13)
	ab              = new(big.Int).Mul(a, b)
	A               = new(big.Int).Exp(g, a, q)
	B               = new(big.Int).Exp(g, b, q)
	C               = new(big.Int).Exp(g, ab, q)
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

	// g ^ answerToChallenge % q
	a1 := verifyExp(*g, *big.NewInt(in.S), *q)
	slog.InfoContext(ctx, "DEBUG!", slog.Int64("a1", a1.Int64())) // Change to DebugContext

	// A^challenge * y1
	a2 := verifyExpMod(A, big.NewInt(user.Y1), big.NewInt(challenge))
	slog.InfoContext(ctx, "DEBUG!", slog.Int64("a2", a2.Int64())) // Change to DebugContext

	// B ^ answerToChallenge % q
	b1 := verifyExp(*B, *big.NewInt(in.S), *q)
	slog.InfoContext(ctx, "DEBUG!", slog.Int64("b1", b1.Int64())) // Change to DebugContext

	// C^challenge * y2
	b2 := verifyExpMod(C, big.NewInt(user.Y2), big.NewInt(challenge))
	slog.InfoContext(ctx, "DEBUG!", slog.Int64("b2", b2.Int64())) // Change to DebugContext

	// In order to be verified: a1 == a2 && b1 == b2
	if a1.Cmp(a2) == 0 && b1.Cmp(b2) == 0 {
		slog.InfoContext(ctx, "Session verified for", slog.String("user", in.AuthId))
		return &pb.AuthenticationAnswerResponse{SessionId: "ValidSession"}, nil
	}

	// If not verified, return an invalid session.
	return &pb.AuthenticationAnswerResponse{SessionId: "NotValid"}, nil
}

// x ^ y * z
func verifyExp(x, y, z big.Int) *big.Int {
	return new(big.Int).Exp(&x, &y, &z)
}

// Verify the following cases
// A^challenge * y1
// C^challenge * y2
func verifyExpMod(N, Y, challenge *big.Int) *big.Int {
	// Calculate the exponent.
	expResult := N.Exp(N, challenge, nil)

	// Multiply
	mulResult := expResult.Mul(expResult, Y)

	// Mod and return
	return new(big.Int).Mod(mulResult, q)
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
