package gtools

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// GenRsaKey 生成 rsa 密钥对
func GenRsaKey(bits int) (pubKey string, priKey string, err error) {
	// gen private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	priKey = string(pem.EncodeToMemory(priBlock))
	// gen public key
	publicKey := &privateKey.PublicKey
	ret, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return
	}
	pubBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: ret,
	}
	pubKey = string(pem.EncodeToMemory(pubBlock))
	return
}

// LoadRSAPrivateKeyFile 加载RSA私钥
func LoadRSAPrivateKeyFile(priKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(priKey))
	if block == nil {
		return nil, errors.New("rsa private key error")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// LoadRSAPublicKeyFile 加载RSA公钥
func LoadRSAPublicKeyFile(pubKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubKey))
	if block == nil {
		return nil, errors.New("rsa public key error")
	}
	pubKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	if pk, ok := pubKeyInterface.(*rsa.PublicKey); !ok {
		return nil, errors.New("type assertion error")
	} else {
		return pk, nil
	}
}

// RSAPublicEncrypt 使用RSA公钥加密
func RSAPublicEncrypt(pubKey string, data string) (string, error) {
	publicKey, err := LoadRSAPublicKeyFile(pubKey)
	if err != nil {
		return "", err
	}
	buffer := bytes.NewBufferString("")
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(data))
	if err != nil {
		return "", err
	}
	buffer.Write(encrypted)
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

// RSAPrivateDecrypt 使用RSA私钥解密
func RSAPrivateDecrypt(priKey string, encrypted string) (string, error) {
	privateKey, err := LoadRSAPrivateKeyFile(priKey)
	if err != nil {
		return "", err
	}
	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	buffer := bytes.NewBufferString("")
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
	if err != nil {
		return "", err
	}
	buffer.Write(decrypted)
	return buffer.String(), nil
}
