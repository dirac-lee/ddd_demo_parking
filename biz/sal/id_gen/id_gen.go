package id_gen

import (
	"context"
	"math"
	"math/rand"
	"time"
)

type IDGen struct {
}

func (gen IDGen) GetID(ctx context.Context) int64 {
	var newLocalId uint64
	timestamp := time.Now().Unix()
	timestampBits := uint64(timestamp << 32)
	randomIdInfoMask := uint64(math.Pow(2, 18)-1)<<14 + uint64(math.Pow(2, 4)-1)
	restRandomBits := uint64(rand.Int63())>>31 | uint64(rand.Int63())<<32&randomIdInfoMask
	newLocalId = newLocalId | timestampBits
	newLocalId = newLocalId | restRandomBits
	return int64(newLocalId)
}

func New() IDGen {
	return IDGen{}
}
