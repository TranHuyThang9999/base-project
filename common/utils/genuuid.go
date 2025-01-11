package utils

import (
	"crypto/rand"
	"math/big"
	"sync/atomic"
	"time"
)

type UUID struct {
	counter int64
}

func NewUUID() *UUID {
	return &UUID{
		counter: 0,
	}
}

func (u *UUID) GenUUID() int64 {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	count := atomic.AddInt64(&u.counter, 1)

	return (timestamp << 20) | (count & 0xFFFFF)
}

func (u *UUID) GenTime() time.Time {
	now := time.Now().UTC()

	return now
}

func (u *UUID) GenPassWord() int64 {
	max := big.NewInt(99999999 - 10000000 + 1)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0
	}

	return n.Int64() + 10000000
}
