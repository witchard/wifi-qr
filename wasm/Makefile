wasm: wasm.go
	GOARCH=wasm GOOS=js go build

wasm_exec.js:
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" .

serve: wasm wasm_exec.js
	go run serve/main.go

clean:
	rm wasm wasm_exec.js

.PHONY: serve clean
