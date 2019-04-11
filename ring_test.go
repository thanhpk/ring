package ring

import (
	"testing"
)

func compareArr(arr, brr []interface{}) bool {
	if len(arr) != len(brr) {
		return false
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != brr[i] {
			return false
		}
	}
	return true
}

func TestRing(t *testing.T) {
	r := NewRing()
	if r.Size() != 0 {
		t.Fatalf("should be zero, got %d", r.Size())
	}

	for i := 0; i < 100100; i++ {
		r.Append(i)
	}
	if r.Size() != 100100 {
		t.Fatalf("should be 10100, got %d", r.Size())
	}

	for i := 0; i < 100; i++ {
		r.Drop()
	}
	if r.Size() != 100000 {
		t.Fatalf("should be 100000, got %d", r.Size())
	}

	i := 100
	r.Each(func(v interface{}) {
		if i != v.(int) {
			t.Fatalf("expect %d, got %d", i, v.(int))
		}
		i++
	})
}

func TestRingExpand(t *testing.T) {
	r := NewRing()
	for i := 0; i < 10; i++ {
		r.Append(i)
	}

	if r.Cap() != 16 {
		t.Fatalf("expect 16, got %d", r.Cap())
	}

	for i := 0; i < 100; i++ {
		r.Append(i)
	}

	if r.Cap() != 128 {
		t.Fatalf("expect 128, got %d", r.Cap())
	}
}

func TestRingShrink(t *testing.T) {
	r := NewRing()
	for i := 0; i < 100; i++ {
		r.Append(i)
	}
	if r.Cap() != 128 {
		t.Fatalf("expect 128, got %d", r.Cap())
	}

	for i := 0; i < 90; i++ {
		r.Drop()
	}

	if r.Cap() != 16 {
		t.Fatalf("expect 16, got %d", r.Cap())
	}
}

func TestRingRotate(t *testing.T) {
	r := NewRing()
	for i := 0; i < 100; i++ {
		r.Append(i)
	}
	if r.Cap() != 128 {
		t.Fatalf("expect 128, got %d", r.Cap())
	}

	for i := 0; i < 10000; i++ {
		r.Drop()
		r.Drop()
		r.Append(i)
		r.Append(i)
	}

	if r.Cap() != 128 {
		t.Fatalf("expect 128, got %d", r.Cap())
	}
}

func TestRotate(t *testing.T) {
	tcs := []struct {
		in  []interface{}
		out []interface{}
		n   int
	}{
		{
			[]interface{}{6, 7, 8, 9, 0, 1, 2, 3, 4, 5},
			[]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			-4,
		},
		{
			[]interface{}{6, 7, 8, 9, 0, 1, 2, 3, 4, 5},
			[]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			-24,
		},
		{
			[]interface{}{17, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			[]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
			17,
		},
	}

	for _, tc := range tcs {
		rotate(tc.in, tc.n)
		if !compareArr(tc.in, tc.out) {
			t.Fatalf("expect %v, got %v", tc.out, tc.in)
		}
	}
}
