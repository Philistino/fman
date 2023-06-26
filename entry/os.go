//go:build !windows

package entry

import (
	"fmt"
	"os"
	// "os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"
	// "golang.org/x/sys/unix"
)

func isHidden(fullPath string) (bool, error) {
	hidden := false
	name := filepath.Base(fullPath)
	if strings.HasPrefix(name, ".") && !strings.HasPrefix(name, "..") {
		hidden = true
	}
	return hidden, nil
}

func userName(f os.FileInfo) string {
	if stat, ok := f.Sys().(*syscall.Stat_t); ok {
		if u, err := user.LookupId(fmt.Sprint(stat.Uid)); err == nil {
			return fmt.Sprintf("%v ", u.Username)
		}
	}
	return ""
}

func groupName(f os.FileInfo) string {
	if stat, ok := f.Sys().(*syscall.Stat_t); ok {
		if g, err := user.LookupGroupId(fmt.Sprint(stat.Gid)); err == nil {
			return fmt.Sprintf("%v ", g.Name)
		}
	}
	return ""
}

// func linkCount(f os.FileInfo) string {
// 	if stat, ok := f.Sys().(*syscall.Stat_t); ok {
// 		return fmt.Sprintf("%v ", stat.Nlink)
// 	}
// 	return ""
// }
