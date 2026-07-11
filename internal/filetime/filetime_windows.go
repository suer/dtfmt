//go:build windows

package filetime

import (
	"os"
	"syscall"
	"time"
)

func Extract(fi os.FileInfo) Times {
	times := baseTimes(fi)

	st, ok := fi.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return times
	}

	times.Atime = time.Unix(0, st.LastAccessTime.Nanoseconds())

	birthtime := time.Unix(0, st.CreationTime.Nanoseconds())
	times.Birthtime = &birthtime

	return times
}
