package main

import (
	"fmt"
	"github.com/bwrightson/euler/problems"
	"time"
)

// elapsed prints the amount of time in seconds since t
func elapsed(t time.Time) {
	fmt.Printf("The program took %6.4f seconds to run.\n", time.Duration.Seconds(time.Since(t)))
}

func main() {
	var start = time.Now()
	fmt.Printf("The answer is %d.\n", problems.Solve001())
	elapsed(start)
}
