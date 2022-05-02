build-api:
	go build -o build/api main.go
dev: build-api
	./build/api
wire:
	wire di/wire.go
installdeps:
	go install github.com/google/wire/cmd/wire@v0.5.0