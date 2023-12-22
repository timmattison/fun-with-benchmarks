package golang_printing

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkPrintln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("Output")
	}
}

func BenchmarkPrintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Printf("Output\n")
	}
}

func BenchmarkFprintfBufferWriteToNewEachTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer

		fmt.Fprintf(&buffer, "Output\n")
	}
}

func BenchmarkFprintfBufferWriteToReused(b *testing.B) {
	var buffer bytes.Buffer

	for i := 0; i < b.N; i++ {
		fmt.Fprintf(&buffer, "Output\n")
	}
}

func BenchmarkFprintfBufferWriteToReusedAndReset(b *testing.B) {
	var buffer bytes.Buffer

	for i := 0; i < b.N; i++ {
		fmt.Fprintf(&buffer, "Output\n")
		buffer.Reset()
	}
}

func BenchmarkFprintfBufferFmtPrintStringNewEachTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer

		fmt.Fprintf(&buffer, "Output\n")
	}
}

func BenchmarkFprintfBufferFmtPrintStringReused(b *testing.B) {
	var buffer bytes.Buffer

	for i := 0; i < b.N; i++ {
		fmt.Fprintf(&buffer, "Output\n")
	}
}

func BenchmarkFprintfBufferFmtPrintStringReusedAndReset(b *testing.B) {
	var buffer bytes.Buffer

	for i := 0; i < b.N; i++ {
		fmt.Fprintf(&buffer, "Output\n")
		buffer.Reset()
	}
}

func BenchmarkFmtPrintStringBuilderNewEachTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var output strings.Builder

		output.WriteString("Output\n")
	}
}

func BenchmarkFmtPrintStringBuilderReused(b *testing.B) {
	var output strings.Builder

	for i := 0; i < b.N; i++ {
		output.WriteString("Output\n")
	}
}

func BenchmarkFmtPrintStringBuilderReusedAndReset(b *testing.B) {
	var output strings.Builder

	for i := 0; i < b.N; i++ {
		output.WriteString("Output\n")
		output.Reset()
	}
}
