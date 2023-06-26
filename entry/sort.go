package entry

import (
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/afero"
)

type SortMethod uint8

const (
	NaturalSort SortMethod = iota
	NameSort
	SizeSort
	MtimeSort
	ExtSort
)

type SortOrder struct {
	method     SortMethod
	dirsFirst  bool
	dirsOnly   bool
	showHidden bool
	reverse    bool
	ignoreDiac bool
	ignoreCase bool
}

func normalize(s1, s2 string, ignorecase, ignoredia bool) (string, string) {
	if ignorecase {
		s1 = strings.ToLower(s1)
		s2 = strings.ToLower(s2)
	}
	if ignoredia {
		s1 = afero.NeuterAccents(s1)
		s2 = afero.NeuterAccents(s2)
	}
	return s1, s2
}

// sortEntries sorts the entries according to the given sortType.
func sortEntries(dirPath string, entries []Entry, sortT SortOrder) []Entry {
	ignorecase, ignoredia := sortT.ignoreCase, sortT.ignoreDiac
	switch sortT.method {
	case NaturalSort:
		sort.SliceStable(entries, func(i, j int) bool {
			s1, s2 := normalize(entries[i].Name(), entries[j].Name(), ignorecase, ignoredia)
			return naturalLess(s1, s2)
		})
	case NameSort:
		sort.SliceStable(entries, func(i, j int) bool {
			s1, s2 := normalize(entries[i].Name(), entries[j].Name(), ignorecase, ignoredia)
			return s1 < s2
		})
	case SizeSort:
		sort.SliceStable(entries, func(i, j int) bool {
			return entries[i].SizeInt < entries[j].SizeInt
		})
	case MtimeSort:
		sort.SliceStable(entries, func(i, j int) bool {
			return entries[i].ModTime().Before(entries[j].ModTime())
		})
	case ExtSort:
		sort.SliceStable(entries, func(i, j int) bool {
			ext1, ext2 := normalize(filepath.Ext(entries[i].Name()), filepath.Ext(entries[j].Name()), ignorecase, ignoredia)

			// if the extension could not be determined (directories, files without)
			// use a zero byte so that these files can be ranked higher
			if ext1 == "" {
				ext1 = "\x00"
			}
			if ext2 == "" {
				ext2 = "\x00"
			}

			name1, name2 := normalize(entries[i].Name(), entries[j].Name(), ignorecase, ignoredia)

			// in order to also have natural sorting with the filenames
			// combine the name with the ext but have the ext at the front
			return ext1 < ext2 || ext1 == ext2 && name1 < name2
		})
	}

	if sortT.dirsFirst {
		sort.SliceStable(entries, func(i, j int) bool {
			if entries[i].IsDir() == entries[j].IsDir() {
				return i < j
			}
			return entries[i].IsDir()
		})
	}

	if sortT.reverse {
		for i, j := 0, len(entries)-1; i < j; i, j = i+1, j-1 {
			entries[i], entries[j] = entries[j], entries[i]
		}
	}

	// TODO: SHOULD THESE NEXT TWO BE MOVED INTO A DIFFERENT FUNCTION?

	// when dironly option is enabled, we move files to the beginning of our file
	// list and then set the beginning of displayed files to the first directory
	// in the list
	if sortT.dirsOnly {
		sort.SliceStable(entries, func(i, j int) bool {
			if !entries[i].IsDir() && !entries[j].IsDir() {
				return i < j
			}
			return !entries[i].IsDir()
		})
		entries = func() []Entry {
			for i, f := range entries {
				if f.IsDir() {
					return entries[i:]
				}
			}
			return entries[len(entries):]
		}()
	}

	// when hidden option is disabled, we move hidden files to the
	// beginning of our file list and then set the beginning of displayed
	// files to the first non-hidden file in the list
	if !sortT.showHidden {
		sort.SliceStable(entries, func(i, j int) bool {
			if entries[i].IsHidden == entries[j].IsHidden {
				return i < j
			}
			return entries[i].IsHidden
		})
		for i, f := range entries {
			if !f.IsHidden {
				entries = entries[i:]
				break
			}
		}
		if len(entries) > 0 && entries[len(entries)-1].IsHidden {
			entries = entries[len(entries):]
		}
	}

	// dir.ind = max(dir.ind, 0)
	// dir.ind = min(dir.ind, len(entries)-1)

	return entries
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

// This function compares two strings for natural sorting which takes into
// account values of numbers in strings. For example, '2' is less than '10',
// and similarly 'foo2bar' is less than 'foo10bar', but 'bar2bar' is greater
// than 'foo10bar'.
func naturalLess(s1, s2 string) bool {
	lo1, lo2, hi1, hi2 := 0, 0, 0, 0
	for {
		if hi1 >= len(s1) {
			return hi2 != len(s2)
		}

		if hi2 >= len(s2) {
			return false
		}

		isDigit1 := isDigit(s1[hi1])
		isDigit2 := isDigit(s2[hi2])

		for lo1 = hi1; hi1 < len(s1) && isDigit(s1[hi1]) == isDigit1; hi1++ {
		}

		for lo2 = hi2; hi2 < len(s2) && isDigit(s2[hi2]) == isDigit2; hi2++ {
		}

		if s1[lo1:hi1] == s2[lo2:hi2] {
			continue
		}

		if isDigit1 && isDigit2 {
			num1, err1 := strconv.Atoi(s1[lo1:hi1])
			num2, err2 := strconv.Atoi(s2[lo2:hi2])

			if err1 == nil && err2 == nil {
				return num1 < num2
			}
		}

		return s1[lo1:hi1] < s2[lo2:hi2]
	}
}
