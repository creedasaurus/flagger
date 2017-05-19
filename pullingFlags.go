// Some golang to download a bunch of flags, either serially, or using goroutines. For practice.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

const urlstrn string = "https://www.cia.gov/library/publications/resources/the-world-factbook/graphics/flags/large/"

var flags = [...]string{"aa", "ac", "ae", "af", "ag", "aj", "al", "am", "an", "ao", "aq", "ar", "as", "at", "au", "av", "ax",
	"ba", "bb", "bc", "bd", "be", "bf", "bg", "bh", "bk", "bl", "bm", "bn", "bo", "bp", "bq", "br", "bt", "bu", "bv", "bx", "by",
	"ca", "cb", "cc", "cd", "ce", "cf", "cg", "ch", "ci", "cj", "ck", "cm", "cn", "co", "cr", "cs", "ct", "cu", "cv", "cy",
	"da", "dj", "do", "dr", "dx",
	"ec", "ee", "eg", "ei", "ek", "en", "er", "es", "et", "ez",
	"fi", "fj", "fk", "fm", "fo", "fp", "fr", "fs",
	"ga", "gb", "gg", "gh", "gi", "gj", "gk", "gl", "gm", "gq", "gr", "gt", "gv", "gy",
	"ha", "hk", "hm", "ho", "hr", "hu",
	"ic", "id", "im", "in", "io", "ip", "ir", "is", "it", "iv", "iz",
	"ja", "je", "jm", "jn", "jo",
	"ke", "kg", "kn", "kr", "ks", "kt", "ku", "kv", "kz",
	"la", "le", "lg", "lh", "li", "lo", "ls", "lt", "lu", "ly",
	"ma", "mc", "md", "mg",
	"nc", "ne", "nf", "ng", "nh", "ni", "nl", "no", "np", "nr", "ns", "nu", "nz",
	"od",
	"pa", "pc", "pe", "pk", "pl", "pm", "po", "pp", "pu",
	"qa",
	"ri", "rm", "rn", "ro", "rp", "rq", "rs", "rw",
	"sa", "sb", "sc", "se", "sf", "sg", "sh", "si", "sk", "sl", "sm", "sn", "so", "sp", "st", "su", "sv", "sw", "sx", "sy", "sz",
	"tb", "td", "th", "ti", "tk", "tl", "tn", "to", "tp", "ts", "tt", "tu", "tv", "tw", "tx", "tz",
	"ug", "uk", "um", "up", "us", "uv", "uy", "uz",
	"vc", "ve", "vi", "vm", "vq", "vt",
	"wa", "wf", "wq", "ws", "wz",
	"ym",
	"za", "zi"}

func getAndSaveFlag(flg string) (size int64) {
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

func runGoLike(flagsCut []string) {
	// var bts uint64
	var wg sync.WaitGroup
	fmt.Println("got to runGoLike")
	for _, flg := range flagsCut {
		wg.Add(1)
		go func(flgs string) {
			defer wg.Done()
			getAndSaveFlag(flgs)
		}(flg)
	}
	wg.Wait()
}

func runSerially(flagsCut []string) (count int64) {
	var bts int64
	for _, elm := range flagsCut {
		bts += getAndSaveFlag(elm)
	}
	return bts
}

func getFlagsSkele(runType string, count int) {
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
	defer os.Chdir("..")

	fmt.Println("Runtype: " + runType)

	switch runType {
	case "serial":
		start := time.Now()
		byteCount = runSerially(flags[:count])
		fmt.Println(byteCount)
		fmt.Println()
		elapsed := time.Since(start)
		fmt.Println("Serial took: ", elapsed)
		break
	case "goroutine":
		start := time.Now()
		runGoLike(flags[:count])
		elapsed := time.Since(start)
		fmt.Println("Goroutines took: ", elapsed)
		break
	default:
		fmt.Println("Not a proper Run Type. Use: serial | goroutine")
		break
	}

}

func main() {

	getFlagsSkele("goroutine", len(flags))
	// getAndSaveFlag("us")

}
