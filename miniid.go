package miniid

import "sync/atomic"

// Generator can encode integers into base62 strings.
type Generator struct {
	currID atomic.Int64
}

// NewGenerator creates a new generator using the provided init number.
func New(initID uint64) *Generator {
	var g Generator
	g.currID.Store(int64(initID))
	return &g
}

// Next returns a next ID in base62.
func (g *Generator) Next() string {
	newID := g.currID.Add(1)
	return encodeBase62(uint64(newID))
}

func encodeBase62(num uint64) string {
	const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	const base = uint64(len(alphabet))

	if num == 0 {
		return string(alphabet[0])
	}

	const maxLen = 16
	var result [maxLen]byte
	i := maxLen - 1

	for ; num > 0; i-- {
		rem := num % base
		result[i] = alphabet[rem]
		num /= base
	}
	return string(result[i+1:])
}
