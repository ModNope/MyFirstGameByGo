package UI

import (
	"image"
	"MyProjectSrc/elements/privat"
	"MyProjectSrc/elements/zone"
	"MyProjectSrc/elements/line"
	"MyProjectSrc/elements/plan"
	"MyProjectSrc/elements/button"
	"MyProjectSrc/elements/word"
	"MyProjectSrc/context/active"
	"MyProjectSrc/elements"
	"MyProjectSrc/elements/do"
	"MyProjectSrc/elements/sum"
	"MyProjectSrc/elements/image"
	"MyProjectSrc/context"
	"MyProjectSrc/data"
	//"fmt"
	"time"
	"github.com/hajimehoshi/ebiten"
	"MyProjectSrc/bisnes"
)

func MainUiMake(Stop *bool)(e element.Element, many *int, manyD func()()){
	
	//fmt.Print("GenUi")
	org := bisnes.Organisation(150)
	many = &bisnes.Many
	manyD = func(){
		bisnes.ManyMutex.Lock()
		bisnes.Many -= 2;
		bisnes.ManyMutex.Unlock()
	}
	Market, list, Pole, poleList:= MakeWind("Выбор предложения:", 280, 285, 3)
	poster, vis, workers, add, PromtLauerUI := org.Gen(10, 280-3*2, 3, func(e element.Element)(){
		list.AppendBody(e)
	})
	list.AppendBody(poster...)
	(*poleList) = append((*poleList), vis...)
	go func(){
		for{
			time.Sleep(90 * time.Second)
			//fmt.Println("Poster moor")
			add<-bisnes.Person{}
		}
	}()
	
	butt := elementButton.MakeTextButtonFont(
		"Стоп", 
		3, 
		3,
		func(button *elementButton.Button){
			*Stop = (1==1)
		},
		nil,
		nil,
		nil,
		&data.MainFontUse,
	)
	
	e = elementSum.MakeSum(
		elementPlan.MakePlane(
			elementPlan.MakePose(
				elementImage.MakeImage(data.OldFon, 800,700),
				0,
				0,
			),
			elementPlan.MakePose(
				Pole,
				800-700,
				300,
			),
			elementPlan.MakePose(
				Market,
				800-290,
				10,
			),
			elementPlan.MakePose(
				workers,
				800-750,
				145,
			),
			elementPlan.MakePose(
				butt,
				0,
				700-butt.GetMetrics().Y,
			),
		),
		PromtLauerUI,
	)
	return
}

func Up(){
	contextActive.Do()
}

func Clamp(xMin, xMax int, c int)(r int){
	if xMax < c{
		r = xMax
		return
	}
	if xMin > c{
		r = xMin
		return
	}
	r = c
	return
}
func Min(min ...int)(r int){
	for i:=range min{
		if min[i] < r{
			r = min[i]
		}
	}
	return
}
func Max(max ...int)(r int){
	for i:=range max{
		if max[i] > r{
			r = max[i]
		}
	}
	return
}

func MakeWind(name string, sx, sy, sg int)(e element.Element, lists *elementLine.Line, Pole element.Element, poleList *elementSum.Sum ){
	poleList = elementSum.MakeSum()
	poleListPoseR := elementPlan.MakePose(poleList, 0, 0)
	poleListPoseRE := elementPlan.MakePlane(poleListPoseR)
	activeHolst := contextActive.MakeActive()
	activeHolst.Handler = func(effect contextActive.Effect){
		if effect != contextActive.Def{
			var H = -Max( 0, - 400 + poleList.GetMetrics().Y )
			var _, addX = ebiten.Wheel()
			(*poleListPoseR).Pos.Y = Clamp(H, 0, (*poleListPoseR).Pos.Y+int(addX*50) )
		}
	}
	doiHolst := elementDo.MakeDo(
		func(contexte *context.Context){
			(activeHolst).Up(
				image.Rectangle{
					Min:contexte.Bound.Min, 
					Max:image.Point{600, 400}.Add(contexte.Bound.Min),
				},
				contexte,
			)
		})
	Pole = elementSum.MakeSum(
		elementZone.MakeZoneUse(sg, image.Point{600, 400}, &data.DefultModD),
		elementPlan.MakePoseSimple(elementPrivat.MakePrivatZone(image.Point{600-sg*2, 400-sg*2},elementSum.MakeSum(poleListPoseRE, doiHolst)),sg,sg),
	)
	zone := elementZone.MakeZoneUse(sg, image.Point{sx, sy}, nil)
	z := elementPlan.MakePose(zone,0,0)
	text := elementWord.MakeLongWord(sg, name)
	l := elementPlan.MakePose(text,(sx-text.GetMetrics().X)/2,sg)
	lists = elementLine.MakeLine(0==1,0==0, 0, nil)
	WinElH := sy-text.GetMetrics().Y-sg*6
	SkrolPose := elementPlan.MakePose(lists ,0,0)
	contentRePose := elementPlan.MakePlane(SkrolPose)
	phoneZone := elementZone.MakeZoneUse(sg, image.Point{X : sx-sg*2, Y : WinElH+sg,}, &data.DefultModD)
	active := contextActive.MakeActive()
	active.Handler = func(effect contextActive.Effect){
		if effect != contextActive.Def{
			var H = -Max( 0, -WinElH+ lists.GetMetrics().Y )
			var _, addX = ebiten.Wheel()
			SkrolPose.Pos.Y = Clamp(H, 0, SkrolPose.Pos.Y+int(addX*50) )
		}
	}
	doi := elementDo.MakeDo(
		func(contexte *context.Context){
			(active).Up(
				image.Rectangle{
					Min:contexte.Bound.Min, 
					Max:image.Point{
						X : sx-sg*4, 
						Y : sy-text.GetMetrics().Y-sg*6,
					}.Add(contexte.Bound.Min),
				},
				contexte,
			)
		})
	render := elementSum.MakeSum(contentRePose, doi)
	contentZone := elementPrivat.MakePrivatZone( image.Point{
		X : sx-sg*4, 
		Y : WinElH-sg,
	}, render)
	contentZoneP := elementPlan.MakePoseSimple(contentZone, sg, sg)
	sumZoneR := elementSum.MakeSum(phoneZone, contentZoneP)
	c := elementPlan.MakePose(sumZoneR , sg, text.GetMetrics().Y+sg*3)
	e = elementPlan.MakePlane(z, l, c)
	return
}