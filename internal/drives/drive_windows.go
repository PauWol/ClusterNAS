package drives

import (
	"strings"

	"github.com/gentlemanautomaton/volmgmt/volume"
	"github.com/shirou/gopsutil/disk"
)

// getVolumeInfo on windows uses volmgmt to get label and removable/fixed info.
// returns (label, type)
func getVolumeInfo(p disk.PartitionStat) (string, string) {
	// Try to open by partition device string returned by gopsutil.
	// p.Device may be "C:" or "\\\\.\\Volume{GUID}\\" or similar.
	candidates := []string{p.Device}

	// if p.Device is like "C:" add \\.\C:
	if strings.HasSuffix(p.Device, ":") {
		candidates = append(candidates,  "\\.\\" + p.Device)
	}

	// try the candidates
	for _, c := range candidates {
		v, err := volume.New(c)
		if err != nil {
			continue
		}
		defer v.Close()

		// Label() gives volume label (if any)
		label, _ := v.Label()

		// RemovableMedia() is straightforward
		if v.RemovableMedia() {
			return label, "Removable"
		}

		// DeviceType and BusType may provide hints - map common cases
		// DeviceType() returns a numeric code; map unknown -> Fixed
		// For simplicity map everything non-removable -> "Fixed"
		return label, "Fixed"
	}

	// fallback: no info found
	return "", "Unknown"
}
