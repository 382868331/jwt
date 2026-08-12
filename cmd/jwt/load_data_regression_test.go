package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDataReadsFileBeforeClosing(t *testing.T) {
	path := filepath.Join(t.TempDir(), "token.txt")
	want := []byte("header.payload.signature")
	if err := os.WriteFile(path, want, 0o600); err != nil {
		t.Fatal(err)
	}

	got, err := loadData(path)
	if err != nil {
		t.Fatalf("loadData returned an error: %v", err)
	}
	if string(got) != string(want) {
		t.Fatalf("loadData returned %q, want %q", got, want)
	}
}
