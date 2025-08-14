package miniid

import "sync/atomic"

// Generator can encode integers into base62 strings.
type Generator struct {
	currID atomic.Int64
}

// NewGenerator creates a new generator using the provided init number.
func New(initID int64) *Generator {
	var g Generator
	g.currID.Store(initID)
	return &g
}

// Next returns a next ID in base62.
func (g *Generator) Next() string {
	newID := g.currID.Add(1)
	return encodeBase62(newID)
}

func encodeBase62(num int64) string {
	const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	const base = int64(len(alphabet))

	if num == 0 {
		return string(alphabet[0])
	}

	result := make([]byte, 0)
	for num > 0 {
		rem := num % base
		result = append([]byte{alphabet[rem]}, result...)
		num /= base
	}
	return string(result)
}
