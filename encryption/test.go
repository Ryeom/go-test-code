package encryption

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func GetUUID() string {
	var r = rand.Reader

	a, _ := NewRandomFromReader(r)
	fmt.Println(a.String())
	return a.String()
}

/* UUID 연습 */

func NewRandomFromReader(r io.Reader) (UUID, error) {
	var uuid UUID
	_, err := io.ReadFull(r, uuid[:])
	if err != nil {
		return uuid, err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid, nil
}

type UUID [16]byte

func (uuid UUID) String() string {
	var buf [36]byte
	encodeHex(buf[:], uuid)
	return string(buf[:])
}

func encodeHex(dst []byte, uuid UUID) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}

/* @function 단방향 암호화 */

func GetSHA256(bv []byte) string {
	hasher := sha256.New()
	hasher.Write(bv)
	sha := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

/* @function @param message를 @param key로 암호화 */

func EncryptAES(key []byte, message string) (encryptText string, err error) {
	plainText := []byte(message)
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	encryptText = base64.URLEncoding.EncodeToString(cipherText)
	return
}

/* @function 암호화 된 @param encryptText를 @param key로 복호화 */

func DecryptAES(key []byte, secure string) (decoded string, err error) {
	cipherText, err := base64.URLEncoding.DecodeString(secure)
	if err != nil {
		return
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText), err
}

/* @function RSA Key Fair 생성 */

func GenerateRSAKeyFair() []string {
	var result []string
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024) // 개인 키와 공개키 생성
	if err != nil {
		return result
	}
	publicKey := &privateKey.PublicKey // 개인 키 변수 안에 공개 키가 들어있음
	t, _ := x509.MarshalPKCS8PrivateKey(privateKey)
	privateBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: t,
		},
	)
	marshaledPublickey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return result
	}

	publicBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: marshaledPublickey,
	})

	result = append(result, string(privateBytes))
	result = append(result, string(publicBytes))

	return result
}

func StringToRSAKey(s string, t string) interface{} {
	fmt.Println("key 암호화 해제", s)
	r := strings.NewReader(s)
	pemBytes, err := ioutil.ReadAll(r)
	if err != nil {
	}
	decoded, _ := pem.Decode(pemBytes)
	parsed, _ := x509.ParsePKCS8PrivateKey(decoded.Bytes)
	return parsed
}

func BytesToPrivateKey(p []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(p)
	var err error

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
	}
	return key
}
func BytesToPublicKey(p []byte) *rsa.PublicKey {
	block, _ := pem.Decode(p)
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

func EncryptWithPublicKey(msg string, pub *rsa.PublicKey) string {
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, []byte(msg), nil)
	if err != nil {
	}
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func DecryptWithPrivateKey(cipherText string, priv *rsa.PrivateKey) (string, error) {
	ct, _ := base64.StdEncoding.DecodeString(cipherText)
	hash := sha256.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ct, nil)
	if err != nil {
	}
	return string(plaintext), err
}

func Sign(privatekey *rsa.PrivateKey, data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	d := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, privatekey, crypto.SHA256, d)
}

func Verify(publickey *rsa.PublicKey, message []byte, sig []byte) error {
	h := sha256.New()
	h.Write(message)
	d := h.Sum(nil)
	return rsa.VerifyPKCS1v15(publickey, crypto.SHA256, d, sig)
}
