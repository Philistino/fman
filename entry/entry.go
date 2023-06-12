package entry

import (
	"io/fs"
	"mime"
	"os"
	"path/filepath"
	"strconv"

	"github.com/djherbis/times"
	"github.com/dustin/go-humanize"
)

type Entry struct {
	fs.FileInfo
	SizeStr     string
	ModifyTime  string
	ChangeTime  string
	MimeType    string
	SymlinkName string
	SymLinkPath string
	timeStats   times.Timespec
	IsHidden    bool
}

// Name returns the basename of the file
//
// this method is provided because the ui creates
// zeroed Entry structs in one place (which it really shouldn't)
func (e Entry) Name() string {
	if e.FileInfo == nil {
		return ""
	}
	return e.FileInfo.Name()
}

// IsDir returns whether the file is a directory
//
// this method is provided because the ui creates
// zeroed Entry structs in one place (which it really shouldn't)
func (e Entry) IsDir() bool {
	if e.FileInfo == nil {
		return false
	}
	return e.FileInfo.IsDir()
}

func newEntry(info fs.FileInfo, path string, hidden bool) (Entry, error) {
	var size string

	if info.IsDir() {
		// get count of entries under this directory
		_entries, err := os.ReadDir(path)
		if err != nil {
			return Entry{FileInfo: info}, err
		}
		lenEntries := len(_entries)
		switch {
		case lenEntries == 0:
			size = "Empty Folder"
		case lenEntries == 1:
			size = "1 entry"
		default:
			size = strconv.Itoa(lenEntries) + " entries"
		}
	} else {
		size = humanize.Bytes(uint64(info.Size()))
	}

	return Entry{
		FileInfo:   info,
		SizeStr:    size,
		MimeType:   mime.TypeByExtension(filepath.Ext(info.Name())),
		ModifyTime: humanize.Time(info.ModTime()),
		IsHidden:   hidden,
	}, nil
}

func GetEntries(path string, showHidden bool, dirsMixed bool) ([]Entry, error) {
	os.Chdir(path)
	cwd, _ := os.Getwd()

	files, err := os.ReadDir(cwd)
	if err != nil {
		return nil, err
	}

	entries := make([]Entry, 0, len(files))
	errMap := make(map[string]error, len(files)) // TODO: use this
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			errMap[file.Name()] = err
			continue
		}

		fullPath := cwd + "/" + file.Name()

		hidden := isHidden(info, cwd, []string{"."})
		if hidden && !showHidden {
			continue
		}

		entry, err := newEntry(info, fullPath, hidden)
		if err != nil {
			errMap[file.Name()] = err
			continue
		}

		// Handle Symlinks
		if info.Mode()&os.ModeSymlink != 0 {
			fullPath, err = os.Readlink(fullPath)
			if err != nil {
				errMap[file.Name()] = err
				continue
			}

			symInfo, err := os.Stat(fullPath)
			if err != nil {
				errMap[file.Name()] = err
				continue
			}

			entry, err = newEntry(symInfo, fullPath, hidden)
			if err != nil {
				errMap[file.Name()] = err
				continue
			}

			entry.SymlinkName = info.Name()
			entry.SymLinkPath = fullPath
		}

		entries = append(entries, entry)
	}

	var dirsFirst sortOption
	if dirsMixed {
		dirsFirst = noneSort
	} else {
		dirsFirst = dirfirstSort
	}

	entries = sortE(
		path,
		entries,
		sortType{
			method: naturalSort,
			option: dirsFirst,
		},
		true,
		true,
		[]string{".*"},
	)
	return entries, nil
}
