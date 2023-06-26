package entry

import (
	"github.com/shirou/gopsutil/v3/disk"
)

// GetMounts returns a slice of Mounts.
func GetMounts() ([]string, error) {
	parts, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}
	// Probably actually want the disk.PartitionStat structs
	mounts := make([]string, len(parts))
	for i, part := range parts {
		print(part.String())
		mounts[i] = part.Mountpoint
	}
	return mounts, nil
}
