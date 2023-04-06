go version
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=auto
go mod tidy
#go build -ldflags="-s -w" -o ZeroBot-Plugin.exe
go run main.go 

if [[ $? -eq 0 ]]; then
  sh run.sh
else
  read -p "Press any key to continue..." -n1 -s
  fi