package main

import (
	"crypto/rand"
	"testing"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/tjfoc/gmsm/sm4"
)

// 基准测试 - 当前性能基线
func BenchmarkSM2KeyGen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = sm2.GenerateKey(rand.Reader)
	}
}

func BenchmarkSM2Sign(b *testing.B) {
	priv, _ := sm2.GenerateKey(rand.Reader)
	msg := []byte("test message for signing")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = priv.Sign(rand.Reader, msg, nil)
	}
}

func BenchmarkSM2Verify(b *testing.B) {
	priv, _ := sm2.GenerateKey(rand.Reader)
	msg := []byte("test message for signing")
	sig, _ := priv.Sign(rand.Reader, msg, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = priv.PublicKey.Verify(msg, sig)
	}
}

func BenchmarkSM2Encrypt(b *testing.B) {
	priv, _ := sm2.GenerateKey(rand.Reader)
	msg := []byte("test data for encryption")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sm2.Encrypt(&priv.PublicKey, msg, rand.Reader, sm2.C1C3C2)
	}
}

func BenchmarkSM2Decrypt(b *testing.B) {
	priv, _ := sm2.GenerateKey(rand.Reader)
	msg := []byte("test data for encryption")
	encrypted, _ := sm2.Encrypt(&priv.PublicKey, msg, rand.Reader, sm2.C1C3C2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sm2.Decrypt(priv, encrypted, sm2.C1C3C2)
	}
}

func BenchmarkSM3Hash(b *testing.B) {
	data := []byte("test data for hashing")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sm3.Sum(data)
	}
}

func BenchmarkSM4Encrypt(b *testing.B) {
	key := make([]byte, 16)
	rand.Read(key)
	data := make([]byte, 64)
	rand.Read(data)
	
	cipher, _ := sm4.NewCipher(key)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out := make([]byte, 64)
		cipher.Encrypt(out, data)
	}
}

func BenchmarkSM4Decrypt(b *testing.B) {
	key := make([]byte, 16)
	rand.Read(key)
	data := make([]byte, 64)
	rand.Read(data)
	
	cipher, _ := sm4.NewCipher(key)
	encrypted := make([]byte, 64)
	cipher.Encrypt(encrypted, data)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out := make([]byte, 64)
		cipher.Decrypt(out, encrypted)
	}
}