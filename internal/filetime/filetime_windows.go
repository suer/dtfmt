//go:build windows

package filetime

import (
	"os"
	"syscall"
	"time"
)

func Extract(fi os.FileInfo) Times {
	mtime := fi.ModTime()
	times := Times{Mtime: mtime, Atime: mtime}

	st, ok := fi.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return times
	}

	times.Atime = time.Unix(0, st.LastAccessTime.Nanoseconds())

	birthtime := time.Unix(0, st.CreationTime.Nanoseconds())
	times.Birthtime = &birthtime

	return times
}
