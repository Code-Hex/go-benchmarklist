package main

import "testing"

type Person struct {
	Name   string
	Age    int
	Belong string
}

func BenchmarkNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := new(Person)
		p.Name = "CodeHex"
		p.Age = 10
		p.Belong = "College"
	}
}

func BenchmarkNewNormal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := &Person{}
		p.Name = "CodeHex"
		p.Age = 10
		p.Belong = "College"
	}
}

func BenchmarkNewSimple(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = &Person{
			Name:   "CodeHex",
			Age:    10,
			Belong: "College",
		}
	}
}
