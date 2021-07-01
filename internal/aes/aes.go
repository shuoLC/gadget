package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// 定义密钥和 iv 向量
const (
	iv  = "4c00ead2d3a039d3" //Go 的 iv 向量为必传参数
	key = "b49b38ae0a13e02d"
)

// Encrypt 加密
func Encrypt (encryptStr string) (string, error) {
	encryptBytes := []byte(encryptStr)
	block, err   := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encryptBytes = PKCS5Padding(encryptBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	encrypted := make([]byte, len(encryptBytes))
	blockMode.CryptBlocks(encrypted, encryptBytes)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// Decrypt 解密
func Decrypt (decryptStr string) (string, error) {
	decryptBytes, err := base64.StdEncoding.DecodeString(decryptStr)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	decrypted := make([]byte, len(decryptBytes))

	blockMode.CryptBlocks(decrypted, decryptBytes)
	decrypted = PKCS5UnPadding(decrypted)
	return string(decrypted), nil
}

func PKCS5Padding (cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func PKCS5UnPadding (decrypted []byte) []byte {
	length := len(decrypted)
	unPadding := int(decrypted[length-1])
	return decrypted[:(length - unPadding)]
}

