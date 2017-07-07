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

type Params struct {
    searchDir string
    verboseOutput bool
}

func main() {

    dirToSize := flag.String("d", "", "Path you wish to return the size of.")
    verbose := flag.Bool("v", false, "Output verbose directory size details.")
    flag.Parse()

    clParam := Params{*dirToSize, *verbose}

    if clParam.searchDir == "" {
        // Quit out the program
        flag.PrintDefaults()
        os.Exit(0)
    }

    if DirExists(clParam.searchDir) {
        DirSize(clParam)
    }

}

func DirExists(path string) (bool) {
    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        fmt.Println("The supplied directory does not exist.")
        fmt.Printf("Directory entered: %s", path)
        return false
    } else {
        return true
    }
}

func DirSize(clParams Params) {
    var totalSize int64
    var fileCount int64
    var dirName string
    var dirCount int64

    dirMap := make(map[string]int64) // map to store directory sizes in

    err := filepath.Walk(clParams.searchDir, func(pth string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            totalSize += info.Size()
            fileCount += 1
            dirMap[dirName] += info.Size()
        } else {
            dirCount += 1
            dirName = pth
            _, ok := dirMap[pth]
            if  ok == false {
                dirMap[pth] = 0
            } 
        }
        return err
    })

    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }   
    
    if clParams.verboseOutput {
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

    return value, unit
}
