FROM golang:latest

WORKDIR /golang-coursework

RUN mkdir backend
RUN mkdir backend/resource 

COPY go.mod go.sum ./
RUN go mod download

COPY backend/resource ./backend/resource

RUN go build -o backend/resource/run ./backend/resource/cmd/app/main.go

ENV LOG_DIR=/backend/resource/logs

RUN mkdir -p $LOG_DIR

CMD ["./backend/resource/run"]