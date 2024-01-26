package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnoprstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
