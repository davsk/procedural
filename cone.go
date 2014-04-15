// Copyright 2014 Davsk. All rights reserved.

package procedural

import (
	"azul3d.org/v1/gfx"
	"math"
)

// Cone builds and returns an new 3D cone mesh. Minimum of 3 steps.
func Cone(scale float32, steps int, capped bool) *gfx.Mesh {
	var (
		angle, inc float64
		uvX, uvY   float32
	)

	x := scale

	if steps < 3 {
		steps = 3
	}

	maxVertices := steps + 1

	if capped {
		maxVertices++
	}

	maxIndices := 3 * steps

	if capped {
		maxIndices *= 2
	}

	vertices := make([]gfx.Vec3, maxVertices)
	indices := make([]uint32, maxIndices)
	textureCoords := make([]gfx.TexCoord, maxVertices)

	vertices[0] = gfx.Vec3{0, 0, x}

	if capped {
		vertices[maxVertices-1] = gfx.Vec3{0, 0, -x}
	}

	angle = 0.0
	inc = (2.0 * math.Pi) / float64(steps)

	for i := 0; i < steps; i++ {
		s, c := math.Sincos(angle)
		angle += inc
		vertices[i+1] = gfx.Vec3{x * float32(s), x * float32(c), -x}

		if i == 0 {
			indices[0] = 0
			indices[1] = uint32(steps)
			indices[2] = 1
			if capped {
				indices[steps*3] = uint32(maxVertices) - 1
				indices[steps*3+1] = uint32(steps)
				indices[steps*3+2] = 1
			}
		} else {
			indices[3*i] = 0
			indices[(3*i)+1] = uint32(i)
			indices[(3*i)+2] = uint32(i) + 1
			if capped {
				indices[(steps+i)*3] = uint32(maxVertices) - 1
				indices[((steps+i)*3)+1] = uint32(i)
				indices[((steps+i)*3)+2] = uint32(i) + 1
			}
		}

		uvX = float32(i - (2 * (i / 2)))
		uvY = float32((i / 2) - (2 * (i / 4)))
		textureCoords[i] = gfx.TexCoord{uvX, uvY}
	}

	return &gfx.Mesh{
		Indices:  indices,
		Vertices: vertices,
		TexCoords: []gfx.TexCoordSet{
			gfx.TexCoordSet{Slice: textureCoords},
		},
	}
}
