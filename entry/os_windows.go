//go:build windows

package entry

import (
	"io/fs"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows"
)

var (
	envOpener = os.Getenv("OPENER")
	envEditor = os.Getenv("VISUAL")
	envPager  = os.Getenv("PAGER")
	envShell  = os.Getenv("SHELL")
)

var envPathExt = os.Getenv("PATHEXT")

var (
	gDefaultShell     = "cmd"
	gDefaultShellFlag = "/c"
)

var (
	gUser        *user.User
	gConfigPaths []string
	gColorsPaths []string
	gIconsPaths  []string
	gFilesPath   string
	gTagsPath    string
	gMarksPath   string
	gHistoryPath string
)

func init() {
	if envOpener == "" {
		envOpener = `start ""`
	}

	if envEditor == "" {
		envEditor = os.Getenv("EDITOR")
		if envEditor == "" {
			envEditor = "notepad"
		}
	}

	if envPager == "" {
		envPager = "more"
	}

	if envShell == "" {
		envShell = "cmd"
	}

	u, err := user.Current()
	if err != nil {
		log.Printf("user: %s", err)
	}
	gUser = u

	// remove domain prefix
	gUser.Username = strings.Split(gUser.Username, `\`)[1]

	data := os.Getenv("LOCALAPPDATA")

	gConfigPaths = []string{
		filepath.Join(os.Getenv("ProgramData"), "lf", "lfrc"),
		filepath.Join(data, "lf", "lfrc"),
	}

	gColorsPaths = []string{
		filepath.Join(os.Getenv("ProgramData"), "lf", "colors"),
		filepath.Join(data, "lf", "colors"),
	}

	gIconsPaths = []string{
		filepath.Join(os.Getenv("ProgramData"), "lf", "icons"),
		filepath.Join(data, "lf", "icons"),
	}

	gFilesPath = filepath.Join(data, "lf", "files")
	gMarksPath = filepath.Join(data, "lf", "marks")
	gTagsPath = filepath.Join(data, "lf", "tags")
	gHistoryPath = filepath.Join(data, "lf", "history")
}

func detachedCommand(name string, arg ...string) *exec.Cmd {
	cmd := exec.Command(name, arg...)
	cmd.SysProcAttr = &windows.SysProcAttr{CreationFlags: 8}
	return cmd
}

func shellSetPG(cmd *exec.Cmd) {
}

func shellKill(cmd *exec.Cmd) error {
	return cmd.Process.Kill()
}

func setUserUmask() {}

func isExecutable(f fs.FileInfo) bool {
	exts := strings.Split(envPathExt, string(filepath.ListSeparator))
	for _, e := range exts {
		if strings.HasSuffix(strings.ToLower(f.Name()), strings.ToLower(e)) {
			log.Print(f.Name(), e)
			return true
		}
	}
	return false
}

func isHidden(f fs.FileInfo, dirPath string, _ []string) bool {
	ptr, err := windows.UTF16PtrFromString(filepath.Join(dirPath, f.Name()))
	if err != nil {
		return false
	}
	attrs, err := windows.GetFileAttributes(ptr)
	if err != nil {
		return false
	}
	return attrs&windows.FILE_ATTRIBUTE_HIDDEN != 0
}

func userName(f os.FileInfo) string {
	return ""
}

func groupName(f os.FileInfo) string {
	return ""
}

func linkCount(f os.FileInfo) string {
	return ""
}

func errCrossDevice(err error) bool {
	return err.(*os.LinkError).Err.(windows.Errno) == 17
}

// func exportFiles(f string, fs []string, pwd string) {
// 	envFile := fmt.Sprintf(`"%s"`, f)

// 	var quotedFiles []string
// 	for _, f := range fs {
// 		quotedFiles = append(quotedFiles, fmt.Sprintf(`"%s"`, f))
// 	}
// 	envFiles := strings.Join(quotedFiles, gOpts.filesep)

// 	envPWD := fmt.Sprintf(`"%s"`, pwd)

// 	os.Setenv("f", envFile)
// 	os.Setenv("fs", envFiles)
// 	os.Setenv("PWD", envPWD)

// 	if len(fs) == 0 {
// 		os.Setenv("fx", envFile)
// 	} else {
// 		os.Setenv("fx", envFiles)
// 	}
// }
