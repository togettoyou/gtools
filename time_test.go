package gtools

import (
	"encoding/json"
	"testing"
	"time"
)

type Model struct {
	CreatedAt FormatTime `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func TestTimeTool(t *testing.T) {
	t.Log(time.Now().Format(TimeFormat))
	m := &Model{
		CreatedAt: FormatTime{time.Now()},
		UpdatedAt: time.Now(),
	}
	bytes, err := json.Marshal(&m)
	ErrExit(t, err)
	t.Log(string(bytes))
}
