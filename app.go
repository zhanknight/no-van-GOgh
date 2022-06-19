package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// This is a starry night ascii generator.

var numLines int
var words = strings.Fields("~ ~ ~ * ~ ~ * ~ ~ / * / * O ~ ~ ~ ~")
var wg sync.WaitGroup

func main() {
	numLines = 8
	for i := 0; i < numLines; i++ {
		// without spawning a go routine, each line takes 100ms
		// with routines, the whole thing takes 100ms no matter how many lines
		// because we create a new routine for each line and they
		// start and finish almost simultaneously. Cool.
		wg.Add(1)
		go func() {
			output := lineGenerator()
			fmt.Println(output)
			wg.Done()
		}()
	}
	// Print the countryside
	wg.Wait()
	fmt.Print("/\\   /\\  /\\/\\ __  /\\   __  /\\/\\ __\n||   ||  ||||/  \\ ||  /  \\ ||||/  \\\n")
}

// generate the sky by shuffling the string declared up top
func lineGenerator() string {
	shuffledWords := words
	rand.Shuffle(len(shuffledWords), func(i, j int) {
		shuffledWords[i], shuffledWords[j] = shuffledWords[j], shuffledWords[i]
	})
	var output = strings.Join(shuffledWords, " ")
	// adding a delay to simulate a task that takes time
	time.Sleep(100 * time.Millisecond)
	return output
}

// TODO
// make all routines add their line to a var/struct/something then print that whole thing
// find a color package, import, use
// make the countryside shuffle too
