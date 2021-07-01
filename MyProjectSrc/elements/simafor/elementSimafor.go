package elementSimafor
import (
	"MyProjectSrc/context"
	"MyProjectSrc/elements"
	
)
type Simafor struct{
	Body []element.Element
	Using int
}
func (simafor *Simafor) Render(context *context.Context){
	if len(simafor.Body)-1 >= (simafor.Using){
		simafor.Body[simafor.Using].Render(context)
	}
	
}
func (simafor *Simafor) GetMetrics()(element.Metrick){
	if len(simafor.Body)-1 >= (simafor.Using){
		return simafor.Body[simafor.Using].GetMetrics()
	}else{
		return element.Metrick{}
	}
}
func(simafor *Simafor) Set(fgs int){
	if (fgs >= 0) {
		if (fgs < len(simafor.Body)){
			simafor.Using = fgs
		}
	}
}
func MakeSimafor(elements ...element.Element)(simafor *Simafor){
	simafor = &Simafor{
		Body : elements,
		Using : 0,
	}
	return
}