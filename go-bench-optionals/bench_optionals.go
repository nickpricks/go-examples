// bench_optionals.go

package main

import (
	"fmt"
	"strings"
	"testing"
)

/**
 * doValue logs required params and any number of optional string params (value type).
 * @param a Required string
 * @param b Required string
 * @param optionals Optional strings (variadic)
 */
func doValue(a, b string, optionals ...string) string {
	// Simulate some work: join all optionals
	return fmt.Sprintf("a=%s b=%s optionals=%s", a, b, strings.Join(optionals, ","))
}

/**
 * doPointer logs required params and any number of optional string params (pointer type).
 * @param a Required string
 * @param b Required string
 * @param optionals Optional *string (variadic)
 */
func doPointer(a, b string, optionals ...*string) string {
	// Simulate some work: join all non-nil optionals
	var vals []string
	for _, s := range optionals {
		if s != nil {
			vals = append(vals, *s)
		}
	}
	return fmt.Sprintf("a=%s b=%s optionals=%s", a, b, strings.Join(vals, ","))
}

// Helper to get *string from string
func strp(s string) *string { return &s }

/**
 * Benchmark: doValue with no optionals
 */
func BenchmarkDoValueNone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doValue("Nick", "Kachroo")
	}
}

/**
 * Benchmark: doValue with 3 optionals
 */
func BenchmarkDoValueSome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doValue("Nick", "Kachroo", "niteshk@simsaw.com", "Architect", "India")
	}
}

/**
 * Benchmark: doPointer with no optionals
 */
func BenchmarkDoPointerNone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doPointer("Nick", "Kachroo")
	}
}

/**
 * Benchmark: doPointer with 3 optionals
 */
func BenchmarkDoPointerSome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doPointer("Nick", "Kachroo", strp("niteshk@simsaw.com"), strp("Architect"), strp("India"))
	}
}

/**
 * Benchmark: doPointer with 3 optionals, one nil
 */
func BenchmarkDoPointerSomeNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doPointer("Nick", "Kachroo", strp("niteshk@simsaw.com"), nil, strp("India"))
	}
}
