package input

import (
	"fmt"
	"os"
	"time"

	"github.com/itlightning/dateparse"
)

type Kind int

const (
	KindFile Kind = iota
	KindTimestamp
	KindDatetime
)

type Result struct {
	Kind          Kind
	Path          string
	FileInfo      os.FileInfo
	Time          time.Time
	TimestampUnit Unit
}

func Detect(arg string) (Result, error) {
	if fi, err := os.Stat(arg); err == nil {
		return Result{Kind: KindFile, Path: arg, FileInfo: fi}, nil
	}

	if unit, v, ok := parseUnixTimestamp(arg); ok {
		return Result{
			Kind:          KindTimestamp,
			Time:          unixToTime(v, unit),
			TimestampUnit: unit,
		}, nil
	}

	// ParseLocal (not ParseAny) so timezone-less strings resolve to local time.
	t, err := dateparse.ParseLocal(arg)
	if err != nil {
		return Result{}, fmt.Errorf("cannot parse %q as a file path, unix timestamp, or date/time string: %w", arg, err)
	}
	return Result{Kind: KindDatetime, Time: t}, nil
}

func unixToTime(v int64, unit Unit) time.Time {
	switch unit {
	case UnitSeconds:
		return time.Unix(v, 0)
	case UnitMilliseconds:
		return time.UnixMilli(v)
	case UnitMicroseconds:
		return time.UnixMicro(v)
	case UnitNanoseconds:
		return time.Unix(0, v)
	default:
		panic(fmt.Sprintf("unknown timestamp unit %q", unit))
	}
}
