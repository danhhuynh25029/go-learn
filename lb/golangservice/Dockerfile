# /golangService/dockerfile
FROM golang:1.21.4-alpine

# defining current working directory 
WORKDIR /app

# coping go.mod and go.sum into current working directory which is "/app"
COPY go.mod ./
# COPY go.sum ./

# downloading dependencies
RUN go mod download

# coping all go files in current directory
COPY *.go ./

# building go project
RUN go build -o /main

# exposing server to port 5000
EXPOSE 5000

# running the go binary 
CMD [ "/main" ]
