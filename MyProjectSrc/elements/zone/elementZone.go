package elementZone
import (
	"MyProjectSrc/context"
	"MyProjectSrc/elements"
	"image"
	"MyProjectSrc/data"
)
type Zone struct{
	Border int
	Size image.Point
	Mod *element.ColorMod
}
func(zone Zone)Render(context *context.Context){
	data.Options.GeoM.Scale( float64( zone.Size.X ), float64( zone.Size.Y ) )
	data.Options.GeoM.Translate( float64(context.Bound.Min.X - context.Sub.X), float64(context.Bound.Min.Y - context.Sub.Y) )
	context.Img.DrawImage(zone.Mod.Acolor, data.Options)
	data.Options.GeoM.Reset()
	
	data.Options.GeoM.Scale( float64( zone.Size.X-zone.Border*2 ), float64( zone.Size.Y-zone.Border*2 ) )
	data.Options.GeoM.Translate( float64(context.Bound.Min.X+zone.Border - context.Sub.X), float64(context.Bound.Min.Y+zone.Border - context.Sub.Y) )
	context.Img.DrawImage(zone.Mod.Bcolor, data.Options)
	data.Options.GeoM.Reset()
}
func (zone *Zone) GetMetrics()(m element.Metrick){
	m.X = zone.Size.X
	m.Y = zone.Size.Y
	return
}
func MakeZone(Border int, size image.Point, Mod *element.ColorMod) (z *Zone){
	z = new(Zone)
	if Mod == nil{
		Mod = &data.DefultModA
	}
	*z = Zone{
		Border : Border,
		Size : size,
		Mod : Mod,
	}
	return
}
func MakeZoneUse(Border int, size image.Point, Mod *element.ColorMod) (z *Zone){
	z = new(Zone)
	if Mod == nil{
		Mod = &data.DefultModA
	}
	*z = Zone{
		Border : Border,
		Size : size,
		Mod : Mod,
	}
	return
}