// dice_test.go

package procedural

import (
	"azul3d.org/v0/scene"
	"azul3d.org/v0/scene/geom"
	"testing"
)

func ExampleDice() {
	geomNode := scene.New("Dice")
	dice := Dice(1.0, geom.Dynamic)
	geom.Add(geomNode, dice)
	// texture.Set(geomNode, layDice, texDice)

}

func BenchmarkDice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Dice(1.0, geom.Dynamic)
	}
}
