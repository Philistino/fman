package entry

import (
	"errors"
	"io/fs"
	"mime"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/spf13/afero"
	"golang.org/x/sync/errgroup"
)

type Dir struct {
	Path    string
	Entries []Entry
	ModTime time.Time
	SortO   SortOrder
}

func (d *Dir) Sort(method SortMethod, reverse bool) {
	d.SortO.method = method
	if reverse {
		d.SortO.reverse = !d.SortO.reverse
	}
	sortEntries(d.Path, d.Entries, d.SortO)
}

// Entry represents a file or directory
type Entry struct {
	fs.FileInfo        // file info
	SizeStr     string // human readable size string
	ModifyTime  string // last modified time
	MimeType    string // expected mime type from file extension
	SymlinkName string // name of the symlink
	SymLinkPath string // path of the symlink
	IsHidden    bool   // whether the file is hidden
	SizeInt     int64  // either size in bytes or count of entries in directory
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
		size = "1 item"
	} else {
		size = strconv.Itoa(lenEntries) + " items"
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
	entries := make([]Entry, len(files))
	errorS := make([]error, len(files))
	errMap := make(map[string]error, len(files)) // TODO: use this

	group := errgroup.Group{}
	group.SetLimit(10)
	for i, file := range files {
		i := i
		file := file
		group.Go(func() error {
			entries[i], errorS[i] = createEntry(fsys, dirPath, file)
			return nil
		})
		// TODO Do something with errors
		// if err != nil {
		// 	errMap[file.Name()] = err
		// }
	}
	group.Wait()

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

// Check For Changes in Directory
func CheckForChanges(fsys afero.Fs, dir Dir) (Dir, map[string]error, error) {
	dirInfo, err := fsys.Stat(dir.Path)
	if err != nil {
		return Dir{}, nil, err
	}
	if !dirInfo.IsDir() {
		return Dir{}, nil, errors.New("not a directory")
	}
	if dirInfo.ModTime() != dir.ModTime {
		entries, errMap, err := GetEntries(fsys, dir.Path, dir.SortO.showHidden, !dir.SortO.dirsOnly)
		dir = Dir{
			Path:    dir.Path,
			ModTime: dirInfo.ModTime(),
			Entries: entries,
			SortO:   dir.SortO,
		}
		return dir, errMap, err
	}

	returnEntries := make([]Entry, 0, len(dir.Entries))
	errMap := make(map[string]error, len(dir.Entries))
	for _, entry := range dir.Entries {
		fullPath := filepath.Join(dir.Path, entry.Name())
		file, err := fsys.Stat(fullPath)
		if errors.Is(err, fs.ErrNotExist) {
			continue
		}
		if err != nil {
			errMap[entry.Name()] = err
			continue
		}
		if entry.ModTime() != file.ModTime() {
			entry, err := createEntry(fsys, dir.Path, file)
			if err != nil {
				errMap[file.Name()] = err
			}
			returnEntries = append(returnEntries, entry)
			continue
		}
		returnEntries = append(returnEntries, entry)
	}
	dir.Entries = returnEntries
	return dir, errMap, nil
}
