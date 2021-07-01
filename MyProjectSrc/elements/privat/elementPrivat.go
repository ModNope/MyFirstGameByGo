package elementPrivat
import (
	"MyProjectSrc/context"
	"MyProjectSrc/elements"
	"image"
	//"image/color"
	"MyProjectSrc/data"
	"github.com/hajimehoshi/ebiten"
	//"fmt"
	
)
var(
	img *ebiten.Image
	col uint8
)
type Privat struct{
	Bound image.Point
	Use element.Element
}
func (privat Privat) Render(context *context.Context)(){
	copys := context.Copy(nil)
	if privat.Use != nil{
		copyContext := copys.CopyLocal()
		
		
		copyContext.Img,_ = ebiten.NewImage(privat.Bound.X, privat.Bound.Y, ebiten.FilterNearest)
		copyContext.Sub = context.Bound.Min
		copyContext.Bound.Max = context.Bound.Min.Add(privat.Bound)
		privat.Use.Render(copyContext)
		
		data.Options.GeoM.Translate( 
			float64(
				context.Bound.Min.X-context.Sub.X,
			), 
			float64(
				context.Bound.Min.Y-context.Sub.Y,
			),
		)
		context.Img.DrawImage(copyContext.Img, data.Options)
		data.Options.GeoM.Reset()
		copyContext.Img.Dispose()
	}
}
func (privat Privat) GetMetrics()(element.Metrick){
	return element.Metrick(privat.Bound)
}
func MakePrivatZone(bound image.Point, element element.Element) (p *Privat){
	privat := Privat{Bound : bound, Use : element}
	p = &privat
	return
}