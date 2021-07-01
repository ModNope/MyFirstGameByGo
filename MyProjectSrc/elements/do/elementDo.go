package elementDo
import (
	"MyProjectSrc/context"
	"MyProjectSrc/elements"
	
)
type Do struct{
	Renders []func(context *context.Context)()
	GetMetricss func()(element.Metrick)
}
func (do Do) Render(context *context.Context){
	copys := context.Copy(nil)
	for i:=range do.Renders{
		do.Renders[i](copys.CopyLocal())
	}
}
func (do Do) GetMetrics()(E element.Metrick){
	if do.GetMetricss != nil{
		return do.GetMetricss()
	}
	return
}
func MakeDo(doinge ...func(*context.Context)())(do *Do){
	do = &Do{
		Renders : doinge,
	}
	return
}
func MakeDoing(doinge func(*context.Context)(), met func()(element.Metrick))(do *Do){
	do = &Do{
		Renders : []func(context *context.Context)(){doinge},
		GetMetricss : met,
	}
	return
}