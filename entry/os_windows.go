//go:build windows

package entry

import (
	"os/exec"

	"golang.org/x/sys/windows"
)

func detachedCommand(name string, arg ...string) *exec.Cmd {
	cmd := exec.Command(name, arg...)
	cmd.SysProcAttr = &windows.SysProcAttr{CreationFlags: 8}
	return cmd
}

func shellKill(cmd *exec.Cmd) error {
	return cmd.Process.Kill()
}

func isHidden(filePath string) (bool, error) {
	ptr, err := windows.UTF16PtrFromString(filePath)
	if err != nil {
		return false, err
	}
	attrs, err := windows.GetFileAttributes(ptr)
	if err != nil {
		return false, err
	}
	return attrs&windows.FILE_ATTRIBUTE_HIDDEN != 0, nil
}
