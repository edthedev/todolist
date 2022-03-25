
# https://github.com/trstringer/cli-debugging-cheatsheets/blob/master/go.md
#
build:
	go build minion.go

run:
	go run todolist.go

todoFile: 
	go run .\minion.go -todo -file test.md

debug:
	dlv debug -- --list

setup: install_go
	choco install golang
	go get -u github.com/derekparker/delve/cmd/dlv
	go get github.com/docopt/docopt-go

