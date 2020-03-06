 Expample from the scratch of microservice using golang , mongodb , go-gin as http router 

1 - Delete go mod file then go mod init github.com/ <YourUserName> / <project-name>
 
2 - Change the dockerfile as below

```

FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o <project-name>

# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/<project-name> .

CMD ["./<project-name>"]
 
```

3 - Change the env variables values in .env file
4 - Docker-compose build
5 - Docker-compose up -d
