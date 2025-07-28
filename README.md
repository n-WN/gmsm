# gmsm

A Go implementation of Chinese National Cryptographic Standards (SM2, SM3, SM4)

[![Build Status](https://dev.azure.com/n-WN/gmsm/_apis/build/status/n-WN.gmsm?branchName=master)](https://dev.azure.com/n-WN/gmsm/_build/latest?definitionId=1&branchName=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/n-WN/gmsm)](https://goreportcard.com/report/github.com/n-WN/gmsm)
[![GoDoc](https://godoc.org/github.com/n-WN/gmsm?status.svg)](https://godoc.org/github.com/n-WN/gmsm)

## Overview

This library provides Go implementations of the Chinese National Cryptographic Standards:
- **SM2**: Elliptic curve digital signature algorithm and public key encryption
- **SM3**: Cryptographic hash function
- **SM4**: Block cipher algorithm

The library is designed to be compatible with Go's standard crypto interfaces and follows Go idioms for ease of use.

## Features

### SM2 (Digital Signature and Public Key Encryption)
- Key generation, signing, and verification
- PEM file support (encrypted and unencrypted)
- Certificate generation and parsing (compatible with RSA and ECDSA interfaces)
- Certificate chain operations
- Implements `crypto.Signer` interface

### SM3 (Hash Function)
- Basic SM3 sum operations
- Implements `hash.Hash` interface

### SM4 (Block Cipher)
- Key generation, encryption, and decryption
- Implements `cipher.Block` interface
- PEM file support (encrypted and unencrypted)

## Installation

```bash
go get github.com/n-WN/gmsm
```

## Quick Start

### SM2 Example

```go
package main

import (
    "crypto/rand"
    "fmt"
    "github.com/tjfoc/gmsm/sm2"
)

func main() {
    // Generate key pair
    priv, err := sm2.GenerateKey(rand.Reader)
    if err != nil {
        panic(err)
    }
    
    // Sign message
    message := []byte("hello world")
    signature, err := priv.Sign(rand.Reader, message, nil)
    if err != nil {
        panic(err)
    }
    
    // Verify signature
    valid := priv.Public().(*sm2.PublicKey).Verify(message, signature)
    fmt.Printf("Signature valid: %v\n", valid)
}
```

### SM3 Example

```go
package main

import (
    "fmt"
    "github.com/tjfoc/gmsm/sm3"
)

func main() {
    data := []byte("hello world")
    hash := sm3.Sum(data)
    fmt.Printf("SM3 hash: %x\n", hash)
}
```

### SM4 Example

```go
package main

import (
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "github.com/tjfoc/gmsm/sm4"
)

func main() {
    key := []byte("1234567890abcdef") // 16 bytes for SM4
    block, err := sm4.NewCipher(key)
    if err != nil {
        panic(err)
    }
    
    plaintext := []byte("hello world")
    ciphertext := make([]byte, len(plaintext))
    
    // ECB mode (for demonstration - use proper mode in production)
    block.Encrypt(ciphertext, plaintext)
    
    fmt.Printf("Encrypted: %x\n", ciphertext)
}
```

## Documentation

For detailed API documentation and usage examples, see the [API documentation](./API使用说明.md) or visit the [GoDoc](https://godoc.org/github.com/n-WN/gmsm).

## Requirements

- Go 1.21 or later

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

Copyright 2017-2024 Suzhou Tongji Fintech Research Institute

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.