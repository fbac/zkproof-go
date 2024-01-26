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

// Used to generate X^Y MOD Z.
func GenerateY(x, y, z *big.Int) *big.Int {
	return new(big.Int).Exp(x, y, z)
}

// Used to generate Y1 and Y2 based on a provided password.
func GenerateYPair(userPassword *big.Int) (*big.Int, *big.Int) {
	// g^b
	GB := new(big.Int).Exp(G, B, P)
	Y1 := GenerateY(G, userPassword, P)
	Y2 := GenerateY(GB, userPassword, P)
	return Y1, Y2
}

// Challenge is just a random number [0,1000].
// This would be ideally a configurable setting.
func Challenge() int64 {
	r := rand.Int63n(1000)
	return r
}

// Challenge answer is ans = password + a x challenge MOD prime
func ChallengeAnswer(userPassword, challenge int64) *int64 {
	x := (userPassword + A.Int64()*challenge) % P.Int64()
	return &x
}

func Verify(Y1, Y2, challengeAnswer, challenge *big.Int) bool {
	GA := new(big.Int).Exp(G, A, P)         // g^a
	GB := new(big.Int).Exp(G, B, P)         // g^b
	GAB := new(big.Int).Exp(G, AB, P)       // g^ab
	a1 := GenerateY(G, challengeAnswer, P)  // Calculate g ^ answerToChallenge % q
	a2 := verifyExpMod(GA, Y1, challenge)   // Calculate A^challenge * y1
	b1 := GenerateY(GB, challengeAnswer, P) // Calculate B ^ answerToChallenge % q
	b2 := verifyExpMod(GAB, Y2, challenge)  // Calculate C^challenge * y2

	// In order to be verified: a1 == a2 && b1 == b2
	if a1.Cmp(a2) == 0 && b1.Cmp(b2) == 0 {
		return true
	}

	return false
}

// Verify the following cases
// A^challenge * y1
// C^challenge * y2
func verifyExpMod(N, Y, challenge *big.Int) *big.Int {
	// Note: the following calculations can be executed in only one line.
	// I'm avoiding this to make it more legible.

	// Calculate the exponent.
	expResult := N.Exp(N, challenge, nil)

	// Multiply
	mulResult := expResult.Mul(expResult, Y)

	// Mod and return
	return new(big.Int).Mod(mulResult, P)
}
