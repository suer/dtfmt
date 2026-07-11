//go:build linux

package filetime

import (
	"os"
	"syscall"
	"time"
)

func Extract(fi os.FileInfo) Times {
	mtime := fi.ModTime()
	times := Times{Mtime: mtime, Atime: mtime}

	st, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return times
	}

	times.Atime = time.Unix(st.Atim.Sec, st.Atim.Nsec)

	ctime := time.Unix(st.Ctim.Sec, st.Ctim.Nsec)
	times.Ctime = &ctime

	return times
}
