package miniid

import (
	"sync"
	"testing"
)

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
