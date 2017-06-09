package main

// RunGoLike - runs the gets with goroutines
// TODO: figure out if returning values from goroutines is ok
func RunGoLike(flagsCut []string, ch chan int64) {
	// var bts uint64
	// var wg sync.WaitGroup

	for _, flg := range flagsCut {
		// wg.Add(1)
		go func(flgs string) {
			// defer wg.Done()
			GetAndSaveFlag(flgs)
		}(flg)
	}

	// fmt.Println("before wait()")
	// wg.Wait()
	// fmt.Println("after wait()")
}

// RunSerially - gets all the flags one-by-one
func RunSerially(flagsCut []string) (count int64) {
	var bts int64
	for _, elm := range flagsCut {
		bts += GetAndSaveFlag(elm)
	}
	return bts
}
