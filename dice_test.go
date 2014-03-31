// dice_test.go

package procedural

import (
	"azul3d.org/v0/scene"
	"azul3d.org/v0/scene/geom"
)

func ExampleDice() {
	geomNode := scene.New("Dice")
	dice := Dice(1.0, geom.Dynamic)
	geom.Add(geomNode, dice)
	// texture.Set(geomNode, layDice, texDice)

}
