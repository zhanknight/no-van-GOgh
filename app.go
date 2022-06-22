package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

// This is a starry night ascii generator.

const (
	words    string = "~ ~ ~ * ~ ~ * ~ ~ / * / * O ~ ~ ~ ~"
	numLines int    = 8
)

var (
	wg           sync.WaitGroup
	mutex        sync.Mutex
	finalString  string
	stringPieces []string
)

func main() {
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
	output := strings.Join(shuffledWords, " ")
	// adding a delay to simulate a task that takes time
	// time.Sleep(100 * time.Millisecond)
	return output
}

// TODO
// find a color package, import, use
// make the countryside shuffle too
