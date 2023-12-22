# Fun with benchmarks

This repository contains a collection of benchmarks for various things I've worked on. When I really
want to know if it's worth it to optimize I put my tests here.

# Chapter 1: Is print really that slow?

TL;DR: Yes, it is.

Try it yourself: `./run-printing-benchmark.sh`

Results:

```
BenchmarkPrintln-8                                      1816780             641.0 ns/op
BenchmarkPrintf-8                                       1842417             618.5 ns/op
BenchmarkFprintfBufferWriteToNewEachTime-8             	19127950	        59.59 ns/op
BenchmarkFprintfBufferWriteToReused-8                  	52082672	        23.89 ns/op
BenchmarkFprintfBufferWriteToReusedAndReset-8          	56520682	        21.17 ns/op
BenchmarkFprintfBufferFmtPrintStringNewEachTime-8      	18596625	        61.30 ns/op
BenchmarkFprintfBufferFmtPrintStringReused-8           	52533654	        22.28 ns/op
BenchmarkFprintfBufferFmtPrintStringReusedAndReset-8   	56728758	        20.84 ns/op
BenchmarkFmtPrintStringBuilderNewEachTime-8            	86669625	        14.35 ns/op
BenchmarkFmtPrintStringBuilderReused-8                 	100000000	        12.32 ns/op
BenchmarkFmtPrintStringBuilderReusedAndReset-8         	82426843	        14.60 ns/op
```

I was skeptical and I was wrong. I have a program that needs to print out a lot of data. I was
doing individual `fmt.Println` calls for each line. And the program was pretty fast but I felt
that it could be faster. So I switched to a string builder on a whim and it was faster. But then
I needed the real info to make sure I wasn't crazy.

At ~12.5 ns/op for a reused string builder we're almost 50x faster than using `fmt.Printf`. A
big win.
