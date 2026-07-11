package input

import "testing"

func TestParseUnixTimestamp(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		wantUnit Unit
		wantV    int64
		wantOK   bool
	}{
		{"zero", "0", UnitSeconds, 0, true},
		{"10 digits seconds", "1700000000", UnitSeconds, 1700000000, true},
		{"boundary 10 digits seconds", "9999999999", UnitSeconds, 9999999999, true},
		{"boundary 11 digits milliseconds", "10000000000", UnitMilliseconds, 10000000000, true},
		{"13 digits milliseconds", "1700000000000", UnitMilliseconds, 1700000000000, true},
		{"boundary 13 digits milliseconds", "9999999999999", UnitMilliseconds, 9999999999999, true},
		{"boundary 14 digits microseconds", "10000000000000", UnitMicroseconds, 10000000000000, true},
		{"16 digits microseconds", "1700000000000000", UnitMicroseconds, 1700000000000000, true},
		{"boundary 16 digits microseconds", "9999999999999999", UnitMicroseconds, 9999999999999999, true},
		{"boundary 17 digits nanoseconds", "10000000000000000", UnitNanoseconds, 10000000000000000, true},
		{"19 digits nanoseconds", "1700000000000000000", UnitNanoseconds, 1700000000000000000, true},
		{"negative seconds", "-1700000000", UnitSeconds, -1700000000, true},
		{"negative zero", "-0", UnitSeconds, 0, true},
		{"leading zeros normalize digit count", "0001700000000", UnitSeconds, 1700000000, true},
		{"decimal not a timestamp", "12.34", "", 0, false},
		{"trailing letter not a timestamp", "12a", "", 0, false},
		{"empty string", "", "", 0, false},
		{"whitespace", "  123", "", 0, false},
		{"overflow 20 digits", "12345678901234567890", "", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unit, v, ok := parseUnixTimestamp(tt.in)
			if ok != tt.wantOK {
				t.Fatalf("ok = %v, want %v", ok, tt.wantOK)
			}
			if !ok {
				return
			}
			if unit != tt.wantUnit {
				t.Errorf("unit = %q, want %q", unit, tt.wantUnit)
			}
			if v != tt.wantV {
				t.Errorf("v = %d, want %d", v, tt.wantV)
			}
		})
	}
}
