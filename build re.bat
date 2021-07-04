set Name=Game
set MainFile="./"
set GOOS=windows
set GOARCH=amd64
go-winres make --in="./WinDataRes/WindowsDataRes.json" --out="rsrc_%Name%" -arch=%GOARCH%
go build -o="./cmd/%Name% %GOOS% %GOARCH%.exe" -ldflags="-H=windowsgui -v" %MainFile%
pause
set GOOS=windows
set GOARCH=arm
go-winres make --in="./WinDataRes/WindowsDataRes.json" --out="rsrc_%Name%" -arch=%GOARCH%
go build -o="./cmd/%Name% %GOOS% %GOARCH%.exe" -ldflags="-H=windowsgui -v" %MainFile%
set GOOS=windows
set GOARCH=386
go-winres make --in="./WinDataRes/WindowsDataRes.json" --out="rsrc_%Name%" -arch=%GOARCH%
go build -o="./cmd/%Name% %GOOS% %GOARCH%.exe" -ldflags="-H=windowsgui -v" %MainFile%
set GOOS=windows
set GOARCH=amd64p32
go-winres make --in="./WinDataRes/WindowsDataRes.json" --out="rsrc_%Name%" -arch=%GOARCH%
go build -o="./cmd/%Name% %GOOS% %GOARCH%.exe" -ldflags="-H=windowsgui -v" %MainFile%
pause

