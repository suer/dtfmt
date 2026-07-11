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

// baseTimes returns the portable subset of Times available from os.FileInfo
// alone. Platform-specific Extract implementations start from this and then
// fill in Atime/Ctime/Birthtime from OS-specific stat data.
func baseTimes(fi os.FileInfo) Times {
	mtime := fi.ModTime()
	return Times{Mtime: mtime, Atime: mtime}
}
