package bisnes
import (
	"MyProjectSrc/elements"
	"MyProjectSrc/elements/zone"
	"MyProjectSrc/elements/plan"
	"MyProjectSrc/elements/word"
	"MyProjectSrc/elements/sum"
	"MyProjectSrc/elements/line"
	"MyProjectSrc/elements/button"
	"MyProjectSrc/elements/image"
	"MyProjectSrc/elements/use"
	"MyProjectSrc/context"
	"MyProjectSrc/elements/do"
	"MyProjectSrc/elements/que"
	"MyProjectSrc/data"
	"MyProjectSrc/elements/prompt"
	"MyProjectSrc/elements/border"
	"image"
	"math/rand"
	"strings"
	"time"
	"strconv"
	"github.com/itchyny/timefmt-go"
	"sync"
	
	//"fmt"
)
var (
	RandSourse = rand.NewSource(0)
	RandContext  = rand.New(RandSourse)
	
	SeedI int64 = int64(time.Now().Nanosecond())
	
	Personal chan Person
	LenghtWay int = 600-40
	TimeStart = time.Now()
	TimeNow = time.Now()
	//VorcerCount int = GetInt(11)+8
	//MassCount int = GetInt(16)+20
	VorcerCount int = 20
	MassCount int = 30
	Many int = 40000
	MaxMass chan Person
	ManyMutex sync.Mutex
	OneRand   sync.Mutex
	
	ManyRecord int = 0
	
	
	
	
	Speed time.Duration = 20000
	
	namat = elementWord.MakeLongWord(3,"")
)

func Start(name string){
	namat.Reset("Ф.И.О: "+name)
}

func GetDataStrD(t time.Time)(string){
	return timefmt.Format( t, `%d.%m.%Y`)
}
func GetDataStrT(t time.Time)(string){
	return timefmt.Format( t, `%H:%M`)
}
func init(){
	RandContext.Seed(int64(time.Now().Sub(time.Time{})))
	Personal = make(chan Person, VorcerCount)
	for i:=0;i<VorcerCount;i++{
		Personal <- Person{}
	}
	MaxMass = make(chan Person, MassCount)
	//fmt.Println(MassCount)
	go func(){
		for{
			time.Sleep(25 * time.Millisecond)
			TimeNow = TimeStart.Add(time.Now().Sub(TimeStart)*Speed )
			
			
			if ManyRecord < Many{
				ManyRecord = Many
			}
		}
	}()
}
func GetInt(max int)(i int){  // RenderMutex.Lock()
	OneRand.Lock()
	SeedI += 100+int64(max*3)
	RandContext.Seed(SeedI)
	i = RandContext.Intn(max)
	OneRand.Unlock()
	return
}
type Person struct{}
type Company struct{
	Name string
}
type Buisnes []*Company
type Obj struct{
	Name string
	Count int
}
type Mutex struct{}
type Poster struct{
	Company *Company
	Name string
	Descript string
	Mass []Obj
	StartData time.Time
	EndData time.Time
}
func (buisnes *Buisnes) GetPoster()(poster *Poster) {
	poster = new(Poster)
	poster.Company = ((*buisnes)[GetInt(len(*buisnes))])
	var tovarTypes = []string{
		"Пакет",
		"Тележка",
		"Ящик",
		"Балон",
		"Паддон",
	}
	for i:=0;i<GetInt(7)+3;i++{
		massT := GetInt(len(tovarTypes))
		poster.Mass = append(poster.Mass, Obj{
			Count : 1,
			Name : tovarTypes[massT],
		})
	}
	poster.StartData = TimeNow.AddDate(0,0,GetInt(10)+4)
	poster.EndData = poster.StartData.AddDate(0,0,GetInt(10)+4)
	var nameF = [][]string{
		[]string{
			"Контракт на хранение",
			"Требуется помещение для хранения",
			"Нужно место для хранения",
			"Ищем помещение для хранения",
		},
		[]string{
			GetDataStrD(poster.StartData),
		},
		[]string{
			"на переод до",
			"до",
			"до даты",
			"к",
		},
		[]string{
			GetDataStrD(poster.EndData),
		},
		[]string{
			"",
			"",
			"",
			"Срочно",
		},
	}
	for i:=range nameF{
		poster.Name += nameF[i][GetInt(len(nameF[i]))]
		poster.Name += " "
	}
	poster.Name = strings.TrimSpace(poster.Name)
	poster.Name += "."
	return
}
func Organisation(count int)(buisnes *Buisnes){
	buisnes = new(Buisnes)
	(*buisnes) = make(Buisnes, 0, count)
	for i := 0; i< count;i++ {
		(*buisnes) = append((*buisnes), Set())
	}
	return
}
func Set()(company *Company){
	company = new(Company)
	var nameGen = [][]string{
		[]string{
			"Гор",
			"Росс",
			"Мега",
			"Тор",
			"Воз",
			"Нано",
		},
		[]string{
			"Бод",
			"Строй",
			"Воз",
			"Билд",
			"Шифт",
			"Спейс",
			"Энджин",
			"Товары",
			"Город",
			"",
			"",
			"",
		},
		[]string{
			"Накт",
			"Найт",
			"Сейф",
			"Люкс",
			"",
			"",
			"",
			"",
			"",
			"",
		},
		[]string{
			"Лэст",
			"Эф",
			"-Промо",
			"-Ру",
			"",
			"",
		},
	}
	company.Name += `«`
	for i:= range (nameGen){
		if GetInt(2) == 0 {
			company.Name += strings.ToLower( 
				nameGen[i][GetInt( len(nameGen[i]) )],
			)
		}else{
			company.Name += nameGen[i][GetInt( len(nameGen[i]) )]
		}
	}
	company.Name += `»`
	return
}
func MakeTextBoxe(texts string, sg int, t1, t2 bool)(element.Element){
	lists := elementPlan.MakePoseSimple(
		elementLine.MakeLine(
			t1, 
			t2, 
			100, 
			elementWord.MakeAbc(
				sg, 
				" ", 
				texts,
			)...,
		), 
		sg, 
		sg,
	)
	Zons := elementZone.MakeZoneUse(sg, image.Point(lists.GetMetrics()).Add(image.Point{sg,sg}), nil)
	return elementSum.MakeSum(Zons,lists)
}
func (buisnes *Buisnes) Gen(
	count, hight, sg int, 
	addeders func(element.Element)(),
)(
	elements []element.Element, 
	vis []element.Element, 
	worker element.Element, 
	adch chan Person, 
	PromtLauerUI *elementSum.Sum,
){
	PromtLauerUI = elementSum.MakeSum()
	listG := elementLine.MakeLine(
		0==1,
		0==1, 
		90, 
		nil,
	)
	vis = append(vis,listG)
	datE := elementWord.MakeLongWord(sg, "ДАТА")
	datT := elementWord.MakeLongWord(sg, "Время")
	datM := elementWord.MakeLongWord(sg, "Деньги")
	datR := elementWord.MakeLongWord(sg, "Рабочие")
	datS := elementWord.MakeLongWord(sg, "Склад")
	
	
	promtPoster0, promtBody0 := elementQue.MakeQue(
		elementBorder.MakeBorder(
			elementSum.MakeSum(
				elementPlan.MakePoseSimple(
					elementWord.MakeLongWord(3, "Данные о игроке"),
					0,
					0,
				),
			),
			4,
		),
	)
	promtPoster1, promtBody1 := elementQue.MakeQue(
		elementBorder.MakeBorder(
			elementSum.MakeSum(
				elementPlan.MakePoseSimple(
					elementWord.MakeLongWord(3, "Показатель времени : дата"),
					0,
					0,
				),
			),
			4,
		),
	)
	promtPoster2, promtBody2 := elementQue.MakeQue(
		elementSum.MakeSum(
		elementBorder.MakeBorder(
			elementSum.MakeSum(
				elementPlan.MakePoseSimple(
					elementWord.MakeLongWord(3, "Показатель времени : часы и минуты"),
					0,
					0,
				),
			),
			4,
		),
		),
	)
	promtPoster3, promtBody3 := elementQue.MakeQue(
		elementSum.MakeSum(
		elementBorder.MakeBorder(
			elementSum.MakeSum(
				elementPlan.MakePoseSimple(
					elementWord.MakeLongWord(3, "Показатель времени : сумма денег на счете организации"),
					0,
					0,
				),
			),
			4,
		),
		),
	)
	promtPoster4, promtBody4 := elementQue.MakeQue(
		elementSum.MakeSum(
		elementBorder.MakeBorder(
			elementSum.MakeSum(
				elementPlan.MakePoseSimple(
					elementWord.MakeLongWord(3, "Показатель времени : загруженность рабочих склада"),
					0,
					0,
				),
			),
			4,
		),
		),
	)
	promtPoster5, promtBody5 := elementQue.MakeQue(
		elementSum.MakeSum(
		elementBorder.MakeBorder(
			elementSum.MakeSum(
				elementPlan.MakePoseSimple(
					elementWord.MakeLongWord(3, "Показатель времени : загруженность склада"),
					0,
					0,
				),
			),
			4,
		),
		),
	)
	PromtLauerUI.Add(promtBody1,promtBody2,promtBody3,promtBody4,promtBody5, promtBody0)
	
	worker = elementSum.MakeSum(
		elementZone.MakeZoneUse(sg, image.Point{400,150}, nil),
		
		elementPlan.MakePoseSimple(
				elementLine.MakeLine(
					0==1,
					0==0, 
					900,
					elementLine.MakeLine(
						0==1,
						0==0, 
						900,
						namat,
						promtPoster0,
					),
					elementLine.MakeLine(
						0==1,
						0==0, 
						900,
						elementSum.MakeSum(
							elementImage.MakeImage(data.Time2, 20,20), 
							elementPlan.MakePoseSimple(
								datE, 
								18, 
								0,
							),
						),
						promtPoster1,
					),
					elementLine.MakeLine(
						0==1,
						0==0, 
						900,
						elementSum.MakeSum(
							elementImage.MakeImage(data.Time1, 20,20), 
							elementPlan.MakePoseSimple(
								datT, 
								18, 
								0,
							),
						),
						promtPoster2,
					),
					elementLine.MakeLine(
						0==1,
						0==0, 
						900,
						elementSum.MakeSum(
							elementImage.MakeImage(data.Many, 20,20), 
							elementPlan.MakePoseSimple(
								datM, 
								18, 
								0,
							),
						),
						promtPoster3,
					),
					elementLine.MakeLine(
						0==1,
						0==0, 
						900,
						elementSum.MakeSum(
							elementPlan.MakePoseSimple(
								datR, 
								0, 
								0,
							),
						),
						promtPoster4,
					),
					elementLine.MakeLine(
						0==1,
						0==0, 
						900,
						elementSum.MakeSum(
							elementPlan.MakePoseSimple(
								datS, 
								0, 
								0,
							),
						),
						promtPoster5,
					),
			),
			10,
			10,
		),
		elementDo.MakeDo(
			func(contexte *context.Context){
				datE.Reset("Дата :"+GetDataStrD(TimeNow))
				datT.Reset("Время :"+GetDataStrT(TimeNow))
				datM.Reset("Бюджет :"+strconv.Itoa(Many)+".р")
				datR.Reset("•Занято рабочих :"+strconv.Itoa(VorcerCount-len(Personal))+" из "+strconv.Itoa(VorcerCount))
				datS.Reset("•Занято склада :"+strconv.Itoa( int(float64(len(MaxMass))/float64(MassCount)*100) )+"%")
			},
		),
	)
	var wait = func(count int){
		time.Sleep(17 * time.Millisecond*time.Duration(int64(GetInt(6)+1)))
		ManyMutex.Lock()
		Many -= int(float64(count)*1.04)
		ManyMutex.Unlock()
	}
	PromtFotMainButton := func()(ee element.Element){
		ee = elementSum.MakeSum(
			elementPlan.MakePoseSimple(
				elementWord.MakeLongWord(3, "Эта кнопка для подписания договора"),
				0,
				0,
			),
			elementPlan.MakePoseSimple(
				elementImage.MakeImage(
					data.Tutorial1, 
					100,
					100,
				),
				0,
				30,
			),
		)
		ee = elementBorder.MakeBorder(ee, 10)
		return 
	}()
	adde := func(){
		//fmt.Println("---")
		uiPromt1, controlPromt1 := elementPromt.MakePromt(PromtFotMainButton, 1)
		post := buisnes.GetPoster()
		Zon := elementZone.MakeZoneUse(sg, image.Point{hight, 150}, nil)
		texNameCompany := MakeTextBoxe(post.Company.Name,sg,0==1,0==0)
		texName := MakeTextBoxe(post.Name,sg,0==1,0==1)
		del := elementUse.MakeUse(elementDo.MakeDo())
		go func(){
			time.Sleep( post.StartData.Sub(TimeNow)/Speed )
			//fmt.Println("Poster dell")
			del.Use(1==0)
			adch <- Person{}
		}()
		ok := elementButton.MakeTextButton(
			"Принять", 
			sg, 
			sg,
			func(button *elementButton.Button){
				controlPromt1.Del()
				del.Use(1==0)
				for m:=range post.Mass {
					col := elementZone.MakeZoneUse(
						sg, 
						image.Point{
							40,
							40,
						}, 
						nil,
					)
					var mesh []element.Element
					for b:=0;b<post.Mass[m].Count;b++ {
						mesh = append(
							mesh,
							elementPlan.MakePoseSimple(elementImage.MakeImage(data.ImageMashine,40,40),40*(b+1),0),
						)
					}
					
					images := elementImage.MakeImage(
						data.Wait,
						34,
						34,
					)
		
					rd := elementUse.MakeUse(elementSum.MakeSum(mesh...))
					elmP := elementPlan.MakePose(
						elementSum.MakeSum(
							col, 
							rd,
							elementPlan.MakePoseSimple(
								images,
								3,
								3,
								),
							),
						LenghtWay,0,
					)
					forDel := elementUse.MakeUse(
						elementSum.MakeSum(
							elementPlan.MakePoseSimple(elementZone.MakeZoneUse(sg, image.Point{600,10}, nil), -sg, 20),
							elementPlan.MakePoseSimple(elementWord.MakeLongWord(sg, post.Name), 0, -5),
							elementPlan.MakePlane(elmP),
						),
					)
					listG.AppendBody(forDel)
					m:=m
					go func(){
						
						//fmt.Println("Time to wait: ",post.StartData.Sub(time.Now())/Speed)
						//fmt.Println("Time after: ", TimeNow)
						//fmt.Println("Time wait: ", (time.Now()).Add(post.StartData.Sub(time.Now())))
						forDel.Use(1==1)
						time.Sleep( post.StartData.Sub(TimeNow)/Speed )
						//fmt.Println("Time now: ", TimeNow)
						//fmt.Println("Time needed waiting:", post.StartData)
						t1 := time.Now()
						col.Mod = &data.DefultModC
						images.Img = data.Need
						for f := 0; f < post.Mass[m].Count; f++{
							<-Personal
						}
						//fmt.Println("Go")
						rd.Use(1==1)
						col.Mod = &data.DefultModB
						images.Img = data.Go
						for l1:=0;l1<LenghtWay-sg*2;l1++{
							wait(post.Mass[m].Count)
							(*elmP).Pos.X -= 1
						}
						for f := 0; f < post.Mass[m].Count; f++{
							Personal<-Person{}
						}
						//fmt.Println("Stay")
						rd.Use(1==0)
						col.Mod = &data.DefultModA
						MaxMass <- Person{}
						t2 := time.Now()
						images.Img = data.Wait
						
						dgg := make(chan int)
						
						go func(){
							for {
								select{
									case <-dgg:
										return
									default:
										time.Sleep( time.Second/10 )
										ManyMutex.Lock()
										Many -= 1;
										//fmt.Println("Mass -")
										ManyMutex.Unlock()
								}
							}
						}()
						time.Sleep( post.EndData.Sub(TimeNow)/Speed )
						images.Img = data.Need
						col.Mod = &data.DefultModC
						dgg <- int(1)
						<-MaxMass
						t12 := time.Now()
						rd.Use(1==1)
						for f := 0; f < post.Mass[m].Count; f++{
							<-Personal
						}
						col.Mod = &data.DefultModB
						images.Img = data.Go
						for l2:=0;l2<LenghtWay-sg*2;l2++{
							wait(post.Mass[m].Count)
							(*elmP).Pos.X += 1
						}
						t22 := time.Now()
						images.Img = data.Wait
						for f := 0; f < post.Mass[m].Count; f++{
							Personal<-Person{}
						}
						forDel.Use(1==0)
						ManyMutex.Lock()
						
						s := (0.7+float64((t2.Sub(t1) + t22.Sub(t12))*Speed/time.Second/60/60)/260)
						
						s *= s
						
						s *= 460
						
						s = 3800+float64(100*post.Mass[m].Count) - s
						
						Many += int(s);
						ManyMutex.Unlock()
						
					}()
				}
			},
			nil,nil,nil,
			controlPromt1,
		)
		oks := elementPlan.MakePoseSimple(ok, hight-ok.GetMetrics().X-sg*2, sg)
		Deskr := MakeTextBoxe(post.Descript,sg,0==1,0==1)
		texDeskr := elementPlan.MakePoseSimple(
			Deskr,
			hight-Deskr.GetMetrics().X,
			sg+ok.GetMetrics().Y,
		)
		del = elementUse.MakeUse(
			elementSum.MakeSum(
				Zon, 
				texNameCompany,
				elementPlan.MakePoseSimple(
					elementSum.MakeSum(
						texName, 
						oks, 
						texDeskr,
					), 0, texNameCompany.GetMetrics().Y,
				),
			),
		)
		del.Use(1==1)
		addeders(del)
		PromtLauerUI.Add(uiPromt1)
	}
	adch = make(chan Person, 1)
	//fmt.Println()
	go func(){
		for {
			<- adch
			adde()
		}
	}()
	for i:=0;i<count;i++{
		//fmt.Println("Poster create")
		adch <- Person{}
	}
	return
}