package format

import (
	"math"
	"time"
)

type UnixFormats struct {
	Seconds      int64  `json:"seconds"`
	Milliseconds int64  `json:"milliseconds"`
	Microseconds int64  `json:"microseconds"`
	Nanoseconds  *int64 `json:"nanoseconds"`
}

type NamedFormats struct {
	RFC3339     string `json:"rfc3339"`
	RFC3339Nano string `json:"rfc3339_nano"`
	RFC822      string `json:"rfc822"`
	RFC822Z     string `json:"rfc822z"`
	RFC850      string `json:"rfc850"`
	RFC1123     string `json:"rfc1123"`
	RFC1123Z    string `json:"rfc1123z"`
	ANSIC       string `json:"ansic"`
	UnixDate    string `json:"unix_date"`
	RubyDate    string `json:"ruby_date"`
	Kitchen     string `json:"kitchen"`
	Stamp       string `json:"stamp"`
	StampMilli  string `json:"stamp_milli"`
	StampMicro  string `json:"stamp_micro"`
	StampNano   string `json:"stamp_nano"`
	DateTime    string `json:"date_time"`
	DateOnly    string `json:"date_only"`
	TimeOnly    string `json:"time_only"`
}

type Formats struct {
	Unix  UnixFormats  `json:"unix"`
	Local NamedFormats `json:"local"`
	UTC   NamedFormats `json:"utc"`
}

func Build(t time.Time) Formats {
	return Formats{
		Unix:  buildUnix(t),
		Local: buildNamed(t.Local()),
		UTC:   buildNamed(t.UTC()),
	}
}

func buildUnix(t time.Time) UnixFormats {
	return UnixFormats{
		Seconds:      t.Unix(),
		Milliseconds: t.UnixMilli(),
		Microseconds: t.UnixMicro(),
		Nanoseconds:  safeUnixNano(t),
	}
}

var (
	minNanoTime = time.Unix(0, math.MinInt64)
	maxNanoTime = time.Unix(0, math.MaxInt64)
)

func safeUnixNano(t time.Time) *int64 {
	if t.Before(minNanoTime) || t.After(maxNanoTime) {
		return nil
	}
	n := t.UnixNano()
	return &n
}

func buildNamed(t time.Time) NamedFormats {
	return NamedFormats{
		RFC3339:     t.Format(time.RFC3339),
		RFC3339Nano: t.Format(time.RFC3339Nano),
		RFC822:      t.Format(time.RFC822),
		RFC822Z:     t.Format(time.RFC822Z),
		RFC850:      t.Format(time.RFC850),
		RFC1123:     t.Format(time.RFC1123),
		RFC1123Z:    t.Format(time.RFC1123Z),
		ANSIC:       t.Format(time.ANSIC),
		UnixDate:    t.Format(time.UnixDate),
		RubyDate:    t.Format(time.RubyDate),
		Kitchen:     t.Format(time.Kitchen),
		Stamp:       t.Format(time.Stamp),
		StampMilli:  t.Format(time.StampMilli),
		StampMicro:  t.Format(time.StampMicro),
		StampNano:   t.Format(time.StampNano),
		DateTime:    t.Format(time.DateTime),
		DateOnly:    t.Format(time.DateOnly),
		TimeOnly:    t.Format(time.TimeOnly),
	}
}
