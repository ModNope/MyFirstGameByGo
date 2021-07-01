package contextActive
import (
	"MyProjectSrc/context"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image"
	//"fmt"
)
var (
	Last []*ActiveZone
	All []*ActiveZone
	x,y int
	lastI *ActiveZone
)
var (
	Def Effect = 0
	Colise Effect = 1
	Click Effect = 2
	ClickDo Effect = 3
)
type Effect int
type ActiveZone struct{
	Zone, ContextZone image.Rectangle
	Update bool
	lv, lvf int
	Handler func(Effect)()
}
func (active *ActiveZone) Up(zone image.Rectangle, context *context.Context){
	active.Update = 1==1
	active.lv = context.Lv
	active.lvf = context.Lvf
	active.Zone = zone
	active.ContextZone = context.Bound
	
}
func (active ActiveZone) ZoneMouse(p image.Point){
	if (active.Update){
		if (p.X > active.Zone.Min.X)&&(p.Y > active.Zone.Min.Y)&&(p.X < active.Zone.Max.X)&&(p.Y < active.Zone.Max.Y){
			if (p.X > active.ContextZone.Min.X)&&(p.Y > active.ContextZone.Min.Y)&&(p.X < active.ContextZone.Max.X)&&(p.Y < active.ContextZone.Max.Y){
				Last = append(Last, &active)
				active.Update = 1==0
			}
		}
	}
}
func Do(){
	//fmt.Println("<<<")
	if lastI != nil{
		if lastI.Handler != nil{
			lastI.Handler(Def)
		}
	}
	lastI = nil
	
	x,y = ebiten.CursorPosition()
	//fmt.Println(image.Point{x,y})
	Last = make([]*ActiveZone, 0, len(All))
	var lastLv, lastLvf int
	for i := range All{
		All[i].ZoneMouse(image.Point{x,y})
	}
	//if len(Last) != 0{
		//fmt.Println(">>>", Last[0])
	//}
	
	for i := range Last{
		//fmt.Println(i)
		//fmt.Println(Last[i].lv,Last[i].lvf)
		//if (Last[i].Update){
			if lastLv <= Last[i].lv{
				lastLv = Last[i].lv
				if lastLvf <= Last[i].lvf{
					
					lastLvf = Last[i].lvf
					lastI = Last[i]
				}
			}
		//}
	}
	
	if lastI != nil{
		//fmt.Println(&lastI)
		if lastI.Handler != nil{
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft){
				lastI.Handler(ClickDo)
			}else{
				if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft){
					lastI.Handler(Click)
				}else{
					lastI.Handler(Colise)
				}
			}
			
		}
	}
}
func MakeActive()(*ActiveZone){
	var ac ActiveZone
	All = append(All, &ac)
	return &ac
}