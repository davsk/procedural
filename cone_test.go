// cone_test.go

package procedural

import (
	"fmt"
	"testing"
)

func ExampleCone() {
	pyramid := Cone(1.0, 4, true)
	hat := Cone(1.0, 8, false)
	cone := Cone(1.0, 32, true)
	fmt.Println(len(pyramid.Indices))
	fmt.Println(len(hat.Indices))
	fmt.Println(len(cone.Indices))
	fmt.Printf("pyramid := %v\n", pyramid)
	fmt.Printf("vertices := %v\n", pyramid.Vertices)
	fmt.Printf("indices := %v\n", pyramid.Indices)
	// Output:
	// 24
	// 24
	// 192
	// pyramid := &{{{0 0} 0 0 0 0} <nil> false false false {Vec3(X=0.000000, Y=0.000000, Z=0.000000) Vec3(X=0.000000, Y=0.000000, Z=0.000000)} [0 4 1 0 1 2 0 2 3 0 3 4 5 4 1 5 1 2 5 2 3 5 3 4] false [{0 0 1} {0 1 -1} {1 6.123234e-17 -1} {1.2246469e-16 -1 -1} {-1 -1.8369701e-16 -1} {0 0 -1}] false [] false [{[{0 0} {1 0} {0 1} {1 1} {0 0} {0 0}] false}]}
	// vertices := [{0 0 1} {0 1 -1} {1 6.123234e-17 -1} {1.2246469e-16 -1 -1} {-1 -1.8369701e-16 -1} {0 0 -1}]
	// indices := [0 4 1 0 1 2 0 2 3 0 3 4 5 4 1 5 1 2 5 2 3 5 3 4]

}

func BenchmarkCone4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Cone(1.0, 4, true)
	}
}

func BenchmarkCone16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Cone(1.0, 16, true)
	}
}

func BenchmarkCone64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Cone(1.0, 64, true)
	}
}
