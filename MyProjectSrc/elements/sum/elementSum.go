package elementSum
import (
	"MyProjectSrc/context"
	"MyProjectSrc/elements"
)


type Sum []element.Element
func (sum *Sum) Render(context *context.Context)(){
	//copys := context.Copy(nil)
	for i := range *sum{
		(*sum)[i].Render( context.CopyLocal() )
	}
}
func (sum *Sum) GetMetrics()(e element.Metrick){
	var p element.Metrick
	for i := range *sum{
		p = (*sum)[i].GetMetrics()
		if p.X > e.X {
			e.X = p.X
		}
		if p.Y > e.Y {
			e.Y = p.Y
		}
	}
	return 
}
func (sum *Sum) Add(elements ...element.Element)(){
	*sum = append(*sum, elements...)
}
func MakeSum(elements ...element.Element) (s *Sum){
	s = new(Sum)
	*s = elements
	return
}