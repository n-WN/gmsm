package benchmarks

import (
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/tjfoc/gmsm/sm4"
)

// BenchmarkSM2KeyGeneration benchmarks SM2 key generation
func BenchmarkSM2KeyGeneration(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := sm2.GenerateKey(rand.Reader)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSM2Sign benchmarks SM2 signing
func BenchmarkSM2Sign(b *testing.B) {
	msg := []byte("benchmark test message for SM2 signing")
	priv, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := priv.Sign(rand.Reader, msg, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSM2Verify benchmarks SM2 verification
func BenchmarkSM2Verify(b *testing.B) {
	msg := []byte("benchmark test message for SM2 verification")
	priv, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		b.Fatal(err)
	}
	signature, err := priv.Sign(rand.Reader, msg, nil)
	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if !priv.PublicKey.Verify(msg, signature) {
			b.Fatal("verification failed")
		}
	}
}

// BenchmarkSM2Encrypt benchmarks SM2 encryption
func BenchmarkSM2Encrypt(b *testing.B) {
	data := []byte("benchmark test data for encryption")
	priv, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		b.Fatal(err)
	}
	pub := &priv.PublicKey
	
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := sm2.Encrypt(pub, data, rand.Reader, sm2.C1C3C2)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSM2Decrypt benchmarks SM2 decryption
func BenchmarkSM2Decrypt(b *testing.B) {
	data := []byte("benchmark test data for decryption")
	priv, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		b.Fatal(err)
	}
	pub := &priv.PublicKey
	encrypted, err := sm2.Encrypt(pub, data, rand.Reader, sm2.C1C3C2)
	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := sm2.Decrypt(priv, encrypted, sm2.C1C3C2)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSM3 benchmarks SM3 hashing
func BenchmarkSM3(b *testing.B) {
	data := []byte("benchmark test data for SM3 hashing")
	
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = sm3.Sum(data)
	}
}

// BenchmarkSM3LargeData benchmarks SM3 hashing with large data
func BenchmarkSM3LargeData(b *testing.B) {
	data := make([]byte, 1024*1024) // 1MB of data
	_, err := rand.Read(data)
	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = sm3.Sum(data)
	}
}

// BenchmarkSM4Encrypt benchmarks SM4 encryption
func BenchmarkSM4Encrypt(b *testing.B) {
	key := []byte("1234567890abcdef")
	data := make([]byte, 1024)
	_, err := rand.Read(data)
	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := sm4.EncryptWithKey(key, data, sm4.ECB)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSM4Decrypt benchmarks SM4 decryption
func BenchmarkSM4Decrypt(b *testing.B) {
	key := []byte("1234567890abcdef")
	data := make([]byte, 1024)
	_, err := rand.Read(data)
	if err != nil {
		b.Fatal(err)
	}
	encrypted, err := sm4.EncryptWithKey(key, data, sm4.ECB)
	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := sm4.DecryptWithKey(key, encrypted, sm4.ECB)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSM4CBC benchmarks SM4 CBC mode
func BenchmarkSM4CBC(b *testing.B) {
	key := []byte("1234567890abcdef")
	data := make([]byte, 1024)
	_, err := rand.Read(data)
	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := sm4.EncryptWithKey(key, data, sm4.CBC)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// RunAllBenchmarks runs all benchmarks and prints results
func RunAllBenchmarks() {
	fmt.Println("Running all SM2/SM3/SM4 benchmarks...")
	
	benchmarks := []testing.InternalBenchmark{
		{Name: "SM2KeyGeneration", F: BenchmarkSM2KeyGeneration},
		{Name: "SM2Sign", F: BenchmarkSM2Sign},
		{Name: "SM2Verify", F: BenchmarkSM2Verify},
		{Name: "SM2Encrypt", F: BenchmarkSM2Encrypt},
		{Name: "SM2Decrypt", F: BenchmarkSM2Decrypt},
		{Name: "SM3", F: BenchmarkSM3},
		{Name: "SM3LargeData", F: BenchmarkSM3LargeData},
		{Name: "SM4Encrypt", F: BenchmarkSM4Encrypt},
		{Name: "SM4Decrypt", F: BenchmarkSM4Decrypt},
		{Name: "SM4CBC", F: BenchmarkSM4CBC},
	}
	
	for _, bm := range benchmarks {
		result := testing.Benchmark(bm.F)
		fmt.Printf("%-20s: %10d ns/op %10d B/op %6d allocs/op\n",
			bm.Name, result.NsPerOp(), result.AllocedBytesPerOp(), result.AllocsPerOp())
	}
}

// MemoryProfile runs memory profiling for all algorithms
func MemoryProfile() {
	fmt.Println("\nMemory profiling...")
	
	// Test with different data sizes
	sizes := []int{64, 1024, 10240, 102400}
	
	for _, size := range sizes {
		fmt.Printf("\nData size: %d bytes\n", size)
		
		// SM3
		data := make([]byte, size)
		testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = sm3.Sum(data)
			}
		})
		
		// SM4
		key := []byte("1234567890abcdef")
		testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = sm4.EncryptWithKey(key, data, sm4.ECB)
			}
		})
	}
}

// ComparePerformance compares performance with other algorithms
func ComparePerformance() {
	fmt.Println("\nPerformance comparison...")
	
	// This would typically compare with other algorithms
	// For now, we'll just run our benchmarks
	RunAllBenchmarks()
}

// ValidateCorrectness validates that all algorithms work correctly
func ValidateCorrectness() error {
	fmt.Println("\nValidating correctness...")
	
	// Test SM2
	priv, pub, err := sm2.NewKeyPair()
	if err != nil {
		return fmt.Errorf("SM2 key generation failed: %w", err)
	}
	
	data := []byte("test data")
	signature, err := sm2.SignData(priv, data)
	if err != nil {
		return fmt.Errorf("SM2 signing failed: %w", err)
	}
	
	if !sm2.VerifySignature(pub, data, signature) {
		return fmt.Errorf("SM2 verification failed")
	}
	
	encrypted, err := sm2.EncryptData(pub, data)
	if err != nil {
		return fmt.Errorf("SM2 encryption failed: %w", err)
	}
	
	decrypted, err := sm2.DecryptData(priv, encrypted)
	if err != nil {
		return fmt.Errorf("SM2 decryption failed: %w", err)
	}
	
	if string(decrypted) != string(data) {
		return fmt.Errorf("SM2 encryption/decryption mismatch")
	}
	
	// Test SM3
	hash := sm3.Sum(data)
	if len(hash) != 32 {
		return fmt.Errorf("SM3 hash length incorrect")
	}
	
	// Test SM4
	key := []byte("1234567890abcdef")
	encrypted, err = sm4.EncryptWithKey(key, data, sm4.ECB)
	if err != nil {
		return fmt.Errorf("SM4 encryption failed: %w", err)
	}
	
	decrypted, err = sm4.DecryptWithKey(key, encrypted, sm4.ECB)
	if err != nil {
		return fmt.Errorf("SM4 decryption failed: %w", err)
	}
	
	if string(decrypted) != string(data) {
		return fmt.Errorf("SM4 encryption/decryption mismatch")
	}
	
	fmt.Println("All algorithms validated successfully")
	return nil
}

// RunFullTestSuite runs the complete test suite
func RunFullTestSuite() {
	fmt.Println("Running full test suite...")
	
	if err := ValidateCorrectness(); err != nil {
		fmt.Printf("Validation failed: %v\n", err)
		return
	}
	
	RunAllBenchmarks()
	MemoryProfile()
	ComparePerformance()
	
	fmt.Println("\nTest suite completed successfully")
}

func init() {
	// Pre-warm pools to avoid cold start effects
	for i := 0; i < 100; i++ {
		_, _ = sm2.GenerateKey(rand.Reader)
		_ = sm3.Sum([]byte("warmup"))
		_, _ = sm4.EncryptWithKey([]byte("1234567890abcdef"), []byte("warmup"), sm4.ECB)
	}
}

// Example usage:
//
// import "github.com/tjfoc/gmsm/benchmarks"
//
// func main() {
//     benchmarks.RunFullTestSuite()
// }