package gtools

import "testing"

func TestFilterTool(t *testing.T) {
	arrS := make([]string, 0)
	arrI := make([]int, 0)
	for i := 0; i < 100000; i++ {
		arrS = append(arrS, "hello")
		arrS = append(arrS, "world")
		arrI = append(arrI, i)
		arrI = append(arrI, i)
	}
	t.Log(RemoveDuplicateStr(arrS))
	t.Log(RemoveDuplicateInt(arrI))
}
