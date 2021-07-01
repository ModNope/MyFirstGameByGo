package elementPromt
import (
	"MyProjectSrc/context"
	"MyProjectSrc/elements"
	"MyProjectSrc/elements/zone"
	"MyProjectSrc/data"
	"MyProjectSrc/elements/sum"
	"MyProjectSrc/elements/floot"
	"MyProjectSrc/elements/privat"
	"MyProjectSrc/elements/traingle"
	"image"
)
var (
	lenght int = 30
	waiter int = 0
	waiterPointer element.Element
)
const WaintTime int = 30
type Promt struct{
	Body element.Element
	Control PromtCall
	Traingle [3]image.Point
}
type PromtCall struct{
	cordPoint image.Point
	isCall bool
	del bool
}

func WaitCaLL(el element.Element)(bool){
	if el != waiterPointer{
		waiterPointer = el
		waiter = 0
		return 0==1
	}else{
		if waiter < WaintTime {
			waiter += 1
			return 0==1
		}else{
			return 1==1
		}
	}
	
}

func (promtCall *PromtCall) ContextSet(context *context.Context){
	promtCall.cordPoint = context.Bound.Min
}

func (promtCall *PromtCall) Del(){
	promtCall.del = 0==0
}
func (promtCall *PromtCall) NoUse(el element.Element){
	if el != waiterPointer{
		waiter = 0
		//print(string(10), el)
	}
	promtCall.isCall = 1==0
	
	
}
func (promtCall *PromtCall) IsUse(){
	promtCall.isCall = 1==1
}

func (promt *Promt) Render(context *context.Context){
	if !promt.Control.del{
		if promt.Control.isCall{
			promt.Traingle[0] = image.Point{-10, 0}
			promt.Traingle[1] = image.Point{-40, 0}
			promt.Traingle[2] = image.Point{ 0, 15}
			
			copyContext := context.CopyLocal()
			copyContext.Bound.Min = promt.Control.cordPoint
			promt.Body.Render(copyContext.CopyLocal())
		}
	}
}
func (promt *Promt) GetMetrics()(m element.Metrick){
	return 
}
func MakePromt(body element.Element, flootX float64) (p *Promt, c *PromtCall){
	p = new(Promt)
	p.Body = elementSum.MakeSum(
			elementFloot.MakeFloot(
				flootX,
				1,
				elementSum.MakeSum(
					elementZone.MakeZoneUse(
						3, 
						image.Point(
							body.GetMetrics(),
						), 
						&data.DefultModD,
					),
					elementPrivat.MakePrivatZone( 
						image.Point(
							body.GetMetrics(),
						),
						body,
					),
				),
			),
			elementTraingle.MakeTraingle(
				[3]*image.Point{
					&p.Traingle[0],
					&p.Traingle[1],
					&p.Traingle[2],
				},
			),
	)
	c = &p.Control
	
	return
}