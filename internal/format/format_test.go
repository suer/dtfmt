package format

import (
	"testing"
	"time"
)

func TestBuildUnix(t *testing.T) {
	tm := time.Date(2023, 11, 14, 22, 13, 20, 0, time.UTC)
	f := Build(tm)

	if f.Unix.Seconds != 1700000000 {
		t.Errorf("Seconds = %d, want 1700000000", f.Unix.Seconds)
	}
	if f.Unix.Milliseconds != 1700000000000 {
		t.Errorf("Milliseconds = %d, want 1700000000000", f.Unix.Milliseconds)
	}
	if f.Unix.Microseconds != 1700000000000000 {
		t.Errorf("Microseconds = %d, want 1700000000000000", f.Unix.Microseconds)
	}
	if f.Unix.Nanoseconds == nil || *f.Unix.Nanoseconds != 1700000000000000000 {
		t.Errorf("Nanoseconds = %v, want 1700000000000000000", f.Unix.Nanoseconds)
	}
}

func TestBuildUnixNanoOverflow(t *testing.T) {
	tm := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)
	f := Build(tm)

	if f.Unix.Nanoseconds != nil {
		t.Errorf("Nanoseconds = %v, want nil for out-of-range year", *f.Unix.Nanoseconds)
	}
	if f.Unix.Seconds == 0 {
		t.Error("Seconds should still be populated for out-of-range year")
	}
}

func TestBuildNamedTimezones(t *testing.T) {
	jst := time.FixedZone("JST", 9*3600)
	tm := time.Date(2023, 11, 15, 15, 13, 20, 0, jst)
	f := Build(tm)

	wantLocal := tm.Local().Format(time.RFC3339)
	if f.Local.RFC3339 != wantLocal {
		t.Errorf("Local.RFC3339 = %q, want %q", f.Local.RFC3339, wantLocal)
	}

	wantUTC := "2023-11-15T06:13:20Z"
	if f.UTC.RFC3339 != wantUTC {
		t.Errorf("UTC.RFC3339 = %q, want %q", f.UTC.RFC3339, wantUTC)
	}
}
