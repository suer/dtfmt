package output

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/suer/dtfmt/internal/input"
)

func TestBuildTimestamp(t *testing.T) {
	r := input.Result{
		Kind:          input.KindTimestamp,
		Time:          time.Unix(1700000000, 0),
		TimestampUnit: "seconds",
	}
	doc := Build("1700000000", r)

	if doc.Input.Type != "timestamp" {
		t.Errorf("Input.Type = %q, want timestamp", doc.Input.Type)
	}
	if doc.Input.Unit != "seconds" {
		t.Errorf("Input.Unit = %q, want seconds", doc.Input.Unit)
	}
	vt, ok := doc.Times.(*ValueTimes)
	if !ok {
		t.Fatalf("Times = %T, want *ValueTimes", doc.Times)
	}
	if vt.Value.Unix.Seconds != 1700000000 {
		t.Errorf("Value.Unix.Seconds = %d, want 1700000000", vt.Value.Unix.Seconds)
	}
}

func TestBuildDatetimeOmitsUnit(t *testing.T) {
	r := input.Result{Kind: input.KindDatetime, Time: time.Now()}
	doc := Build("2023-11-15T09:00:00Z", r)

	b, err := json.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(b), `"unit"`) {
		t.Errorf("expected no unit field, got %s", b)
	}
}

func TestBuildFileKeyOrderAndNullHandling(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")
	if err := os.WriteFile(path, []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}
	fi, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	r := input.Result{Kind: input.KindFile, Path: path, FileInfo: fi}
	doc := Build(path, r)

	ft, ok := doc.Times.(*FileTimes)
	if !ok {
		t.Fatalf("Times = %T, want *FileTimes", doc.Times)
	}

	b, err := json.Marshal(doc.Times)
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)

	mtimeIdx := strings.Index(s, `"mtime"`)
	atimeIdx := strings.Index(s, `"atime"`)
	ctimeIdx := strings.Index(s, `"ctime"`)
	birthtimeIdx := strings.Index(s, `"birthtime"`)
	if mtimeIdx >= atimeIdx || atimeIdx >= ctimeIdx || ctimeIdx >= birthtimeIdx {
		t.Errorf("expected key order mtime < atime < ctime < birthtime, got indices %d %d %d %d", mtimeIdx, atimeIdx, ctimeIdx, birthtimeIdx)
	}

	if ft.Ctime == nil && !strings.Contains(s, `"ctime":null`) {
		t.Errorf("expected ctime:null when unavailable, got %s", s)
	}
	if ft.Birthtime == nil && !strings.Contains(s, `"birthtime":null`) {
		t.Errorf("expected birthtime:null when unavailable, got %s", s)
	}
}
