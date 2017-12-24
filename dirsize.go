package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	byte     = 1.0
	kilobyte = 1024.0 * byte
	megabyte = 1024.0 * kilobyte
	gigabyte = 1024.0 * megabyte
	terabyte = 1024.0 * gigabyte
)

type params struct {
	searchdir     string
	verboseoutput bool
}

func main() {

	dirToSize := flag.String("d", "", "Path you wish to return the size of.")
	verbose := flag.Bool("v", false, "Output verbose directory size details.")
	flag.Parse()

	clParam := params{*dirToSize, *verbose}

	if clParam.searchdir == "" {
		// Quit out the program
		flag.PrintDefaults()
		os.Exit(0)
	}

	if dirExists(clParam.searchdir) {
		dirSize(clParam)
	}

}

func dirExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Println("The supplied directory does not exist.")
		fmt.Printf("Directory entered: %s", path)
		return false
	}
	return true
}

func dirSize(clParams params) {
	var totalSize int64
	var fileCount int64
	var dirName string
	var dirCount int64

	dirMap := make(map[string]int64) // map to store directory sizes in

	err := filepath.Walk(clParams.searchdir, func(pth string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			totalSize += info.Size()
			fileCount++
			dirMap[dirName] += info.Size()
		} else {
			dirCount++
			dirName = pth
			_, ok := dirMap[pth]
			if ok == false {
				dirMap[pth] = 0
			}
		}
		return err
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if clParams.verboseoutput {
		for key, value := range dirMap {
			dSize, sizeUnit := outputStats(value)
			fmt.Printf("Dir: %s \t Size: %.2f %s \n", key, dSize, sizeUnit)
		}
	}

	dSize, sizeUnit := outputStats(totalSize)
	fmt.Printf("Dir size  : %.2f%s \n", dSize, sizeUnit)
	fmt.Printf("Dir Count : %d\n", dirCount)
	fmt.Printf("File count: %d\n", fileCount)

}

func outputStats(bytes int64) (float64, string) {
	var value float64
	var unit string

	switch {
	case bytes >= terabyte:
		unit = "TB"
		value = float64(bytes) / terabyte
	case bytes >= gigabyte:
		unit = "GB"
		value = float64(bytes) / gigabyte
	case bytes >= megabyte:
		unit = "MB"
		value = float64(bytes) / megabyte
	case bytes >= kilobyte:
		unit = "KB"
		value = float64(bytes) / kilobyte
	default:
		unit = "Bytes"
		value = float64(bytes)
	}

	return value, unit
}
