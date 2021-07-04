go mod init Fast
go install -v "github.com/tc-hib/go-winres@latest"
go get -u -v "github.com/hajimehoshi/ebiten"
go get -u -v "github.com/hajimehoshi/ebiten/ebitenutil"
go get -u -v "github.com/hajimehoshi/ebiten/v2/text"
go get -u -v "golang.org/x/image/font"
go get -u -v "golang.org/x/image/font/opentype"
go get -u -v "github.com/itchyny/timefmt-go"
pause
"./build re.bat"
