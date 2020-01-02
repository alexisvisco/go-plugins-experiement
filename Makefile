build_plugin:
	go build -buildmode=plugin -o addons/gerard.so  addon-gerard/main.go

run: build_plugin
	@go run cli/main.go gerard


