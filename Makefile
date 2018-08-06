setup:
	go get golang.org/dl/go1.11beta3 && go1.11beta3 download

.PHONY: build.example.%
build.example.%:
	GOROOT=~/gowasm GOARCH=wasm GOOS=js go1.11beta3 build -o bin/$*.wasm `go list ./example/$*`

.PHONY: run.example.%
run.example.%:
	$(MAKE) build.example.$*
	go run cmd/server/main.go -wasm-file=bin/$*.wasm