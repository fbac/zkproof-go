package zk

import (
	"context"
	"log/slog"
	"math/big"
)

var (
	// Define a prime number
	// This could be random generated
	P = big.NewInt(10009)

	// Data agreed between server <-> client for zkproof verifications.
	// Ideally, the values would be loaded from an encrypted secret / k8s configmap.
	G  = big.NewInt(3)
	A  = big.NewInt(10)
	B  = big.NewInt(13)
	AB = new(big.Int).Mul(A, B)

	// g^a, g^b,g^ab
	GA  = new(big.Int).Exp(G, A, P)
	GB  = new(big.Int).Exp(G, B, P)
	GAB = new(big.Int).Exp(G, AB, P)
)

func GenerateY(g, userPassword, p *big.Int) *big.Int {
	return new(big.Int).Exp(g, userPassword, p)
}

func Verify(ctx context.Context, Y1, Y2, challengeAnswer, challenge *big.Int) bool {
	// g ^ answerToChallenge % q
	a1 := verifyExp(G, challengeAnswer, P)
	slog.InfoContext(ctx, "DEBUG!", slog.Int64("a1", a1.Int64())) // Change to DebugContext

	// A^challenge * y1
	a2 := verifyExpMod(GA, Y1, challenge)
	slog.InfoContext(ctx, "DEBUG!", slog.Int64("a2", a2.Int64())) // Change to DebugContext

	// B ^ answerToChallenge % q
	b1 := verifyExp(GB, challengeAnswer, P)
	slog.InfoContext(ctx, "DEBUG!", slog.Int64("b1", b1.Int64())) // Change to DebugContext

	// C^challenge * y2
	b2 := verifyExpMod(GAB, Y2, challenge)
	slog.InfoContext(ctx, "DEBUG!", slog.Int64("b2", b2.Int64())) // Change to DebugContext

	// In order to be verified: a1 == a2 && b1 == b2
	if a1.Cmp(a2) == 0 && b1.Cmp(b2) == 0 {
		return true
	}

	return false
}

// x ^ y * z
func verifyExp(x, y, z *big.Int) *big.Int {
	return new(big.Int).Exp(x, y, z)
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
	return new(big.Int).Mod(mulResult, P)
}
