package zk

import (
	"math/big"
	"math/rand"
)

/*
Calculations and data sourced from https://crypto.stackexchange.com/questions/99262/chaum-pedersen-protocol
*/
type ZKClient struct{}

type ZKProver interface {
	GenerateYPair(*big.Int) (*big.Int, *big.Int)
	ChallengeAnswer(int64, int64) int64
}

type ZKServer struct{}

type ZKVerifier interface {
	Challenge() int64
	Verify(*big.Int, *big.Int, *big.Int, *big.Int) bool
}

var (
	// Ideally the following variables would be stored securely and not hardcoded.
	// Some ideas:
	// Negotiated in a similar fashion as TLS Cipher suites.
	// Stored as ciphered secrets or configmap keys.
	// G, H, P, Q should be randomly selected and validated?

	// Define a prime number shared server <-> client
	P = big.NewInt(10003)

	// Data agreed between server <-> client for zkproof verifications.
	G  = big.NewInt(5)
	A  = big.NewInt(10)
	B  = big.NewInt(15)
	AB = new(big.Int).Mul(A, B)
)

// Used to generate Y1 and Y2 based on a provided password.
func (s *ZKClient) GenerateYPair(userPassword *big.Int) (*big.Int, *big.Int) {
	// g^b
	GB := new(big.Int).Exp(G, B, P)
	Y1 := generateExp(G, userPassword, P)
	Y2 := generateExp(GB, userPassword, P)
	return Y1, Y2
}

// Used to generate X^Y MOD Z.
func generateExp(x, y, z *big.Int) *big.Int {
	return new(big.Int).Exp(x, y, z)
}

// Challenge answer is ans = password + a x challenge MOD prime
func (s *ZKClient) ChallengeAnswer(userPassword, challenge int64) int64 {
	x := (userPassword + A.Int64()*challenge) % P.Int64()
	return x
}

/* ZK Server functions */

// Challenge is just a random number [0,1000].
// This would be ideally a configurable setting.
func (s *ZKServer) Challenge() int64 {
	r := rand.Int63n(1000)
	return r
}

func (s *ZKServer) Verify(Y1, Y2, challengeAnswer, challenge *big.Int) bool {
	GA := generateExp(G, A, P)                // g^a
	GB := generateExp(G, B, P)                // g^b
	GAB := generateExp(G, AB, P)              // g^ab
	a1 := generateExp(G, challengeAnswer, P)  // Calculate g ^ answerToChallenge % q
	a2 := verifyExpMod(GA, Y1, challenge)     // Calculate A^challenge * y1
	b1 := generateExp(GB, challengeAnswer, P) // Calculate B ^ answerToChallenge % q
	b2 := verifyExpMod(GAB, Y2, challenge)    // Calculate C^challenge * y``

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
	// Avoiding an oneliner here to make the code more legible.

	// Calculate the exponent.
	expResult := N.Exp(N, challenge, nil)

	// Multiply
	mulResult := expResult.Mul(expResult, Y)

	// Mod and return
	return new(big.Int).Mod(mulResult, P)
}
