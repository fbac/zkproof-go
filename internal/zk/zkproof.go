package zk

import (
	"math/big"
	"math/rand"
)

var (
	// Define a prime number
	// This could be random generated
	P = big.NewInt(100027)

	// Data agreed between server <-> client for zkproof verifications.
	// Ideally, the values would be loaded from an encrypted secret / k8s configmap.
	G  = big.NewInt(5)
	A  = big.NewInt(10)
	B  = big.NewInt(15)
	AB = new(big.Int).Mul(A, B)
)

func GenerateY(x, y, z *big.Int) *big.Int {
	return new(big.Int).Exp(x, y, z)
}

func GenerateYPair(userPassword *big.Int) (Y1 *big.Int, Y2 *big.Int) {
	// g^b
	GB := new(big.Int).Exp(G, B, P)
	Y1 = GenerateY(G, userPassword, P)
	Y2 = GenerateY(GB, userPassword, P)
	return Y1, Y2
}

func Challenge() *int64 {
	r := rand.Int63n(1000)
	return &r
}

func ChallengeAnswer(userPassword, challenge int64) *int64 {
	x := (userPassword + A.Int64()*challenge) % P.Int64()
	return &x
}

func Verify(Y1, Y2, challengeAnswer, challenge *big.Int) bool {
	// Calculate g^a, g^b,g^ab
	GA := new(big.Int).Exp(G, A, P)
	GB := new(big.Int).Exp(G, B, P)
	GAB := new(big.Int).Exp(G, AB, P)

	// Calculate g ^ answerToChallenge % q
	a1 := verifyExp(G, challengeAnswer, P)
	//slog.InfoContext(ctx, "DEBUG!", slog.Int64("a1", a1.Int64())) // Change to DebugContext

	// Calculate A^challenge * y1
	a2 := verifyExpMod(GA, Y1, challenge)
	//slog.InfoContext(ctx, "DEBUG!", slog.Int64("a2", a2.Int64())) // Change to DebugContext

	// Calculate B ^ answerToChallenge % q
	b1 := verifyExp(GB, challengeAnswer, P)
	//slog.InfoContext(ctx, "DEBUG!", slog.Int64("b1", b1.Int64())) // Change to DebugContext

	// Calculate C^challenge * y2
	b2 := verifyExpMod(GAB, Y2, challenge)
	//slog.InfoContext(ctx, "DEBUG!", slog.Int64("b2", b2.Int64())) // Change to DebugContext

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
