//go:build windows

package storage

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type StorageInfo struct {
	FreeSpace      uint64
	TotalSpace     uint64
	AvailableSpace uint64
}

func GetStorageInfo() (StorageInfo, error) {
	dll := windows.NewLazyDLL("kernel32.dll")
	proc := dll.NewProc("GetDiskFreeSpaceExW")
	info := StorageInfo{}
	_, _, err := proc.Call(uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("."))),
		uintptr(unsafe.Pointer(&info.AvailableSpace)),
		uintptr(unsafe.Pointer(&info.TotalSpace)),
		uintptr(unsafe.Pointer(&info.FreeSpace)))

	// err always returns non nil from proc.Call so check that the values out are ok
	// and return the error if the values are incorrect
	if info.FreeSpace == 0 || info.TotalSpace == 0 || info.AvailableSpace == 0 {
		return info, err
	}

	return info, nil
}
