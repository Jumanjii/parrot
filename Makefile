SOURCES = main.go\
	  error.go\
	  handler.go\
	  parser.go\
	  store.go\
	  view.go

all: bin/parrot

bin/parrot: $(SOURCES)
	CGO_ENABLED=0 go build -o bin/parrot

run: $(SOURCES)
	CGO_ENABLED=0 go run $(SOURCES)


.PHONY: all run
