package gtools

import "testing"

func TestMD5Tool(t *testing.T) {
	v := "123456"
	t.Log(MD5LowercaseEncode(v))
	t.Log(MD5Lowercase16Encode(v))
	t.Log(MD5CapitalEncode(v))
	t.Log(MD5Capital16Encode(v))
}
