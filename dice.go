// dice.go

package procedural

import (
	"azul3d.org/v0/scene/geom"
	"azul3d.org/v0/scene/texture"
)

// Dice builds and returns an new 3D cube mesh
// with texture mapped to 6 parts based on western right-handed or counterclockwise.
func Dice(scale float32, hint geom.Hint) *geom.Mesh {
	s := scale

	vertices := []geom.Vertex{
		// Bottom
		geom.Vertex{-s, s, -s},
		geom.Vertex{-s, -s, -s},
		geom.Vertex{s, -s, -s},
		geom.Vertex{-s, s, -s},
		geom.Vertex{s, -s, -s},
		geom.Vertex{s, s, -s},

		// Front
		geom.Vertex{-s, -s, s},
		geom.Vertex{-s, -s, -s},
		geom.Vertex{s, -s, -s},
		geom.Vertex{-s, -s, s},
		geom.Vertex{s, -s, -s},
		geom.Vertex{s, -s, s},

		// Left
		geom.Vertex{-s, s, s},
		geom.Vertex{-s, s, -s},
		geom.Vertex{-s, -s, -s},
		geom.Vertex{-s, s, s},
		geom.Vertex{-s, -s, -s},
		geom.Vertex{-s, -s, s},

		// Back
		geom.Vertex{-s, s, s},
		geom.Vertex{-s, s, -s},
		geom.Vertex{s, s, -s},
		geom.Vertex{-s, s, s},
		geom.Vertex{s, s, -s},
		geom.Vertex{s, s, s},

		// Right
		geom.Vertex{s, s, s},
		geom.Vertex{s, s, -s},
		geom.Vertex{s, -s, -s},
		geom.Vertex{s, s, s},
		geom.Vertex{s, -s, -s},
		geom.Vertex{s, -s, s},

		// Top
		geom.Vertex{-s, s, s},
		geom.Vertex{-s, -s, s},
		geom.Vertex{s, -s, s},
		geom.Vertex{-s, s, s},
		geom.Vertex{s, -s, s},
		geom.Vertex{s, s, s},
	}

	textureCoords := []texture.Coord{
		// Bottom 6
		texture.UV(5.0/6.0, 1),
		texture.UV(5.0/6.0, 0),
		texture.UV(1, 0),
		texture.UV(5.0/6.0, 1),
		texture.UV(1, 0),
		texture.UV(1, 1),

		// Front 2
		texture.UV(1.0/6.0, 0),
		texture.UV(1.0/6.0, 1),
		texture.UV(2.0/6.0, 1),
		texture.UV(1.0/6.0, 0),
		texture.UV(2.0/6.0, 1),
		texture.UV(2.0/6, 0),

		// Left 3
		texture.UV(2.0/6.0, 0),
		texture.UV(2.0/6.0, 1),
		texture.UV(3.0/6.0, 1),
		texture.UV(2.0/6.0, 0),
		texture.UV(3.0/6.0, 1),
		texture.UV(3.0/6.0, 0),

		// Back 5
		texture.UV(5.0/6.0, 0),
		texture.UV(5.0/6.0, 1),
		texture.UV(4.0/6.0, 1),
		texture.UV(5.0/6.0, 0),
		texture.UV(4.0/6.0, 1),
		texture.UV(4.0/6.0, 0),

		// Right 4
		texture.UV(4.0/6.0, 0),
		texture.UV(4.0/6.0, 1),
		texture.UV(3.0/6.0, 1),
		texture.UV(4.0/6.0, 0),
		texture.UV(3.0/6.0, 1),
		texture.UV(3.0/6.0, 0),

		// Top 1
		texture.UV(0, 0),
		texture.UV(0, 1),
		texture.UV(1.0/6.0, 1),
		texture.UV(0, 0),
		texture.UV(1.0/6.0, 1),
		texture.UV(1.0/6.0, 0),
	}

	return &geom.Mesh{
		Hint:     hint,
		Vertices: vertices,
		TextureCoords: [][]texture.Coord{
			textureCoords,
		},
	}
}
