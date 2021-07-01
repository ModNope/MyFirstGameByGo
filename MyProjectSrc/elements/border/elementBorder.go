package elementBorder
import (
	"MyProjectSrc/context"
	"MyProjectSrc/elements"
	"image"
)
type Border struct{
	Body element.Element
	BorderSize int
}
func (border *Border) Render(context *context.Context)(){
	context.Bound = context.Bound.Add(
		image.Point{
			border.BorderSize,
			border.BorderSize,
		},
	)
	border.Body.Render(context)
}
func (border *Border) GetMetrics()(e element.Metrick){
	return element.Metrick(
		image.Point(
				border.Body.GetMetrics(),
			).Add(
				image.Point{
					border.BorderSize*2, 
					border.BorderSize*2,
				},
			),
		)
}
func MakeBorder(elements element.Element, s int) (border *Border){
	border = &Border{
		Body : elements,
		BorderSize : s,
	}
	return
}