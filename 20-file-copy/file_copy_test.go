package main

import "testing"

func BenchmarkReadFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		useReadFile()
	}
}

func BenchmarkReadAll(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		useReadAll()
	}
}

func BenchmarkIteration(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		useIteration()
	}
}

func BenchmarkCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		useCopy()
	}
}

func BenchmarkBufIO(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		useBufIO()
	}
}
