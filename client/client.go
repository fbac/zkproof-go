package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"os"
	"time"

	"github.com/fbac/zkproof-grpc/pkg/check"
	"github.com/fbac/zkproof-grpc/pkg/zk"
	pb "github.com/fbac/zkproof-grpc/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	userPassword             int64
	userName, hostName, port string
)

type zkClient struct {
	pb.AuthClient
	zk.ZKProver
}

func main() {
	flag.StringVar(&hostName, "host", "localhost", "connect to hostname")
	flag.StringVar(&port, "port", "50051", "TCP port")
	flag.StringVar(&userName, "user", "testUser", "Username")
	flag.Int64Var(&userPassword, "password", 1, "Password")
	flag.Parse()

	if !check.IsValidPort(port) {
		log.Printf("port %s is invalid", port)
		os.Exit(1)
	}

	if !check.IsValidString(hostName) || !check.IsValidString(port) || !check.IsValidString(userName) {
		log.Fatalf("invalid input provided: host %s, port %s, user %s", hostName, port, userName)
	}

	// Create the context for this connection, it will carry all the info
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Open the socket served where the server is listening
	host := fmt.Sprintf("%s:%s", hostName, port)
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()
	slog.InfoContext(ctx, "Connected to", slog.String("host", host))

	// Create the grpc client
	zkClient := zkClient{
		pb.NewAuthClient(conn),
		&zk.ZKClient{},
	}

	// We generate Y1 and Y2 calculated with the provided userPassword
	Y1, Y2 := zkClient.GenerateYPair(big.NewInt(userPassword))

	// Call the register endpoint to register the username, Y1 and Y2
	// This will register the user if it doesn't exist.
	_, err = zkClient.Register(ctx, &pb.RegisterRequest{
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
	resp, err := zkClient.CreateAuthenticationChallenge(ctx, &pb.AuthenticationChallengeRequest{
		User: userName,
	})
	if err != nil {
		slog.ErrorContext(ctx, "Error creating authentication challenge", slog.String("message", err.Error()))
		os.Exit(1)
	}
	slog.InfoContext(ctx, "Server challenged user with", slog.Int64("challenge", resp.C))

	// On this step we calculate the challenge answer based on Chaum-Pedersen
	// answerToChallenge is calculated from the password and server's challenge
	answerToChallenge := zkClient.ChallengeAnswer(userPassword, resp.C)

	// Invoke the verify auth endpoint for this user and answerToChallenge
	verify, err := zkClient.VerifyAuthentication(ctx, &pb.AuthenticationAnswerRequest{
		AuthId: userName,
		S:      answerToChallenge})
	if err != nil {
		slog.ErrorContext(ctx, "Error", slog.String("message", err.Error()))
		os.Exit(1)
	}

	slog.InfoContext(ctx, "Authentication succesful", slog.String("sessionId", verify.SessionId))
	os.Exit(0)
}
