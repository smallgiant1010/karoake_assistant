FROM golang:1.25

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin ./cmd/api-server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /bin .

COPY --from=build /app/etc /app/etc

EXPOSE 8080

CMD ["./bin"]
