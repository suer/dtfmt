package filetime

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func TestExtractMtimeAtime(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")
	if err := os.WriteFile(path, []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}

	mtime := time.Date(2024, 3, 15, 10, 30, 0, 0, time.UTC)
	atime := time.Date(2024, 3, 16, 11, 0, 0, 0, time.UTC)
	if err := os.Chtimes(path, atime, mtime); err != nil {
		t.Fatal(err)
	}

	fi, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	times := Extract(fi)
	if !times.Mtime.Equal(mtime) {
		t.Errorf("Mtime = %v, want %v", times.Mtime, mtime)
	}
	if !times.Atime.Equal(atime) {
		t.Errorf("Atime = %v, want %v", times.Atime, atime)
	}
}

func TestExtractDarwinCtimeBirthtime(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("ctime/birthtime assertions are darwin-specific")
	}

	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")
	if err := os.WriteFile(path, []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}

	fi, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	times := Extract(fi)
	if times.Ctime == nil {
		t.Fatal("Ctime is nil")
	}
	if times.Birthtime == nil {
		t.Fatal("Birthtime is nil")
	}
	if times.Birthtime.After(times.Mtime.Add(time.Second)) {
		t.Errorf("Birthtime %v is after Mtime %v", times.Birthtime, times.Mtime)
	}
}
