FROM golang:1.19.0

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy
# RUN go build -o main ./cmd/main.go

# EXPOSE 3000
# CMD [ "/app/main" ]