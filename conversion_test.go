package gtools

import "testing"

func TestConversionTool(t *testing.T) {
	t.Log(
		StructToMap(
			struct {
				Code int `map:"code"`
				Msg  string
			}{
				Code: 0,
				Msg:  "hello",
			}),
	)
	t.Log(ArrayToString([]string{"hello", "world"}), ArrayToString([]int{1, 2}))
	t.Log(StrToFloat64("99.989", 2))
}
