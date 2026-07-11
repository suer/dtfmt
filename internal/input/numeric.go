package input

import (
	"regexp"
	"strconv"
)

type Unit string

const (
	UnitSeconds      Unit = "seconds"
	UnitMilliseconds Unit = "milliseconds"
	UnitMicroseconds Unit = "microseconds"
	UnitNanoseconds  Unit = "nanoseconds"
)

var numericRe = regexp.MustCompile(`^-?[0-9]+$`)

func parseUnixTimestamp(s string) (unit Unit, v int64, ok bool) {
	if !numericRe.MatchString(s) {
		return "", 0, false
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return "", 0, false
	}

	abs := v
	if abs < 0 {
		abs = -abs
	}
	digits := len(strconv.FormatInt(abs, 10))

	switch {
	case digits <= 10:
		unit = UnitSeconds
	case digits <= 13:
		unit = UnitMilliseconds
	case digits <= 16:
		unit = UnitMicroseconds
	default:
		unit = UnitNanoseconds
	}
	return unit, v, true
}
