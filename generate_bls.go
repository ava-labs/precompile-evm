package main

import (
    "encoding/hex"
    "fmt"
    "log"

    "github.com/ava-labs/avalanchego/utils/crypto/bls"
)

func main() {
    // Generate a new BLS private key
    sk, err := bls.NewSecretKey()
    if err != nil {
        log.Fatalf("failed to generate BLS secret key: %v", err)
    }

    // Get the public key
    pk := bls.PublicFromSecretKey(sk)
    pkBytes := bls.PublicKeyToCompressedBytes(pk)
    fmt.Printf("Public Key: 0x%s\n", hex.EncodeToString(pkBytes))

    // Message to sign
    message := []byte("Hello Ava Labs, I have a new kitten, his name is Sunny and he's an orange car")
    fmt.Printf("\nOriginal Message: %s\n", message)
    fmt.Printf("Message in hex: 0x%s\n", hex.EncodeToString(message))

    // Sign the message
    sig := bls.Sign(sk, message)
    sigBytes := bls.SignatureToBytes(sig)
    fmt.Printf("\nSignature: 0x%s\n", hex.EncodeToString(sigBytes))

    // Verify the signature
    isValid := bls.Verify(pk, sig, message)
    fmt.Printf("\nSignature is valid: %v\n", isValid)

    fmt.Printf("\nComplete signed message components:\n")
    fmt.Printf("----------------------------------------\n")
    fmt.Printf("Public Key: 0x%s\n", hex.EncodeToString(pkBytes))
	fmt.Printf("Message: 0x%s\n", hex.EncodeToString([]byte(message))) 
    fmt.Printf("Signature: 0x%s\n", hex.EncodeToString(sigBytes))
}