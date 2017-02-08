package main

import (
	"regexp"
	"strings"
	"testing"
)

const TEXT = "Hello, Hello, Hello!!"

func BenchmarkReplacer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strings.NewReplacer("Hello", "World").Replace(TEXT)
	}
}

func BenchmarkReplacerOnce(b *testing.B) {
	b.ResetTimer()
	r := strings.NewReplacer("Hello", "World")
	for i := 0; i < b.N; i++ {
		r.Replace(TEXT)
	}
}

func BenchmarkStrings(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strings.Replace(TEXT, "Hello", "World", -1)
	}
}

func BenchmarkRegex(b *testing.B) {
	b.ResetTimer()
	reg := regexp.MustCompile(`Hello`)
	for i := 0; i < b.N; i++ {
		reg.ReplaceAllString(TEXT, "World")
	}
}
