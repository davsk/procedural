// dice.go

package procedural

import (
	"azul3d.org/gfx.v1"
)

// Dice builds and returns an new 3D cube mesh
// with texture mapped to 6 parts based on western right-handed or counterclockwise.
func Dice(scale float32) *gfx.Mesh {
	s := scale

	vertices := []gfx.Vec3{
		// Bottom
		gfx.Vec3{-s, s, -s},
		gfx.Vec3{-s, -s, -s},
		gfx.Vec3{s, -s, -s},
		gfx.Vec3{-s, s, -s},
		gfx.Vec3{s, -s, -s},
		gfx.Vec3{s, s, -s},

		// Front
		gfx.Vec3{-s, -s, s},
		gfx.Vec3{-s, -s, -s},
		gfx.Vec3{s, -s, -s},
		gfx.Vec3{-s, -s, s},
		gfx.Vec3{s, -s, -s},
		gfx.Vec3{s, -s, s},

		// Left
		gfx.Vec3{-s, s, s},
		gfx.Vec3{-s, s, -s},
		gfx.Vec3{-s, -s, -s},
		gfx.Vec3{-s, s, s},
		gfx.Vec3{-s, -s, -s},
		gfx.Vec3{-s, -s, s},

		// Back
		gfx.Vec3{-s, s, s},
		gfx.Vec3{-s, s, -s},
		gfx.Vec3{s, s, -s},
		gfx.Vec3{-s, s, s},
		gfx.Vec3{s, s, -s},
		gfx.Vec3{s, s, s},

		// Right
		gfx.Vec3{s, s, s},
		gfx.Vec3{s, s, -s},
		gfx.Vec3{s, -s, -s},
		gfx.Vec3{s, s, s},
		gfx.Vec3{s, -s, -s},
		gfx.Vec3{s, -s, s},

		// Top
		gfx.Vec3{-s, s, s},
		gfx.Vec3{-s, -s, s},
		gfx.Vec3{s, -s, s},
		gfx.Vec3{-s, s, s},
		gfx.Vec3{s, -s, s},
		gfx.Vec3{s, s, s},
	}

	textureCoords := []gfx.TexCoord{
		// Bottom 6
		gfx.TexCoord{5.0 / 6.0, 1},
		gfx.TexCoord{5.0 / 6.0, 0},
		gfx.TexCoord{1, 0},
		gfx.TexCoord{5.0 / 6.0, 1},
		gfx.TexCoord{1, 0},
		gfx.TexCoord{1, 1},

		// Front 2
		gfx.TexCoord{1.0 / 6.0, 0},
		gfx.TexCoord{1.0 / 6.0, 1},
		gfx.TexCoord{2.0 / 6.0, 1},
		gfx.TexCoord{1.0 / 6.0, 0},
		gfx.TexCoord{2.0 / 6.0, 1},
		gfx.TexCoord{2.0 / 6, 0},

		// Left 3
		gfx.TexCoord{2.0 / 6.0, 0},
		gfx.TexCoord{2.0 / 6.0, 1},
		gfx.TexCoord{3.0 / 6.0, 1},
		gfx.TexCoord{2.0 / 6.0, 0},
		gfx.TexCoord{3.0 / 6.0, 1},
		gfx.TexCoord{3.0 / 6.0, 0},

		// Back 5
		gfx.TexCoord{5.0 / 6.0, 0},
		gfx.TexCoord{5.0 / 6.0, 1},
		gfx.TexCoord{4.0 / 6.0, 1},
		gfx.TexCoord{5.0 / 6.0, 0},
		gfx.TexCoord{4.0 / 6.0, 1},
		gfx.TexCoord{4.0 / 6.0, 0},

		// Right 4
		gfx.TexCoord{4.0 / 6.0, 0},
		gfx.TexCoord{4.0 / 6.0, 1},
		gfx.TexCoord{3.0 / 6.0, 1},
		gfx.TexCoord{4.0 / 6.0, 0},
		gfx.TexCoord{3.0 / 6.0, 1},
		gfx.TexCoord{3.0 / 6.0, 0},

		// Top 1
		gfx.TexCoord{0, 0},
		gfx.TexCoord{0, 1},
		gfx.TexCoord{1.0 / 6.0, 1},
		gfx.TexCoord{0, 0},
		gfx.TexCoord{1.0 / 6.0, 1},
		gfx.TexCoord{1.0 / 6.0, 0},
	}

	return &gfx.Mesh{
		Vertices: vertices,
		TexCoords: []gfx.TexCoordSet{
			gfx.TexCoordSet{Slice: textureCoords},
		},
	}
}
