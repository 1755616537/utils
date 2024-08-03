package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

// 加密
func JiaMi(rsaPublicKey string) string {
	secretMessage := "send reinforcements, we're going to advance"

	cipherdata := EncryptOAEP(secretMessage)

	ciphertext := base64.StdEncoding.EncodeToString([]byte(cipherdata))

	return ciphertext
}

// 解密
func JieMi(ciphertext string, rsaPrivateKey string) string {
	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)

	plaintext := DecryptOAEP(string(cipherdata))

	return plaintext
}

// 加密
func EncryptOAEP(text string) string {
	rsaPublicKey := ParsePKIXPublicKey()
	secretMessage := []byte(text)
	rng := rand.Reader
	cipherdata, err := rsa.EncryptOAEP(sha1.New(), rng, rsaPublicKey, secretMessage, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return ""
	}
	ciphertext := base64.StdEncoding.EncodeToString(cipherdata)
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	return ciphertext
}

// 解密
func DecryptOAEP(ciphertext string) string {
	rsaPrivateKey := ParsePKCS1PrivateKey()
	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, rsaPrivateKey, cipherdata, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return ""
	}
	fmt.Printf("Plaintext: %s\n", string(plaintext))
	return string(plaintext)
}

// 解析公钥
func ParsePKIXPublicKey() *rsa.PublicKey {
	publicKey, err := ioutil.ReadFile("static/cert/apiclient_cert.pem")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	block, _ := pem.Decode(publicKey)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pubInterface.(*rsa.PublicKey)
}

// 解析私钥
func ParsePKCS1PrivateKey() *rsa.PrivateKey {
	privateKey, err := ioutil.ReadFile("static/cert/apiclient_key.pem")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	block, _ := pem.Decode(privateKey)
	privateInterface, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return privateInterface
}
