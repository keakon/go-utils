package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
	"testing"
)

func TestSum(t *testing.T) {
	_, err := Sum(sha256.New(), "")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	f, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	defer os.RemoveAll(f.Name())

	data := bytes.Repeat([]byte{'a'}, fileBufferSize)
	h := sha256.New()
	h.Write(data)
	wantHash := hex.EncodeToString(h.Sum(nil))

	f.Write(data)
	hash, err := Sum(sha256.New(), f.Name())
	if err != nil {
		t.Fatal(err)
	}
	if hash != wantHash {
		t.Errorf("want %s, got %s", wantHash, hash)
	}

	h.Write([]byte{'a'})
	wantHash = hex.EncodeToString(h.Sum(nil))
	f.Write([]byte{'a'})
	hash, err = Sum(sha256.New(), f.Name())
	if err != nil {
		t.Fatal(err)
	}
	if hash != wantHash {
		t.Errorf("want %s, got %s", wantHash, hash)
	}
}

func TestSHA256(t *testing.T) {
	f, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	defer os.RemoveAll(f.Name())

	data := bytes.Repeat([]byte{'a'}, fileBufferSize)
	f.Write(data)
	hash, err := SHA256(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	if hash == "" {
		t.Error("Failed to get SHA256")
	}
}

func TestMD5(t *testing.T) {
	f, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	defer os.RemoveAll(f.Name())

	data := bytes.Repeat([]byte{'a'}, fileBufferSize)
	f.Write(data)
	hash, err := MD5(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	if hash == "" {
		t.Error("Failed to get MD5")
	}
}
