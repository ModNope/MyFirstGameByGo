package End
import (
	"MyProjectSrc/elements"
	//"MyProjectSrc/elements/button"
	//"MyProjectSrc/elements/image"
	"MyProjectSrc/elements/plan"
	//"MyProjectSrc/elements/sum"
	"MyProjectSrc/elements/word"
	"MyProjectSrc/data"
)
func END()(elm element.Element){
	text := elementWord.MakeLongWordFont(3, "Игра завалена!", &data.MainFontUseMeny)
	elm = elementPlan.MakePoseSimple(
		text,
		(800-text.GetMetrics().X)/2,
		(700-text.GetMetrics().Y)/2,
	)
	return
}