package elementImage
import (
	"MyProjectSrc/context"
	"MyProjectSrc/elements"
	"MyProjectSrc/data"
	"github.com/hajimehoshi/ebiten"
	
)
type Image struct{
	Img *ebiten.Image
	x,y int
}
func (img *Image) Render(context *context.Context){
	x,y := img.Img.Size()
	data.Options.GeoM.Scale( float64( img.x )/float64( x ), float64( img.y )/float64( y ))
	data.Options.GeoM.Translate( float64(context.Bound.Min.X - context.Sub.X), float64(context.Bound.Min.Y - context.Sub.Y) )
	context.Img.DrawImage(img.Img, data.Options)
	data.Options.GeoM.Reset()
}
func (img *Image) GetMetrics()(E element.Metrick){
	E = element.Metrick{
		X : img.x,
		Y : img.y,
	}
	return 
}
func MakeImage(image *ebiten.Image, x,y int)(img *Image){
	img = new(Image)
	img.Img = image
	img.x = x
	img.y = y
	return
}