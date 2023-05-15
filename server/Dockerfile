FROM --platform=linux/amd64 golang:latest as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./server

FROM --platform=linux/amd64 alpine:latest

WORKDIR /root/

COPY --from=builder /usr/src/app/server .

ENV SERVER_PORT 8080
EXPOSE $SERVER_PORT

CMD ["./server"]