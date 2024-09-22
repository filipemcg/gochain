package gochain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func GeneratePrivavateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func Verify(pub *ecdsa.PublicKey, hash []byte, r, s *big.Int) bool {
	return ecdsa.Verify(pub, hash, r, s)
}

func Sign(priv *ecdsa.PrivateKey, hash []byte) (*big.Int, *big.Int, error) {
	return ecdsa.Sign(rand.Reader, priv, hash)
}
