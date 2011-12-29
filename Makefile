include $(GOROOT)/src/Make.inc

TARG=github.com/griffy/htmlfiller
GOFMT=gofmt -s -spaces=true -tabindent=false -tabwidth=4

GOFILES=\
  htmlfiller.go\

include $(GOROOT)/src/Make.pkg

format:
	${GOFMT} -w ${GOFILES}

