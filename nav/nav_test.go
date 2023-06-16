package nav

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"unsafe"
)

func TestHandleCursor(t *testing.T) {
	testcases := []struct {
		name    string
		srcPath string
		dstPath string
		want    string
	}{
		{
			name:    "Parent",
			srcPath: "/a/b",
			dstPath: "/a",
			want:    "b",
		},
		{
			name:    "Windows",
			srcPath: "C:/Users/Jimbo/Documents/GitHub",
			dstPath: "C:/Users/Jimbo",
			want:    "Documents",
		},
		{
			name:    "Windows fwd slash",
			srcPath: "C:\\Users\\Jimbo\\Documents\\GitHub",
			dstPath: "C:\\Users\\Jimbo",
			want:    "Documents",
		},
		{
			name:    "Unix",
			srcPath: "/a/b/c/d/e",
			dstPath: "/a/b",
			want:    "c",
		},
		{
			name:    "Different root",
			srcPath: "/a/b/c/d/e",
			dstPath: "/d/e/f/g/h",
			want:    "",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			n := NewNav(true, true, "Bingo", nil)
			n.currentPath = tc.srcPath
			got := n.handleCursor(tc.dstPath)
			if got != tc.want {
				t.Errorf("handleCursor(%s, %s) = %s; want %s", tc.srcPath, tc.dstPath, got, tc.want)
			}
		})
	}
}

func TestHandleCursorWithHistory(t *testing.T) {
	testcases := []struct {
		name    string
		srcPath string
		dstPath string
		want    string
	}{
		{
			name:    "Windows",
			srcPath: "C:/Users/Jimbo",
			dstPath: "C:/Users/Jimbo/Documents/GitHub",
			want:    "Github",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			n := NewNav(true, true, "Bingo", nil)
			n.cursorHist[tc.dstPath] = tc.want
			n.currentPath = tc.srcPath
			got := n.handleCursor(tc.dstPath)
			if got != tc.want {
				t.Errorf("handleCursor(%s, %s) = %s; want %s", tc.srcPath, tc.dstPath, got, tc.want)
			}
		})
	}
}

const PathSeparator = filepath.Separator

// Canonicalize checks that the file path is valid and returns it in the "canonical" form:
// - /foo/bar -> foo/bar
// - / -> "."
func Canonicalize(file string) (string, error) {
	const pathSep = string(PathSeparator)

	if strings.HasPrefix(file, pathSep+pathSep) {
		// The relative path may pretend to be an absolute path within
		// the root, but the double path separator on Windows implies
		// something else and is out of spec.
		return "", errors.New("invalid path")
	}

	// The relative path should be clean from internal dotdots and similar
	// funkyness.
	file = filepath.Clean(file)
	// file = filepath.ToSlash(file)
	// It is not acceptable to attempt to traverse upwards.
	if file == ".." {
		return "", errors.New("errPathTraversingUpwards")
	}
	if strings.HasPrefix(file, ".."+pathSep) {
		return "", errors.New("errPathTraversingUpwards")
	}

	if strings.HasPrefix(file, pathSep) {
		if file == pathSep {
			return ".", nil
		}
		return file[1:], nil
	}

	return file, nil
}

func TestCanonicalize(t *testing.T) {
	type testcase struct {
		path     string
		expected string
		ok       bool
	}
	cases := []testcase{
		// Valid cases
		{"/bar", "bar", true},
		{"/bar/baz", "bar/baz", true},
		{"bar", "bar", true},
		{"bar/baz", "bar/baz", true},

		// Not escape attempts, but oddly formatted relative paths
		{"", ".", true},
		{"/", ".", true},
		{"/..", ".", true},
		{"./bar", "bar", true},
		{"./bar/baz", "bar/baz", true},
		{"bar/../baz", "baz", true},
		{"/bar/../baz", "baz", true},
		{"./bar/../baz", "baz", true},

		// Results in an allowed path, but does it by probing. Disallowed.
		{"../foo", "", false},
		{"../foo/bar", "", false},
		{"../foo/bar", "", false},
		{"../../baz/foo/bar", "", false},
		{"bar/../../foo/bar", "", false},
		{"bar/../../../baz/foo/bar", "", false},

		// Escape attempts.
		{"..", "", false},
		{"../", "", false},
		{"../bar", "", false},
		{"../foobar", "", false},
		{"bar/../../quux/baz", "", false},
	}

	for _, tc := range cases {
		res, err := Canonicalize(tc.path)
		if tc.ok {
			if err != nil {
				t.Errorf("Unexpected error for Canonicalize(%q): %v", tc.path, err)
				continue
			}
			exp := filepath.FromSlash(tc.expected)
			if res != exp {
				t.Errorf("Unexpected result for Canonicalize(%q): %q != expected %q", tc.path, res, exp)
			}
		} else if err == nil {
			t.Errorf("Unexpected pass for Canonicalize(%q) => %q", tc.path, res)
			continue
		}
	}
}

func getFinalPathName(in string) (string, error) {
	// Return the normalized path
	// Wrap the call to GetFinalPathNameByHandleW
	// The string returned by this function uses the \?\ syntax
	// Implies GetFullPathName + GetLongPathName
	kernel32, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		return "", err
	}
	GetFinalPathNameByHandleW, err := kernel32.FindProc("GetFinalPathNameByHandleW")
	// https://github.com/golang/go/blob/ff048033e4304898245d843e79ed1a0897006c6d/src/internal/syscall/windows/syscall_windows.go#L303
	if err != nil {
		return "", err
	}
	inPath, err := syscall.UTF16PtrFromString(in)
	if err != nil {
		return "", err
	}
	// Get a file handler
	h, err := syscall.CreateFile(inPath,
		syscall.GENERIC_READ,
		syscall.FILE_SHARE_READ,
		nil,
		syscall.OPEN_EXISTING,
		uint32(syscall.FILE_FLAG_BACKUP_SEMANTICS),
		0)
	if err != nil {
		return "", err
	}
	defer syscall.CloseHandle(h)
	// Call GetFinalPathNameByHandleW
	var VOLUME_NAME_DOS uint32 = 0x0      // not yet defined in syscall
	var bufSize uint32 = syscall.MAX_PATH // 260
	for i := 0; i < 2; i++ {
		buf := make([]uint16, bufSize)
		var ret uintptr
		ret, _, err = GetFinalPathNameByHandleW.Call(
			uintptr(h),                       // HANDLE hFile
			uintptr(unsafe.Pointer(&buf[0])), // LPWSTR lpszFilePath
			uintptr(bufSize),                 // DWORD  cchFilePath
			uintptr(VOLUME_NAME_DOS),         // DWORD  dwFlags
		)
		// The returned value is the actual length of the norm path
		// After Win 10 build 1607, MAX_PATH limitations have been removed
		// so it is necessary to check newBufSize
		newBufSize := uint32(ret) + 1
		if ret == 0 || newBufSize > bufSize*100 {
			break
		}
		if newBufSize <= bufSize {
			return syscall.UTF16ToString(buf), nil
		}
		bufSize = newBufSize
	}
	return "", err
}

func TestGetFinalPathNameByHandleW(t *testing.T) {
	testCases := []struct {
		input         string
		expectedPath  string
		eqToEvalSyml  bool
		ignoreMissing bool
	}{
		{
			input:         `c:\`,
			expectedPath:  `C:\`,
			eqToEvalSyml:  true,
			ignoreMissing: false,
		},
		{
			input:         `\\?\c:\`,
			expectedPath:  `C:\`,
			eqToEvalSyml:  false,
			ignoreMissing: false,
		},
		{
			input:         `c:\wInDows\sYstEm32`,
			expectedPath:  `C:\Windows\System32`,
			eqToEvalSyml:  true,
			ignoreMissing: false,
		},

		{
			input:         `c:\parent\child`,
			expectedPath:  `C:\parent\child`,
			eqToEvalSyml:  false,
			ignoreMissing: true,
		},
	}

	for _, testCase := range testCases {
		out, err := getFinalPathName(testCase.input)
		if err != nil {
			if testCase.ignoreMissing && os.IsNotExist(err) {
				continue
			}
			t.Errorf("getFinalPathName failed at %q with error %s", testCase.input, err)
		}
		// Trim UNC prefix
		if strings.HasPrefix(out, `\\?\UNC\`) {
			out = `\` + out[7:]
		} else {
			out = strings.TrimPrefix(out, `\\?\`)
		}
		if out != testCase.expectedPath {
			t.Errorf("getFinalPathName got wrong path: %q (expected %q)", out, testCase.expectedPath)
		}
		if testCase.eqToEvalSyml {
			evlPath, err1 := filepath.EvalSymlinks(testCase.input)
			if err1 != nil || out != evlPath {
				t.Errorf("EvalSymlinks got different results %q %s", evlPath, err1)
			}
		}
	}
}
