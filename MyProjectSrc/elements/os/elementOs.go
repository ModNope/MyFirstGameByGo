package elementOs
import (
	"MyProjectSrc/elements"
	"MyProjectSrc/context"
	"image"
	"fmt"
)
type Os struct{
	Body element.Element
	xy image.Point
	n int
}
func (os *Os) Render(context *context.Context){
	copy := context.CopyLocal()
	a := image.Point(os.Body.GetMetrics())
	if os.n != 0{
		fmt.Println(os.Body.GetMetrics())
	}
	copy.Bound = copy.Bound.Sub(
		a.Sub(
			os.xy,
		).Div(
			2,
		),
	)
	(os.Body).Render(copy)
}
func (os *Os) GetMetrics()(element.Metrick){
	return element.Metrick(os.xy)
}

func MakeOs(X,Y int,el element.Element)(os *Os){
	return &Os{
		Body : el,
		xy : image.Point{
			X,
			Y,
		},
	}
}
func MakeOs2(X,Y int,el element.Element, n int)(os *Os){
	return &Os{
		Body : el,
		xy : image.Point{
			X,
			Y,
		},
		n : n,
	}
}