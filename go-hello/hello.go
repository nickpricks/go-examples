package main

import (
	"fmt"
	"os"

	rscQuote "rsc.io/quote"
	rscSampler "rsc.io/sampler"
)

func main() {
	// Temporarily set the LANG environment variable
	os.Setenv("LANG", "hi_IN.UTF-8") // for english use en_US.UTF-8

	// Verify the setting
	lang := os.Getenv("LANG")
	fmt.Println("LANG is set to:", lang)

	fmt.Println("Bye World.", rscSampler.Hello())
	fmt.Printf("glass => %s", rscQuote.Glass())
	fmt.Println("go => ", rscQuote.Go())
	fmt.Println("hello => ", rscQuote.Hello())
	fmt.Println("opt => ", rscQuote.Opt())
}
