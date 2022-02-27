build-api:
	go build -o build/api main.go
dev: build-api
	./build/api
gen:
	wire di/wire.go