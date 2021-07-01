package elementTraingle
import (
	"MyProjectSrc/context"
	"MyProjectSrc/elements"
	"image"
	"MyProjectSrc/data"
	"github.com/hajimehoshi/ebiten"
)
var(
	pixel *ebiten.Image
	Options *ebiten.DrawTrianglesOptions
)
type Traingle struct{
	Body [3]*image.Point
}
func init(){
	Options = new(ebiten.DrawTrianglesOptions)
	pixel = data.DefultModD.Acolor
}
func (traingle Traingle) Render(context *context.Context)(){
	
	if (traingle.Body[0] != nil) && (traingle.Body[1] != nil) && (traingle.Body[2] != nil) {
		context.Img.DrawTriangles(
			[]ebiten.Vertex{
				ebiten.Vertex{
					DstX : float32( traingle.Body[0].X+context.Bound.Min.X-context.Sub.X ),
					DstY : float32( traingle.Body[0].Y+context.Bound.Min.Y-context.Sub.Y ),
					ColorR : 1,
					ColorG : 1,
					ColorB : 1,
					ColorA : 1,
				},
				ebiten.Vertex{
					DstX : float32( traingle.Body[1].X+context.Bound.Min.X-context.Sub.X ),
					DstY : float32( traingle.Body[1].Y+context.Bound.Min.Y-context.Sub.Y ),
					ColorR : 1,
					ColorG : 1,
					ColorB : 1,
					ColorA : 1,
				},
				ebiten.Vertex{
					DstX : float32( traingle.Body[2].X+context.Bound.Min.X-context.Sub.X ),
					DstY : float32( traingle.Body[2].Y+context.Bound.Min.Y-context.Sub.Y ),
					ColorR : 1,
					ColorG : 1,
					ColorB : 1,
					ColorA : 1,
				},
			},
			[]uint16{
				0,
				1,
				2,
			},
			pixel,
			Options,
		)
		
	}
}
func (traingle Traingle) GetMetrics()(e element.Metrick){
	return 
}
func MakeTraingle(body [3]*image.Point) (traingle *Traingle){
	traingle = &Traingle{
		Body : body, 
	}
	return
}