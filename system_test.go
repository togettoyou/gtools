package gtools

import "testing"

func TestSystemTool(t *testing.T) {
	t.Log(GetCurrentIP().String())
}
