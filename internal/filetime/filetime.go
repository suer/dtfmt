package filetime

import (
	"os"
	"time"
)

type Times struct {
	Mtime     time.Time
	Atime     time.Time
	Ctime     *time.Time
	Birthtime *time.Time
}

func baseTimes(fi os.FileInfo) Times {
	mtime := fi.ModTime()
	return Times{Mtime: mtime, Atime: mtime}
}
