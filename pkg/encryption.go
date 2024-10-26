package gochain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

// Example using this package
// func Demo() {
//  // Generate ECDSA keys
//  privateKey, err := GeneratePrivavateKey()
//  if err != nil {
//     fmt.Println("Error generating ECDSA keys:", err)
//     return
//  }
//  publicKey := &privateKey.PublicKey
//
//  // Sign a message
//  message := "Hello, Bitcoin!"
//  hash := sha256.Sum256([]byte(message))
//
//  r, s, err := Sign(privateKey, hash[:])
//  if err != nil {
//     fmt.Println("Error signing message:", err)
//     return
//  }
//  fmt.Printf("Signature: (r: %s, s: %s)\n", r.String(), s.String())
//
// // Verify the signature
//   valid := Verify(publicKey, hash[:], r, s)
//   fmt.Printf("Signature valid: %v\n", valid)
// }

func GeneratePrivavateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func Verify(pub *ecdsa.PublicKey, hash []byte, r, s *big.Int) bool {
	return ecdsa.Verify(pub, hash, r, s)
}

func Sign(priv *ecdsa.PrivateKey, hash []byte) (*big.Int, *big.Int, error) {
	return ecdsa.Sign(rand.Reader, priv, hash)
}
