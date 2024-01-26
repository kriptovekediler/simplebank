package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnoprstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}
