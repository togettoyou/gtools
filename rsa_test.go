package gtools

import "testing"

func TestRSATool(t *testing.T) {
	pubKey, priKey, err := GenRsaKey(2048)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("公钥: %s\n私钥: %s", pubKey, priKey)
	data, err := RSAPublicEncrypt(pubKey, "123456")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("加密: %s", data)
	out, err := RSAPrivateDecrypt(priKey, data)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("解密: %s", out)
}
