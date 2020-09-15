start-ws:
	go run cmd/ws/main.go

start-api:
	go run cmd/api/main.go

start:
	make -j2 start-ws start-api
