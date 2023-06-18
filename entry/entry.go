package entry

import (
	"errors"
	"io/fs"
	"mime"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/spf13/afero"
)

type Entry struct {
	fs.FileInfo
	SizeStr     string
	ModifyTime  string
	MimeType    string
	SymlinkName string
	SymLinkPath string
	IsHidden    bool
	SizeInt     int64
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

func handleSymlink(fsys afero.Fs, fullPath string, file fs.FileInfo) (string, fs.FileInfo, error) {
	var symLinkPath string

	reader, ok := fsys.(afero.LinkReader) // check if filesystem supports reading symlinks
	if !ok {
		return symLinkPath, file, afero.ErrNoSymlink
	}
	symLinkPath, err := reader.ReadlinkIfPossible(fullPath)
	if err != nil {
		return symLinkPath, file, err
	}
	symInfo, err := fsys.Stat(symLinkPath)
	if err != nil {
		return symLinkPath, file, err
	}
	return symLinkPath, symInfo, nil
}

func getSize(fsys afero.Fs, file fs.FileInfo, filePath string) (string, int64, error) {
	if !file.IsDir() {
		return humanize.Bytes(uint64(file.Size())), file.Size(), nil
	}

	// get count of entries under this directory
	entries, err := afero.ReadDir(fsys, filePath)
	if errors.Is(err, fs.ErrPermission) {
		return "Access Denied", 0, nil
	}
	if err != nil {
		return "", 0, err
	}
	var size string
	lenEntries := len(entries)
	if lenEntries == 1 {
		size = "1 entry"
	} else {
		size = strconv.Itoa(lenEntries) + " entries"
	}
	return size, int64(lenEntries), nil
}

func createEntry(fsys afero.Fs, dirPath string, file fs.FileInfo) (Entry, error) {
	var symLinkName string
	var symLinkPath string

	fullPath := filepath.Join(dirPath, file.Name())
	hidden, err := isHidden(fullPath)
	if err != nil {
		hidden = false
	}
	// Handle Symlinks
	if file.Mode()&os.ModeSymlink != 0 {
		reader, ok := fsys.(afero.LinkReader) // check if filesystem supports reading symlinks
		if !ok {
			return Entry{FileInfo: file}, afero.ErrNoSymlink
		}
		linkedPath, err := reader.ReadlinkIfPossible(fullPath)
		if err != nil {
			return Entry{FileInfo: file}, err
		}
		symInfo, err := fsys.Stat(linkedPath)
		if err != nil {
			return Entry{FileInfo: file}, err
		}
		symLinkName = file.Name()
		symLinkPath = linkedPath
		file = symInfo
		fullPath = linkedPath
	}
	sizeStr, sizeInt, err := getSize(fsys, file, fullPath)
	if err != nil {
		return Entry{FileInfo: file}, err
	}
	entry := Entry{
		FileInfo:    file,
		SizeStr:     sizeStr,
		SizeInt:     sizeInt,
		ModifyTime:  humanize.Time(file.ModTime()),
		MimeType:    mime.TypeByExtension(filepath.Ext(file.Name())),
		SymlinkName: symLinkName,
		SymLinkPath: symLinkPath,
		IsHidden:    hidden,
	}
	return entry, err
}

func GetEntries(fsys afero.Fs, dirPath string, showHidden bool, dirsMixed bool) ([]Entry, map[string]error, error) {
	files, err := afero.ReadDir(fsys, dirPath)
	if err != nil {
		return nil, nil, err
	}
	entries := make([]Entry, 0, len(files))
	errMap := make(map[string]error, len(files)) // TODO: use this

	for _, file := range files {
		entry, err := createEntry(fsys, dirPath, file)
		if err != nil {
			errMap[file.Name()] = err
		}
		entries = append(entries, entry)
	}

	sortT := SortOrder{
		method:     NaturalSort,
		dirsFirst:  !dirsMixed,
		dirsOnly:   false,
		showHidden: showHidden,
		reverse:    false,
		ignoreDiac: true,
		ignoreCase: true,
	}

	entries = sortEntries(dirPath, entries, sortT)
	return entries, errMap, nil
}
