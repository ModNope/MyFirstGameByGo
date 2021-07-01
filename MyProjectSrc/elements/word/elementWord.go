package elementWord
import (
	"MyProjectSrc/elements"
	"MyProjectSrc/context"
	"MyProjectSrc/data"
	"image"
	"github.com/hajimehoshi/ebiten/text"
	"strings"
	"golang.org/x/image/font"
)
type Words struct{
	Str string
	Metrick element.Metrick
	border image.Rectangle
	bordSize int
	font *font.Face
}
func (texts *Words) Reset(str string){
	texts.Str = str
	textBound := text.BoundString(*texts.font, str)
	texts.Metrick = element.Metrick{
		X : textBound.Dx()+texts.bordSize*2,
		Y : textBound.Dy()+texts.bordSize*2,
	}
	texts.border = textBound
	
}
func (texts Words) Render(context *context.Context){
	text.Draw(
		context.Img, 
		texts.Str, 
		*texts.font, 
		context.Bound.Min.X - texts.border.Min.X + texts.bordSize - context.Sub.X, 
		context.Bound.Min.Y + texts.border.Max.Y - texts.border.Min.Y + texts.bordSize - context.Sub.Y, 
		data.TextColor,
	)
}
func (text Words) GetMetrics()(element.Metrick){
	return text.Metrick
}
func MakeLongWord(sizeDorder int, str string)(*Words){
	textBound := text.BoundString(data.MainFontUse, str)
	var g = Words{ 
		Str : str, 
		Metrick : 
		element.Metrick{
			X : textBound.Dx()+sizeDorder*2,
			Y : textBound.Dy()+sizeDorder*2,
		},
		border : textBound,
		bordSize : sizeDorder,
		font : &data.MainFontUse,
	}
	return &g
}
func MakeWord(sizeDorder int, str ...string)(e []Words){
	for i,_:=range str{
		for _,v := range strings.Split(str[i], " ") {
			e = append(e,*MakeLongWord(sizeDorder, v))
		}
	}
	return
}
func MakeAbc(sizeDorder int, r string, str ...string)(e []element.Element){
	for i,_:=range str{
		for _,v := range strings.Split(str[i], r) {
			e = append(e,*MakeLongWord(sizeDorder, v))
		}
	}
	return
}

func MakeLongWordFont(sizeDorder int, str string, f *font.Face)(*Words){
	textBound := text.BoundString(*f, str)
	if f == nil{
		f = &data.MainFontUse
	}
	var g = Words{ 
		Str : str, 
		Metrick : 
		element.Metrick{
			X : textBound.Dx()+sizeDorder*2,
			Y : textBound.Dy()+sizeDorder*2,
		},
		border : textBound,
		bordSize : sizeDorder,
		font : f,
	}
	return &g
}