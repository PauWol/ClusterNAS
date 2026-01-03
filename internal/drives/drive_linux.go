package drives

import (
	"fmt"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/shirou/gopsutil/disk"
)

// getVolumeInfo for linux uses lsblk(1) for LABEL and /sys/block/<dev>/queue/rotational for SSD/HDD
func getVolumeInfo(p disk.PartitionStat) (string, string) {
	// 1) LABEL via lsblk
	label := ""
	out, err := exec.Command("lsblk", "-no", "LABEL", p.Device).Output()
	if err == nil {
		label = strings.TrimSpace(string(out))
	}

	// 2) derive base device: /dev/sda1 -> sda
	base := path.Base(p.Device)
	// strip partition numbers (e.g. sda1 -> sda, nvme0n1p1 -> nvme0n1)
	re := regexp.MustCompile(`^([a-zA-Z0-9]+?)\d*$`)
	m := re.FindStringSubmatch(base)
	if len(m) >= 2 {
		base = m[1]
	}

	// try rotational flag
	rotPath := fmt.Sprintf("/sys/block/%s/queue/rotational", base)
	out, err = exec.Command("cat", rotPath).Output()
	if err == nil {
		val := strings.TrimSpace(string(out))
		if val == "0" {
			return label, "SSD"
		}
		if val == "1" {
			return label, "HDD"
		}
	}

	// fallback removable check
	remPath := fmt.Sprintf("/sys/block/%s/removable", base)
	out, err = exec.Command("cat", remPath).Output()
	if err == nil && strings.TrimSpace(string(out)) == "1" {
		return label, "Removable"
	}

	// final fallback
	return label, "Unknown"
}
