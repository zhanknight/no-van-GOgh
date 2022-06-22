package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// This is a starry night ascii generator.
const words string = "~ ~ ~ * ~ ~ * ~ ~ / * / * O ~ ~ ~ ~"
const numLines int = 8

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var finalString string
	var stringPieces []string
	for i := 0; i < numLines; i++ {
		wg.Add(1)
		go func() {
			output := lineGenerator()
			mutex.Lock()
			stringPieces = append(stringPieces, output)
			mutex.Unlock()
			wg.Done()
		}()
	}
	// Print the countryside
	wg.Wait()
	finalString = strings.Join(stringPieces, "\n")
	fmt.Print(finalString)
	fmt.Print("\n/\\   /\\  /\\/\\ __  /\\   __  /\\/\\ __\n||   ||  ||||/  \\ ||  /  \\ ||||/  \\\n")
}

// generate the sky by shuffling the string declared up top
func lineGenerator() string {
	shuffledWords := strings.Fields(words)
	rand.Shuffle(len(shuffledWords), func(i, j int) {
		shuffledWords[i], shuffledWords[j] = shuffledWords[j], shuffledWords[i]
	})
	var output = strings.Join(shuffledWords, " ")
	// adding a delay to simulate a task that takes time
	time.Sleep(100 * time.Millisecond)
	return output
}

// TODO
// find a color package, import, use
// make the countryside shuffle too
