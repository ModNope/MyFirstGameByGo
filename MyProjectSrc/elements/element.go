package element
import (
	"MyProjectSrc/context"
	"github.com/hajimehoshi/ebiten"
	"image"
)
type Element interface {
	Render(*context.Context)
	GetMetrics()(Metrick)
}
type Metrick image.Point

type HardElement struct{
	El Element
	Me Metrick
}
type ColorMod struct {
	Acolor, Bcolor *ebiten.Image
}