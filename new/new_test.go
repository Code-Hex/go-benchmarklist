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

func BenchmarkNew(b *testing.B) {
	b.ResetTimer()
	t := Test{
		Title:   "Japanese",
		Score:   80,
		Average: 64.5,
	}
	for i := 0; i < b.N; i++ {
		p := new(Person)
		p.Name = "CodeHex"
		p.Age = 10
		p.Belong = "College"
		p.Rank = 'A'
		p.Test = t
	}
}

func BenchmarkNewNormal(b *testing.B) {
	b.ResetTimer()
	t := Test{
		Title:   "Japanese",
		Score:   80,
		Average: 64.5,
	}
	for i := 0; i < b.N; i++ {
		p := &Person{}
		p.Name = "CodeHex"
		p.Age = 10
		p.Belong = "College"
		p.Rank = 'A'
		p.Test = t
	}
}

func BenchmarkNewSimple(b *testing.B) {
	b.ResetTimer()
	t := Test{
		Title:   "Japanese",
		Score:   80,
		Average: 64.5,
	}
	for i := 0; i < b.N; i++ {
		_ = &Person{
			Name:   "CodeHex",
			Age:    10,
			Belong: "College",
			Rank:   'A',
			Test:   t,
		}
	}
}
