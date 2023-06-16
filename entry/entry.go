package entry

import (
	"io/fs"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/djherbis/times"
	"github.com/dustin/go-humanize"
)

type Entry struct {
	fs.FileInfo
	SizeStr     string
	ModifyTime  string
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

func newEntry(fsys fs.FS, info fs.FileInfo, path string, hidden bool) (Entry, error) {
	var size string

	if info.IsDir() {
		// get count of entries under this directory
		_entries, err := fs.ReadDir(fsys, path)
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
		FileInfo:    info,
		SizeStr:     size,
		ModifyTime:  humanize.Time(info.ModTime()),
		MimeType:    mime.TypeByExtension(filepath.Ext(info.Name())),
		SymlinkName: "",
		SymLinkPath: path,
		timeStats:   nil,
		IsHidden:    hidden,
	}, nil
}

func GetEntries(fsys fs.FS, path string, showHidden bool, dirsMixed bool) ([]Entry, map[string]error, error) {
	files, err := fs.ReadDir(fsys, path)
	time.Sleep(time.Second * 10)
	if err != nil {
		return nil, nil, err
	}

	entries := make([]Entry, 0, len(files))
	errMap := make(map[string]error, len(files)) // TODO: use this
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			errMap[file.Name()] = err
			continue
		}

		fullPath := path + "/" + file.Name()

		hidden := isHidden(info, path, []string{"."})
		if hidden && !showHidden {
			continue
		}

		entry, err := newEntry(fsys, info, fullPath, hidden)
		if err != nil {
			log.Println("errored here ++++++++++++++++++++++++")
			errMap[file.Name()] = err
			continue
		}

		// Handle Symlinks
		if info.Mode()&os.ModeSymlink != 0 {
			fullPath, err = os.Readlink(fullPath)
			if err != nil {
				log.Println("errored in symlink")
				errMap[file.Name()] = err
				continue
			}

			symInfo, err := fs.Stat(fsys, fullPath)
			if err != nil {
				log.Println("errored stating symlink")
				errMap[file.Name()] = err
				continue
			}

			entry, err = newEntry(fsys, symInfo, fullPath, hidden)
			if err != nil {
				log.Println("errored in newEntry")
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
	return entries, errMap, nil
}
