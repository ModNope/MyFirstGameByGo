package elementPlan
import (
	"MyProjectSrc/context"
	"MyProjectSrc/elements"
	"image"
)
type Plane []*PosRender
type PosRender struct{
	Pos image.Point
	Use element.Element
}
func (plane Plane) Render(context *context.Context){
	copys := context.Copy(nil)
	for i := range plane{
		newe := copys.CopyLocal()
		newe.Bound.Min = newe.Bound.Min.Add(plane[i].Pos)
		newe.Bound.Max = newe.Bound.Max.Add(plane[i].Pos)
		plane[i].Use.Render(newe)
	}
}
func (plane Plane) GetMetrics()(m element.Metrick){
	for i:=range plane{
		newe := plane[i].Pos.Add(image.Point(plane[i].Use.GetMetrics()))
		if newe.X > m.X{
			m.X = newe.X
		}
		if newe.Y > m.Y{
			m.Y = newe.Y
		}
	}
	return 
}
func MakePose(i element.Element, px,py int) (r *PosRender){
	r = new(PosRender)
	r.Pos = image.Point{px,py}
	r.Use = i
	return
}
func MakePlane(use ...*PosRender) (p *Plane){
	p = new(Plane)
	*p = use
	return
}
func MakePoseSimple(i element.Element, px,py int)(p *Plane){
	p = new(Plane)
	r := new(PosRender)
	r.Pos = image.Point{px,py}
	r.Use = i
	*p = Plane{r}
	return
}