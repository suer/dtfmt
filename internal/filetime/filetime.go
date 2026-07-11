package filetime

import "time"

type Times struct {
	Mtime     time.Time
	Atime     time.Time
	Ctime     *time.Time
	Birthtime *time.Time
}
