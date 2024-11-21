# Stage 1
FROM golang:alpine

WORKDIR /app

COPY go.mod /app/

RUN go install github.com/air-verse/air@latest

RUN go mod download && go mod verify

EXPOSE 8090

CMD ["air", "-c", ".air.toml"]

