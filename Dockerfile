FROM --platform=linux/amd64 amd64/golang:1.17-alpine3.15 as base

WORKDIR /app

RUN apk update \
 && apk add --no-cache bash curl jq openssl \
 && rm -rf /var/cache/apk/*
 
COPY . .

RUN export PATH=$(go env GOPATH)/bin:$PATH
 

RUN cd cmd/highway-cli && go mod tidy

RUN cd cmd/highway-cli && go build -o highwayd

ENTRYPOINT [ "./highwayd" "serve"]
