// sphere.go

package procedural

import (
	gmath "azul3d.org/v0/math"
	"azul3d.org/v0/scene/geom"
	"azul3d.org/v0/scene/texture"
	"math"
)

// almostEqualFloat32 because a == b does not work close to 0
func almostEqualFloat32(a, b float32) bool {
	var c float32

	if a > b {
		c = a - b
	} else {
		c = b - a
	}

	return c < 0.00000000001
}

// almostEqualVertex because a == b does not work close to 0
func almostEqualVertex(a, b geom.Vertex) bool {
	if !almostEqualFloat32(a.X, b.X) {
		return false
	}
	if !almostEqualFloat32(a.Y, b.Y) {
		return false
	}
	if !almostEqualFloat32(a.Z, b.Z) {
		return false
	}

	return true
}

// findVertice so we can reuse the same vertice
func findVertice(s []geom.Vertex, x geom.Vertex) (i uint32) {
	for i, v := range s {
		if almostEqualVertex(v, x) {
			return uint32(i)
		}
	}

	return 0
}

// betweenPoints is a vector of 1 between these two vectors
func betweenPoints(a geom.Vertex, b geom.Vertex) (c geom.Vertex) {
	var vA, vB, vC gmath.Vec3
	vA.X = float64(a.X)
	vA.Y = float64(a.Y)
	vA.Z = float64(a.Z)
	vB.X = float64(b.X)
	vB.Y = float64(b.Y)
	vB.Z = float64(b.Z)
	vC, _ = vA.Add(vB).Normalized()
	c.X = float32(vC.X)
	c.Y = float32(vC.Y)
	c.Z = float32(vC.Z)

	return c
}

// recurseSphere should only be called by Sphere
func recurseSphere(a uint32, b uint32, c uint32, vertices []geom.Vertex, textureCoords []texture.Coord, indices []uint32, steps int) (v []geom.Vertex, t []texture.Coord, i []uint32) {
	if steps == 0 {
		indices = append(indices, a)
		indices = append(indices, b)
		indices = append(indices, c)
	} else {
		dV := betweenPoints(vertices[a], vertices[b])
		eV := betweenPoints(vertices[b], vertices[c])
		fV := betweenPoints(vertices[c], vertices[a])

		d := findVertice(vertices, dV)
		e := findVertice(vertices, eV)
		f := findVertice(vertices, fV)

		if d == 0 {
			d = uint32(len(vertices))
			vertices, textureCoords = appendVerticesTextureCoords(vertices, textureCoords, dV)
			if textureCoords[d].U == 0.0 {
				if textureCoords[a].U == 1.0 {
					d++
				} else if textureCoords[b].U == 1.0 {
					d++
				}
			}
		}

		if e == 0 {
			e = uint32(len(vertices))
			vertices, textureCoords = appendVerticesTextureCoords(vertices, textureCoords, eV)
			if textureCoords[e].U == 0.0 {
				if textureCoords[b].U == 1.0 {
					e++
				} else if textureCoords[c].U == 1.0 {
					e++
				}
			}
		}

		if f == 0 {
			f = uint32(len(vertices))
			vertices, textureCoords = appendVerticesTextureCoords(vertices, textureCoords, fV)
			if textureCoords[f].U == 0.0 {
				if textureCoords[a].U == 1.0 {
					f++
				} else if textureCoords[c].U == 1.0 {
					f++
				}
			}
		}

		vertices, textureCoords, indices = recurseSphere(a, d, f, vertices, textureCoords, indices, steps-1)
		vertices, textureCoords, indices = recurseSphere(b, d, e, vertices, textureCoords, indices, steps-1)
		vertices, textureCoords, indices = recurseSphere(c, e, f, vertices, textureCoords, indices, steps-1)
		vertices, textureCoords, indices = recurseSphere(d, e, f, vertices, textureCoords, indices, steps-1)
	}
	return vertices, textureCoords, indices
}

// appendVerticesTextureCoords appends geom.Vertex to slice and match UV.Coord to texture slice.
func appendVerticesTextureCoords(vertices []geom.Vertex, textureCoords []texture.Coord, v geom.Vertex) ([]geom.Vertex, []texture.Coord) {
	vertices = append(vertices, v)
	uvX, uvY := UvLatLng(v)
	textureCoords = append(textureCoords, texture.UV(uvX, uvY))

	if uvX == 0.0 {
		vertices = append(vertices, v)
		uvX = 1.0
		textureCoords = append(textureCoords, texture.UV(uvX, uvY))
	}

	return vertices, textureCoords
}

// Sphere reports a mesh representing a regular octahedron which has been divided by a number of steps.
// Minimum of 0 steps. Time and memory used expands exponentially as the number of steps increases.
// Vertices on the dateline are duplicated with textures mapped east or west to match the slope of the face.
func Sphere(steps int, hint geom.Hint) *geom.Mesh {
	if steps < 0 {
		steps = 0
	}

	maxVertices := 4 ^ steps

	maxIndices := 3 * maxVertices

	vertices := make([]geom.Vertex, 0, maxVertices)
	indices := make([]uint32, 0, maxIndices)
	textureCoords := make([]texture.Coord, 0, maxVertices)

	vertices, textureCoords = appendVerticesTextureCoords(vertices, textureCoords, geom.Vertex{0, 0, 1})
	vertices, textureCoords = appendVerticesTextureCoords(vertices, textureCoords, geom.Vertex{0, 0, -1})
	vertices, textureCoords = appendVerticesTextureCoords(vertices, textureCoords, geom.Vertex{1, 0, 0})
	vertices, textureCoords = appendVerticesTextureCoords(vertices, textureCoords, geom.Vertex{0, 1, 0})
	vertices, textureCoords = appendVerticesTextureCoords(vertices, textureCoords, geom.Vertex{-1, 0, 0})
	vertices, textureCoords = appendVerticesTextureCoords(vertices, textureCoords, geom.Vertex{0, -1, 0})

	vertices, textureCoords, indices = recurseSphere(0, 2, 3, vertices, textureCoords, indices, steps)
	vertices, textureCoords, indices = recurseSphere(0, 3, 4, vertices, textureCoords, indices, steps)
	vertices, textureCoords, indices = recurseSphere(0, 4, 6, vertices, textureCoords, indices, steps)
	vertices, textureCoords, indices = recurseSphere(0, 5, 2, vertices, textureCoords, indices, steps)
	vertices, textureCoords, indices = recurseSphere(1, 2, 3, vertices, textureCoords, indices, steps)
	vertices, textureCoords, indices = recurseSphere(1, 3, 4, vertices, textureCoords, indices, steps)
	vertices, textureCoords, indices = recurseSphere(1, 4, 6, vertices, textureCoords, indices, steps)
	vertices, textureCoords, indices = recurseSphere(1, 5, 2, vertices, textureCoords, indices, steps)

	return &geom.Mesh{
		Hint:     hint,
		Indices:  indices,
		Vertices: vertices,
		TextureCoords: [][]texture.Coord{
			textureCoords,
		},
	}
}

// UvLatLng reports UV Coords 0.0-1.0
func UvLatLng(p geom.Vertex) (u, v float32) {
	var flag bool

	c := gmath.CoordSysZUpRight
	d := gmath.Vec2{0.0, -1.0}
	vec2 := gmath.Vec2{float64(p.X), float64(p.Y)}
	vec2, flag = vec2.Normalized()
	if flag {
		u = float32(d.Angle(vec2) / math.Pi)
	} else {
		u = 0.5
	}
	vec3 := gmath.Vec3{float64(p.X), float64(p.Y), float64(p.Z)}
	v = float32(c.Down().Angle(vec3) / math.Pi)

	return u, v
}
