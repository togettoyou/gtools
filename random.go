package gtools

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

type random struct {
	rnd *rand.Rand
	mx  sync.Mutex
}

func NewRandom() *random {
	return &random{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
		mx:  sync.Mutex{},
	}
}

// GenString 生成指定长度的随机字符串
func (r *random) GenString(n int, allowedChars ...[]rune) string {
	r.mx.Lock()
	defer r.mx.Unlock()
	var letters []rune
	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.rnd.Intn(len(letters))]
	}
	return string(b)
}

// GenCode 生成指定长度的随机数字
func (r *random) GenCode(length int) string {
	r.mx.Lock()
	defer r.mx.Unlock()
	var container string
	for i := 0; i < length; i++ {
		container += fmt.Sprintf("%01v", r.rnd.Int31n(10))
	}
	return container
}

// GenNum 生成指定范围内 [min,max) 随机值
func (r *random) GenNum(min, max int) int {
	r.mx.Lock()
	defer r.mx.Unlock()
	return r.rnd.Intn(max-min) + min
}
