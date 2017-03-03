package main

import "testing"

type Person struct {
	Name   string
	Age    int
	Belong string
	Rank   rune
	Test
}

type Test struct {
	Title   string
	Score   uint
	Average float64
}

func New() *Person {
	return &Person{
		Name:   "CodeHex",
		Age:    10,
		Belong: "College",
		Rank:   'A',
		Test: Test{
			Title:   "Japanese",
			Score:   80,
			Average: 64.5,
		},
	}

}

func BenchmarkNormal(b *testing.B) {
	b.ResetTimer()
	people := make(map[int]*Person)
	for i := 0; i < b.N; i++ {
		people[i] = New()
	}
}

func BenchmarkAllocated(b *testing.B) {
	b.ResetTimer()
	people := make(map[int]*Person, b.N)
	for i := 0; i < b.N; i++ {
		people[i] = New()
	}
}
