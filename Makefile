start-ws:
	go run ws.go

start-api:
	go run api.go

start:
	make -j2 start-ws start-api
