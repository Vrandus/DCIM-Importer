package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

type filePath struct {
	name string
	path string
}

func main() {
	var files []filePath
	var totalBytes int64

	if len(os.Args) < 3 {
		fmt.Println("Missing arguments")
		fmt.Println("Usage: import <srcPath> <dstPath>")
		return
	}
	srcPath := os.Args[1]

	err := filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fi := filePath{name: info.Name(), path: path}
			files = append(files, fi)
			log.Printf("scanning: " + path)
			totalBytes += info.Size()
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	var totalInGB = float32(totalBytes) / 1000000000
	var transferredBytes int64
	var timeSpent float32
	var totalTime float32
	var transferSpeed float32
	var existingFile int64
	then := time.Now().UnixNano()
	for i, file := range files {
		fileInfo, err := os.Stat(file.path)
		if err != nil {
			log.Fatal(err)
		}

		clear()
		photoDate := fileInfo.ModTime()
		dstPath := os.Args[2] + strconv.Itoa(photoDate.Year()) + "/" + strconv.Itoa(int(photoDate.Month())) + " - " + photoDate.Month().String() + "/"
		timeLeft := (float32(totalBytes-transferredBytes) / 1000000) / transferSpeed
		timeElapsed := (time.Now().UnixNano() - then) / 1000000000

		fmt.Printf("Copying %d of %d\n", i+1, len(files))
		fmt.Printf("%.3f GB / %.3f GB @ %.3f MB/s\n", float32(transferredBytes)/1000000000, totalInGB, transferSpeed)
		fmt.Printf("Time Elapsed: %02d : %02d : %02d; Time Left: %02d : %02d : %02d\n", timeElapsed/3600, timeElapsed/60%60, timeElapsed%60, int64(timeLeft)/3600, int64(timeLeft)/60%60, int64(timeLeft)%60)
		fmt.Println(dstPath + file.name)
		transferredBytes += fileInfo.Size()

		os.MkdirAll(dstPath, 0700)
		timeSpent = copyFile(file.path, dstPath+file.name)
		if timeSpent == -1 {
			existingFile += fileInfo.Size()
		}
		if timeSpent != -1 {
			totalTime += timeSpent
			transferSpeed = (float32(transferredBytes-existingFile) / 1000000) / totalTime
		}
	}

}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func copyFile(src string, dst string) float32 {
	_, err := os.Stat(dst)
	if err == nil {
		log.Printf("File already exists.")
		return -1
	}
	sourceFile, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer destinationFile.Close()

	then := time.Now().UnixNano()
	bytesWritten, err := io.Copy(destinationFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now().UnixNano()

	var timeSpent float32 = (float32(now - then)) / 1000000000

	log.Printf("Copied %d bytes in %.3f seconds", bytesWritten, timeSpent)

	err = destinationFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
	return timeSpent
}
