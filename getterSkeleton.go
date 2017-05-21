package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// GetFlagsSkele - A skeleton structure for running each of the different run styles
func GetFlagsSkele(runType string, count int) {
	var byteCount int64

	if count < 1 && count > len(flags) {
		fmt.Println("Not a proper number of flags to get")
		return
	}

	// Create directory name and move into it
	if _, err := os.Stat(runType); os.IsNotExist(err) {
		os.Mkdir(runType, 0755)
	}
	os.Chdir(runType)
	// Defer the backing out of that new directory
	defer os.Chdir("..")

	fmt.Println("Runtype: " + runType)
	switch runType {

	case "serial":
		start := time.Now()
		byteCount = RunSerially(flags[:count])
		fmt.Println(byteCount)
		fmt.Println()
		elapsed := time.Since(start)
		fmt.Println("Serial took: ", elapsed)
		break

	case "goroutine":
		start := time.Now()
		RunGoLike(flags[:count])
		elapsed := time.Since(start)
		fmt.Println("Goroutines took: ", elapsed)
		break

	default:
		fmt.Println("Not a proper Run Type. Use: serial | goroutine")
		break
	}
}

// GetAndSaveFlag - This is the single function that will make the
// Get() call and save it
func GetAndSaveFlag(flg string) (size int64) {
	flgstrg := flg + "-lgflag.gif"
	outfile, err := os.Create(flgstrg)
	if err != nil {
		fmt.Println("Error while making file")
		return
	}
	defer outfile.Close()

	resp, err := http.Get(urlstrn + flgstrg)
	if err != nil {
		fmt.Println("Error in Get")
		return
	}
	defer resp.Body.Close()

	bts, err := io.Copy(outfile, resp.Body)
	if err != nil {
		fmt.Println("error copying file to os")
		return
	}
	return bts
}
