package elementQue
import (
	"MyProjectSrc/elements"
	"MyProjectSrc/context"
	"MyProjectSrc/elements/prompt"
	"MyProjectSrc/elements/image"
	"MyProjectSrc/context/active"
	"MyProjectSrc/data"
	"image"
)
type Que struct{
	Body element.Element
	CursoreZone *contextActive.ActiveZone
	PromtCall *elementPromt.PromtCall
	
}
func (que *Que) Render(context *context.Context){
	que.PromtCall.ContextSet(context)
	que.Body.Render(context.CopyLocal())
	f := que.GetMetrics()
	x,y := f.X, f.Y
	que.CursoreZone.Up(
		image.Rectangle{
			Min:context.Bound.Min, 
			Max:image.Point{
				X:context.Bound.Min.X+x, 
				Y:context.Bound.Min.Y+y,
			},
		},
		context,
	)
}
func (que *Que) GetMetrics()(element.Metrick){
	return que.Body.GetMetrics()
}
func MakeQue(elements element.Element)(que *Que, promts *elementPromt.Promt){
	que = new(Que)
	
	promts, _ = elementPromt.MakePromt(elements, 0.4)
	que.PromtCall = &promts.Control
	que.CursoreZone = contextActive.MakeActive()
	que.CursoreZone.Handler = func(effect contextActive.Effect){
		if effect == contextActive.Colise{
			if elementPromt.WaitCaLL(que) {
				promts.Control.IsUse()
			}else{
			}
		}else{
			if effect == contextActive.Def {
				promts.Control.NoUse(que)
			}
		}
	}
	que.Body = elementImage.MakeImage(
		data.Que, 
		18,
		18,
	)
	
	
	return
}