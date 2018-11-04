package main

import (
	"fmt"
)

// RunGoLike - runs the gets with goroutines
// TODO: figure out if returning values from goroutines is ok
func RunGoLike(flagsCut []string) {
	receivedFlags := make(chan *Flag, len(flagsCut))
	// errChan := make(chan error, len(flagsCut))
	for _, flag := range flagsCut {
		go func(flg string) {
			flagData, err := GetFlag(flg)
			if err != nil {
				fmt.Println(err)
				return
			}
			receivedFlags <- flagData
		}(flag)
	}

	for i := 0; i < len(flagsCut); i++ {
		flagType := <-receivedFlags
		go func(flg *Flag) {
			err := SaveFlag(flg)
			if err != nil {
				fmt.Println(err)
				return
			}
		}(flagType)
	}
}

// RunSerially - gets all the flags one-by-one
func RunSerially(flagsCut []string) (count int64) {
	var bts int64
	for _, elm := range flagsCut {
		bts += GetAndSaveFlag(elm)
	}
	return bts
}
