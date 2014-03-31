// cone_test.go

package procedural

import (
	"azul3d.org/v0/scene/geom"
	"fmt"
)

func ExampleCone() {
	pyramid := Cone(1.0, 4, true, geom.Static)
	hat := Cone(1.0, 8, false, geom.Dynamic)
	cone := Cone(1.0, 32, true, geom.Static)
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
	// pyramid := Mesh(Static, Hidden=false, 6 Vertices)
	// vertices := [{0 0 1} {0 1 -1} {1 6.123234e-17 -1} {1.2246469e-16 -1 -1} {-1 -1.8369701e-16 -1} {0 0 -1}]
	// indices := [0 4 1 0 1 2 0 2 3 0 3 4 5 4 1 5 1 2 5 2 3 5 3 4]

}
