
# https://github.com/trstringer/cli-debugging-cheatsheets/blob/master/go.md
#
build:
	go build todolist.go

run:
	go run todolist.go

todoFile: 
	go run .\todolist.go -file ./test.md

debug:
	dlv debug -- --list

setup: install_go
	choco install golang
	go get -u github.com/derekparker/delve/cmd/dlv
	go get github.com/docopt/docopt-go

