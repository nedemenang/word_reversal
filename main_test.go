package main

import (
	"testing"
)

func TestWordReversal(t *testing.T) {
	var tests = []struct {
		s        string
		expected string
	}{
		{"I love green flowers", "flowers green love I"},
		{"Australia is a country and continent surrounded by the Indian and Pacific oceans", "oceans Pacific and Indian the by surrounded continent and country a is Australia"},
	}
	for _, test := range tests {
		if result := WordReversal(test.s); result != test.expected {
			t.Error("Test Failed: Expected {} to equal {}, recieved: {}", test.s, test.expected, result)
		} else {
			t.Logf("Test passed")
		}
	}
}


func benchmarkWordReversal(s string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		WordReversal(s)
	}
}

func BenchmarkWordReversal1(b *testing.B) { benchmarkWordReversal("I love green flowers", b)}
func BenchmarkWordReversal2(b *testing.B) { benchmarkWordReversal("Australia is a country and continent surrounded by the Indian and Pacific oceans", b)}



func TestLetterReversal(t *testing.T) {
	var tests = []struct {
		s        string
		expected string
	}{
		{"I love green flowers", "srewolf neerg evol I"},
		{"Australia is a country and continent surrounded by the Indian and Pacific oceans", "snaeco cificaP dna naidnI eht yb dednuorrus tnenitnoc dna yrtnuoc a si ailartsuA"},
	}
	for _, test := range tests {
		if result := LetterReversal(test.s); result != test.expected {
			t.Error("Test Failed: Expected {} to equal {}, recieved: {}", test.s, test.expected, result)
		} else {
			t.Logf("Test passed")
		}
	}
}


func benchmarkLetterReversal(s string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		WordReversal(s)
	}
}

func BenchmarkLetterReversal1(b *testing.B) { benchmarkLetterReversal("I love green flowers", b)}
func BenchmarkLetterReversal2(b *testing.B) { benchmarkLetterReversal("Australia is a country and continent surrounded by the Indian and Pacific oceans", b)}
