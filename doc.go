// doc.go

// Package procedural adds graphic primitives that can be used by azul3d graphics engine.
//
// Version v0.1.0 is compatable with package azul3d.org/v0.
//
// Current status of project is at http://github.com/go-azul3d/procedural
//
// Installation
//
// Instructions for Azul3d are at https://code.google.com/p/azul3d/ .
//  go get azul3d.org/v0
//  go get gopkg.in/go-azul3d/procedural.v0
//  import "gopkg.in/go-azul3d/procedural.v0"
//
// Benchmarks
//
// Note that recursive routines eat up time and memory so use with care.
//  [ `go test -bench . -benchmem` | done: 1m10.521025674s ]
// 	PASS
// 	BenchmarkCone4-4	 1000000	      2916 ns/op	     631 B/op	       5 allocs/op
// 	BenchmarkCone16-4	  500000	      5955 ns/op	    1253 B/op	       5 allocs/op
// 	BenchmarkCone64-4	  100000	     16173 ns/op	    3876 B/op	       5 allocs/op
// 	BenchmarkDice-4 	  500000	      4523 ns/op	    1382 B/op	       4 allocs/op
// 	BenchmarkSphereT0-4	  500000	      6104 ns/op	     838 B/op	       8 allocs/op
// 	BenchmarkSphereT1-4	  100000	     25421 ns/op	    2300 B/op	      12 allocs/op
// 	BenchmarkSphereT2-4	   10000	    140341 ns/op	   10114 B/op	      18 allocs/op
// 	BenchmarkSphereT3-4	    2000	   1133368 ns/op	   47312 B/op	      24 allocs/op
// 	BenchmarkSphereT4-4	     100	  12962877 ns/op	  241232 B/op	      46 allocs/op
// 	BenchmarkSphereT5-4	      10	 176748449 ns/op	 1186128 B/op	      61 allocs/op
// 	BenchmarkSphereT6-4	       1	2762356353 ns/op	 4716832 B/op	      76 allocs/op
// 	BenchmarkSphereT7-4	       1	43677949329 ns/op	19033360 B/op	      91 allocs/op
// 	ok  	github.com/go-azul3d/procedural	69.528s
//
package procedural
