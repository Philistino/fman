package entry

import (
	"path/filepath"
	"sort"
	"strings"
)

type sortMethod byte

const (
	naturalSort sortMethod = iota
	nameSort
	sizeSort
	timeSort
	atimeSort
	ctimeSort
	extSort
)

type sortOption byte

const (
	dirfirstSort sortOption = iota
	hiddenSort
	reverseSort
	noneSort
)

type sortType struct {
	method sortMethod
	option sortOption
}

func normalize(s1, s2 string, ignorecase, ignoredia bool) (string, string) {
	if ignorecase {
		s1 = strings.ToLower(s1)
		s2 = strings.ToLower(s2)
	}
	if ignoredia {
		s1 = removeDiacritics(s1)
		s2 = removeDiacritics(s2)
	}
	return s1, s2
}

func searchMatch(name, pattern string, ignorecase, ignoredia, smartcase, smartdia, globsearch, dironly bool) (matched bool, err error) {
	if ignorecase {
		lpattern := strings.ToLower(pattern)
		if !smartcase || lpattern == pattern {
			pattern = lpattern
			name = strings.ToLower(name)
		}
	}
	if ignoredia {
		lpattern := removeDiacritics(pattern)
		if !smartdia || lpattern == pattern {
			pattern = lpattern
			name = removeDiacritics(name)
		}
	}
	if globsearch {
		return filepath.Match(pattern, name)
	}
	return strings.Contains(name, pattern), nil
}

// func isFiltered(f os.FileInfo, filter []string) bool {
// 	for _, pattern := range filter {
// 		matched, err := searchMatch(f.Name(), strings.TrimPrefix(pattern, "!"))
// 		if err != nil {
// 			log.Printf("Filter Error: %s", err)
// 			return false
// 		}
// 		if strings.HasPrefix(pattern, "!") && matched {
// 			return true
// 		} else if !strings.HasPrefix(pattern, "!") && !matched {
// 			return true
// 		}
// 	}
// 	return false
// }

// TODO: RENAME!!!!
func sortE(dirPath string, entries []Entry, sortT sortType, ignorecase, ignoredia bool, hiddenfiles []string) []Entry {

	switch sortT.method {
	case naturalSort:
		sort.SliceStable(entries, func(i, j int) bool {
			s1, s2 := normalize(entries[i].Name(), entries[j].Name(), ignorecase, ignoredia)
			return naturalLess(s1, s2)
		})
	case nameSort:
		sort.SliceStable(entries, func(i, j int) bool {
			s1, s2 := normalize(entries[i].Name(), entries[j].Name(), ignorecase, ignoredia)
			return s1 < s2
		})
	case sizeSort:
		sort.SliceStable(entries, func(i, j int) bool {
			return entries[i].Size() < entries[j].Size()
		})
	case timeSort:
		sort.SliceStable(entries, func(i, j int) bool {
			return entries[i].timeStats.ModTime().Before(entries[j].timeStats.ModTime())
		})
	case atimeSort:
		sort.SliceStable(entries, func(i, j int) bool {
			return entries[i].timeStats.AccessTime().Before(entries[j].timeStats.AccessTime())
		})
	case ctimeSort:
		sort.SliceStable(entries, func(i, j int) bool {
			return entries[i].timeStats.ChangeTime().Before(entries[j].timeStats.ChangeTime())
		})
	case extSort:
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

	if sortT.option == reverseSort {
		for i, j := 0, len(entries)-1; i < j; i, j = i+1, j-1 {
			entries[i], entries[j] = entries[j], entries[i]
		}
	}

	if sortT.option == dirfirstSort {
		sort.SliceStable(entries, func(i, j int) bool {
			if entries[i].IsDir() == entries[j].IsDir() {
				return i < j
			}
			return entries[i].IsDir()
		})
	}

	// when dironly option is enabled, we move files to the beginning of our file
	// list and then set the beginning of displayed files to the first directory
	// in the list
	// THIS CAUSES A PANIC IN THE ENTRYINFO VIEW METHOD IN EMPTY DIRECTORIES
	// if dironly {
	// 	sort.SliceStable(entries, func(i, j int) bool {
	// 		if !entries[i].IsDir() && !entries[j].IsDir() {
	// 			return i < j
	// 		}
	// 		return !entries[i].IsDir()
	// 	})
	// 	entries = func() []Entry {
	// 		for i, f := range entries {
	// 			if f.IsDir() {
	// 				return entries[i:]
	// 			}
	// 		}
	// 		return entries[len(entries):]
	// 	}()
	// }

	// when hidden option is disabled, we move hidden files to the
	// beginning of our file list and then set the beginning of displayed
	// files to the first non-hidden file in the list
	if sortT.option == hiddenSort {
		sort.SliceStable(entries, func(i, j int) bool {
			if isHidden(entries[i], dirPath, hiddenfiles) && isHidden(entries[j], dirPath, hiddenfiles) {
				return i < j
			}
			return isHidden(entries[i], dirPath, hiddenfiles)
		})
		for i, f := range entries {
			if !isHidden(f, dirPath, hiddenfiles) {
				entries = entries[i:]
				break
			}
		}
		if len(entries) > 0 && isHidden(entries[len(entries)-1], dirPath, hiddenfiles) {
			entries = entries[len(entries):]
		}
	}

	// if len(dir.filter) != 0 {
	// 	sort.SliceStable(entries, func(i, j int) bool {
	// 		if isFiltered(entries[i], filter) && isFiltered(entries[j], filter) {
	// 			return i < j
	// 		}
	// 		return isFiltered(entries[i], filter)
	// 	})
	// 	for i, f := range entries {
	// 		if !isFiltered(f, filter) {
	// 			entries = entries[i:]
	// 			break
	// 		}
	// 	}
	// 	if len(entries) > 0 && isFiltered(entries[len(entries)-1], filter) {
	// 		entries = entries[len(entries):]
	// 	}
	// }

	// dir.ind = max(dir.ind, 0)
	// dir.ind = min(dir.ind, len(entries)-1)

	return entries
}
