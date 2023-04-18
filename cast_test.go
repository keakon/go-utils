package utils

import (
	"bytes"
	"strings"
	"testing"
)

var s = strings.Repeat("a", 1<<20)
var bs = []byte(s)

func TestStringToBytes(t *testing.T) {
	b := StringToBytes(s)
	if !bytes.Equal(b, bs) {
		t.Fail()
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToBytes(s)
	}
}

func BenchmarkConvertStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

func TestBytesToString(t *testing.T) {
	ss := BytesToString(bs)
	if ss != s {
		t.Fail()
	}
}

func BenchmarkBytesToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesToString(bs)
	}
}

func BenchmarkConvertBytesToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(bs)
	}
}
