package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnoprstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
