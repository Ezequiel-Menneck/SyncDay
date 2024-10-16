FROM golang:1.22.4-alpine

WORKDIR /syncday

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

WORKDIR /syncday/cmd/syncday

RUN go build -o /syncday/bin/syncday .

EXPOSE 8080
ENTRYPOINT [ "/syncday/bin/syncday" ]