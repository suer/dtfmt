//go:build !darwin && !linux && !windows

package filetime

import "os"

func Extract(fi os.FileInfo) Times {
	mtime := fi.ModTime()
	return Times{Mtime: mtime, Atime: mtime}
}
