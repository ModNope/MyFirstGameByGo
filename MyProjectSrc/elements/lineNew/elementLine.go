package elementLine
import (
	"MyProjectSrc/elements"
	"MyProjectSrc/context"
	//"image"
)
type Line struct{
	Body []element.Element
	Direction bool
	resize bool
}

func (line Line) Render(context *context.Context){
	line.BuildLines()
	
	cop := context.CopyLocal()
	for i,_:= range line.Body{
		pe := line.Body[i].GetMetrics()
		if line.Direction{
			cop.Bound.Min.X += pe.X
		}else{
			cop.Bound.Min.X += pe.Y
		}
		line.Body[i].Render(cop.CopyLocal())
	}
}

func (line *Line) GetMetrics()(m element.Metrick){
	return 
}

func(line *Line)BuildLines()(){
	if line.resize{
		bufer []element.Element
		
		for i,_ := range line.Body{
			
		}
		
	}
}
func(line *Line)AppendBody(elements ...element.Element)(s, count int){
	s = len(line.Body)
	count = len(elements)
	line.Body = append(line.Body,elements...)
	line.BuildLines()
	return
}
/*
func(line *Line)GetBody()(elements []element.Element){
	elements = line.Body
	return
}

func(line *Line)SetBody(elements ...element.Element)(){
	line.Body = elements
	line.BuildLines()
}

func(line *Line)GetBodyId(i int)(elements element.Element){
	elements = line.Body[i]
	return
}

func(line *Line)DelId(i int)(){
	line.Body = append(line.Body[:i],line.Body[i+1:]...)
	line.BuildLines()
}
func(line *Line)Len()(i int){
	i = len(line.Body)
	return
}
*/
// resize If true - make new line.
// vector If true - xorisontal line
func MakeLine(resize, vector bool, lenght int, elements ...element.Element)(l *Line){
	l = &Line{
		Body : elements,
		Direction : vector,
		resize : resize,
	}
	return
}