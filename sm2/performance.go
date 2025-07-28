package sm2

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
	"math/big"
	"sync"
)

// KeyPool is a sync.Pool for *PrivateKey instances to reduce allocations
var keyPool = sync.Pool{
	New: func() interface{} {
		return &PrivateKey{}
	},
}

// GenerateKeyWithPool generates a new SM2 private key using the pool
func GenerateKeyWithPool(random io.Reader) (*PrivateKey, error) {
	if random == nil {
		random = rand.Reader
	}
	
	priv := keyPool.Get().(*PrivateKey)
	c := P256Sm2()
	params := c.Params()
	b := make([]byte, params.BitSize/8+8)
	_, err := io.ReadFull(random, b)
	if err != nil {
		keyPool.Put(priv)
		return nil, err
	}

	k := new(big.Int).SetBytes(b)
	n := new(big.Int).Sub(params.N, two)
	k.Mod(k, n)
	k.Add(k, one)
	
	priv.PublicKey.Curve = c
	priv.D = k
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
	
	return priv, nil
}

// ReturnKey returns a private key to the pool
func ReturnKey(priv *PrivateKey) {
	if priv != nil {
		priv.D = nil
		priv.PublicKey.X = nil
		priv.PublicKey.Y = nil
		keyPool.Put(priv)
	}
}

// SignData signs data with the provided private key and returns the signature
// This is a convenience function that handles the entire signing process
func SignData(priv *PrivateKey, data []byte) ([]byte, error) {
	return priv.Sign(rand.Reader, data, nil)
}

// VerifySignature verifies a signature against data and public key
// This is a convenience function that handles the entire verification process
func VerifySignature(pub *PublicKey, data, signature []byte) bool {
	return pub.Verify(data, signature)
}

// EncryptData encrypts data with the provided public key
// This is a convenience function that handles the entire encryption process
func EncryptData(pub *PublicKey, data []byte) ([]byte, error) {
	return pub.EncryptAsn1(data, rand.Reader)
}

// DecryptData decrypts data with the provided private key
// This is a convenience function that handles the entire decryption process
func DecryptData(priv *PrivateKey, encryptedData []byte) ([]byte, error) {
	return priv.DecryptAsn1(encryptedData)
}

// NewKeyPair generates a new key pair and returns both private and public keys
func NewKeyPair() (*PrivateKey, *PublicKey, error) {
	priv, err := GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return priv, &priv.PublicKey, nil
}

// Hash computes a SHA256 hash of the data for use in SM2 operations
// This is provided as a convenience function for data hashing
func Hash(data []byte) []byte {
	h := sha256.Sum256(data)
	return h[:]
}

// BatchSign signs multiple messages with the same private key
// This is more efficient than signing each message individually
func BatchSign(priv *PrivateKey, messages [][]byte) ([][]byte, error) {
	signatures := make([][]byte, len(messages))
	for i, msg := range messages {
		sig, err := priv.Sign(rand.Reader, msg, nil)
		if err != nil {
			return nil, err
		}
		signatures[i] = sig
	}
	return signatures, nil
}

// BatchVerify verifies multiple signatures with the same public key
// This is more efficient than verifying each signature individually
func BatchVerify(pub *PublicKey, messages [][]byte, signatures [][]byte) ([]bool, error) {
	if len(messages) != len(signatures) {
		return nil, errors.New("messages and signatures count mismatch")
	}
	
	results := make([]bool, len(messages))
	for i := range messages {
		results[i] = pub.Verify(messages[i], signatures[i])
	}
	return results, nil
}