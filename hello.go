package main

import (
	"fmt"
	"os"

	rsc "rsc.io/quote"
	rscSampler "rsc.io/sampler"
)

func main() {
	// Temporarily set the LANG environment variable
	os.Setenv("LANG", "hi_IN.UTF-8") // for english use en_US.UTF-8

	// Verify the setting
	lang := os.Getenv("LANG")
	fmt.Println("LANG is set to:", lang)

	fmt.Println("Hello World.", rscSampler.Hello())
	fmt.Println("glass => ", rsc.Glass())
	fmt.Println("go => ", rsc.Go())
	fmt.Println("hello => ", rsc.Hello())
	fmt.Println("opt => ", rsc.Opt())
}
