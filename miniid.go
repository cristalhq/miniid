package miniid

import "sync/atomic"

const (
	Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base     = uint64(len(Alphabet))
)

var index = func() [256]int {
	var index [256]int
	for i := range index {
		index[i] = -1
	}
	for i, c := range Alphabet {
		index[c] = i
	}
	return index
}()

// Encode number in base62.
func Encode(n uint64) string {
	return encodeBase62(n)
}

// EncodeFixed number in base62.
func EncodeFixed(n uint64, length int) string {
	return encodeFixedBase62(n, length)
}

// Decode converts a base62-encoded string back into a uint64 number.
// It panics if the input contains characters not in the base62 alphabet.
func Decode(s string) uint64 {
	var n uint64
	for i := 0; i < len(s); i++ {
		val := index[s[i]]
		if val < 0 {
			panic("invalid base62 character")
		}
		n = n*base + uint64(val)
	}
	return n
}

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
	if num == 0 {
		return string(Alphabet[0])
	}

	const maxLen = 16
	var result [maxLen]byte
	i := maxLen - 1

	for ; num > 0; i-- {
		rem := num % base
		result[i] = Alphabet[rem]
		num /= base
	}
	return string(result[i+1:])
}

func encodeFixedBase62(num uint64, length int) string {
	buf := make([]byte, length)
	for i := length - 1; i >= 0; i-- {
		buf[i] = Alphabet[num%base]
		num /= base
	}
	return string(buf)
}
