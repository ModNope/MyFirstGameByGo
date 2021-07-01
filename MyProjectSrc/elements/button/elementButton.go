package elementButton
import (
	"MyProjectSrc/context"
	"MyProjectSrc/context/active"
	"MyProjectSrc/elements"
	"MyProjectSrc/elements/zone"
	"MyProjectSrc/elements/word"
	"image"
	"MyProjectSrc/data"
	"golang.org/x/image/font"
	"MyProjectSrc/elements/prompt"
	
	//"fmt"
)
type Button struct{
	Print element.Element
	Zone  elementZone.Zone
	BordSize int
	OnColizeColor, NoClickColor, OnClickColor *element.ColorMod
	CursoreZone *contextActive.ActiveZone
	
	PromtCall *elementPromt.PromtCall
}

func (button Button)Render(context *context.Context){
	f := button.Zone.GetMetrics()
	x,y := f.X, f.Y
	button.CursoreZone.Up(
		image.Rectangle{
			Min:context.Bound.Min, 
			Max:image.Point{
				X:context.Bound.Min.X+x, 
				Y:context.Bound.Min.Y+y,
			},
		},
		context,
	)
	button.Zone.Render(context)
	
	context.Bound.Min.X += button.BordSize
	context.Bound.Min.Y += button.BordSize
	button.Print.Render(context)
	context.Bound.Min.X -= button.BordSize
	context.Bound.Min.Y -= button.BordSize
	
	if nil != button.PromtCall{
		button.PromtCall.ContextSet(context)
	}
	
	
}
func (button Button)GetMetrics()(element.Metrick){
	return button.Zone.GetMetrics()
}
func MakeTextButton(text string, borderSize, zoneSize int, doint func(*Button), modA, modB, modC *element.ColorMod, promtCall *elementPromt.PromtCall) (b *Button){
	if modA == nil{
		modA = &data.DefultModA
	}
	if modB == nil{
		modB = &data.DefultModB
	}
	if modC == nil{
		modC = &data.DefultModC
	}
	
	textElement := elementWord.MakeLongWord(zoneSize, text)
	var textMetrick = textElement.GetMetrics()
	textMetrick.X += (borderSize+zoneSize)*2
	textMetrick.Y += (borderSize+zoneSize)*2
	
	zoneElement := elementZone.MakeZoneUse(borderSize, image.Point(textMetrick), modA)
	button := Button{}
	b = &button
	b.PromtCall = promtCall
	b.Print = textElement
	b.Zone = *zoneElement
	b.BordSize = borderSize+zoneSize
	b.OnColizeColor = modA
	b.NoClickColor = modB
	b.OnClickColor = modC
	
	b.CursoreZone = contextActive.MakeActive()
	b.CursoreZone.Handler = func(effect contextActive.Effect){
		var obj = b
		var arr = []*element.ColorMod{ obj.OnColizeColor, obj.NoClickColor, obj.OnClickColor, obj.OnClickColor }
		obj.Zone.Mod = arr[effect]
		if effect == contextActive.ClickDo{
			if doint != nil {
				doint(b)
			}
		}else{
			if effect == contextActive.Colise{
				if elementPromt.WaitCaLL(b) {
					promtCall.IsUse()
				}else{
				}
			}else{
				if effect == contextActive.Def {
					promtCall.NoUse(b)
				}
			}
		}
	}
	
	return
}



func MakeTextButtonFont(text string, borderSize, zoneSize int, doint func(*Button), modA, modB, modC *element.ColorMod, font *font.Face) (b *Button){
	if modA == nil{
		modA = &data.DefultModA
	}
	if modB == nil{
		modB = &data.DefultModB
	}
	if modC == nil{
		modC = &data.DefultModC
	}
	textElement := elementWord.MakeLongWordFont(zoneSize, text, font)
	var textMetrick = textElement.GetMetrics()
	textMetrick.X += (borderSize+zoneSize)*2
	textMetrick.Y += (borderSize+zoneSize)*2
	
	zoneElement := elementZone.MakeZoneUse(borderSize, image.Point(textMetrick), modA)
	button := Button{}
	b = &button
	b.Print = textElement
	b.Zone = *zoneElement
	b.BordSize = borderSize+zoneSize
	b.OnColizeColor = modA
	b.NoClickColor = modB
	b.OnClickColor = modC
	
	b.CursoreZone = contextActive.MakeActive()
	b.CursoreZone.Handler = func(effect contextActive.Effect){
		var obj = b
		var arr = []*element.ColorMod{ obj.OnColizeColor, obj.NoClickColor, obj.OnClickColor, obj.OnClickColor }
		obj.Zone.Mod = arr[effect]
		if effect == contextActive.ClickDo{
			if doint != nil {
				doint(b)
			}
		}
	}
	return
}