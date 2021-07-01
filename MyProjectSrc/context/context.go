package context
import (
	"github.com/hajimehoshi/ebiten"
	"image"
)
type Context struct{
	Img *ebiten.Image
	Bound image.Rectangle
	Sub image.Point
	Lv, Lvf int
}
func (context *Context) Copy(n *image.Rectangle)(Ncontext *Context){
	if n == nil{
		n = &context.Bound
	}
	var con = Context{
		Img : context.Img,
		Bound : *n,
		Sub : context.Sub,
		Lv : context.Lv+1,
		Lvf : context.Lvf,
	}
	Ncontext = &con
	return 
}
func (context *Context) CopyLocal()(Ncontext *Context){
	context.Lvf += 1
	var con = Context{
		Img : context.Img,
		Bound : context.Bound,
		Sub : context.Sub,
		Lv : context.Lv,
		Lvf : context.Lvf,
	}
	Ncontext = &con
	return 
}
func New(img *ebiten.Image, n *image.Rectangle)(context *Context){
	if n == nil{
		if img != nil{
			x,y := img.Size()
			n = &image.Rectangle{Max : image.Point{X:x,Y:y}}
		}
	}
	context = &Context{
		Img : img, 
		Bound : *n,
	}
	return
}