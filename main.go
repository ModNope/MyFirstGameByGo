package main
import (
	"github.com/hajimehoshi/ebiten"
	"MyProjectSrc/ui"
	"MyProjectSrc/elements"
	"MyProjectSrc/context"
	"MyProjectSrc/level/end"
	"MyProjectSrc/level/start"
	"MyProjectSrc/bisnes"
	"fmt"
	"os"
	"log"
	"path"
	"image"
	"time"
	"strconv"
	"MyProjectSrc/data"
	
)
var (
	game Prog
	many *int
	play *bool
	nalogi func()()
	sub int
	Stop *bool = new(bool)
	
	Name *[]rune
	
	end = func(){
		SaveInf := path.Join(".","Данные о игроке")
		if err := os.MkdirAll(SaveInf, os.ModePerm); err != nil {
			fmt.Println("Can`t make save dir")
			log.Fatal(err)
		}
		newe, err := os.Create(path.Join(SaveInf,"Информация о игроке "+string(*Name)+".txt"))
		if err != nil {
			log.Fatal(err)
		}
		defer newe.Close()
		time.Sleep(time.Second*1)
		bisnes.ManyMutex.Lock()
		newe.Write([]byte(
			"Ф.И.О: "+string(*Name)+
			string(10)+
			"Период игры с: "+bisnes.GetDataStrD(bisnes.TimeStart)+
			string(10)+
			"По: "+bisnes.GetDataStrD(bisnes.TimeNow)+
			string(10)+
			"Максимальный бюджет: "+strconv.Itoa(bisnes.ManyRecord)+
			string(10)+
			"Бюджет на момент конца игры: "+strconv.Itoa(bisnes.Many)+
			string(10),
		))
	}
	
	
)
type Prog struct{
	Screen element.Element
	Levels []element.Element
}
func (prog *Prog) Draw(scren *ebiten.Image) {
	var context = context.New(scren, nil)
	if prog.Screen != nil{
		prog.Screen.Render(context)
	}
}
func (prog *Prog) Update(scren *ebiten.Image) error {
	
	UI.Up()
	
	
	if *play{
		*play = 1==0
		*many = 40000
		game.Screen = game.Levels[1]
		bisnes.Start(string(*Name))
	}
	sub+=1
	if sub==2{
		nalogi()
		sub = 0
	}
	if (*many <= 0)||(*Stop){
		game.Screen = game.Levels[2]
		if end != nil{
			end()
		}
		end = nil
		
		go func(){
			time.Sleep(time.Second * 10)
			os.Exit(0)
		}()
	}
	return nil
}
func (prog *Prog) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
func init(){
	game.Levels = make([]element.Element, 3)
	game.Levels[0], play, Name = Start.Meny()
	game.Levels[1], many, nalogi = UI.MainUiMake(Stop)
	game.Levels[2] = End.END()
	game.Screen = game.Levels[0]
	
	
	
	ebiten.SetWindowIcon([]image.Image{data.Icon})
}
func main(){
	ebiten.SetWindowTitle("Игра")
	ebiten.SetWindowSize(800,700)
	ebiten.SetWindowResizable(1==0)
	ebiten.RunGame(&game)
}