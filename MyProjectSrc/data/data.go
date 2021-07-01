package data
import (
	"MyProjectSrc/elements"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"image"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
	"io/ioutil"
)
var(
	MainFontUse, MainFontUseMeny, MainFontPrinter font.Face
	
	TextColor = color.White
	
	Options = &ebiten.DrawImageOptions{}
	
	_,ColorModImg,_ = ebitenutil.NewImageFromFile("textures/ColorMod1.png", ebiten.FilterDefault)
	
	DefultModA = SetMod(0,0, &ColorModImg)
	DefultModB = SetMod(1,0, &ColorModImg)
	DefultModC = SetMod(2,0, &ColorModImg)
	DefultModD = SetMod(3,0, &ColorModImg)
	
	ImageMashine,_,_ = ebitenutil.NewImageFromFile("textures/ImageMashine.png", ebiten.FilterDefault)
	Many,_,_ = ebitenutil.NewImageFromFile("textures/Many.png", ebiten.FilterDefault)
	Time1,_,_ = ebitenutil.NewImageFromFile("textures/Time1.png", ebiten.FilterDefault)
	Time2,_,_ = ebitenutil.NewImageFromFile("textures/Time2.png", ebiten.FilterDefault)
	OldFon,_,_ = ebitenutil.NewImageFromFile("textures/OldFon.png", ebiten.FilterDefault)
	Tutorial1,_,_ = ebitenutil.NewImageFromFile("textures/tutorial 1-1-1.png", ebiten.FilterDefault)
	Que,_,_ = ebitenutil.NewImageFromFile("textures/que.png", ebiten.FilterDefault)
	ReUp,_,_ = ebitenutil.NewImageFromFile("textures/ReUp.png", ebiten.FilterDefault)
	ReDown,_,_ = ebitenutil.NewImageFromFile("textures/ReDown.png", ebiten.FilterDefault)
	_,Icon,_ = ebitenutil.NewImageFromFile("textures/IconL.png", ebiten.FilterDefault)
	
	Wait,_,_ = ebitenutil.NewImageFromFile("textures/Wait.png", ebiten.FilterDefault)
	Go,_,_ = ebitenutil.NewImageFromFile("textures/Go.png", ebiten.FilterDefault)
	Need,_,_ = ebitenutil.NewImageFromFile("textures/Need.png", ebiten.FilterDefault)
	
)
func init(){
	mainFontGet, err := ioutil.ReadFile("20805.ttf")
	if err != nil {
		log.Println(err)
	}
	mainFontPars, err := opentype.Parse(mainFontGet)
	if err != nil {
		log.Fatal(err)
	}
	MainFontUse, err = opentype.NewFace(mainFontPars, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	MainFontUseMeny, err = opentype.NewFace(mainFontPars, &opentype.FaceOptions{
		Size:    60,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	MainFontPrinter, err = opentype.NewFace(mainFontPars, &opentype.FaceOptions{
		Size:    14,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}
func SetMod(x,y int, img *image.Image)(m element.ColorMod){
	m.Acolor,_ = ebiten.NewImage(1,1, ebiten.FilterNearest)
	m.Acolor.Set(0,0, (*img).At(x,y) )
	m.Bcolor,_ = ebiten.NewImage(1,1, ebiten.FilterNearest)
	m.Bcolor.Set(0,0, (*img).At(x,y+1) )
	return
}
func init(){
}