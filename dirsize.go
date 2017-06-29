package main

import (
	"fmt"
	"os"
	"path/filepath"
	"flag"
)

const (
	BYTE     = 1.0
	KILOBYTE = 1024.0 * BYTE
	MEGABYTE = 1024.0 * KILOBYTE
	GIGABYTE = 1024.0 * MEGABYTE
	TERABYTE = 1024.0 * GIGABYTE
)

func main() {
	var size int64
	var fileCount int64
	var err error

	dirToSize := flag.String("d", "", "Path you wish to return the size of.")
	verbose := flag.Bool("v", false, "Output verbode directory size details.")
	flag.Parse()

	searchDir := *dirToSize
	verboseOutput := *verbose

	if searchDir == "" {
		// Quit out the program
		flag.PrintDefaults()
		os.Exit(0)
	}

	if !DirExists(searchDir) {
		fmt.Println("The supplied directory does not exist.")
		fmt.Printf("Directory entered: %s", searchDir)
	} else {
		size, fileCount, err = DirSize(searchDir, verboseOutput)
		if err != nil {
			fmt.Println(err)
		} else {
			outputStats(size, fileCount)
		}
	}

}

func DirExists(path string) (bool) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func DirSize(path string, opt bool) (int64, int64, error) {
    var totalSize int64
    var dirSize int64
    var fileCount int64
    m := make(map[string]int64)	// map to store directory sizes in
    var dirName string

    err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            totalSize += info.Size()
            fileCount += 1
            dirSize += info.Size()
            m[dirName] += dirSize
        } else {
        	dirName = p
        	_, ok := m[p]
    		if  ok == false {
    			m[p] = 0
    		} 
        }
        dirSize = 0
        return err
    })

    if opt {
    	for key, value := range m {
    		fmt.Printf("Dir: %s \t Size: %d Bytes \n", key, value)
		}
    }
    return totalSize, fileCount, err
}

func outputStats(bytes int64, fcount int64) {
	var value float64
	var unit string

	switch {
	case bytes >= TERABYTE:
		unit = "TB"
		value = float64(bytes) / TERABYTE
	case bytes >= GIGABYTE:
		unit = "GB"
		value = float64(bytes) / GIGABYTE
	case bytes >= MEGABYTE:
		unit = "MB"
		value = float64(bytes) / MEGABYTE
	case bytes >= KILOBYTE:
		unit = "KB"
		value = float64(bytes) / KILOBYTE
	default:
		unit = "Bytes"
		value = float64(bytes)
	}

	fmt.Printf("Dir size: %.2f%s \n", value, unit)
	fmt.Printf("File count: %d\n", fcount)
}
