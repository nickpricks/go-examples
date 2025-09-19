// main.go

package main

import (
	"fmt"
)

// logIndices logs the current outer and inner loop indices.
func logIndices(outer, inner int) {
	fmt.Printf("Outer: %d, Inner: %d\n", outer, inner)
}

// caseContinue demonstrates skipping a log when a condition is met using 'continue'.
// Expected: For each outer loop, if inner == 5, skip logging that pair.
func caseContinue() {
	fmt.Println("Case 1: continue (skip inner==5)")
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if j == 5 {
				continue // Skip logging when inner == 5
			}
			logIndices(i, j)
		}
	}
	fmt.Printf("Case 1 done.\n.\n.\n")
}

// caseBreak demonstrates breaking out of the inner loop when a condition is met.
// Expected: For each outer loop, when inner == 5, stop logging further inner values for that outer value.
func caseBreak() {
	fmt.Println("Case 2: break (stop inner loop at inner==5)")
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if j == 5 {
				break // Stop inner loop when inner == 5
			}
			logIndices(i, j)
		}
	}
	fmt.Printf("Case 2 done.\n.\n.\n")
}

// caseReturn demonstrates returning from the function when a condition is met.
// Expected: When outer == 3 and inner == 5, exit the function immediately (no further logs).
func caseReturn() {
	fmt.Println("Case 3: return (exit all loops at outer==3, inner==5)")
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i == 3 && j == 5 {
				fmt.Println("Return triggered at Outer: 3, Inner: 5")
				return // Exit the function immediately
			}
			logIndices(i, j)
		}
	}
	fmt.Printf("Case 3 done.\n.\n.\n")
}

func main() {
	caseContinue()
	caseBreak()
	caseReturn()
}
