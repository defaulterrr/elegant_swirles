GO_BIN=$(shell go env GOPATH)/bin
BUF_EXE=$(GO_BIN)/buf$(shell go env GOEXE)

.PHONY: generate
generate: 
	$(BUF_EXE) generate

.PHONY: run
run: 
	go run ./cmd/main.go