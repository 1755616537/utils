package utils

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"log"
)

// AES - ECB解密
func AesDecrypt(ciphertext string) string {
	key := []byte("UJATIMB38cHO5X4ABekT4FZT0V7O0Pv3") // 加密的密钥
	encrypted, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	genKey := make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}

	cipher, _ := aes.NewCipher(genKey)
	decrypted := make([]byte, len(encrypted))

	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	decrypted = decrypted[:trim]

	log.Println("解密结果：", string(decrypted))
	return string(decrypted)
}
