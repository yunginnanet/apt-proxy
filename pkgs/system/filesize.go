package system

import (
	"math"
	"os"
	"path/filepath"
	"strconv"

	"golang.org/x/sys/unix"
)

func DiskAvailable() (int64, error) {
	var stat unix.Statfs_t
	wd, err := os.Getwd()

	if err != nil {
		return 0, err
	}

	unix.Statfs(wd, &stat)
	return int64(stat.Bavail * uint64(stat.Bsize)), nil
}

// https://stackoverflow.com/questions/32482673/how-to-get-directory-total-size
func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// https://hakk.dev/docs/golang-convert-file-size-human-readable/
func fileSizeRound(val float64, roundOn float64, places int) float64 {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}

func HumanFileSize(input int64) string {
	size := float64(input)

	var units = []string{
		0: "B",
		1: "KB",
		2: "MB",
		3: "GB",
		4: "TB",
	}

	level := math.Log(size) / math.Log(1024)
	fileSize := fileSizeRound(math.Pow(1024, level-math.Floor(level)), .5, 2)
	unit := units[int(math.Floor(level))]

	return strconv.FormatFloat(fileSize, 'f', -1, 64) + " " + string(unit)
}
