package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestRsa(test *testing.T) {
	bits := 1024
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	publicKey := privateKey.PublicKey
	var marshaledPrivateKey []byte
	marshaledPrivateKey, _ = x509.MarshalPKCS8PrivateKey(privateKey) // if bits == 2048;x509.MarshalPKCS1PrivateKey(privateKey)
	privateBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: marshaledPrivateKey,
		},
	)
	marshaledPublickey, _ := x509.MarshalPKIXPublicKey(&publicKey)

	publicBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: marshaledPublickey,
	})

	fmt.Println(strings.Repeat("-", 50) + "privateBytes" + strings.Repeat("-", 50))
	pbk := bytesToPublicKey(string(privateBytes))

	fmt.Println(strings.Repeat("-", 50) + "publicBytes" + strings.Repeat("-", 50))
	prk := bytesToPrivateKey(string(publicBytes))

	fmt.Println(prk, pbk)
	fmt.Println("Private Key (1024) :  ", *privateKey)
	fmt.Println("Public key (1024) ", publicKey)
	secretMessage := "hi hello 안녕 히히 1000-5-0"
	encryptedMessage := EncryptOAEP(secretMessage, publicKey)
	fmt.Println("Cipher Text  ", encryptedMessage)
	DecryptOAEP(encryptedMessage, *privateKey)
	//이걸로간다

}
func bytesToPrivateKey(k string) *rsa.PrivateKey {
	fmt.Println(k)
	block, _ := pem.Decode([]byte(k))
	b := block.Bytes
	var err error

	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		fmt.Println(err)
	}
	return key
}
func bytesToPublicKey(k string) *rsa.PublicKey {
	fmt.Println(k)
	block, _ := pem.Decode([]byte(k))
	if block == nil {
		fmt.Println(block)
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	key, ok := pub.(*rsa.PublicKey)
	if !ok {
	}
	return key
}

func EncryptOAEP(secretMessage string, pubkey rsa.PublicKey) string {
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &pubkey, []byte(secretMessage), label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return "Error from encryption"
	}
	return base64.StdEncoding.EncodeToString(ciphertext)
}
func DecryptOAEP(cipherText string, privKey rsa.PrivateKey) string {
	ct, _ := base64.StdEncoding.DecodeString(cipherText)
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, &privKey, ct, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return "Error from Decryption"
	}
	fmt.Printf("Plaintext: %s\n", string(plaintext))

	return string(plaintext)
}
