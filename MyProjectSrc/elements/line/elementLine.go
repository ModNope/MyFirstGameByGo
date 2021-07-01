package elementLine
import (
	"MyProjectSrc/elements"
	"MyProjectSrc/context"
	"image"
	
	"sync"
	//"fmt"
)
var(
	Mainx   sync.Mutex
)
type Line struct{
	Body []element.Element
	AddBufer []element.Element
	BodyLines []Lines
	resize, vector bool
	Lenght int
}
type Lines struct{
	Body []element.Element
	Hight element.Metrick
}

func (line Line) Render(context *context.Context){
	line.BuildLines()
	var res, h int
	var leng = line.GetMetrics().X
	copyL := context.Copy(nil)
	for i := range line.BodyLines{
		res = 0
		for f := range line.BodyLines[i].Body{
			var copyContext = copyL.CopyLocal()
			siz := line.BodyLines[i].Body[f].GetMetrics()
			if line.vector{
				if line.resize{
					copyContext.Bound.Min.X += int( float64(res)/float64(line.BodyLines[i].Hight.X)*float64(leng) )
				}else{
					copyContext.Bound.Min.X += res
				}
				
				copyContext.Bound.Min.Y += h
				res += siz.X
			}else{
				if line.resize{
					copyContext.Bound.Min.Y += int( float64(res)/float64(line.BodyLines[i].Hight.Y)*float64(leng) )
				}else{
					copyContext.Bound.Min.Y += res
				}
				copyContext.Bound.Min.X += h
				res += siz.Y
			}
			copyContext.Bound.Max = copyContext.Bound.Min.Add(image.Point(line.BodyLines[i].Body[f].GetMetrics()))
			line.BodyLines[i].Body[f].Render(copyContext)
		}
		if line.vector{
			h += line.BodyLines[i].Hight.Y
		}else{
			h += line.BodyLines[i].Hight.X
		}
	}
	
}

func (line *Line) GetMetrics()(m element.Metrick){
	//line.BuildLines()
	var hight, lenght int
	lenght = line.Lenght
	for i := range line.BodyLines{
		metrick := line.BodyLines[i].Hight
		if line.vector{
			hight += metrick.Y
			if lenght < metrick.X{
				lenght = metrick.X
			}
		}else{
			hight += metrick.X
			if lenght < metrick.Y{
				lenght = metrick.Y
			}
		}
		
	}
	if !line.vector{
		lenght, hight = hight, lenght
	}
	m.X, m.Y = lenght, hight
	
	return 
}
func(line *Line)BuildLines()(){
	//fmt.Println("Start 1")
	Mainx.Lock()
	//fmt.Println("Start 2")
	line.Body = append(line.Body, line.AddBufer...)
	line.AddBufer = []element.Element{}
	if (len(line.Body) > 0){
		var os int
		var id int
		line.BodyLines = make([]Lines, 1, len(line.Body))
		
		ElSize := line.Body[0].GetMetrics()
		line.BodyLines[0].Body = append(line.BodyLines[0].Body, line.Body[0]) // ---------------------
		line.BodyLines[0].Hight = ElSize
		if line.vector{
			os += ElSize.X
			stp := line.Body[1:]
			for i := range stp{
				ElSize = stp[i].GetMetrics()
				if os+ElSize.X > line.Lenght{
					id++
					line.BodyLines = append(line.BodyLines, Lines{})
					os = ElSize.X
				}
				line.BodyLines[id].Body = append(line.BodyLines[id].Body, stp[i])
				if line.BodyLines[id].Hight.Y < ElSize.Y{
					line.BodyLines[id].Hight.Y = ElSize.Y
				}
				line.BodyLines[id].Hight.X += ElSize.X
			}
		}else{
			os += ElSize.Y
			stp := line.Body[1:]
			for i := range stp{
				ElSize = stp[i].GetMetrics()
				if os+ElSize.Y > line.Lenght{
					id++
					line.BodyLines = append(line.BodyLines, Lines{})
					os = ElSize.Y
				}
				line.BodyLines[id].Body = append(line.BodyLines[id].Body, stp[i]) // 0--------
				if line.BodyLines[id].Hight.X < ElSize.X{
					line.BodyLines[id].Hight.X = ElSize.X
				}
				line.BodyLines[id].Hight.Y += ElSize.Y
			}
		}
	}
	//fmt.Println("End 1")
	Mainx.Unlock()
	//fmt.Println("End 2")
}

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

func(line *Line)AppendBody(elements ...element.Element)(s, count int){
	s = len(line.Body)
	count = len(elements)
	Mainx.Lock()
	line.AddBufer = append(line.AddBufer,elements...)
	Mainx.Unlock()
	line.BuildLines()
	return
}
func(line *Line)DelId(i int)(){
	Mainx.Lock()
	line.Body = append(line.Body[:i],line.Body[i+1:]...)
	Mainx.Unlock()
	line.BuildLines()
}
func(line *Line)Len()(i int){
	i = len(line.Body)
	return
}

// resize If true - make new line.
// vector If true - xorisontal line
func MakeLine(resize, vector bool, lenght int, elements ...element.Element)(l *Line){
	var newElm = make([]element.Element,0,len(elements))
	for i := range elements{
		if  elements[i] != nil{
			newElm = append(newElm, elements[i])
		}
	}
	
	l = &Line{
		resize : resize,
		vector : vector,
		Body : newElm,
		Lenght : lenght,
	}
	l.BuildLines()
	return
}