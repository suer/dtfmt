package input

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestDetectFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")
	if err := os.WriteFile(path, []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}

	r, err := Detect(path)
	if err != nil {
		t.Fatal(err)
	}
	if r.Kind != KindFile {
		t.Fatalf("Kind = %v, want KindFile", r.Kind)
	}
	if r.Path != path {
		t.Errorf("Path = %q, want %q", r.Path, path)
	}
	if r.FileInfo == nil {
		t.Error("FileInfo is nil")
	}
}

func TestDetectTimestamp(t *testing.T) {
	r, err := Detect("1700000000")
	if err != nil {
		t.Fatal(err)
	}
	if r.Kind != KindTimestamp {
		t.Fatalf("Kind = %v, want KindTimestamp", r.Kind)
	}
	if r.TimestampUnit != "seconds" {
		t.Errorf("TimestampUnit = %q, want seconds", r.TimestampUnit)
	}
	if !r.Time.Equal(time.Unix(1700000000, 0)) {
		t.Errorf("Time = %v, want %v", r.Time, time.Unix(1700000000, 0))
	}
}

func TestDetectDatetime(t *testing.T) {
	r, err := Detect("2023-11-15T09:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	if r.Kind != KindDatetime {
		t.Fatalf("Kind = %v, want KindDatetime", r.Kind)
	}
	want := time.Date(2023, 11, 15, 9, 0, 0, 0, time.UTC)
	if !r.Time.Equal(want) {
		t.Errorf("Time = %v, want %v", r.Time, want)
	}
}

func TestDetectUnrecognized(t *testing.T) {
	_, err := Detect("not a valid input at all")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
