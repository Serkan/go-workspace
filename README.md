##Go Study Lab
This repository contains codes while learning to code in Go. Feel free to send pull requests or 
ask for to become committer. All discussions and contributions are most welcomed.


##TODO in IDEA plugin
1. alt+insert -> select interface and struct to implement and interface
2. Object literal should automaticly prints all fields to fill out or ask which fields to fill
3. Indicate unused or unimported packages by an icon

##Discussions
- [Gitter to discuss what to do next](https://gitter.im/gostudylab/Lobby#)
- [Asana to order and assign task among gophers](https://app.asana.com/0/166591653546069/list)

##Useful things
- [logging techniques for debugging](http://changelog.ca/log/2015/03/09/golang)
- [Effective Go](https://golang.org/doc/effective_go.html) 


## Setup Workspace for IDE support
// gocode is used by many editors to provide intellisense
go get github.com/nsf/gocode

// goimports is something you should run when saving code to fill in import paths
go get golang.org/x/tools/cmd/goimports

// gorename is used by many editors to provide identifier rename support
go get golang.org/x/tools/cmd/gorename

// oracle is a tool that help with code navigation and search
go get golang.org/x/tools/cmd/oracle

// golint should be run after every build to check your code
go get github.com/golang/lint/golint
