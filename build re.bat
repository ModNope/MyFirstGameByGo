set Name=Game
set MainFile="./"
set CMD="cmd"
set GOOS=windows
set GOARCH=amd64
go-winres make --in="./WinDataRes/WindowsDataRes.json" --out="rsrc_%Name%" -arch=%GOARCH%
go build -o="./%CMD%/%Name% %GOOS% %GOARCH%.exe" -ldflags="-H=windowsgui -v" %MainFile%
cd %CMD%
go run "../."
pause
cd ../
set GOOS=windows
set GOARCH=arm
go-winres make --in="./WinDataRes/WindowsDataRes.json" --out="rsrc_%Name%" -arch=%GOARCH%
go build -o="./%CMD%/%Name% %GOOS% %GOARCH%.exe" -ldflags="-H=windowsgui -v" %MainFile%
set GOOS=windows
set GOARCH=386
go-winres make --in="./WinDataRes/WindowsDataRes.json" --out="rsrc_%Name%" -arch=%GOARCH%
go build -o="./%CMD%/%Name% %GOOS% %GOARCH%.exe" -ldflags="-H=windowsgui -v" %MainFile%
set GOOS=windows
set GOARCH=amd64p32
go-winres make --in="./WinDataRes/WindowsDataRes.json" --out="rsrc_%Name%" -arch=%GOARCH%
go build -o="./%CMD%/%Name% %GOOS% %GOARCH%.exe" -ldflags="-H=windowsgui -v" %MainFile%
pause

