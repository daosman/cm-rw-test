.PHONY: all
all: cm-test

cm-test: cm-test.go
	GOOS=linux go build -o cm-test cm-test.go

.PHONY: install
install:
	cp -p cm-test /usr/local/bin/

.PHONY: clean 
clean:
	rm -f cm-test
