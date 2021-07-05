package gtools

import (
	"runtime"
	"testing"
)

func TestExecTool(t *testing.T) {
	cmdClient := NewCmdClient(func(out string) {
		t.Log(out)
	})
	ErrExit(cmdClient.Command("go", "version"))

	if runtime.GOOS == `windows` {
		ErrExit(cmdClient.Command(
			"cmd.exe",
			"/c",
			`chcp 65001 && for /l %n in (1, 1, 5) do start /wait timeout /T 1 && echo hello`,
		))
	} else {
		ErrExit(cmdClient.Command(
			"/bin/bash",
			"-c",
			`#!/bin/bash
for ((i=1; i<=5; i++))
do
  sleep 1
  echo hello
done`,
		))
	}
}
