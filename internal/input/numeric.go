package input

import (
	"regexp"
	"strconv"
)

var numericRe = regexp.MustCompile(`^-?[0-9]+$`)

func parseUnixTimestamp(s string) (unit string, v int64, ok bool) {
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
		unit = "seconds"
	case digits <= 13:
		unit = "milliseconds"
	case digits <= 16:
		unit = "microseconds"
	default:
		unit = "nanoseconds"
	}
	return unit, v, true
}
