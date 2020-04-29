# Playing with Golang -- Downloading Flags

I decided I wanted to learn a little bit of golang, just for kicks. Looks like a kind of cool language. I tried writing this really small program that would just pull down a ton of these little flag gifs. 

#### To run:
```sh
go run ./...
```

Current results on 4-29-2020:
```
➜  go_get_flags git:(master) ✗ go run ./...
Runtype: serial
3091147 bytes
Serial took:  19.27957685s
Runtype: goroutine
3091147 bytes
Goroutines took:  3.312755528s
```

This was big improvement on the first version, but really I have barely scratched the surface. I basically didn't use
any go idioms when I first wrote this so I'll have to revisit it sometime and clean it up for fun.

