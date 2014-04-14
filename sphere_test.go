// sphere_test.go

package procedural

import (
	"azul3d.org/v0/scene/geom"
	"fmt"
	"testing"
)

func ExampleSphere() {
	ball := Sphere(1, geom.Static)
	fmt.Printf("ball := %v\n", ball)
	fmt.Printf("Vertices := %v\n", ball.Vertices)
	fmt.Printf("Indices := %v\n", ball.Indices)
	fmt.Printf("TextureCoords := %v\n", ball.TextureCoords)
	// Output:
	// ball := Mesh(Static, Hidden=false, 21 Vertices)
	// Vertices := [{0 0 1} {0 0 -1} {1 0 0} {0 1 0} {-1 0 0} {0 -1 0} {0 -1 0} {0.70710677 0 0.70710677} {0.70710677 0.70710677 0} {0 0.70710677 0.70710677} {-0.70710677 0.70710677 0} {-0.70710677 0 0.70710677} {-0.70710677 -0.70710677 0} {0 -0.70710677 0.70710677} {0 -0.70710677 0.70710677} {0.70710677 -0.70710677 0} {0.70710677 0 -0.70710677} {0 0.70710677 -0.70710677} {-0.70710677 0 -0.70710677} {0 -0.70710677 -0.70710677} {0 -0.70710677 -0.70710677}]
	// Indices := [0 7 9 2 7 8 3 8 9 7 8 9 0 9 11 3 9 10 4 10 11 9 10 11 0 11 14 4 11 12 6 12 14 11 12 14 0 13 7 5 13 15 2 15 7 13 15 7 1 16 17 2 16 8 3 8 17 16 8 17 1 17 18 3 17 10 4 10 18 17 10 18 1 18 20 4 18 12 6 12 20 18 12 20 1 19 16 5 19 15 2 15 16 19 15 16]
	// TextureCoords := [[{0.5 1 1 1} {0.5 0 1 1} {0.25 0.5 1 1} {0.5 0.5 1 1} {0.75 0.5 1 1} {0 0.5 1 1} {1 0.5 1 1} {0.25 0.75 1 1} {0.375 0.5 1 1} {0.5 0.75 1 1} {0.625 0.5 1 1} {0.75 0.75 1 1} {0.875 0.5 1 1} {0 0.75 1 1} {1 0.75 1 1} {0.125 0.5 1 1} {0.25 0.25 1 1} {0.5 0.25 1 1} {0.75 0.25 1 1} {0 0.25 1 1} {1 0.25 1 1}]]
	//
}

func TestUvLatLng(t *testing.T) {
	vA := geom.Vertex{0, 0, 1}
	vB := geom.Vertex{0, -1, 0}
	vC := betweenPoints(vA, vB)
	x, y := UvLatLng(vA)
	t.Log(vA, x, y)
	x, y = UvLatLng(vB)
	t.Log(vB, x, y)
	x, y = UvLatLng(vC)
	t.Log(vC, x, y)
}

func BenchmarkSphereT0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Sphere(0, geom.Static)
	}
}

func BenchmarkSphereT1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Sphere(1, geom.Static)
	}
}

func BenchmarkSphereT2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Sphere(2, geom.Static)
	}
}

func BenchmarkSphereT3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Sphere(3, geom.Static)
	}
}

func BenchmarkSphereT4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Sphere(4, geom.Static)
	}
}

func BenchmarkSphereT5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Sphere(5, geom.Static)
	}
}

func BenchmarkSphereT6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Sphere(6, geom.Static)
	}
}

func BenchmarkSphereT7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Sphere(7, geom.Static)
	}
}
