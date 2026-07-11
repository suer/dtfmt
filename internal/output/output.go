package output

import (
	"fmt"

	"github.com/suer/dtfmt/internal/filetime"
	"github.com/suer/dtfmt/internal/format"
	"github.com/suer/dtfmt/internal/input"
)

type InputInfo struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Unit  string `json:"unit,omitempty"`
}

type FileTimes struct {
	Mtime     format.Formats  `json:"mtime"`
	Atime     format.Formats  `json:"atime"`
	Ctime     *format.Formats `json:"ctime"`
	Birthtime *format.Formats `json:"birthtime"`
}

type ValueTimes struct {
	Value format.Formats `json:"value"`
}

type Document struct {
	Input InputInfo   `json:"input"`
	Times interface{} `json:"times"`
}

func Build(arg string, r input.Result) Document {
	switch r.Kind {
	case input.KindFile:
		ft := filetime.Extract(r.FileInfo)
		times := &FileTimes{
			Mtime: format.Build(ft.Mtime),
			Atime: format.Build(ft.Atime),
		}
		if ft.Ctime != nil {
			f := format.Build(*ft.Ctime)
			times.Ctime = &f
		}
		if ft.Birthtime != nil {
			f := format.Build(*ft.Birthtime)
			times.Birthtime = &f
		}
		return Document{Input: InputInfo{Type: "file", Value: arg}, Times: times}

	case input.KindTimestamp:
		return Document{
			Input: InputInfo{Type: "timestamp", Value: arg, Unit: string(r.TimestampUnit)},
			Times: &ValueTimes{Value: format.Build(r.Time)},
		}

	case input.KindDatetime:
		return Document{
			Input: InputInfo{Type: "datetime", Value: arg},
			Times: &ValueTimes{Value: format.Build(r.Time)},
		}

	default:
		panic(fmt.Sprintf("unknown input kind %v", r.Kind))
	}
}
