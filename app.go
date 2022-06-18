package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// This is a starry night ascii generator.

var numLines int
var words = strings.Fields("~ ~ ~ * ~ ~ * ~ ~ / * / * O ~ ~ ~ ~")

func main() {

	// fmt.Println("How many lines should we print?")
	// fmt.Scanln(&numLines)
	// don't ask for now, just generate 8 lines of starry night.
	numLines = 8
	for i := 0; i < numLines; i++ {
		// without spawning a go routine, each line takes 100ms
		// with routines, the whole thing takes 100ms no matter how many lines
		// because we create a new routine for each line and they
		// start and finish almost simultaneously. Cool.
		go func() {
			output := lineGenerator()
			fmt.Println(output)
		}()
		// output := lineGenerator()
		// fmt.Println(output)
	}
	time.Sleep(150 * time.Millisecond)
	// Print the countryside
	// need this to wait for routines to finish, using simple sleep for now
	fmt.Println("/\\   /\\  /\\/\\ __  /\\   __  /\\/\\ __")
	fmt.Println("||   ||  ||||/  \\ ||  /  \\ ||||/  \\")
}

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
// make the countryside wait for all routines to complete before printing itself
// make all routines add their line to a var/struct/something then print that whole thing
// find a color package, import, use
