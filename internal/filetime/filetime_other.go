//go:build !darwin && !linux && !windows

package filetime

import "os"

func Extract(fi os.FileInfo) Times {
	return baseTimes(fi)
}
