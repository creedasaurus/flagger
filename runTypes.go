package main

import (
	"context"
	"log"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

// RunGoLike - runs the gets with goroutines
// TODO: figure out if returning values from goroutines is ok
func RunGoLike(flagsCut []string) int64 {
	producer := make(chan string, len(flagsCut))
	for _, flag := range flagsCut {
		producer <- flag
	}
	close(producer)

	var bts uint64
	gr, _ := errgroup.WithContext(context.Background())
	workers := 6
	for i := 0; i < workers; i++ {
		//workerNum := i
		gr.Go(func() error {
			for fl := range producer {
				//fmt.Printf("worker %d, getting flag: %s\n", workerNum, fl)
				siz := GetAndSaveFlag(fl)
				atomic.AddUint64(&bts, uint64(siz))
			}
			return nil
		})
	}

	if err := gr.Wait(); err != nil {
		log.Fatal(err)
	}
	return int64(bts)
}

// RunSerially - gets all the flags one-by-one
func RunSerially(flagsCut []string) (count int64) {
	var bts int64
	for _, elm := range flagsCut {
		//fmt.Println("getting flag: ", elm)
		bts += GetAndSaveFlag(elm)
	}
	return bts
}
