//go:build darwin

package filetime

import (
	"os"
	"syscall"
	"time"
)

func Extract(fi os.FileInfo) Times {
	times := baseTimes(fi)

	st, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return times
	}

	times.Atime = time.Unix(st.Atimespec.Sec, st.Atimespec.Nsec)

	ctime := time.Unix(st.Ctimespec.Sec, st.Ctimespec.Nsec)
	times.Ctime = &ctime

	birthtime := time.Unix(st.Birthtimespec.Sec, st.Birthtimespec.Nsec)
	times.Birthtime = &birthtime

	return times
}
