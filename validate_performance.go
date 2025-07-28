package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/tjfoc/gmsm/sm4"
)

func main() {
	fmt.Println("=== GMSM Performance Validation ===")
	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Printf("Architecture: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("CPU Cores: %d\n", runtime.NumCPU())
	fmt.Println()

	// Test SM3 Performance
	fmt.Println("=== SM3 Performance Test ===")
	testSM3()
	fmt.Println()

	// Test SM4 Performance
	fmt.Println("=== SM4 Performance Test ===")
	testSM4()
	fmt.Println()

	// Test SM2 Performance
	fmt.Println("=== SM2 Performance Test ===")
	testSM2()
	fmt.Println()

	fmt.Println("=== Performance Validation Complete ===")
}

func testSM3() {
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i % 256)
	}

	// Test convenience function
	start := time.Now()
	hash := sm3.Sum(data)
	elapsed := time.Since(start)
	fmt.Printf("SM3.Sum() - 1KB data: %v (hash: %x...)\n", elapsed, hash[:8])

	// Test throughput
	largeData := make([]byte, 1024*1024) // 1MB
	start = time.Now()
	hash = sm3.Sum(largeData)
	elapsed = time.Since(start)
	throughput := float64(len(largeData)) / float64(elapsed.Microseconds()) * 1000
	fmt.Printf("SM3 Throughput: %.2f MB/s\n", throughput)
}

func testSM4() {
	key := []byte("1234567890abcdef")
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i % 256)
	}

	// Test convenience function
	start := time.Now()
	encrypted, err := sm4.EncryptWithKey(key, data, sm4.ECB)
	if err != nil {
		fmt.Printf("SM4 encryption error: %v\n", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("SM4 EncryptWithKey(ECB) - 1KB data: %v\n", elapsed)

	// Test decryption
	start = time.Now()
	decrypted, err := sm4.DecryptWithKey(key, encrypted, sm4.ECB)
	if err != nil {
		fmt.Printf("SM4 decryption error: %v\n", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("SM4 DecryptWithKey(ECB) - 1KB data: %v\n", elapsed)

	// Verify correctness
	if string(decrypted) != string(data) {
		fmt.Println("SM4 ECB decryption verification failed!")
	} else {
		fmt.Println("SM4 ECB decryption verified!")
	}

	// Test throughput
	largeData := make([]byte, 1024*1024) // 1MB
	start = time.Now()
	encrypted, err = sm4.EncryptWithKey(key, largeData, sm4.ECB)
	if err != nil {
		fmt.Printf("SM4 encryption error: %v\n", err)
		return
	}
	elapsed = time.Since(start)
	throughput := float64(len(largeData)) / float64(elapsed.Microseconds()) * 1000
	fmt.Printf("SM4 Throughput: %.2f MB/s\n", throughput)
}

func testSM2() {
	// Generate key pair
	start := time.Now()
	priv, pub, err := sm2.NewKeyPair()
	if err != nil {
		fmt.Printf("SM2 key generation error: %v\n", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("SM2 Key Generation: %v\n", elapsed)

	// Test signing
	message := []byte("Hello, World!")
	start = time.Now()
	signature, err := sm2.SignData(priv, message)
	if err != nil {
		fmt.Printf("SM2 signing error: %v\n", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("SM2 Signing: %v\n", elapsed)

	// Test verification
	start = time.Now()
	valid := sm2.VerifySignature(pub, message, signature)
	elapsed = time.Since(start)
	fmt.Printf("SM2 Verification: %v (valid: %t)\n", elapsed, valid)

	// Test batch operations
	messages := [][]byte{
		[]byte("Message 1"),
		[]byte("Message 2"),
		[]byte("Message 3"),
	}
	start = time.Now()
	signatures, err := sm2.BatchSign(priv, messages)
	if err != nil {
		fmt.Printf("SM2 batch signing error: %v\n", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("SM2 Batch Signing (3 messages): %v\n", elapsed)

	// Test batch verification
	start = time.Now()
	results, err := sm2.BatchVerify(pub, messages, signatures)
	if err != nil {
		fmt.Printf("SM2 batch verification error: %v\n", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("SM2 Batch Verification (3 messages): %v\n", elapsed)
	fmt.Printf("All batch verifications valid: %t\n", allTrue(results))
}

func allTrue(results []bool) bool {
	for _, r := range results {
		if !r {
			return false
		}
	}
	return true
}