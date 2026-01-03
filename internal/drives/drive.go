package drives

import (
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/disk"
)

type Drive struct {
	Name       string  `json:"name"`  // device path (e.g. /dev/sda1, C:)
	Label      string  `json:"label"` // volume label (user visible)
	Type       string  `json:"type"`  // SSD / HDD / Removable / Unknown
	Size       string  `json:"size"`
	Format     string  `json:"format"`
	Mountpoint string  `json:"mountpoint"`
	Used       string  `json:"used"`
	Free       string  `json:"free"`
	Percent    float64 `json:"percent"`
}

// prettyNameFromDevice derives a friendly name from device path
func prettyNameFromDevice(dev string) string {
	// e.g. "/dev/sda1" -> "sda1" or "C:" remains "C:"
	if idx := strings.LastIndex(dev, "/"); idx != -1 {
		return dev[idx+1:]
	}
	return dev
}

// formatBytes converts bytes to a human-readable string (e.g., "1.2 GB")
func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	value := float64(bytes) / float64(div)
	units := []string{"KB", "MB", "GB", "TB", "PB", "EB"}
	if exp >= len(units) {
		return fmt.Sprintf("%.1f B", value*float64(div))
	}
	return fmt.Sprintf("%.1f %s", value, units[exp])
}

// DriveProvider interface
type DriveProvider interface {
	GetDrives() ([]Drive, error)
}

type Provider struct{}

func (Provider) GetDrives() ([]Drive, error) {

	// Get all partitions
	partitions, err := disk.Partitions(false)
	
	if err != nil {
		return nil, err
	}

	// Enrich partitions with label/type info
	var drives []Drive
	for _, p := range partitions {
		// Try to get usage (skip pseudo or inaccessible mounts)
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			// If mountpoint empty (sometimes on Windows) try usage on device path
			// but skip if still failing
			continue
		}

		label, devType := getVolumeInfo(p) // platform-specific implementation

		// If label empty, attempt to derive a nicer name from device
		displayName := p.Device
		if label != "" {
			displayName = label
		} else {
			// On linux devices like /dev/sda1 -> sda
			displayName = prettyNameFromDevice(p.Device)
		}

		drives = append(drives, Drive{
			Name:       p.Device, //Name = device identifier
			Label:      label,
			Type:       devType,
			Size:       formatBytes(usage.Total),
			Format:     p.Fstype,
			Mountpoint: p.Mountpoint,
			Used:       formatBytes(usage.Used),
			Free:       formatBytes(usage.Free),
			Percent:    usage.UsedPercent,
		})

		// optionally override Name to show friendly name:
		_ = displayName // unused here, but you can set drives[len-1].Name = displayName if desired
	}

	return drives, nil
}

func NewProvider() DriveProvider {
	return Provider{}
}
