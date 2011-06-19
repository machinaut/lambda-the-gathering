include $(GOROOT)/src/Make.inc

TARG=ajr
GOFILES=ajr.go card.go

all: gofmt ajr

gofmt:
	gofmt -spaces -tabindent=false -tabwidth=4 -w card.go 
	gofmt -spaces -tabindent=false -tabwidth=4 -w ajr.go

include $(GOROOT)/src/Make.cmd
