package examples

import (
	"fmt"
	"log"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/tjfoc/gmsm/sm4"
)

// HighPerformanceUsage demonstrates optimized usage patterns for the GMSM library
func HighPerformanceUsage() {
	fmt.Println("=== High-Performance GMSM Usage Examples ===")

	// Example 1: Efficient SM3 hashing with pooling
	fmt.Println("1. Efficient SM3 Hashing:")
	data := []byte("Hello, this is test data for SM3 hashing")
	hash := sm3.Sum(data)
	fmt.Printf("   SM3 Hash: %x\n", hash)
	fmt.Printf("   Hash length: %d bytes\n\n", len(hash))

	// Example 2: High-performance SM4 encryption
	fmt.Println("2. High-Performance SM4 Encryption:")
	key := []byte("1234567890abcdef")
	plaintext := []byte("This is confidential data")
	
	// Encrypt using convenience function
	encrypted, err := sm4.EncryptWithKey(key, plaintext, sm4.CBC)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("   Original: %s\n", string(plaintext))
	fmt.Printf("   Encrypted: %x\n", encrypted)
	fmt.Printf("   Encrypted length: %d bytes\n", len(encrypted))
	
	// Decrypt
	decrypted, err := sm4.DecryptWithKey(key, encrypted, sm4.CBC)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("   Decrypted: %s\n\n", string(decrypted))

	// Example 3: SM2 key operations with optimized API
	fmt.Println("3. SM2 Key Operations:")
	priv, pub, err := sm2.NewKeyPair()
	if err != nil {
		log.Fatal(err)
	}
	
	message := []byte("Important message to sign")
	
	// Sign with convenience function
	signature, err := sm2.SignData(priv, message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("   Message: %s\n", string(message))
	fmt.Printf("   Signature: %x\n", signature)
	fmt.Printf("   Signature length: %d bytes\n", len(signature))
	
	// Verify with convenience function
	valid := sm2.VerifySignature(pub, message, signature)
	fmt.Printf("   Signature valid: %t\n\n", valid)

	// Example 4: SM2 encryption/decryption
	fmt.Println("4. SM2 Encryption/Decryption:")
	confidential := []byte("This is highly confidential data")
	
	encryptedSM2, err := sm2.EncryptData(pub, confidential)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("   Original: %s\n", string(confidential))
	fmt.Printf("   Encrypted: %x\n", encryptedSM2)
	fmt.Printf("   Encrypted length: %d bytes\n", len(encryptedSM2))
	
	decryptedSM2, err := sm2.DecryptData(priv, encryptedSM2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("   Decrypted: %s\n\n", string(decryptedSM2))

	// Example 5: Batch operations for better performance
	fmt.Println("5. Batch Operations:")
	messages := [][]byte{
		[]byte("Message 1"),
		[]byte("Message 2"),
		[]byte("Message 3"),
	}
	
	// Batch signing
	signatures, err := sm2.BatchSign(priv, messages)
	if err != nil {
		log.Fatal(err)
	}
	
	// Batch verification
	results, err := sm2.BatchVerify(pub, messages, signatures)
	if err != nil {
		log.Fatal(err)
	}
	
	for i, valid := range results {
		fmt.Printf("   Message %d: %s - Valid: %t\n", i+1, string(messages[i]), valid)
	}
	
	// Example 6: Memory-efficient usage patterns
	fmt.Println("\n6. Memory-Efficient Patterns:")
	
	// Large data processing with SM3
	largeData := make([]byte, 1024*1024) // 1MB
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}
	
	hashLarge := sm3.Sum(largeData)
	fmt.Printf("   Large data hash: %x\n", hashLarge)
	fmt.Printf("   Large data size: %d bytes\n", len(largeData))
	
	// Large data encryption with SM4
	largeEncrypted, err := sm4.EncryptWithKey(key, largeData, sm4.ECB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("   Large encrypted size: %d bytes\n", len(largeEncrypted))
	
	largeDecrypted, err := sm4.DecryptWithKey(key, largeEncrypted, sm4.ECB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("   Large decryption successful: %t\n\n", string(largeDecrypted) == string(largeData))

	fmt.Println("=== All examples completed successfully ===")
}

// PerformanceTips demonstrates performance optimization techniques
func PerformanceTips() {
	fmt.Println("=== Performance Optimization Tips ===")
	
	fmt.Println("1. Use convenience functions for common operations:")
	fmt.Println("   - sm3.Sum(data) instead of creating hasher manually")
	fmt.Println("   - sm4.EncryptWithKey() instead of manual cipher setup")
	fmt.Println("   - sm2.SignData() and sm2.VerifySignature()")
	
	fmt.Println("\n2. Use batch operations when processing multiple items:")
	fmt.Println("   - sm2.BatchSign() for multiple signatures")
	fmt.Println("   - sm2.BatchVerify() for multiple verifications")
	
	fmt.Println("\n3. Choose appropriate cipher modes:")
	fmt.Println("   - ECB: Fastest, but less secure for large data")
	fmt.Println("   - CBC: Good balance of security and performance")
	fmt.Println("   - CFB/OFB: Streaming modes for large data")
	
	fmt.Println("\n4. Memory management:")
	fmt.Println("   - Use built-in pooling for better memory efficiency")
	fmt.Println("   - Avoid creating new instances for each operation")
	fmt.Println("   - Reuse buffers when possible")
	
	fmt.Println("\n5. Data size considerations:")
	fmt.Println("   - SM2 is best for small data (< 256 bytes)")
	fmt.Println("   - SM4 is best for large data encryption")
	fmt.Println("   - SM3 is efficient for any data size")
}

// SecurityBestPractices demonstrates security best practices
func SecurityBestPractices() {
	fmt.Println("=== Security Best Practices ===")
	
	fmt.Println("1. Always use cryptographically secure random:")
	fmt.Println("   - Use crypto/rand instead of math/rand")
	
	fmt.Println("\n2. Secure key management:")
	fmt.Println("   - Never hardcode keys in source code")
	fmt.Println("   - Use secure key storage (HSM, key vault)")
	fmt.Println("   - Rotate keys regularly")
	
	fmt.Println("\n3. Algorithm selection:")
	fmt.Println("   - Use SM2 for digital signatures and key exchange")
	fmt.Println("   - Use SM3 for hashing and integrity checking")
	fmt.Println("   - Use SM4 for symmetric encryption")
	
	fmt.Println("\n4. Data handling:")
	fmt.Println("   - Clear sensitive data from memory when done")
	fmt.Println("   - Use authenticated encryption when possible")
	fmt.Println("   - Validate all inputs")
	
	fmt.Println("\n5. Compliance:")
	fmt.Println("   - Follow GM/T standards")
	fmt.Println("   - Ensure regulatory compliance")
	fmt.Println("   - Regular security audits")
}

// IntegrationExample shows how to integrate with existing systems
func IntegrationExample() {
	fmt.Println("=== Integration Examples ===")
	
	// HTTP server integration
	fmt.Println("1. HTTP Server Integration:")
	fmt.Println(`
   // Example HTTP handler using SM2 for API authentication
   func handleSecureAPI(w http.ResponseWriter, r *http.Request) {
       // Verify signature in request header
       signature := r.Header.Get("X-Signature")
       data, _ := io.ReadAll(r.Body)
       
       if !sm2.VerifySignature(publicKey, data, []byte(signature)) {
           http.Error(w, "Invalid signature", http.StatusUnauthorized)
           return
       }
       
       // Process request...
   }`)
	
	// File encryption example
	fmt.Println("2. File Encryption:")
	fmt.Println(`
   // Encrypt a file using SM4
   func encryptFile(filename string, key []byte) error {
       data, err := os.ReadFile(filename)
       if err != nil {
           return err
       }
       
       encrypted, err := sm4.EncryptWithKey(key, data, sm4.CBC)
       if err != nil {
           return err
       }
       
       return os.WriteFile(filename+".enc", encrypted, 0600)
   }`)
	
	// Database integration
	fmt.Println("3. Database Integration:")
	fmt.Println(`
   // Store hashed passwords using SM3
   func hashPassword(password string) string {
       hash := sm3.Sum([]byte(password))
       return hex.EncodeToString(hash[:])
   }`)
}

func main() {
	HighPerformanceUsage()
	PerformanceTips()
	SecurityBestPractices()
	IntegrationExample()
}