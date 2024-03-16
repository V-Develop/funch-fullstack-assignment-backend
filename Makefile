GOCMD := go
GORUN := $(GOCMD) run
GOTEST := $(GOCMD) test

start:
	$(GORUN) main.go