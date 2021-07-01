package elementFloot
import (
	"MyProjectSrc/elements"
	"MyProjectSrc/context"
	"image"
)
type Floot struct{
	Body element.Element
	x,y float64
}
func (floot *Floot) Render(context *context.Context){
	context.Bound = context.Bound.Sub(
		Mult(
			image.Point(
				floot.Body.GetMetrics(),
			),
			floot.x,
			floot.y,
		),
	)
	floot.Body.Render(context)
}
func Mult(p image.Point, x,y float64)(r image.Point){
	r.X = int(float64(p.X) * x)
	r.Y = int(float64(p.Y) * y)
	return
}

func (floot *Floot) GetMetrics()(element.Metrick){
	return floot.Body.GetMetrics()
}

func MakeFloot(X,Y float64,el element.Element)(*Floot){
	return &Floot{
		Body : el,
		x : X,
		y : Y,
	}
	
}