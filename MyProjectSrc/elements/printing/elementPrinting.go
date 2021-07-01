package elementPrinting
import (
	"MyProjectSrc/elements"
	"MyProjectSrc/context"
	"MyProjectSrc/elements/button"
	"MyProjectSrc/elements/line"
	//"MyProjectSrc/elements/plan"
	"MyProjectSrc/elements/sum"
	//"MyProjectSrc/elements/image"
	//"MyProjectSrc/elements/floot"
	"MyProjectSrc/elements/zone"
	"image"
	"MyProjectSrc/elements/plan"
	//"MyProjectSrc/elements/border"
	"MyProjectSrc/elements/os"
	//"MyProjectSrc/elements/simafor"
	"MyProjectSrc/elements/do"
	"MyProjectSrc/elements/word"
	"MyProjectSrc/data"
	"strings"
	"time"
	"unicode/utf8"
	//"fmt"
)
func RemoveLastChar(str string) string {
      for len(str) > 0 {
              _, size := utf8.DecodeLastRuneInString(str)
              return str[:len(str)-size]
      }
      return str
}
func MakePrinting(str *[]rune)(Printing element.Element, UpdateString element.Element){
	
	//reuse := elementSimafor.MakeSimafor()
	
	sizeChar := 40
	ABCru := [][]string{
		[]string{
			"й",
			"ц",
			"у",
			"к",
			"е",
			"н",
			"г",
			"ш",
			"щ",
		},
		[]string{
			"з",
			"х",
			"ъ",
			"ф",
			"ы",
			"в",
			"а",
			"п",
		},
		[]string{
			"р",
			"о",
			"л",
			"д",
			"ж",
			"э",
			"я",
			"ч",
		},
		[]string{
			"с",
			"м",
			"и",
			"т",
			"ь",
			"б",
			"ю",
		},
	}
	
	lime1 := [4][]element.Element{}
	
	
	upReg := elementOs.MakeOs(
		sizeChar,
		sizeChar,
		elementButton.MakeTextButtonFont(
			"<-",
			0,
			5,
			func(button *elementButton.Button){
				*str = []rune(RemoveLastChar(string(*str)))
			},
			nil,nil,nil,
			&data.MainFontPrinter,
		),
	)
	slash := elementOs.MakeOs(
		sizeChar,
		sizeChar,
		elementButton.MakeTextButtonFont(
			"_",
			0,
			5,
			func(button *elementButton.Button){
				*str = append(*str, []rune(" ")...)
			},
			nil,nil,nil,
			&data.MainFontPrinter,
		),
	)
	point := elementOs.MakeOs(
		sizeChar,
		sizeChar,
		elementButton.MakeTextButtonFont(
			".",
			0,
			5,
			func(button *elementButton.Button){
				*str = append(*str, []rune(".")...)
			},
			nil,nil,nil,
			&data.MainFontPrinter,
		),
	)
	
	
	for u,_ := range ABCru {
		for i,_ := range ABCru[u] {
			abc := ABCru[u][i]
			lime1[u] = append(
				lime1[u],
				elementSum.MakeSum(
					elementZone.MakeZoneUse(
						3,
						image.Point{sizeChar*2,sizeChar}, 
						&data.DefultModD,
					),
					elementPlan.MakePoseSimple(
						elementLine.MakeLine(
							0==1,
							0==1, 
							0, 
							elementOs.MakeOs(
								sizeChar,
								sizeChar,
								elementButton.MakeTextButtonFont(
									ABCru[u][i],
									0,
									5,
									func(button *elementButton.Button){
										*str = append(*str, []rune(abc)...)
									},
									nil,nil,nil,
									&data.MainFontPrinter,
								),
							),
							elementOs.MakeOs(
								sizeChar,
								sizeChar,
								elementButton.MakeTextButtonFont(
									strings.Title(ABCru[u][i]),
									0,
									5,
									func(button *elementButton.Button){
										*str = append(*str, []rune(strings.Title(abc))...)
									},
									nil,nil,nil,
									&data.MainFontPrinter,
								),
							),
						),
						0,
						0,
					),
				),
			)
		}
	}
	lime1[2] = append(
		lime1[2],
		elementSum.MakeSum(
			elementZone.MakeZoneUse(
				3,
				image.Point{sizeChar,sizeChar}, 
				&data.DefultModD,
			),
			slash,
		),
	)
	lime1[3] = append(
		lime1[3],
		elementSum.MakeSum(
			elementZone.MakeZoneUse(
				3,
				image.Point{sizeChar*2,sizeChar}, 
				&data.DefultModD,
			),
			elementLine.MakeLine(
				0==1,
				0==1, 
				0, 
				point,
				upReg,
			),
		),
	)
	//lime1[3] = append(lime1[3], upReg)
	
	DFG := elementLine.MakeLine(
		0==1,
		0==0, 
		0, 
		elementOs.MakeOs(
			len(ABCru[1])*(sizeChar+3)+3,
			sizeChar,
			elementLine.MakeLine(
				0==1,
				0==1, 
				0, 
				lime1[0]...,
			),
		),
		elementOs.MakeOs(
			len(ABCru[1])*(sizeChar+3)+3,
			sizeChar,
			elementLine.MakeLine(
				0==1,
				0==1, 
				0, 
				lime1[1]...,
			),
		),
		elementOs.MakeOs(
			len(ABCru[1])*(sizeChar+3)+3,
			sizeChar,
			elementLine.MakeLine(
				0==1,
				0==1, 
				0, 
				lime1[2]...,
			),
		),
		elementOs.MakeOs(
			len(ABCru[1])*(sizeChar+3)+3,
			sizeChar,
			elementLine.MakeLine(
				0==1,
				0==1, 
				0, 
				lime1[3]...,
			),
		),
	)
	
	Printing = elementSum.MakeSum(
		elementZone.MakeZoneUse(
			3,
			image.Point{740,170}, 
			&data.DefultModD,
		),
		elementPlan.MakePoseSimple(
			DFG,
			197,
			4,
		),
	)
	
	word := elementWord.MakeLongWord(3, string(*str))
	
	end := "_"
	
	go func(){
		for {
			time.Sleep(time.Millisecond * 500)
			end = ""
			time.Sleep(time.Millisecond * 500)
			end = "_"
		}
	}()
	
	UpdateString = elementSum.MakeSum(
		elementZone.MakeZoneUse(
			3,
			image.Point{600,30}, 
			&data.DefultModD,
		),
		elementPlan.MakePoseSimple(
			elementDo.MakeDo(
				func(contexte *context.Context){
					if len(*str) != 0 {
						word.Reset("Ваше Ф.И.О: "+ string(*str)+end)
					}else{
						word.Reset("Ваше Ф.И.О: ")
					}
					word.Render(contexte)
				},
			),
			3,
			3,
		),
	)
	
	
	
	
	
	
	return
}