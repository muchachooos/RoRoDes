FROM golang:1.19
WORKDIR /app
COPY . .
RUN go mod vendor && go build cmd/main.go
ENTRYPOINT ["./main"]