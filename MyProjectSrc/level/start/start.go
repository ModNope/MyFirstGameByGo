package Start
import (
	"MyProjectSrc/elements"
	"MyProjectSrc/elements/button"
	"MyProjectSrc/elements/image"
	"MyProjectSrc/elements/plan"
	"MyProjectSrc/elements/sum"
	"MyProjectSrc/elements/printing"
	"MyProjectSrc/data"
	"MyProjectSrc/elements/os"
)
func Meny()(elm element.Element, b *bool, name *[]rune){
	
	name = new([]rune)
	b = new(bool)
	
	button := elementButton.MakeTextButtonFont(
		"Начать",
		3,
		3,
		func(button *elementButton.Button){
			*b = 1==1
		},
		nil,nil,nil,
		&data.MainFontUseMeny,
	)
	
	printerPlata, UpdateString := elementPrinting.MakePrinting(name);
	
	elm = elementSum.MakeSum(
		elementImage.MakeImage(data.OldFon, 800,700),
		elementPlan.MakePoseSimple(
			elementOs.MakeOs(
				800,
				100,
				UpdateString,
			),
			0,
			0,
		),
		elementPlan.MakePoseSimple(
			button,
			800/2-button.GetMetrics().X/2,
			700/2-button.GetMetrics().Y/2,
		),
		elementPlan.MakePoseSimple(
			printerPlata,
			800/2-printerPlata.GetMetrics().X/2,
			700/2-printerPlata.GetMetrics().Y/2+button.GetMetrics().Y*2,
		),
	)
	
	
	return
}