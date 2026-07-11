# dtfmt

A CLI tool that guesses what kind of date/time value its argument is (file path, unix timestamp, or date/time string) and prints it back out in every common date/time format, as JSON.

## Install

```bash
$ go install github.com/suer/dtfmt@latest
```

Or install a prebuilt release binary via [mise](https://mise.jdx.dev/):

```bash
$ mise use -g github:suer/dtfmt
```

## Usage

```bash
$ dtfmt <file-path|unix-timestamp|datetime-string>
```

dtfmt tries to interpret the argument in this order:

1. **File path** — if it exists, outputs `mtime`/`atime`/`ctime`/`birthtime`. `ctime`/`birthtime` are OS-dependent; when unavailable the field is `null`.
2. **Unix timestamp** — a numeric string; the unit (seconds/milliseconds/microseconds/nanoseconds) is guessed from its digit count.
3. **Date/time string** — parsed via [araddon/dateparse](https://github.com/araddon/dateparse); timezone-less strings are interpreted as local time.

Each resolved time is expanded into unix formats (seconds/milliseconds/microseconds/nanoseconds) and named formats (RFC3339, RFC822, RFC1123, etc.), in both local time and UTC.

### Examples

```bash
$ dtfmt 1700000000
```
```json
{
  "input": { "type": "timestamp", "value": "1700000000", "unit": "seconds" },
  "times": {
    "value": {
      "unix": {
        "seconds": 1700000000,
        "milliseconds": 1700000000000,
        "microseconds": 1700000000000000,
        "nanoseconds": 1700000000000000000
      },
      "local": { "rfc3339": "2023-11-15T07:13:20+09:00" },
      "utc": { "rfc3339": "2023-11-14T22:13:20Z" }
    }
  }
}
```

```bash
$ dtfmt go.mod
```
```json
{
  "input": { "type": "file", "value": "go.mod" },
  "times": {
    "mtime": { "unix": { "seconds": 1752192000 }, "local": { "rfc3339": "..." }, "utc": { "rfc3339": "..." } },
    "atime": { "..." : "..." },
    "ctime": { "..." : "..." },
    "birthtime": { "..." : "..." }
  }
}
```

```bash
$ dtfmt "2023-11-15 09:00:00"
```
```json
{
  "input": { "type": "datetime", "value": "2023-11-15 09:00:00" },
  "times": { "value": { "unix": { "seconds": 1700006400 }, "local": { "rfc3339": "2023-11-15T09:00:00+09:00" } } }
}
```

## Build

```bash
$ mise run build
```

## Test

```bash
$ mise run test
```
