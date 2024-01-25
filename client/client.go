package main

import (
	"context"
	"log"
	"log/slog"
	"math/big"
	"os"
	"time"

	zk "github.com/fbac/zkproof-grpc/internal/zk"
	pb "github.com/fbac/zkproof-grpc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// userPassword is hardcoded as this is just a PoC
	// In prod the user has to set this via flag/envvar
	userPassword = big.NewInt(5)
	userName     = "TestUser"

	// We need Y1 and Y2 to prove y1 = gx^1 and y2 = hx^2
	Y1 = zk.GenerateY(zk.G, userPassword, zk.P)
	Y2 = zk.GenerateY(zk.GB, userPassword, zk.P)
)

func main() {
	// Open the socket served where the server is listening
	// Ideally this is configurable via flag using cobra
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()

	// Create the grpc client
	c := pb.NewAuthClient(conn)

	// Create the context for this connection, it will carry all the info
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the register endpoint to register the username, Y1 and Y2
	// This will register the user if it doesn't exist.
	_, err = c.Register(ctx, &pb.RegisterRequest{
		User: userName,
		Y1:   Y1.Int64(),
		Y2:   Y2.Int64(),
	})
	if err != nil {
		slog.ErrorContext(ctx, "Error registering user", slog.String("message", err.Error()))
		os.Exit(1)
	}
	slog.InfoContext(ctx, "Registered user with", slog.String("username", userName))

	// Request an authentication challenge from the server for this user.
	resp, err := c.CreateAuthenticationChallenge(ctx, &pb.AuthenticationChallengeRequest{
		User: userName,
	})
	if err != nil {
		slog.ErrorContext(ctx, "Error creating authentication challenge", slog.String("message", err.Error()))
		os.Exit(1)
	}
	slog.InfoContext(ctx, "Server answered with", slog.Int64("challenge", resp.C))

	// On this step we calculate the challenge answer based on Chaum-Pedersen
	slog.InfoContext(ctx, "Calculating answer to challenge")
	answerToChallenge := (userPassword.Int64() + zk.A.Int64()*resp.C) % zk.P.Int64()

	// Invoke the verify auth endpoint for this user and answerToChallenge
	verify, err := c.VerifyAuthentication(ctx, &pb.AuthenticationAnswerRequest{
		AuthId: userName,
		S:      answerToChallenge})
	if err != nil {
		slog.ErrorContext(ctx, "Error verifying challenge", slog.String("message", err.Error()))
		os.Exit(1)
	}

	slog.InfoContext(ctx, "Authentication succesful", slog.String("sessionId", verify.SessionId))
}
