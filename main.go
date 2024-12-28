package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	"time"
	"unicode"
)

var wg sync.WaitGroup

func labeler(str string) {
	defer wg.Done()
	limit := 2
	for count := 0; count < limit; count++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Count %d from %s\n", count+1, str)
	}
}

const (
	ALPHABET string = "abcdefghijklmnopqrstuvwxyz"
)

func main() {
	start := time.Now()
	wg.Add(len(ALPHABET))
	fmt.Println("Waiting to finish.")
	for _, r := range ALPHABET {
		buf := &bytes.Buffer{}
		buf.WriteRune(unicode.ToUpper(r))
		go labeler(buf.String())
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("Terminating program.\nElapsed period: %v", end.Sub(start))
}
