export: clean wasm_exec
	tinygo build -o ./html/wasm.wasm -target wasm -no-debug ./export/wasm.go
	cp ./export/wasm.js ./html/
	cp ./export/index.html ./html/

callback: clean wasm_exec
	tinygo build -o ./html/wasm.wasm -target wasm ./callback/wasm.go
	cp ./callback/wasm.js ./html/
	cp ./callback/index.html ./html/

main: clean wasm_exec
	tinygo build -o ./html/wasm.wasm -target wasm -no-debug ./main/main.go
	cp ./main/index.html ./html/

wasm_exec:
	cp ../../../targets/wasm_exec.js ./html/

clean:
	rm -rf ./html
	mkdir ./html
