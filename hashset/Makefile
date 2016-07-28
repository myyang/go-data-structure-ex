
CC=go
PROJ=hashset
DEBUGFLAGS=-c -gcflags "-N -l"

build:
	$(CC) build -o  $(PROJ).a .

debug:
	$(CC) test $(DEBUGFLAGS)

test:
	make debug && ./$(PROJ).test -test.v

benchalloc:
	make debug && GODEBUG=allocfreetrace=1 ./$(PROJ).test -test.run=None -test.bench . 2>trace.log

bench:
	make debug && ./$(PROJ).test -test.bench .

coverage:
	go test -coverprofile cover.out && go tool cover -html cover.out
