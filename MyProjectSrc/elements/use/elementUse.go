package elementUse
import (
	"MyProjectSrc/context"
	"MyProjectSrc/elements"
)
type Use struct{
	E element.Element
	b bool
}

func (use *Use) Render(context *context.Context){
	if use.b{
		use.E.Render(context)
	}
	
}
func (use *Use) GetMetrics()(m element.Metrick){
	if use.b{
		m = use.E.GetMetrics()
	}
	return 
}
func (use *Use) Use(t bool)(){
	use.b = t
}


func MakeUse(e element.Element)(use *Use){
	use = new(Use)
	use.E = e
	return
}