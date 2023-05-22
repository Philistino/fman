package entry

import (
	"bytes"
	"io/fs"
	"mime"
	"os"
	"path/filepath"
	"strconv"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/djherbis/times"
	"github.com/dustin/go-humanize"
)

type Entry struct {
	fs.FileInfo

	SizeStr string

	ModifyTime string
	AccessTime string
	ChangeTime string

	Type        string
	SymlinkName string
	SymLinkPath string
	timeStats   times.Timespec
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

func HighlightSyntax(name string, preview string) (string, error) {
	var buffer bytes.Buffer

	lexer := lexers.Match(name)

	if lexer == nil {
		lexer = lexers.Fallback
	}

	style := styles.Get("monokai")
	formatter := formatters.Get("terminal")

	iterator, err := lexer.Tokenise(nil, preview)

	if err != nil {
		return "", err
	}

	if err = formatter.Format(&buffer, style, iterator); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func GetEntry(info fs.FileInfo, path string) (Entry, error) {

	// Get Entry size
	size := humanize.IBytes(uint64(info.Size()))

	// If entry is a folder, get count of entries under this directory
	if info.IsDir() {
		_entries, err := os.ReadDir(path)

		if err != nil {
			return Entry{}, err
		}

		size = strconv.Itoa(len(_entries)) + " entries"

		if len(_entries) == 0 {
			size = "Empty Folder"
		}
	}

	return Entry{
		FileInfo: info,
		SizeStr:  size,

		Type: mime.TypeByExtension(filepath.Ext(info.Name())),

		ModifyTime: humanize.Time(info.ModTime()),
		// ChangeTime: humanize.Time(info.),
		// AccessTime: humanize.Time(timeStats.AccessTime()),
		// timeStats:  timeStats,
	}, nil

}

func GetEntries(path string, showHidden bool, dirsMixed bool) ([]Entry, error) {
	os.Chdir(path)
	newPath, _ := os.Getwd()

	files, err := os.ReadDir(newPath)

	if err != nil {
		return []Entry{}, err
	}

	entries := make([]Entry, 0, len(files))

	for _, file := range files {
		info, err := file.Info()

		if err != nil {
			continue
		}

		fullPath := filepath.Join(newPath, file.Name())

		if err != nil {
			continue
		}

		hidden, err := FileHidden(file.Name())
		if err != nil || (hidden && !showHidden) {
			continue
		}

		entry, err := GetEntry(info, fullPath)

		if err != nil {
			continue
		}

		// Handle Symlinks
		if info.Mode()&os.ModeSymlink != 0 {
			fullPath, err = os.Readlink(fullPath)

			if err != nil {
				continue
			}

			symInfo, err := os.Stat(fullPath)

			if err != nil {
				return []Entry{}, err
			}

			entry, err = GetEntry(symInfo, fullPath)

			if err != nil {
				return []Entry{}, err
			}

			entry.SymlinkName = info.Name()
			entry.SymLinkPath = fullPath
		}

		entries = append(entries, entry)
	}

	entries = sortE(
		path,
		entries,
		sortType{
			method: naturalSort,
			option: dirfirstSort,
		},
		true,
		true,
		[]string{".*"},
	)

	// if !dirsMixed {
	// 	sort.Slice(entries, func(i, j int) bool {
	// 		switch {
	// 		case entries[i].IsDir && !entries[j].IsDir:
	// 			return true
	// 		case !entries[i].IsDir && entries[j].IsDir:
	// 			return false
	// 		default:
	// 			return entries[i].Name < entries[j].Name
	// 		}
	// 	})
	// }
	return entries, nil
}

// func sortEntries(entries []Entry) []Entry {
// 	sort.Slice(entries, func(i, j int) bool {
// 		switch {
// 		case entries[i].IsDir() && !entries[j].IsDir:
// 			return true
// 		case !entries[i].IsDir && entries[j].IsDir:
// 			return false
// 		default:
// 			return entries[i].Name < entries[j].Name
// 		}
// 	})
// 	return entries
// }
