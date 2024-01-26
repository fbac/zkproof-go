package zk

import (
	"math/big"
	"testing"
)

// This functional test should success, as we provide the correct password.
func TestZKProofSuccess(t *testing.T) {
	// The client generate a new random password to register itself.
	var userPassword int64 = 10

	// With the password, it generates a pair Y1, Y2.
	Y1, Y2 := GenerateYPair(big.NewInt(userPassword))

	// The server answer with a random challenge (int64).
	challenge := Challenge()

	// The answer to the challenge is based on the formula:
	// ans = password + a x challenge MOD prime.
	// The prime number is set between both parties at the beginning.
	challengeAnswer := ChallengeAnswer(userPassword, challenge)

	// Verify verifies that the two ecuations are equal.
	if !Verify(Y1, Y2, big.NewInt(*challengeAnswer), big.NewInt(challenge)) {
		t.Fail()
	}
}

// This functional test should fail, as we provide an incorrect password.
func TestZKProofFailure(t *testing.T) {
	// The client generate a new random password to register itself.
	var userPassword int64 = 10

	// With the password, it generates a pair Y1, Y2.
	Y1, Y2 := GenerateYPair(big.NewInt(userPassword))

	// The server answer with a random challenge (int64).
	challenge := Challenge()

	// The answer to the challenge is based on the formula:
	// ans = password + a x challenge MOD prime.
	// The prime number is set between both parties at the beginning.
	// On this case we provide a *WRONG* password.
	challengeAnswer := ChallengeAnswer(11, challenge)

	// Verify verifies that the two ecuations are equal.
	if Verify(Y1, Y2, big.NewInt(*challengeAnswer), big.NewInt(challenge)) {
		t.Fail()
	}
}
