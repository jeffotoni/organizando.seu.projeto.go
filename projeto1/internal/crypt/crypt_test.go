package crypt

import "testing"

var mm = make(map[string]string, 0)

func TestCrc32(t *testing.T) {
	got := Crc32()
	want, ok := mm[got]
	if ok {
		t.Errorf("Crc32() = %q, want %q (já existe)", got, want)
		return
	} else {
		mm[got] = got
	}
}

func TestSha1(t *testing.T) {
	got := Sha1(RandomNew())
	want, ok := mm[got]
	if ok {
		t.Errorf("Sha1() = %q, want %q (já existe)", got, want)
		return
	} else {
		mm[got] = got
	}
}

func BenchmarkSha1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sha1(RandomNew())
	}
}

func BenchmarkCrc32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Crc32()
	}
}
