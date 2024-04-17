PLUGINS := $(basename $(wildcard ./plugins/*.go))

run: ${PLUGINS}
	go run main.go

${PLUGINS}:
	go build -o $@.so -buildmode=plugin $@.go

clean:
	rm -rf out
