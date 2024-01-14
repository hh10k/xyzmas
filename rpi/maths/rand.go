package maths

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
	"time"
)

func newRand() *math_rand.Rand {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	var seed int64
	if err == nil {
		seed = int64(binary.LittleEndian.Uint64(b[:]))
	} else {
		seed = time.Now().UnixNano()
	}

	source := math_rand.NewSource(seed)
	return math_rand.New(source)
}

var Rand = newRand()
