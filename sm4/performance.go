package sm4

import (
	"errors"
)

// CipherMode represents the different cipher modes supported
type CipherMode int

const (
	ECB CipherMode = iota
	CBC
	CFB
	OFB
)

// EncryptWithKey encrypts data using the provided key and returns the encrypted data
// This is a convenience function that handles key setup and cipher creation automatically
func EncryptWithKey(key, data []byte, mode CipherMode) ([]byte, error) {
	if len(key) != BlockSize {
		return nil, errors.New("SM4: invalid key size")
	}
	
	_, err := NewCipher(key)
	if err != nil {
		return nil, err
	}
	
	switch mode {
	case ECB:
		return Sm4Ecb(key, data, true)
	case CBC:
		return Sm4Cbc(key, data, true)
	case CFB:
		return Sm4CFB(key, data, true)
	case OFB:
		return Sm4OFB(key, data, true)
	default:
		return nil, errors.New("SM4: unsupported cipher mode")
	}
}

// DecryptWithKey decrypts data using the provided key and returns the decrypted data
// This is a convenience function that handles key setup and cipher creation automatically
func DecryptWithKey(key, data []byte, mode CipherMode) ([]byte, error) {
	if len(key) != BlockSize {
		return nil, errors.New("SM4: invalid key size")
	}
	
	switch mode {
	case ECB:
		return Sm4Ecb(key, data, false)
	case CBC:
		return Sm4Cbc(key, data, false)
	case CFB:
		return Sm4CFB(key, data, false)
	case OFB:
		return Sm4OFB(key, data, false)
	default:
		return nil, errors.New("SM4: unsupported cipher mode")
	}
}

