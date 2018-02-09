package main

// Benchmark functions
import "testing"

const n = 8

//BenchmarkrandomString - masking improved with source
func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString(n)
	}
}

//BenchmarkRunes - runes
func BenchmarkRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringRunes(n)
	}
}

//BenchmarkBytes - bytes
func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytes(n)
	}
}

//BenchmarkBytesRmndr - Remainder
func BenchmarkBytesRmndr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesRmndr(n)
	}
}

//BenchmarkBytesMask - masking
func BenchmarkBytesMask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMask(n)
	}
}

//BenchmarkBytesMaskImpr - masking improved
func BenchmarkBytesMaskImpr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImpr(n)
	}
}

//BenchmarkBytesMaskImprSrc - masking improved with source
func BenchmarkBytesMaskImprSrc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImprSrc(n)
	}
}

//BenchmarkRandASCIIBytes - masking improved with source
func BenchmarkRandASCIIBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandASCIIBytes(n)
	}
}
