# Stage 1
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

# RUN go build -o main .
RUN go build -ldflags="-s -w" -o main .

# Stage 2

FROM alpine:latest 

WORKDIR /app

COPY --from=builder /app/main . /app/

# EXPOSE 8091

CMD ["./main"]