package gtools

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
)

// JWK2PEM JWK JSON 格式转换为 PEM 格式公钥
func JWK2PEM(js string) (string, error) {
	jwk := map[string]string{}
	err := json.Unmarshal([]byte(js), &jwk)
	if err != nil {
		return "", err
	}
	if jwk["kty"] != "RSA" {
		return "", errors.New(fmt.Sprintf("invalid key type: %s", jwk["kty"]))
	}
	nb, err := base64.RawURLEncoding.DecodeString(jwk["n"])
	if err != nil {
		return "", err
	}

	e := 0
	// The default exponent is usually 65537, so just compare the
	// base64 for [1,0,1] or [0,1,0,1]
	if jwk["e"] == "AQAB" || jwk["e"] == "AAEAAQ" {
		e = 65537
	} else {
		// need to decode "e" as a big-endian int
		return "", errors.New(fmt.Sprintf("need to deocde e: %s", jwk["e"]))
	}

	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nb),
		E: e,
	}

	der, err := x509.MarshalPKIXPublicKey(pk)
	if err != nil {
		return "", err
	}

	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: der,
	}

	var out bytes.Buffer
	err = pem.Encode(&out, block)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
