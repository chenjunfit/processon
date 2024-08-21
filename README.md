# GoFrame Template For SingleRepo

Quick Start: 
- https://goframe.org/pages/viewpage.action?pageId=1114399

#gen dao
gf gen dao -p ./internal/app/server 
#
gf gen service -s internal/app/server/logic/ -d internal/app/server/service/
#linux
export CGO_ENABLED=0 GOOS=linux GOARCH=amd64