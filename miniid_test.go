package miniid

import (
	"sync"
	"testing"
)

func TestEncode(t *testing.T) {
	testCases := []struct {
		num  uint64
		want string
	}{
		{0, "0"},
		{1, "1"},
		{2, "2"},
		{61, "z"},
		{62, "10"},
		{62 * 61, "z0"},
		{100, "1c"},
		{62*62 - 2, "zy"},
		{62*62*62 - 2, "zzy"},
		{62*62*62 - 1, "zzz"},
		{1<<64 - 3, "LygHa16AHYD"},
	}

	for _, tc := range testCases {
		have := Encode(tc.num)

		if have != tc.want {
			t.Errorf("Start at %d want %s, have '%s'",
				tc.num, tc.want, have)
		}
	}
}

func TestEncodeFixed(t *testing.T) {
	testCases := []struct {
		num    uint64
		length int
		want   string
	}{
		{0, 7, "0000000"},
		{1, 7, "0000001"},
		{2, 7, "0000002"},
		{61, 7, "000000z"},
		{62, 7, "0000010"},
		{62 * 61, 7, "00000z0"},
		{100, 7, "000001c"},
		{62*62 - 2, 7, "00000zy"},
		{62*62*62 - 2, 7, "0000zzy"},
		{62*62*62 - 1, 7, "0000zzz"},
		{1<<64 - 3, 7, "a16AHYD"},
	}

	for _, tc := range testCases {
		have := EncodeFixed(tc.num, tc.length)

		if have != tc.want {
			t.Errorf("Start at %d want %s, have '%s'",
				tc.num, tc.want, have)
		}
	}
}

func TestEncodeDecode(t *testing.T) {
	tests := []uint64{
		0, 1, 10, 61, 62, 12345, 999999, 405160095143, 1<<63 - 1,
	}

	for _, n := range tests {
		enc := Encode(n)
		dec := Decode(enc)
		if dec != n {
			t.Fatalf("Encode/Decode failed for %d: got %d (encoded: %s)", n, dec, enc)
		}
	}
}

func TestGenerator(t *testing.T) {
	testCases := []struct {
		num  uint64
		want []string
	}{
		{0, []string{"1", "2", "3", "4", "5"}},
		{1, []string{"2", "3", "4", "5"}},
		{60, []string{"z", "10", "11"}},
		{62 * 61, []string{"z1", "z2", "z3"}},
		{100, []string{"1d", "1e", "1f"}},
		{62*62 - 2, []string{"zz", "100", "101"}},
		{62*62*62 - 2, []string{"zzz", "1000"}},
		{1<<64 - 3, []string{"LygHa16AHYE", "LygHa16AHYF", "0", "1", "2"}},
	}

	for _, tc := range testCases {
		start := tc.num
		g := New(start)

		next := start
		for i, want := range tc.want {
			next++

			have := g.Next()
			if have != want {
				t.Errorf("Start at %d, call #%d: want %s, have '%s'",
					tc.num, i+1, want, have)
			}

			if s := Encode(next); s != have {
				t.Errorf("call #%d: want %s, have '%s'", i+1, want, have)
			}
		}
	}
}

func TestGeneratorConcurrency(t *testing.T) {
	const n = 1000
	g := New(0)

	var wg sync.WaitGroup
	results := make(chan string, n)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			results <- g.Next()
		}()
	}

	wg.Wait()
	close(results)

	seen := make(map[string]struct{}, n)
	for id := range results {
		if _, ok := seen[id]; ok {
			t.Fatalf("Duplicate ID generated: %s", id)
		}
		seen[id] = struct{}{}
	}

	if len(seen) != n {
		t.Fatalf("want %d unique IDs, have %d", n, len(seen))
	}
}

func BenchmarkGenerate(b *testing.B) {
	g := New(1000)

	for b.Loop() {
		s := g.Next()
		if s == "" {
			b.Fail()
		}
	}
}
