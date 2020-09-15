#build API
FROM golang:1.15 as build-go
WORKDIR /scaling-ws
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/main cmd/api/main.go

# final stage
FROM alpine:latest
RUN addgroup -S scaling-ws && adduser -S scaling-ws -G scaling-ws
USER scaling-ws
WORKDIR /home/scaling-ws
COPY --from=build-go /bin/main ./
EXPOSE 8081
ENTRYPOINT ["./main"]
