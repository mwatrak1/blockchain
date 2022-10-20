package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	address    string
}

func NewWallet() *Wallet {
	wallet := new(Wallet)

	// 1. Creating ECDSA provate key (32 bytes) public key (64 bytes)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	wallet.privateKey = privateKey
	wallet.publicKey = &privateKey.PublicKey

	// 2. Perform SHA-256 on the public key
	publicKeySha256Hash := sha256.New()
	publicKeySha256Hash.Write(wallet.PrivateKey().X.Bytes())
	publicKeySha256Hash.Write(wallet.PrivateKey().Y.Bytes())
	publicKeySha256Digest := publicKeySha256Hash.Sum(nil)

	// 3. Perform RIPEMD-160 hashing on the result of SHA-256
	publicKeyRipemdSha256Hash := ripemd160.New()
	publicKeyRipemdSha256Hash.Write(publicKeySha256Digest)
	publicKeyRipemdSha256Digest := publicKeyRipemdSha256Hash.Sum(nil)

	// 4. Add version byte in front of RIPEMD-160 hash (0x00 for Bitcoin Main Network)
	publicKeyDoubleHashedWithVersion := make([]byte, 21)
	publicKeyDoubleHashedWithVersion[0] = 0x00
	copy(publicKeyDoubleHashedWithVersion[1:], publicKeyRipemdSha256Digest[:])

	// 5. Perforn SHA-256 on the versioned double hashed public key
	publicKeyTripleHashed := sha256.New()
	publicKeyTripleHashed.Write(publicKeyDoubleHashedWithVersion)
	publicKeyTripleHashedDigest := publicKeyTripleHashed.Sum(nil)

	// 6. Perform another SHA-256 hash on the previous result
	publicKeyQuadrupleHashed := sha256.New()
	publicKeyQuadrupleHashed.Write(publicKeyTripleHashedDigest)
	publicKeyQuadrupleHashedDigest := publicKeyQuadrupleHashed.Sum(nil)

	// 7. Take the first 4 bytes for the checksum
	checksum := publicKeyQuadrupleHashedDigest[:4]

	// 8. Add the checksum to the current hashed public key
	publicKeyQuadrupleHashedWithChecksum := make([]byte, 25)
	copy(publicKeyQuadrupleHashedWithChecksum[:21], publicKeyQuadrupleHashedDigest[:])
	copy(publicKeyQuadrupleHashedWithChecksum[21:], checksum[:])

	// 9. Convert the hash with checksum to base58 string
	address := base58.Encode(publicKeyQuadrupleHashedWithChecksum)
	wallet.address = address

	return wallet
}

func (wallet *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return wallet.privateKey
}

func (wallet *Wallet) PrivateKeyString() string {
	return fmt.Sprintf("%x", wallet.privateKey.D.Bytes())
}

func (wallet *Wallet) PublicKey() *ecdsa.PublicKey {
	return wallet.publicKey
}

func (wallet *Wallet) PublicKeyString() string {
	return fmt.Sprintf("%x%x", wallet.publicKey.X.Bytes(), wallet.publicKey.Y.Bytes())
}

func (wallet *Wallet) Address() string {
	return wallet.address
}
