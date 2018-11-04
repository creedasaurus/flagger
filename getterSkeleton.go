package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// getFlagsUsingRunType - A skeleton structure for running each of the different run styles
func getFlagsUsingRunType(runType string, count int) {
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

type Flag struct {
	Data     []byte
	Filename string
}

func GetFlag(flg string) (*Flag, error) {
	flag := &Flag{Filename: fmt.Sprintf("%s-lgflag.gif", flg)}
	// fmt.Println("GETting flag:", flag.Filename)
	response, err := http.Get(urlstrn + flag.Filename)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Non 200 response")
	}

	var data bytes.Buffer
	_, err = io.Copy(&data, response.Body)
	if err != nil {
		return nil, err
	}

	flag.Data = data.Bytes()
	return flag, nil
}

func SaveFlag(flag *Flag) error {
	// fmt.Println("Saving Flag:", flag.Filename)
	outFile, err := os.Create(flag.Filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = outFile.Write(flag.Data)
	if err != nil {
		return err
	}

	outFile.Sync()
	// fmt.Println("saved!")
	return nil
}

// GetAndSaveFlag - This is the single function that will make the
func GetAndSaveFlag(flg string) (size int64) {
	// defer fmt.Println("Cleaning up", flg)
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
