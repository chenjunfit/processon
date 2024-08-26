# GoFrame Template For SingleRepo

Quick Start: 
- https://goframe.org/pages/viewpage.action?pageId=1114399

#gen dao
gf gen dao -p ./internal/app/server 
#
gf gen service -s internal/app/server/logic/ -d internal/app/server/service/
#linux
go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64
go env -w CGO_ENABLED=0 GOOS=windows GOARCH=amd64
2、Linux下编译Mac, Windows平台的64位可执行程序：
$ go env -w CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
$ go env -w CGO_ENABLED=0 GOOS=windows GOARCH=amd64
3、Windows下编译Mac, Linux平台的64位可执行程序：
$ go env -w CGO_ENABLED=0 GOOS=darwin3 GOARCH=amd64
$ go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64