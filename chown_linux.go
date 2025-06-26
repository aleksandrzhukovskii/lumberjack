package lumberjack

import (
	"os"
	"syscall"
)

// osChown is a var so we can mock it out during tests.
var osChown = os.Chown

func chown(name string, info os.FileInfo, uid, gid int) error {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
	if err != nil {
		return err
	}
	f.Close()
	stat := info.Sys().(*syscall.Stat_t)
	if uid == 0 {
		uid = stat.Uid
	}
	if gid == 0 {
		gid = stat.Gid
	}
	return osChown(name, int(stat.Uid), int(stat.Gid))
}
